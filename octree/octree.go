//-----------------------------------------------------------------------------

package octree

//-----------------------------------------------------------------------------

const (
	OCTREE_TYPE_PARENT = iota // child nodes are present
	OCTREE_TYPE_SPACE         // empty space
	OCTREE_TYPE_SOLID         // solid
)

type Octree struct {
	vtype int        // octree type
	child [8]*Octree // octree children
}

//-----------------------------------------------------------------------------
