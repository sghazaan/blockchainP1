package blockchain

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

type Blockchain struct {
	head   *Block
	length int
}

func (bChain *Blockchain) NewBlock(transaction MerkleTree, nonce int, prevBlockHash string) *Block {
	tempBlock := &Block{transaction: transaction, nonce: nonce, prevBlockHash: prevBlockHash, next: nil}
	toHash := string(tempBlock.transaction.RootNode.Data) + strconv.Itoa(tempBlock.nonce)
	tempBlock.BlockHash = CalculateHash((toHash))
	if bChain.head == nil {
		bChain.head = tempBlock
	} else {
		tempPtr := bChain.head
		for tempPtr.next != nil {
			tempPtr = tempPtr.next
		}
		tempPtr.next = tempBlock
	}

	// tempPtr := bChain.head
	// for tempPtr != nil {
	// 	tempPtr = tempPtr.next
	// }
	// tempPtr = tempBlock

	bChain.length++
	return tempBlock
}

func (bChain *Blockchain) DisplayBlockchain() {
	tempPtr := bChain.head
	counter := 1
	for tempPtr != nil {
		fmt.Println("BLOCK -> ", counter)
		fmt.Println("TRANSACTION -> ", tempPtr.transaction)
		fmt.Println("NONCE -> ", tempPtr.nonce)
		fmt.Println("PREVIOUS BLOCK HASH -> ", tempPtr.prevBlockHash)
		fmt.Println("BLOCK HASH -> ", tempPtr.BlockHash)
		fmt.Print("\n\n")
		tempPtr = tempPtr.next
		counter++
	}
}

func (bChain *Blockchain) ChangeBlock(targetLoc int, newTransaction MerkleTree, newNonce int) {
	if bChain.length == 0 {
		fmt.Println("BLOCKCHAIN IS EMPTY !!")
		os.Exit(1)
	}

	if targetLoc > bChain.length || targetLoc < 0 {
		fmt.Println("INSERTED TARGET LOCATION IS INVALID !!!")
		os.Exit(1)
	}
	tempPtr := bChain.head
	for i := 0; i < targetLoc; i++ {
		tempPtr = tempPtr.next
	}

	tempBlock1 := tempPtr
	tempBlock1.transaction = newTransaction
	tempBlock1.nonce = newNonce
	tempBlock1.BlockHash = CalculateHash(string(newTransaction.RootNode.Data) + (strconv.Itoa(newNonce)))
}

func (bchain Blockchain) Mining(block *Block) {

}

func (bChain *Blockchain) VerifyChain() {
	tempPtr := bChain.head
	counter := 0
	for tempPtr.next != nil {

		if tempPtr.BlockHash != tempPtr.next.prevBlockHash {
			fmt.Print("BLOCK # ", counter)
			counter++
			fmt.Print("'s HASH AND BLOCK # ", counter)
			fmt.Print("'s PREVIOUS HASH DO NOT MATCH !!!\n")
			break
		} else {
			tempPtr = tempPtr.next
			counter++
		}
	}
}

func CalculateHash(stringToHash string) string {

	var hash = fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
	//	fmt.Println("STRING RECIEVED TO HASH -> ", stringToHash)
	return hash
}
