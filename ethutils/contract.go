package ethutils

import (
	challenge "Go-Assignment/Contract"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func createSession(instance *challenge.Challenge, auth *bind.TransactOpts) *challenge.ChallengeSession {

	session := &challenge.ChallengeSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 3141592,
		},
	}

	return session
}

func Deploy(privateKey string, client *backends.SimulatedBackend) (*challenge.ChallengeSession, *types.Transaction) {

	auth := bind.NewKeyedTransactor(LoadPrivateKey(privateKey))
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)     // in wei
	// auth.GasLimit = uint64(300000) // in units
	// auth.GasPrice = big.NewInt(100000000000)

	address, tx, instance, err := challenge.DeployChallenge(auth, client)
	client.Commit()
	if err != nil {
		log.Fatal(err)
	}

	session := createSession(instance, auth)

	fmt.Println("New Contract Address", address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	return session, tx
}

func Load(address string, privateKey string, client *ethclient.Client) *challenge.ChallengeSession {

	auth := bind.NewKeyedTransactor(LoadPrivateKey(privateKey))

	instance, err := challenge.NewChallenge(common.HexToAddress(address), client)
	if err != nil {
		log.Fatal(err)
	}

	session := createSession(instance, auth)

	return session
}
