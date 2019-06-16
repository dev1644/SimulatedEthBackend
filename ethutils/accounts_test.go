package ethutils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

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

	privateKey3 := LoadPrivateKey("privateKey1")
	t.Log(privateKey3)
	if privateKey3 == nil {
		t.Errorf("Two private keys can't be same")
	}
}
