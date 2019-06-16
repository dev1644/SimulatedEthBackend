package main

import (
	"Go-Assignment/ethutils"
	utils "Go-Assignment/ethutils"
	"Go-Assignment/ipfsutils"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	ipfs "Go-Assignment/ipfsutils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func deleteFiles() {

	err := os.Remove("Contract/Challenge.bin")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("Contract/Challenge.abi")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("Contract/challenge.go")
	if err != nil {
		log.Fatal(err)
	}
}

func PrintReciept(done chan bool) chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case reciept, open := <-ch:
				if open {
					fmt.Println("Reciept has been added ", reciept)

				} else {
					done <- true
				}
			}
		}
	}()
	return ch
}

func CheckForReciepts(hashes []common.Hash, reciepts chan string, client *backends.SimulatedBackend, ipfsManager *ipfsutils.IpfsManager) {

	for _, val := range hashes {
		var reciept *types.Receipt

		for reciept == nil {
			reciept, _ = client.TransactionReceipt(context.Background(), val)
			time.Sleep(14 * time.Second)
		}
		recieptJson, _ := reciept.MarshalJSON()
		hash, _ := ipfsManager.DagPut(recieptJson, "json", "cbor")
		reciepts <- hash
	}

	close(reciepts)
}

func main() {

	// client, err := ethclient.Dial("https://ropsten.infura.io")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer deleteFiles()

	privateKey := ethutils.NewPrivateKey()
	addr := utils.DeriveAddress(privateKey)
	fmt.Println("Account Address", addr)

	client := ethutils.SimlatingClient(privateKey)

	done := make(chan bool)
	hashes := make([]common.Hash, 0)

	challSession, tx := ethutils.Deploy(privateKey, client)
	ipfsManager, err := ipfs.NewManager("localhost:5001", time.Second)

	if err != nil {
		log.Fatal("Having error connecting with IPFS", err)
	}

	var reciept *types.Receipt

	for reciept == nil {
		reciept, _ = client.TransactionReceipt(context.Background(), tx.Hash())
		time.Sleep(14 * time.Second)
	}

	hashes = append(hashes, tx.Hash())
	reciepts := PrintReciept(done)

	uintb, _ := challSession.GetB()
	tx1, _ := challSession.SetB(big.NewInt(1))
	hashes = append(hashes, tx1.Hash())

	addressa, _ := challSession.GetA()
	tx2, _ := challSession.SetA()
	hashes = append(hashes, tx2.Hash())
	fmt.Println("uint b", uintb)
	fmt.Println("address a", addressa.Hex())

	CheckForReciepts(hashes, reciepts, client, ipfsManager)

	<-done

}
