package settings

import "github.com/enjoypi/god/types"

type Node struct {
	Type string
	ID   types.NodeID
	Apps []string
}
