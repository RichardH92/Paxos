package main

type Sender interface {
	SendMessage(msg Message, nodeData NodeData) (err error)
}
