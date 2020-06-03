package crypto

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/pkg/errors"
)

const (
	staticDeployer      = "0x5B501f5dB54A1652F2f58B4bE7aaDAA117dCE622"
	staticSignature     = "0xffb2d0383ab970139a7e0fa9263c446199464b5778b92bf9e7936b5a383a8fd00abababababababababababababababababababababababababababababababa00"
	staticConfigAddress = "0x46e9742C098267122DA466d6b7a3fb844436Ac37"
)

// DeployConfigContract deploys config contract.
// It requires manually transaction build to make sure we are always deploying to the exact address.
// All the static arguments cannot be changed to allow having the exact address.
func DeployConfigContract(privkey string, client *ethclient.Client) (common.Address, error) {
	err := TransferEth(privkey, common.HexToAddress(staticDeployer), big.NewInt(20000000000000000), client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed transfer eth: %w", err)
	}

	auth := &bind.TransactOpts{
		Nonce:    big.NewInt(0),
		GasPrice: big.NewInt(100000000000),
		GasLimit: 200000,
		Value:    big.NewInt(0),
	}

	auth.Signer = func(types.Signer, common.Address, *types.Transaction) (*types.Transaction, error) {
		tx := types.NewContractCreation(0, auth.Value, auth.GasLimit, auth.GasPrice, common.FromHex(bindings.ConfigBin))
		signer := types.NewEIP155Signer(nil)
		sig := common.FromHex(staticSignature)

		return tx.WithSignature(signer, sig)
	}

	actualAddress, _, _, err := bindings.DeployConfig(auth, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy config contract: %w", err)
	}

	expectedAddress := staticConfigAddress

	if !strings.EqualFold(actualAddress.Hex(), common.HexToAddress(expectedAddress).Hex()) {
		return common.Address{}, fmt.Errorf("config deployed to a wrong address. Expected %q, got %q", common.HexToAddress(expectedAddress).Hex(), actualAddress.Hex())
	}

	return common.HexToAddress(staticConfigAddress), nil
}

// SetupConfig adds required configuration options to the deployed config.
// It sets config owner, channel implementation and accountant implementation addresses.
func SetupConfig(opts *bind.TransactOpts, client bind.ContractBackend, owner string, channelImplProxyAddress, channelImplAddress, accountantImplAddress, accountantImplProxyAddress common.Address) error {
	cfg, err := bindings.NewConfig(common.HexToAddress(staticConfigAddress), client)
	if err != nil {
		return fmt.Errorf("failed to get config by provided address: %w", err)
	}

	_, err = cfg.SetOwner(opts, common.HexToAddress(owner))
	if err != nil {
		return fmt.Errorf("failed to set config owner: %w", err)
	}

	_, err = cfg.AddConfig(opts, bytes32(common.FromHex("0x2ef7e7c50e1b6a574193d0d32b7c0456cf12390a0872cf00be4797e71c3756f7")),
		bytes32(channelImplAddress.Hash().Bytes()))
	if err != nil {
		return fmt.Errorf("failed to add channel implementation to config: %w", err)
	}

	_, err = cfg.AddConfig(opts, bytes32(common.FromHex("0x48df65c92c1c0e8e19a219c69bfeb4cf7c1c123e0c266d555abb508d37c6d96e")),
		bytes32(channelImplProxyAddress.Hash().Bytes()))
	if err != nil {
		return fmt.Errorf("failed to add channel implementation proxy to config: %w", err)
	}

	_, err = cfg.AddConfig(opts, bytes32(common.FromHex("0xe6906d4b6048dd18329c27945d05f766dd19b003dc60f82fd4037c490ee55be0")),
		bytes32(accountantImplAddress.Hash().Bytes()))
	if err != nil {
		return fmt.Errorf("failed to add accountant implementation to config: %w", err)
	}

	_, err = cfg.AddConfig(opts, bytes32(common.FromHex("0x52948fa93a94851571e57fddc2be83c51e0a64bb5e9ca55f4f90439b9802b575")),
		bytes32(accountantImplProxyAddress.Hash().Bytes()))
	if err != nil {
		return fmt.Errorf("failed to add accountant implementation proxy to config: %w", err)
	}

	return nil
}

func TransferEth(from string, target common.Address, amount *big.Int, client *ethclient.Client) error {
	pk, addr, err := ParsePrivkey(from)
	if err != nil {
		return errors.Wrap(err, "could not parse private key")
	}

	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return errors.Wrap(err, "could not get nonce")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not get gas price")
	}
	gasLimit := uint64(21000) // in units

	tx := types.NewTransaction(nonce, target, amount, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
	if err != nil {
		return errors.Wrap(err, "could not sign tx")
	}

	return errors.Wrap(client.SendTransaction(context.Background(), signedTx), "could not send transaction")
}

func ParsePrivkey(privkey string) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privkey)
	if err != nil {
		return nil, common.Address{}, errors.Wrap(err, "could not decode private key")
	}

	return privateKey, ParsePubKey(privateKey), nil
}

func ParsePubKey(privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	addr := crypto.PubkeyToAddress(*publicKey)
	return addr
}

func bytes32(arr []byte) (res [32]byte) {
	copy(res[:], arr)
	return res
}
