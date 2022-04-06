package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
)

func main() {
	os.Setenv("PORT", "1234")
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// create server
	server := CreateServer(port)
	rpc.Register(server)
	rpc.HandleHTTP()

	// initial nodes discovery
	server.DiscoverNodes()

	l, e := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	http.Serve(l, nil)
}
