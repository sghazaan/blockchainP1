package blockchain

type Block struct {
	transaction   MerkleTree //  string
	nonce         int
	prevBlockHash string
	BlockHash     string
	next          *Block
}
