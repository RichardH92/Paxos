package main

import "errors"

type Proposor struct {
	NodeData
}

func (proposor Proposor) HandleMessage(msg Message, nodeData NodeData) (err error) {
	switch msg.Type {
	case accepted:
		return handleAcceptedMessage(msg, nodeData)

	case accept:
		//TODO: implement here
		return nil

	default:
		return errors.New("Invalid message type")
	}
}

func handleAcceptedMessage(msg Message, nodeData NodeData) (err error) {
	return nil
}

func handleAcceptMessage(msg Message, nodeData NodeData) (err error) {
	return nil
}

func (proposor Proposor) AddNode(node NodeData) (err error) {
	switch node.Type {
	default:
		return errors.New("Invalid node type")
	}
}

func (proposor Proposor) RemoveNode(node NodeData) (err error) {
	switch node.Type {
	default:
		return errors.New("Invalid node type")
	}
}
