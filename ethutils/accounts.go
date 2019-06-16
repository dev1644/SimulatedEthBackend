package ethutils

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SimlatingClient(privateKey string) *backends.SimulatedBackend {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)

	if err != nil {
		log.Fatal("Invalid Private Key Provided")
	}

	auth := bind.NewKeyedTransactor(privateKeyECDSA)
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}
	blockGasLimit := uint64(4712388)

	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	return client

}

func NewPrivateKey() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)

	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	return hexutil.Encode(privateKeyBytes)[2:]

}

func LoadPrivateKey(privateKey string) *ecdsa.PrivateKey {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("Invalid Private Key Provided")
	}

	return privateKeyECDSA
}

func DeriveAddress(privateKey string) string {

	privateKeyECDSA := LoadPrivateKey(privateKey)

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address
}

func FetchBalance(privateKey string, client *ethclient.Client) *big.Int {
	address := DeriveAddress(privateKey)
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}
