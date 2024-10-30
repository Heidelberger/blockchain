/*
This is a simple example of how to use the blockchain package.

1. Create a new blockchain instance with a mining difficulty of 2.
2. Record transactions on the blockchain for Alice, Bob, and John.
3. Check if the blockchain is valid; expecting true.

Note: The blockchain package is designed to be used in a real-world application,
where the data is securely stored and transmitted across a network.
In this example, we're just demonstrating the basic functionality of the package.
In a real application, additional security measures, such as encryption and authentication,
should be implemented to protect the data and prevent unauthorized access or tampering.
*/

package main

import (
	"fmt"

	"github.com/Heidelberger/blockchain/blockchain"
)

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	myBlockchain := blockchain.CreateBlockchain(2)

	// record transactions on the blockchain for Alice, Bob, and John
	myBlockchain.AddTransaction("Alice", "Bob", 5)
	myBlockchain.AddTransaction("John", "Bob", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(myBlockchain.IsValid())
}
