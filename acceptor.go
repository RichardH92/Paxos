package main

import (
	"errors"
)

func HandleMessage(msg Message) (retMsg Message, err error) {
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

func addLearner(node Node) {

}

func removeLearner(node Node) {

}
