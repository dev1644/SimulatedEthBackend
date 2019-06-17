package ethutils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var testContractAddress common.Address

func TestCheckForReciepts(t *testing.T) {
	testSimulatingClient = SimlatingClient(testPrivatekey)
	instance, contractAddress, _ := Deploy(testPrivatekey, testSimulatingClient)

	testContractAddress = contractAddress

	uintb, _ := instance.GetB()
	addressa, _ := instance.GetA()

	if uintb.Int64() != 0 || addressa.Hex() != "0x0000000000000000000000000000000000000000" {
		t.Errorf("Contract is initialized with different state")
	}
}
