package main

import (
	"encoding/json"
	"time"
)

func CreateServer(port int) *Server {
	server := Server{ID: int(time.Now().UnixNano()), Database: *NewDatabase(), nodes: make([]Node, 0), selfPort: port}
	return &server
}

func (s *Server) DiscoverNodes() {
	// add self
	s.nodes = append(s.nodes, Node{ID: s.ID, port: s.selfPort})

	// query other ports to discover nodes

	// order nodes by ID to form ring structure
}

func (s *Server) GetValue(key *string, reply *string) error {
	// hash key
	keyHash := hash(*key)

	if s.nodes[int(keyHash)%len(s.nodes)].ID != s.ID {
		// key is on different node, query correct node instead
		for _, node := range s.nodes {
			if node.ID == int(keyHash)%len(s.nodes) {
				// query this node
				*reply = QueryNode(node, key)
			}
		}
	} else {
		// key is on this node, return value from local store
		*reply = s.Database.Get(*key)
	}

	return nil
}

func (s *Server) PushValue(pushEventBytes *[]byte, reply *bool) error {
	// decode args
	var pushEvent PushEvent
	err := json.Unmarshal(*pushEventBytes, &pushEvent)
	if err != nil {
		*reply = false
		return err
	}

	// hash key
	keyHash := hash(pushEvent.Key)

	if s.nodes[int(keyHash)%len(s.nodes)].ID != s.ID {
		// key is supposed to be on a different node, send it to the correct node instead
		for _, node := range s.nodes {
			if node.ID == int(keyHash)%len(s.nodes) {
				// send to this node and forward reply to client
				*reply = PushNode(node, pushEventBytes)
			}
		}
	} else {
		// key is supposed to be on this node, add to local store
		s.Database.Put(pushEvent.Key, pushEvent.Value)
		*reply = true
	}

	return nil
}
