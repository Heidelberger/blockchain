/*
Blockchain package

This package contains utility items to create and manage a simple blockchain system.

The package demonstrates:
1. Creation of a new blockchain
2. Adding transactions to the blockchain
3. Validating the integrity of the blockchain

Key components:
- Blockchain initialization with a specified difficulty
- Transaction processing and block mining
- Blockchain validation

Usage:
- myBlockchain := blockchain.CreateBlockchain(2) // create a new blockchain instance with a mining difficulty of 2
- myBlockchain.AddTransaction("Alice", "Bob", 5) // record a transaction on the blockchain
- fmt.Println(myBlockchain.IsValid()) // ensure that the blockchain is valid

Note: Ensure that the blockchain package is properly imported and all
dependencies are correctly managed in the go.mod file.
*/

package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// The "Block" is the basic component of any blockchain.
type Block struct {
	data         map[string]interface{} // transaction data
	hash         string                 // cryptographic hash used as a unique identifier
	previousHash string                 // a secure link to the previous block. This is the "chain" of the blockchain.
	timestamp    time.Time              // creation time
	pow          int                    // the amount of work to derive this block's hash
}

// This holds the blocks of our blockchain
type Blockchain struct {
	genesisBlock Block   // the very first block
	chain        []Block // all other blocks
	difficulty   int     // the amount of work required to mine a new block
}

// This method calculates the cryptographic hash of a block based on its data, previous hash, and timestamp.
// It uses the SHA-256 hashing algorithm to generate a unique hash value for each block.
func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

// This method mines a new block by adjusting the "proof of work" (PoW) value until the hash meets the required difficulty.
// The difficulty is determined by the number of leading zeros in the hash. A higher difficulty requires more computational power to mine a block.
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

// This function creates a new blockchain with a genesis block and an empty chain.
// The difficulty is set to a default value of 2, which means that the hash must start with two leading zeros to be considered valid.
// The difficulty can be adjusted based on the expected time required to mine a new block and the computational power available.
// A higher difficulty will make it more difficult to mine a new block but will also require more computational power.
func CreateBlockchain(difficulty int) Blockchain {
	// Set the hash of our genesis block to "0". Because it is the first block in the blockchain,
	// there is no value for the previous hash, and the data property is empty.
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

// This method adds a new block to the blockchain with the provided transaction data and
// mining it with the specified difficulty.
// The new block's "previousHash" is set to the hash of the last block in the chain, ensuring
// that the blockchain is a linked list of blocks.
// The new block's "hash" is calculated based on the previous hash, the transaction data, and
// the timestamp. The mining process adjusts the "proof of work" (PoW) value until the hash
// meets the required difficulty.
// The transaction data is stored as a map of key-value pairs "blockData", where the keys and
// values are strings and floats, respectively.
// The amount of work required to mine a new block is stored in the "proof of work" (PoW)
// value of the new block.
func (b *Blockchain) AddTransaction(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

// Recalculate the hash of every block on the blockchain, compare them with the stored hash
// values of the other blocks, and check whether the "previousHash" value of every block
// is equal to the hash value of the block before it.
// If any check fail, the blockchain has been tampered with.
func (b Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
