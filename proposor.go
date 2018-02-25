package main

import "errors"

type Proposor struct {
	NodeData
}

func (proposor Proposor) HandleMessage(msg Message) (retMsg Message, err error) {
	switch msg.Type {

	default:
		return msg, errors.New("Invalid message type")
	}
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
