package node

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"github.com/siftikharm/assignment02bca/blockchain"
)

type Node struct {
	Name         string
	IP           string
	Port         string
	DataStream   net.Listener
	Neighbours   []string
	IsBootstrap  bool
	blockchain   blockchain.Blockchain
	Transactions []string
}

type Neighbour struct {
	Neighbours []string
}

var neigh bool = false

func NewNode(port string) *Node {
	node := new(Node)
	node.Port = port
	node.IP = "127.0.0.1"
	node.Neighbours = []string{}
	node.IsBootstrap = false

	return node
}

func (node *Node) LaunchServer() {
	var err error
	for {
		node.DataStream, err = net.Listen("tcp", ":"+(node.Port))
		//fmt.Println("Node", node.IP, ":", node.Port, " running Server")
		if err != nil {
			log.Fatal(err)
		}
		conn, err := node.DataStream.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		node.DataStream.Close()
		//ConnectToNeighbour
		go handleConnection(conn, node)

	}
}

func ConnectWithBootstrap(bootStrap string) *Node {
	conn, err := net.Dial("tcp", ":"+(bootStrap))

	if err != nil {
		log.Fatal(err)
	} else {

		var node Node
		var neigh Neighbour

		gobDecoder := gob.NewDecoder(conn)

		errr := gobDecoder.Decode(&node)
		if errr != nil {
			log.Fatal(errr)
		}

		err := gobDecoder.Decode(&neigh)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("received = %+v", neigh.Neighbours)

		node.Neighbours = append(node.Neighbours, neigh.Neighbours...)
		return &node

	}
	return nil
}

func handleConnection(con net.Conn, node *Node) {
	//	fmt.Println("Client has connected", con.RemoteAddr())

	var message string
	if neigh == true {
		neigh = false
		addNeigh(con, node)
		return
	}
	gobDecoder := gob.NewDecoder(con)
	{
		err := gobDecoder.Decode(&message)
		fmt.Println("received message = " + message)
		//fmt.Println()
		if err != nil {
			log.Fatal(err)
		}

		if message == "addNeigh" {
			neigh = true
		} else {
			if !contains(node.Transactions, message) {
				node.Transactions = append(node.Transactions, message)
				if len(node.Transactions) == 5 {
					fmt.Println(node.Transactions)
					//Todo something
					//tree := blockchain.NewMerkleTree(node.Transactions)
					//block := node.blockchain.NewBlock(*tree , rand.Intn(1500) , node.blockchain.   )

					node.Transactions = []string{}

				}
			}

		}
	}

}
func contains(nodes []string, port string) bool {
	for _, node := range nodes {
		if port == node {
			return true
		}
	}

	return false
}

func addNeigh(con net.Conn, node *Node) {

	var neigh string
	gobDecoder := gob.NewDecoder(con)
	err := gobDecoder.Decode(&neigh)
	//fmt.Println("received Neighbour = " + neigh)
	if err != nil {
		log.Fatal(err)
	}

	node.Neighbours = append(node.Neighbours, neigh)

}

func (node Node) Connect(port string, msg string) {
	conn, err := net.Dial("tcp", ":"+(port))

	if err != nil {
		log.Fatal(err)
	} else {

		gobEncoder := gob.NewEncoder(conn)
		err := gobEncoder.Encode(msg)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func (node *Node) GetNodeIPPort() string {
	return node.IP + ":" + (node.Port)
}
