package transact

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func MakeSigner(signingKey string, chainId int) *bind.TransactOpts {
	// Load your private key (securely)
	privateKey, err := crypto.HexToECDSA(signingKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create an auth object to use for the transaction
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainId)))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	return auth
}
