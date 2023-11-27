package main

import (
	"fmt"

// 	Node "github.com/siftikharm/assignment02bca/node"
// )

func main() {
	// node := Node.ConnectWithBootstrap("4000")
	// go node.LaunchServer()

	for _, neigh := range node.Neighbours {
		node.Connect(neigh, "addNeigh")
		node.Connect(neigh, node.Port)
	}

	for {
		var inp string
		fmt.Print("\nEnter Transaction : ")
		fmt.Scan(&inp)
		node.Transactions = append(node.Transactions, inp)
		for _, neigh := range node.Neighbours {
			node.Connect(neigh, inp)
		}

	}

}

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	Net "github.com/siftikharm/assignment02bca/network"
// )

// func createP2PNetwork(numInitialNodes int) Net.Network {
// 	network := Net.Network{}

// 	i := 0
// 	port := 6000
// 	for ; i < numInitialNodes; i++ {
// 		network.JoinNetwork(strconv.Itoa(port))
// 		port++
// 	}

// 	bootstrapNode := network.MakeBootstrapNode()

// 	for _, node := range network.GetNodes() {

// 		if node.Port != bootstrapNode.Port {
// 			network.AddNodeToNeighs(bootstrapNode.Port, node.Port)
// 			// node.ConnectWithBootstrap(bootstrapNode.Port)
// 			// bootstrapNode.ConnectWithBootstrap(node.Port)
// 		}

// 	}
// 	for _, node := range network.GetNodeWithPort(network.GetBootstrapIP()).Neighbours {
// 		fmt.Print(node + " ")
// 	}
// 	fmt.Println("")

// 	return network
// }

// // func JoinNetwork(port int, network Net.Network) {
// // 	for _, node := range network.GetNodes() {
// // 		if node.Port == port {
// // 			fmt.Print("Node with the IP already exists in the network\n")
// // 			return
// // 		}

// // 	}

// // 	{
// // 		node := network.NewNode(port)
// // 		node.Name = strconv.Itoa(port)
// // 		network.nodes = append(network.nodes, node)
// // 		network.connectWithBootStrap(node)
// // 		//	network.connectWithRandoms(node)
// // 		//node.ConnectWithBootstrap(network.GetNodeWithPort(BootstrapPort))
// // 		go node.LaunchServer()
// // 	}
// // }

// // func joinNetwork(port int, network *Net.Network) {
// // 	network.JoinNetwork(port)
// // 	bootstrapPort := network.GetBootstrapIP()
// // 	conn := network.GetNodeWithPort(port).ConnectWithBootstrap(bootstrapPort)

// // 	var neighbours []*Net.Node
// // 	dec := gob.NewDecoder(conn)
// // 	err := dec.Decode(&neighbours)

// // 	fmt.Print(neighbours)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // }

// func main() {

// 	// data1 := []string{"My Transacation 1", "My Transacation 2", "My Transacation 3", "My Transacation 4"}
// 	// t1 := blockchain.NewMerkleTree(data1)
// 	// data2 := []string{"My Trans 1", "My Trans 2", "My Trans 3", "My Trans 4"}
// 	// t2 := blockchain.NewMerkleTree(data2)
// 	// data3 := []string{"My Transac 1", "My Transac 2", "My Transac 3", "My Transac 4"}
// 	// t3 := blockchain.NewMerkleTree(data3)
// 	// data4 := []string{"My T 1", "My T 2", "My T 3", "My T 4", "My T 5", "My T 6", "My T 7", "My T 8", "My T 9", "My T 10"}
// 	// t4 := blockchain.NewMerkleTree(data4)

// 	// tempBlockchain := new(blockchain.Blockchain)
// 	// block := tempBlockchain.NewBlock(*t1, 5574, "GENESIS BLOCK")
// 	// block = tempBlockchain.NewBlock(*t2, 7717, block.BlockHash)
// 	// block = tempBlockchain.NewBlock(*t3, 6872, block.BlockHash)
// 	// tempBlockchain.DisplayBlockchain()
// 	// tempBlockchain.ChangeBlock(1, *t4, 2340)
// 	// tempBlockchain.displayBlockchain()
// 	// tempBlockchain.VerifyChain()
// 	// blockchain.DisplayMerkleTree(t4.RootNode)

// 	initialNodes := 6
// 	network := createP2PNetwork(initialNodes)

// 	// go network.GetNodeWithPort(6001).ConnectWithBootstrap(4000)
// 	time.Sleep(time.Second * 10)
// 	fmt.Println("---------")
// 	// go network.GetNodeWithPort(6002).ConnectWithBootstrap(6001)
// 	network.JoinNetwork(strconv.Itoa(4000))
// 	network.JoinNetwork(strconv.Itoa(4001))
// 	network.JoinNetwork(strconv.Itoa(4002))
// 	network.JoinNetwork(strconv.Itoa(4003))
// 	time.Sleep(time.Second * 5)

// 	network.DisplayNetwork()
// 	fmt.Println("---------")

// }
