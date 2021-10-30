package stdlib

import "github.com/enjoypi/god/types"

// thread safe
// post message to actor, if no in this nod, then post message to the node of the actor
func Post(sender types.ActorID, receiver types.ActorID, message types.Message) {
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

func (n *Node) Post(sender types.ActorID, receiver types.ActorID, message types.Message) {

}

func GetNode(receiver types.ActorID) *Node {
	return nil
}
