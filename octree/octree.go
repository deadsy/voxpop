//-----------------------------------------------------------------------------

package octree

import (
	"fmt"
	"strings"
)

//-----------------------------------------------------------------------------

type Properties struct {
	Type int // node type
}

type NodeSet [8]*Node

type Node struct {
	Props    *Properties // node properties
	Children *NodeSet    // child node pointers
}

//-----------------------------------------------------------------------------

func (n *Node) String() string {
	if n.Children == nil {
		return "node has no children"
	}
	var s []string
	for i := 0; i < 8; i++ {
		s = append(s, fmt.Sprintf("child %d", i))
	}
	return strings.Join(s, "\n")
}

//-----------------------------------------------------------------------------

// Return a root level octree
func Alloc() *Node {
	var n Node
	return &n
}

//-----------------------------------------------------------------------------

// Add children to a node
func (n *Node) Divide() {
	var child NodeSet
	for i := 0; i < 8; i++ {
		child[i] = Alloc()
	}
	n.Children = &child
}

//-----------------------------------------------------------------------------
