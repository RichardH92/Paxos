package main

import "errors"

type Learner struct {
	NodeData
}

func (learner Learner) HandleMessage(msg Message) (retMsg Message, err error) {
	switch msg.Type {

	case accepted:
		return learner.handleAccepted(msg)

	default:
		return msg, errors.New("Invalid message type")
	}
}

func (learner Learner) handleAccepted(msg Message) (retMsg Message, err error) {

}

func (learner Learner) AddNode(node NodeData) (err error) {
	switch node.Type {
	default:
		return errors.New("Invalid node type")
	}
}
