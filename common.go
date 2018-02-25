package main

type Message struct {
	Type MessageType
}

type MessageType int

// iota reset: it will be 0.
const (
	prepare MessageType = iota
	promise
	propose
	accept
	accepted
)

type Node struct {
	Type NodeType
}

type NodeType int

// iota reset: it will be 0.
const (
	proposor NodeType = iota
	acceptor
	learner
)
