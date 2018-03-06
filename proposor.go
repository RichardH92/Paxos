package main

import "errors"

type Proposor struct {
	NodeData
}

func (proposor *Proposor) HandleMessage(msg Message, nodeData *NodeData) (err error) {
	switch msg.Type {
	case promise:
		return proposor.handlePromiseMessage(msg, nodeData)

	case accept:
		return proposor.handleAcceptMessage(msg, nodeData)

	default:
		return errors.New("Invalid message type")
	}
}

func (proposor *Proposor) handlePromiseMessage(msg Message, nodeData *NodeData) (err error) {
	return nil
}

func (proposor *Proposor) handleAcceptMessage(msg Message, nodeData *NodeData) (err error) {
	return nil
}

func (proposor *Proposor) AddNode(node *NodeData) (err error) {
	switch node.Type {
	default:
		return errors.New("Invalid node type")
	}
}

func (proposor *Proposor) RemoveNode(node *NodeData) (err error) {
	switch node.Type {
	default:
		return errors.New("Invalid node type")
	}
}
