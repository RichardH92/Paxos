package main

import "errors"

type Learner struct {
	NodeData
}

func (learner Learner) HandleMessage(msg Message, nodeData NodeData) (retMsg Message, err error) {
	switch msg.Type {

	case accepted:
		return learner.handleAccepted(msg)

	default:
		return msg, errors.New("Invalid message type")
	}
}

func (learner Learner) handleAccepted(msg Message) (retMsg Message, err error) {
	var ret Message
	return ret, nil
}

func (learner Learner) AddNode(node NodeData) (err error) {
	switch node.Type {
	default:
		return errors.New("Invalid node type")
	}
}

func (learner Learner) RemoveNode(node NodeData) (err error) {
	return nil
}
