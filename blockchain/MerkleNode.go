package blockchain

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  string
	hash  string
}
