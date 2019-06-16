package main

import (
	"Go-Assignment/ethutils"
	utils "Go-Assignment/ethutils"
	"Go-Assignment/ipfsutils"
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	ipfs "Go-Assignment/ipfsutils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
)

/* func deleteFiles() {

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
} */

// PrintReciept executes a go routine that recieves reciepts via channel and prints them to STDOUT
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

//CheckForReciepts is a helper function that generate reciepts from hashes and store them on IPFS.
func CheckForReciepts(hashes []common.Hash, reciepts chan string, client *backends.SimulatedBackend, ipfsManager *ipfsutils.IpfsManager) {

	for _, val := range hashes {

		reciept, _ := client.TransactionReceipt(context.Background(), val)
		recieptJSON, _ := reciept.MarshalJSON()
		hash, _ := ipfsManager.DagPut(recieptJSON, "json", "cbor")
		reciepts <- hash
	}

	close(reciepts)
}

func main() {

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

	hashes = append(hashes, tx.Hash())
	reciepts := PrintReciept(done)

	uintb, _ := challSession.GetB()
	tx1, _ := challSession.SetB(big.NewInt(1))
	client.Commit()
	hashes = append(hashes, tx1.Hash())

	addressa, _ := challSession.GetA()
	tx2, _ := challSession.SetA()
	client.Commit()
	hashes = append(hashes, tx2.Hash())
	fmt.Println("uint b", uintb)
	fmt.Println("address a", addressa.Hex())

	CheckForReciepts(hashes, reciepts, client, ipfsManager)

	<-done

}
