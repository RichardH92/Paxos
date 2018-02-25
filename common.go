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

type NodeData struct {
	Type NodeType
}

type Node interface {
	HandleMessage(msg Message, nodeData NodeData) (retMsg Message, err error)
	AddNode(node NodeData) (err error)
	RemoveNode(node NodeData) (err error)
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
