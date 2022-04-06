package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	ID   int
	port int
}

type PushEvent struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Server struct {
	ID       int
	selfPort int
	Database Database
	nodes    []Node
}

func CreateServer(port int) *Server {
	server := Server{ID: 0, Database: *NewDatabase(), nodes: make([]Node, 0), selfPort: port}
	return &server
}

func (s *Server) DiscoverNodes() {
	// add self
	s.nodes = append(s.nodes, Node{ID: s.ID, port: s.selfPort})

	// query other ports to discover nodes
}

func (s *Server) GetValue(key *string, reply *string) error {
	keyHash := hash(*key)
	fmt.Printf("%v", int(keyHash))

	if int(keyHash)%len(s.nodes) != s.ID {
		// query correct node
		for _, node := range s.nodes {
			if node.ID == int(keyHash)%len(s.nodes) {
				// query this node
				*reply = QueryNode(node, key)
			}
		}
	} else {
		*reply = s.Database.Get(*key)
	}

	return nil
}

func (s *Server) PushValue(pushEventBytes *[]byte, reply *bool) error {
	var pushEvent PushEvent
	err := json.Unmarshal(*pushEventBytes, &pushEvent)
	if err != nil {
		*reply = false
		return err
	}

	keyHash := hash(pushEvent.Key)

	if int(keyHash)%len(s.nodes) != s.ID {
		// query for correct node
		for _, node := range s.nodes {
			if node.ID == int(keyHash)%len(s.nodes) {
				// query this node
				*reply = PushNode(node, pushEventBytes)
			}
		}
	} else {
		s.Database.Put(pushEvent.Key, pushEvent.Value)
		*reply = true
	}

	return nil
}
