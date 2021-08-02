package crypto

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

type PayAndSettleBeneficiaryPayload struct {
	Beneficiary                    common.Address
	Signature                      []byte
	ChainID                        int64
	ProviderChannelIDForWithdrawal string
	Amount                         *big.Int
	R                              [32]byte
}

func NewPayAndSettleBeneficiaryPayload(beneficiary common.Address, chainID int64, providerChannelIDForWithdrawal string, amount *big.Int, r [32]byte) *PayAndSettleBeneficiaryPayload {
	if hasHexPrefix(providerChannelIDForWithdrawal) {
		providerChannelIDForWithdrawal = providerChannelIDForWithdrawal[2:]
	}

	return &PayAndSettleBeneficiaryPayload{
		Beneficiary:                    beneficiary,
		ChainID:                        chainID,
		ProviderChannelIDForWithdrawal: providerChannelIDForWithdrawal,
		Amount:                         amount,
		R:                              r,
	}
}

func (pasp *PayAndSettleBeneficiaryPayload) getMessage() []byte {
	message := []byte{}
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(pasp.ChainID))
	message = append(message, Pad(b, 32)...)
	message = append(message, Pad(common.Hex2Bytes(pasp.ProviderChannelIDForWithdrawal), 32)...)
	message = append(message, Pad(math.U256(pasp.Amount).Bytes(), 32)...)
	message = append(message, pasp.R[:]...)
	message = append(message, pasp.Beneficiary.Bytes()...)
	return message
}

func (pasp *PayAndSettleBeneficiaryPayload) Sign(ks hashSigner, signer common.Address) error {
	message := pasp.getMessage()
	hash := crypto.Keccak256(message)

	signature, err := ks.SignHash(
		accounts.Account{Address: signer},
		hash,
	)
	if err != nil {
		return err
	}

	ReformatSignatureVForBC(signature)

	pasp.Signature = signature
	return nil
}

// RecoverSigner recovers signer address out of promise signature
func (pasp *PayAndSettleBeneficiaryPayload) RecoverSigner() (common.Address, error) {
	sig := make([]byte, 65)
	copy(sig, pasp.Signature)

	err := ReformatSignatureVForRecovery(sig)
	if err != nil {
		return common.Address{}, err
	}

	return RecoverAddress(pasp.getMessage(), sig)
}
