package main

import (
	"fmt"
	"github.com/deadsy/voxpop/octree"
)

func main() {
	t := octree.Alloc()
	fmt.Printf("%s\n", t)
}
