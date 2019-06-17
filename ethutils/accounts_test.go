package ethutils

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var testPrivatekey = "f11b79dcb87e3811bad41cc68b86d07194c1a2fc70e5b9895b29049cc4357c68"
var testAccountAddress = "0x894432959D4A3e11B3355c4ed1cd79C6795308b2"
var testSimulatingClient *backends.SimulatedBackend

func TestSimlatingClient(t *testing.T) {
	testSimulatingClient = SimlatingClient(testPrivatekey)
	privateKeyECDSA := LoadPrivateKey(testPrivatekey)
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	balance, _ := testSimulatingClient.BalanceAt(context.Background(), address, big.NewInt(0))

	expectedBalance := new(big.Int)
	expectedBalance.SetString("10000000000000000000", 10)

	if balance.String() != expectedBalance.String() {
		t.Errorf("Balance should be same")
	}
}

func TestNewPrivateKey(t *testing.T) {
	privateKey1 := NewPrivateKey()
	privateKey2 := NewPrivateKey()
	if privateKey1 == privateKey2 {
		t.Errorf("Two private keys can't be same")
	}
}

func TestLoadPrivateKey(t *testing.T) {
	privateKey1 := NewPrivateKey()
	privateKey2 := LoadPrivateKey(privateKey1)

	if privateKey1 != hexutil.Encode(crypto.FromECDSA(privateKey2))[2:] {
		t.Errorf("Two private keys can't be same")
	}
}

func TestDeriveAddress(t *testing.T) {

	address := DeriveAddress(testPrivatekey)

	if address != testAccountAddress {
		t.Errorf("Address should be same")
	}
}
