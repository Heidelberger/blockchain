package main

import (
	"fmt"

	"github.com/Heidelberger/blockchain/blockchain"
)

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	myBlockchain := blockchain.CreateBlockchain(2)

	// record transactions on the blockchain for Alice, Bob, and John
	myBlockchain.AddBlock("Alice", "Bob", 5)
	myBlockchain.AddBlock("John", "Bob", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(myBlockchain.IsValid())
}
