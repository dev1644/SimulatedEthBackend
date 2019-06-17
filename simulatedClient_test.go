package main

import (
	"Go-Assignment/ethutils"
	ipfs "Go-Assignment/ipfsutils"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var testPrivatekey = "f11b79dcb87e3811bad41cc68b86d07194c1a2fc70e5b9895b29049cc4357c68"
var testAccountAddress = "0x894432959D4A3e11B3355c4ed1cd79C6795308b2"
var testIpfsManager, _ = ipfs.NewManager("localhost:5001", time.Second)
var done = make(chan bool)
var hashes = make([]common.Hash, 0)
var testReciepts chan string

func TestPrintReciept(t *testing.T) {
	testReciepts = PrintReciept(done)
}
func TestCheckForReciepts(t *testing.T) {
	testSimulatingClient := ethutils.SimlatingClient(testPrivatekey)
	instance, _, tx := ethutils.Deploy(testPrivatekey, testSimulatingClient)

	hashes = append(hashes, tx.Hash())

	uintb, _ := instance.GetB()
	tx1, _ := instance.SetB(big.NewInt(1))
	testSimulatingClient.Commit()
	hashes = append(hashes, tx1.Hash())

	addressa, _ := instance.GetA()
	tx2, _ := instance.SetA()
	testSimulatingClient.Commit()
	hashes = append(hashes, tx2.Hash())
	fmt.Println("uint b", uintb)
	fmt.Println("address a", addressa.Hex())

	CheckForReciepts(hashes, testReciepts, testSimulatingClient, testIpfsManager)
}
