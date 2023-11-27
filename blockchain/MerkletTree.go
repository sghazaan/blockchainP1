package blockchain

import (
	"fmt"
	"math"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

func newMerkleNode(left, right *MerkleNode, data string) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		node.Data = data
		hash := CalculateHash(data)
		node.hash = hash
	} else {
		node.Data = (left.hash + right.hash)
		hash := CalculateHash(node.Data)
		node.hash = hash
	}

	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(data []string) *MerkleTree {
	var leafNodes []MerkleNode
	val := math.Log2(float64(len(data)))

	//Making leafs in power of 2
	if val != math.Floor(val) {
		data = append(data, data...)
		data = data[0:int(math.Pow(2, math.Floor(math.Log2(float64(len(data))))))]
	}

	//making leaf nodes
	for _, dat := range data {
		node := newMerkleNode(nil, nil, dat)
		leafNodes = append(leafNodes, *node)
	}

	//generating nodes, starting from leaf to root node
	for i := 0; i < int(math.Log2(float64(len(data)))); i++ {
		var level []MerkleNode
		//generating nodes level by level, starting from leaf to root node
		for j := 0; j < len(leafNodes); j += 2 {
			node := newMerkleNode(&leafNodes[j], &leafNodes[j+1], ((leafNodes[j+1].Data) + (leafNodes[j].Data)))
			level = append(level, *node)
		}
		leafNodes = level
	}
	//final root node
	tree := MerkleTree{&leafNodes[0]}

	return &tree
}

var COUNT = int(10)

func displayMerkleTreeNodes(root *MerkleNode, space int) {
	// Base case
	if root == nil {
		return
	}
	// Increase distance between levels
	space += COUNT

	// Process right child first
	displayMerkleTreeNodes(root.Right, space)

	for i := COUNT; i < space; i++ {
		fmt.Print(" ")
	}
	// Print current node after space
	fmt.Println(" " + root.Data + " ")

	// Process left child
	displayMerkleTreeNodes(root.Left, space)
}

func DisplayMerkleTree(root *MerkleNode) {

	COUNT = 10
	fmt.Println()
	fmt.Println("Displaying Merkle Tree")
	displayMerkleTreeNodes(root, 0)
	// if root != nil {
	// 	fmt.Println(root.Data)
	// }
	// if root == nil {
	// 	return
	// }
	// displayMerkleTree(root.Left)
	// displayMerkleTree(root.Right)

}
