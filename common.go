package main

type Message struct {
	Type MessageType
}

type MessageType int

// iota reset: it will be 0.
const (
	noop MessageType = iota
	prepare
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

var (
	MessageTypeToNodeTypeHandlerMap = map[MessageType]NodeType{
		prepare:  acceptor,
		promise:  proposor,
		propose:  acceptor,
		accept:   proposor,
		accepted: learner,
	}
)
