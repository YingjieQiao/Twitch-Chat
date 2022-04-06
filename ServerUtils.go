package main

import (
	"hash/fnv"
	"log"
	"net/rpc"
	"strconv"
)

// QueryNode queries specific node for key
func QueryNode(node Node, key *string) string {
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(node.port))
	if err != nil {
		log.Fatal("Dialing: ", err)
	}

	var reply string

	err = client.Call("Server.GetValue", key, &reply)
	if err != nil {
		log.Fatal("Server.GetValue error:", err)
	}

	return reply
}

// PushNode queries specific node for key
func PushNode(node Node, key *[]byte) bool {
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(node.port))
	if err != nil {
		log.Fatal("Dialing: ", err)
	}

	var reply bool

	err = client.Call("Server.PushValue", key, &reply)
	if err != nil {
		log.Fatal("Server.GetValue error:", err)
	}

	return reply
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
