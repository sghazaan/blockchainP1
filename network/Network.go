package network

import (
	"fmt"
	"math/rand"
	"time"

	Node "github.com/sghazaan/blockchainP1/node"
)

type Neighbour struct {
	Neighbours []string
}
type Network struct {
	nodes []*Node.Node
}

var BootstrapIP string = ""
var BootstrapPort string = "-1"

// Add a node to the blockchain Peer to peer network
func (network *Network) JoinNetwork(port string) *Node.Node {

	for _, node := range network.GetNodes() {
		if node.Port == port {
			fmt.Print("Node with the IP already exists in the network\n")
			return nil
		}

	}

	node := Node.NewNode(port)
	node.Name = (port)
	network.nodes = append(network.nodes, node)
	//fmt.Println("Node created with port", node.Port)
	//fmt.Println(network.nodes)
	if len(network.nodes) < 2 {
		return node
	}
	fmt.Println("Node ")

	// network.ConnectWithBootStrap(node)
	//	network.connectWithRandoms(node)
	//node.ConnectWithBootstrap(network.GetNodeWithPort(BootstrapPort))
	//go node.LaunchServer()
	return node

}

func (network *Network) ConnectWithBootStrap(node *Node.Node) bool {
	if BootstrapPort == "-1" {
		return false
	}
	//neighs := Node.ConnectWithBootstrap((BootstrapPort))
	network.AddNodeToNeighs(node.Port, BootstrapPort)
	//network.ConnectWithRandoms(neighs, node)
	//network.connectWithRandoms(neigh, node)
	return true
}

func (network *Network) ConnectWithRandoms(neighs []string, node *Node.Node) bool {
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(len(neighs))

	fmt.Println(val, len(neighs))
	slice := make([]int, len(neighs))
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println(slice)
	for i := 0; i < val; i++ {
		network.AddNodeToNeighs(neighs[(i)], node.Port)
	}
	return true
}

// Connect a node to the nodes in the network
func (network *Network) AddNodeToNeighs(to, from string) {
	fromNode := network.GetNodeWithPort(from)
	toNode := network.GetNodeWithPort(to)
	if fromNode == nil || toNode == nil {
		fmt.Print("Node with the port does not exist in the network \n")
	} else if contains(fromNode.Neighbours, to) || contains(toNode.Neighbours, from) {
		fmt.Print("Already Connected\n")
	} else {
		fromNode.Neighbours = append(fromNode.Neighbours, to)
		toNode.Neighbours = append(toNode.Neighbours, from)
	}
}

func (network *Network) MakeBootstrapNode() *Node.Node {
	var node *Node.Node = network.nodes[rand.Intn(len(network.nodes))]
	node.IsBootstrap = true
	BootstrapIP = node.IP
	BootstrapPort = node.Port
	fmt.Println("Node", node.IP, ":", node.Port, " is the Bootstrap Node ")
	return node
}

func (network *Network) GetNodeWithPort(port string) *Node.Node {
	for i, node := range network.nodes {
		if node.Port == port {
			return network.nodes[i]
		}
	}

	fmt.Printf("Node with the Port : %v does not exist in the network \n", port)
	return nil
}

// Display the network as an adjacency list
func (network *Network) DisplayNetwork() {
	fmt.Println("\t\t ** Network **")
	for _, node := range network.nodes {
		fmt.Printf("Node IP %+v ", node.GetNodeIPPort())
		for _, node := range node.Neighbours {
			fmt.Printf(" %+v", node)
		}
		fmt.Println()
	}
}

// IP addresses in the network should be unique
func contains(nodes []string, port string) bool {
	for _, node := range nodes {
		if port == node {
			return true
		}
	}

	return false
}

func (network *Network) GetNodes() []*Node.Node {
	return network.nodes
}

func (network *Network) GetBootstrapIP() string {
	return BootstrapPort
}
