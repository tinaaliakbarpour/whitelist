package main

import (
	"context"
	"fmt"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/blockchain"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/data"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/examples"
	"merkletree/tree"
	logger "github.com/ElrondNetwork/elrond-go-logger"
)

var log = logger.GetOrCreate("elrond-sdk-erdgo/examples/examplesAccount")


//you can use tree package to construct a new merkle tree ...
//you can use elrond sdk for interacting with elrond smart contracts
//smart contract in src folder is in solidity and is used to implement white list feature
//you can compile it with soll , and deploy it in test ...

func main() {
	// build a merkle tree push data in it and get the proof
	tree := merkletree.New()
	tree.SetIndex(1)
	tree.Push([]byte("a"))
	tree.Push([]byte("b"))
	// The merkle root could be obtained by calling tree.Root(), but will also
	// be provided by tree.Prove()
	merkleRoot,set, proof, proofIndex, numLeaves := tree.Prove()

	fmt.Println("merkleRoot:",merkleRoot,"set:",set, "proof:",proof,"proofIndex:", proofIndex, "numLeaves:",numLeaves )
	fmt.Println("________________________________________________________________________")
	fmt.Println(string(proof[0][0]),string(proof[0][1]))
	x :=tree.Root()
	fmt.Println(string(x[:]))


	//////elrond vm query we can use this part to interact with elrond smart contract
	ep := blockchain.NewElrondProxy(examples.TestnetGateway, nil)

	vmRequest := &data.VmValueRequest{
		Address:    "erd1qqqqqqqqqqqqqpgqp699jngundfqw07d8jzkepucvpzush6k3wvqyc44rx",
		FuncName:   "version",
		CallerAddr: "erd1rh5ws22jxm9pe7dtvhfy6j3uttuupkepferdwtmslms5fydtrh5sx3xr8r",
		CallValue:  "",
		Args:       nil,
	}
	response, err := ep.ExecuteVMQuery(context.Background(), vmRequest)
	if err != nil {
		log.Error("error executing vm query", "error", err)
		return
	}

	contractVersion := string(response.Data.ReturnData[0])
	log.Info("response", "contract version", contractVersion)
}
