package main

import (
	"fmt"
	"log"
)

type sn struct {
	parent *sn
	left   *sn
	right  *sn
	val    int
	isLeaf bool
}

func (x sn) String() string {
	switch {
	case x.isLeaf:
		return fmt.Sprintf("%d", x.val)
	default:
		return fmt.Sprintf("[%s,%s]", x.left, x.right)
	}
}

func main() {
	log.Println("day 18")
	test := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"
	root := &sn{}
	NewSn(test, root)
	result := fmt.Sprint(root)
	if result != test {
		panic("tree failed to build")
	}
	log.Printf("%s -> %s\n", test, root)

}

func NewSn(str string, node *sn) {
	if len(str) == 0 {
		return
	}
	c := str[0]
	switch c {
	case '[': // now dealing with a new left
		left := &sn{}
		left.parent = node
		node.left = left
		NewSn(str[1:], left)
	case ',': // now dealing with a new right
		right := &sn{}
		right.parent = node
		node.right = right
		NewSn(str[1:], right)
	case ']': // finished with a pair
		NewSn(str[1:], node.parent)
	default: // have a value
		node.isLeaf = true
		node.val = int(c - '0')
		NewSn(str[1:], node.parent)
	}

}
