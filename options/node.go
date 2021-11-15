package options

import "github.com/enjoypi/god/def"

type Node struct {
	Type string
	ID   def.NodeID
	Apps []string
}
