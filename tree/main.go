package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	V     int
	Right *Tree
}

func CreateNewTree(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func Traverse(t *Tree) {
	if t == nil {
		return
	}
	Traverse(t.Left)
	fmt.Print(t.V, " ")
	Traverse(t.Right)
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.V {
		return t
	}
	if v < t.V {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := CreateNewTree(10)
	fmt.Println("The value of the root of the tree is", tree.V)
	Traverse(tree)
	fmt.Println("=----")
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	Traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.V)
}
