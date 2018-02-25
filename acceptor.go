package main

import (
	"errors"
)

type Acceptor struct {
	NodeData
}

func (acceptor Acceptor) HandleMessage(msg Message, nodeData NodeData) (err error) {
	switch msg.Type {
	case propose:
		return acceptor.handlePropose(msg, nodeData)

	case prepare:
		return acceptor.handlePrepare(msg, nodeData)

	default:
		return errors.New("Invalid message type")
	}
}

func (acceptor Acceptor) handlePropose(msg Message, nodeData NodeData) (err error) {
	return nil
}

func (acceptor Acceptor) handlePrepare(msg Message, nodeData NodeData) (err error) {
	return nil
}

func (acceptor Acceptor) AddNode(addNode NodeData) (err error) {
	switch addNode.Type {
	case learner:
		acceptor.addLearner(addNode)
		return nil

	default:
		return errors.New("Invalid node type")
	}
}

func (acceptor Acceptor) addLearner(addNode NodeData) {

}

func (accetor Acceptor) RemoveNode(node NodeData) (err error) {
	switch node.Type {
	case learner:
		return nil
	default:
		return errors.New("Invalid node type")
	}
}

func removeLearner(node Node) {

}
