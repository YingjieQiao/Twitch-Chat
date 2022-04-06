package main

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
	nodes    []Node // ordered by ascending ID order
}
