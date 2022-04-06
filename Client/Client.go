package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc"
)

type PushEvent struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":8081") // connect to the node
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	var reply bool
	var reply2 string

	args, err := json.Marshal(PushEvent{Key: "Hello", Value: "There"})
	if err != nil {
		log.Fatal("JSON error:", err)
	}

	err = client.Call("Server.PushValue", args, &reply)
	if err != nil {
		log.Fatal("RPC error:", err)
	}
	fmt.Printf("%v\n", reply) // should be true, pushed successfully

	err = client.Call("Server.GetValue", "Hello", &reply2)
	if err != nil {
		log.Fatal("RPC error:", err)
	}
	fmt.Printf("Hello %v\n", reply2) // should be "There" as value was pushed successfully

}
