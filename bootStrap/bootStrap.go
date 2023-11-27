package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"strconv"

	Net "github.com/sghazaan/blockchainP1//network"
)

var ports []int
var network Net.Network

func main() {

	network = Net.Network{}
	node := network.JoinNetwork("4000")
	ports = append(ports, 4000)
	var err error

	for {
		node.DataStream, err = net.Listen("tcp", ":"+(node.Port))
		fmt.Println("Node", node.IP, ":", node.Port, " running Server")
		if err != nil {
			log.Fatal(err)
		}
		conn, err := node.DataStream.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		node.DataStream.Close()
		node.Neighbours = append(node.Neighbours, strconv.Itoa(ports[len(ports)-1]+1))
		//ConnectToNeighbour
		fmt.Println("neighbours =", node.Neighbours)

		go handleConnectionClients(conn, node.Neighbours)

	}

}

func handleConnectionClients(con net.Conn, neighbours []string) {
	fmt.Println("Client has connected", con.RemoteAddr())

	node := network.JoinNetwork(strconv.Itoa(ports[len(ports)-1] + 1))
	ports = append(ports, ports[len(ports)-1]+1)

	//fmt.Printf("neighbours = %+v \n", neighbours)

	neigh := new(Net.Neighbour)
	neigh.Neighbours = append(neigh.Neighbours, "4000")
	neigh.Neighbours = neighbours[:len(neighbours)-1]
	add := &neigh
	gobEncoder := gob.NewEncoder(con)
	erro := gobEncoder.Encode(node)
	if erro != nil {
		log.Fatal(erro)
	}

	err := gobEncoder.Encode(add)
	if err != nil {
		log.Fatal(err)
	}

}
