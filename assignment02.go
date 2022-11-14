package assignment02

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce       int
	BlockData   []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
	chain     *[]Block
}

func GenerateNonce(blockData []Transaction) int {
	return rand.Intn(100)
}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
			blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	var block Block
	var nonce int
	var hash_value string
	nonce = GenerateNonce(blockData) //Generating Nonce
	hash_value = CalculateHash(blockData, int(nonce))
	block.Nonce = nonce
	block.BlockData = blockData
	block.CurrentHash = hash_value
	if chainHead != nil {
		block.PrevHash = chainHead.CurrentHash
		block.PrevPointer = chainHead
	} else {
		block.PrevHash = ""
		block.PrevPointer = nil
	}
	return &block
}

func ListBlocks(chainHead *Block) {
	temp := chainHead
	var total_blocks []Block
	for temp.PrevPointer != nil {
		total_blocks = append(total_blocks, *temp)
		temp = temp.PrevPointer
	}
	total_blocks = append(total_blocks, *temp)
	// Now temp is pointing to first block
	block_num := 0
	for i := len(total_blocks) - 1; i >= 0; i-- {
		println("\n<------------ Block " + strconv.Itoa(block_num) + " ------------>")
		DisplayTransactions(total_blocks[i].BlockData)
		block_num++
	}
}

func DisplayTransactions(blockData []Transaction) {
	for i := 0; i < len(blockData); i++ {
		println("Transaction ID, ", i+1)
		println("Sender , ", blockData[i].Sender)
		println("Receiver , ", blockData[i].Receiver)
		println("Amount , ", blockData[i].Amount)
	}
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	var transaction Transaction
	transaction.Sender = sender
	transaction.Receiver = receiver
	transaction.Amount = amount
	return transaction
}
