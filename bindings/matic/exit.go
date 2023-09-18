package matic

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	pc "github.com/mysteriumnetwork/payments/crypto"
)

// GetExitParams gets the exit params for matic exit.
func GetExitParams(l2client *ethclient.Client, withdrawTxHash common.Hash, timeout time.Duration, network Network) ([]byte, error) {
	ctx, cancel := contextWithTimeout(timeout)
	defer cancel()
	rcpt, err := l2client.TransactionReceipt(ctx, withdrawTxHash)
	if err != nil {
		return nil, err
	}

	api := NewAPI(network, &http.Client{
		Timeout: time.Second * 20,
	})

	blockIncludedResponse, err := api.GetBlockIncludedResponse(rcpt.BlockNumber)
	if err != nil {
		return nil, err
	}

	start, ok := big.NewInt(0).SetString(blockIncludedResponse.Start, 10)
	if !ok {
		return nil, errors.New("could not parse blockIncludedResponse start")
	}
	end, ok := big.NewInt(0).SetString(blockIncludedResponse.End, 10)
	if !ok {
		return nil, errors.New("could not parse blockIncludedResponse end")

	}

	blockProofResponse, err := api.GetFastMerkleProof(start, end, rcpt.BlockNumber)
	if err != nil {
		return nil, err
	}

	block, err := l2client.BlockByNumber(ctx, rcpt.BlockNumber)
	if err != nil {
		return nil, err
	}

	stateSyncTXHash := getStateSyncTxHash(block)

	receiptProof, err := getReceiptProof(l2client, stateSyncTXHash, rcpt, block, timeout)
	if err != nil {
		return nil, err
	}

	idx := 0
	for i, log := range rcpt.Logs {
		if len(rcpt.Logs) < 3 {
			continue
		}
		if strings.EqualFold(log.Topics[0].Hex(), ERC20_TRANSFER_EVENT_SIG) &&
			strings.EqualFold(log.Topics[0].Hex(), "0x0000000000000000000000000000000000000000000000000000000000000000") {
			idx = i
		}
	}

	headerNumber, err := hex.DecodeString(removeLeadingHexIndicator(blockIncludedResponse.HeaderBlockNumber))
	if err != nil {
		return nil, err
	}

	blockProof, err := hex.DecodeString(removeLeadingHexIndicator(blockProofResponse.Proof))
	if err != nil {
		return nil, err
	}

	blockNumber := rcpt.BlockNumber.Bytes()

	timestamp := block.Time()

	transactionsRoot := block.TxHash().Bytes()

	receiptsRoot := block.ReceiptHash().Bytes()

	receiptBytes, err := getReceiptBytes(rcpt, block)
	if err != nil {
		return nil, err
	}

	encoded := hex.EncodeToString(receiptProof.Path)
	newPath := "00" + encoded
	p, err := hex.DecodeString(newPath)
	if err != nil {
		return nil, err
	}

	bts, err := rlp.EncodeToBytes(receiptProof.ParentNodes)
	if err != nil {
		return nil, fmt.Errorf("could not encode to bytes parent nodes %w", err)
	}

	ep := ExitParams{
		HeaderBlockNumber: headerNumber,
		BlockProof:        blockProof,
		BlockNumber:       blockNumber,
		Timestamp:         timestamp,
		TransactionsRoot:  transactionsRoot,
		ReceiptsRoot:      receiptsRoot,
		ReceiptBytes:      receiptBytes,
		Parents:           bts,
		Path:              p,
		LogIndex:          uint64(idx),
	}

	data, err := rlp.EncodeToBytes(ep)
	return data, err
}

func removeLeadingHexIndicator(in string) string {
	return strings.TrimPrefix(in, "0x")
}

func getStateSyncTxHash(block *types.Block) []byte {
	res := make([]byte, 0)
	res = append(res, []byte("matic-bor-receipt-")...)
	res = append(res, pc.Pad(block.Number().Bytes(), 8)...)
	res = append(res, block.Hash().Bytes()...)
	return crypto.Keccak256(res)
}

// ExitParams the struct for matic exit param encoding.
type ExitParams struct {
	HeaderBlockNumber []byte
	BlockProof        []byte
	BlockNumber       []byte
	Timestamp         uint64
	TransactionsRoot  []byte
	ReceiptsRoot      []byte
	ReceiptBytes      []byte
	Parents           []byte
	Path              []byte
	LogIndex          uint64
}

func getReceiptProof(client *ethclient.Client, stateSyncTxHash []byte, receipt *types.Receipt, block *types.Block, timeout time.Duration) (ReceiptProof, error) {
	db := memorydb.New()
	defer db.Close()

	tr, err := trie.New(trie.TrieID(common.Hash{}), trie.NewDatabase(rawdb.NewMemoryDatabase(), trie.HashDefaults))
	if err != nil {
		return ReceiptProof{}, err
	}

	txs := block.Transactions()
	ctx, cancel := contextWithTimeout(timeout)
	defer cancel()
	receiptsToPaths := make(map[string][]byte)
	for i := range txs {
		if bytes.Equal(txs[i].Hash().Bytes(), stateSyncTxHash) {
			// this is state sync - ignore it
			continue
		}

		// get receipt
		rcp, err := client.TransactionReceipt(ctx, txs[i].Hash())
		if err != nil {
			return ReceiptProof{}, err
		}

		path, err := rlp.EncodeToBytes(rcp.TransactionIndex)
		if err != nil {
			return ReceiptProof{}, err
		}

		bts, err := getReceiptBytes(rcp, block)
		if err != nil {
			return ReceiptProof{}, err
		}

		receiptsToPaths[hex.EncodeToString(bts)] = path

		err = tr.Update(path, bts)
		if err != nil {
			return ReceiptProof{}, err
		}
	}

	encodedIndex, err := rlp.EncodeToBytes(receipt.TransactionIndex)
	if err != nil {
		return ReceiptProof{}, err
	}

	proofer := NewProofer()
	err = tr.Prove(encodedIndex, proofer)
	if err != nil {
		return ReceiptProof{}, err
	}

	prf := ReceiptProof{
		ParentNodes: make([][][]byte, 0),
		BlockHash:   receipt.BlockHash.Bytes(),
		Root:        block.Header().ReceiptHash.Bytes(),
		Path:        encodedIndex,
	}

	for i := range proofer.state {
		a := make([][]byte, 0)
		err = rlp.DecodeBytes(proofer.state[i].value, &a)
		if err != nil {
			return prf, err
		}
		prf.ParentNodes = append(prf.ParentNodes, a)
	}

	return prf, nil
}

func contextWithTimeout(t time.Duration) (context.Context, func()) {
	return context.WithTimeout(context.Background(), t)
}

// ReceiptProof represents the receipt proof for withdrawal.
type ReceiptProof struct {
	BlockHash   []byte
	ParentNodes [][][]byte
	Root        []byte
	Path        []byte
	Value       []byte
}

// NewProofer creates a new instance of proofer.
func NewProofer() *Proofer {
	return &Proofer{
		state: make([]proof, 0),
	}
}

type proof struct {
	key, value []byte
}

// Proofer captures proofs emitted by ethereum merkle trie.
type Proofer struct {
	state []proof
}

// Put puts the proof in the internal state.
func (p *Proofer) Put(key []byte, value []byte) error {
	p.state = append(p.state, proof{
		key:   key,
		value: value,
	})

	return nil
}

// Delete does not do anything, not required in our case.
func (p *Proofer) Delete(key []byte) error {
	return nil
}

func getReceiptBytes(rcp *types.Receipt, block *types.Block) ([]byte, error) {
	res := StructForRLP{}
	res.Status = []byte{0}
	if rcp.Status == types.ReceiptStatusSuccessful {
		res.Status = []byte{1}
	}

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, rcp.CumulativeGasUsed)
	res.GasUsed = removeLeadingZeroes(buf)

	res.Bloom = block.Bloom().Bytes()

	for i := range rcp.Logs {
		log := LogsForRLP{
			Topics: make([][]byte, 0),
		}
		log.Address = rcp.Logs[i].Address.Bytes()
		for j := range rcp.Logs[i].Topics {
			log.Topics = append(log.Topics, rcp.Logs[i].Topics[j].Bytes())
		}
		log.Data = rcp.Logs[i].Data
		res.Logs = append(res.Logs, log)
	}

	bts, err := rlp.EncodeToBytes(res)
	return bts, err
}

// StructForRLP struct we use for rlp encoding the receipt.
type StructForRLP struct {
	Status  []byte
	GasUsed []byte
	Bloom   []byte

	Logs []LogsForRLP
}

// LogsForRLP struct we use for rlp encoding of the receipt logs.
type LogsForRLP struct {
	Address []byte
	Topics  [][]byte
	Data    []byte
}

// there's probably a better way of doing this
func removeLeadingZeroes(in []byte) []byte {
	var leadingZeroes = 0
	for i := range in {
		if in[i] == 0 {
			leadingZeroes++
		} else {
			break
		}
	}

	res := make([]byte, len(in)-leadingZeroes)
	for i := 0; i < len(in)-leadingZeroes; i++ {
		res[i] = in[i+leadingZeroes]
	}
	return res
}
