package registry

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/rand"
	"time"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/MysteriumNetwork/payments/registry/generated"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"fmt"
)

//go:generate abigen --sol ../contracts/registry.sol --pkg generated --out generated/registry.go

func init() {
	rand.Seed(time.Now().UnixNano())  //don't do this at home kids, use better random generators :)
}

type MystIdentity struct {
	PrivateKey  *ecdsa.PrivateKey
	PublicKey * ecdsa.PublicKey
	Address common.Address
}

func NewMystIdentity() (*MystIdentity, error) {
	key , err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &MystIdentity{
		key,
		&key.PublicKey,
		crypto.PubkeyToAddress(key.PublicKey),
	}, nil
}

type ProofOfIdentity struct {
	RandomNumber uint64
	Signature    *DecomposedSignature
}

func (proof * ProofOfIdentity)String() string {
	return fmt.Sprintf("Proof: %+v", *proof)
}

func CreateProofOfIdentity(identity * MystIdentity) (*ProofOfIdentity , error) {
	number := rand.Uint64()
	encodedNumber := make([]byte,8,8)
	binary.BigEndian.PutUint64(encodedNumber, number)
	signature , err := crypto.Sign(crypto.Keccak256([]byte("Register prefix:"), encodedNumber), identity.PrivateKey )
	if err != nil {
		return nil ,err
	}

	decSig , err := DecomposeSignature(signature)
	if err != nil {
		return nil , err
	}

	return &ProofOfIdentity{
		number,
		decSig,
	}, nil
}

type Registry struct {
	generated.IdentityRegistrySession
	Address common.Address
}

func DeployRegistry(owner * bind.TransactOpts , erc20address common.Address, backend bind.ContractBackend) (*Registry, error) {

	address , _ , contract , err := generated.DeployIdentityRegistry(owner, backend, erc20address, big.NewInt(1000))
	if err != nil {
		return nil , err
	}

	return &Registry{
		generated.IdentityRegistrySession{
			TransactOpts: *owner,
			CallOpts: bind.CallOpts{},
			Contract: contract,
		},
		address,
	}, nil
}

func (registry * Registry) RegisterIdentity(proof * ProofOfIdentity) ( * types.Transaction, error) {
	signature := proof.Signature
	return registry.IdentityRegistrySession.RegisterIdentity(proof.RandomNumber , signature.V, signature.R , signature.S)
}