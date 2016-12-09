//-----------------------------------------------------------------------------

package octree

//-----------------------------------------------------------------------------

type Properties struct {
	Type int // node type
}

type Node struct {
	Props *Properties // node properties
	Empty bool        // true if the node is empty (no children)
	Child [8]*Node    // child node pointers
}

//-----------------------------------------------------------------------------

func (n *Node) String() string {
	if n.Empty {
		return "node has 0 children"
	}
	return "node has children"
}

//-----------------------------------------------------------------------------

// Return a root level octree
func Alloc() *Node {
	var n Node
	n.Empty = true
	return &n
}

//-----------------------------------------------------------------------------
