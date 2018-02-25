package main

func main() {
	//Get message
	var msg Message
	//Get nodeData
	var nodeData NodeData

	var learnerNode Learner
	var acceptorNode Acceptor
	var proposorNode Proposor

	switch MessageTypeToNodeTypeHandlerMap[msg.Type] {
	case proposor:
		err := proposorNode.HandleMessage(msg, nodeData)

		if err != nil {
			//handle error here
		}

	case acceptor:
		err := acceptorNode.HandleMessage(msg, nodeData)

		if err != nil {
			//handle error here
		}

	case learner:
		err := learnerNode.HandleMessage(msg, nodeData)

		if err != nil {
			//handle error here
		}

	default:
		//TODO: error
	}
}
