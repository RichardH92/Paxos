package main

import "errors"

type Learner struct {
	NodeData
}

func (learner *Learner) HandleMessage(msg Message, nodeData *NodeData) (err error) {
	switch msg.Type {

	case accepted:
		return learner.handleAccepted(msg)

	default:
		return errors.New("Invalid message type")
	}
}

func (learner *Learner) handleAccepted(msg Message) (err error) {
	return nil
}

func (learner *Learner) AddNode(node *NodeData) (err error) {
	switch node.Type {
	default:
		return errors.New("Invalid node type")
	}
}

func (learner *Learner) RemoveNode(node *NodeData) (err error) {
	return nil
}
