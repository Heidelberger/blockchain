/*
This is a simple blockchain application.

1. Create a new blockchain instance
2. Record some transactions on the blockchain
3. Check if the blockchain is valid

In a real application, additional security measures, such as encryption and authentication,
would be implemented to protect the data and prevent unauthorized access or tampering.
*/

package main

import (
	"fmt"

	"github.com/Heidelberger/blockchain/blockchain"
)

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	myBlockchain := blockchain.CreateBlockchain(2)

	// record some transactions on the blockchain
	myBlockchain.AddTransaction("Alice", "Bob", 5)
	myBlockchain.AddTransaction("John", "Bob", 2)

	// check if the blockchain is valid
	fmt.Println(myBlockchain.IsValid())
}
