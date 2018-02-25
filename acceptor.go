package main

import (
	"errors"
)

type Acceptor struct {
	NodeData
}

func (acceptor Acceptor) HandleMessage(msg Message) (retMsg Message, err error) {
	switch msg.Type {
	case propose:
		return handlePropose(msg)

	case prepare:
		return handlePrepare(msg)

	default:
		return msg, errors.New("Invalid message type")
	}
}

func handlePropose(msg Message) (retMsg Message, err error) {

}

func handlePrepare(msg Message) (retMsg Message, err error) {

}

func (acceptor Acceptor) AddNode(node NodeData) (err error) {
	switch node.Type {
	case learner:
		return nil

	default:
		return errors.New("Invalid node type")
	}
}

func addLearner(node Node) {

}

func removeLearner(node Node) {

}
