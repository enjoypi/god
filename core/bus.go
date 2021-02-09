package core

// thread safe
// post message to actor, if no in this nod, then post message to the node of the actor
func Post(sender ActorID, receiver ActorID, message Message) {
	actor := DefaultActorManager.Get(receiver)
	if actor != nil {
		actor.Post(message)
		return
	}

	for i := 0; i < 10; i++ {
		node := GetNode(receiver)
		if node != nil {
			node.Post(sender, receiver, message)
			return
		}
	}
}

type Node struct {
}

func (n *Node) Post(sender ActorID, receiver ActorID, message Message) {

}

func GetNode(receiver ActorID) *Node {
	return nil
}
