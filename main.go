package main

import (
	"fmt"
	"github.com/deadsy/voxpop/octree"
)

func main() {
	t := octree.Alloc()
	t.Divide()
	fmt.Printf("%s\n", t)
}
