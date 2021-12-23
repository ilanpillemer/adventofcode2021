package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type sn struct {
	parent  *sn
	left    *sn
	right   *sn
	before  *sn
	after   *sn
	val     int
	isLeaf  bool
	ordered []*sn
}

func (x sn) String() string {
	switch {
	case x.isLeaf:
		return fmt.Sprintf("%d", x.val)
	default:
		return fmt.Sprintf("[%s,%s]", x.left, x.right)
	}
}

func (x *sn) Height() int {
	if x.parent == nil {
		return 0
	}
	return 1 + x.parent.Height()
}

func Add(left, right *sn) *sn {
	root := &sn{}
	root.left = left
	root.right = right

	left.parent = root
	right.parent = root

	return root

}

//decorate adds pointers to before and after as explode needs to know this
//and its very useful to what what you need to know
//cache Breadth-First or Level Order Traversal:
var ordered []*sn

func (x *sn) Decorate(start bool) {
	if start {
		ordered = []*sn{}
	}

	switch {
	case x.isLeaf:
		ordered = append(ordered, x)
	default:
		x.left.parent = x
		x.right.parent = x
		x.left.Decorate(false)
		x.right.Decorate(false)
	}
	//now that level order is cached, decorate nodes with the edges
	if start {
		x.ordered = ordered
		for i := range ordered {
			switch {
			case i == 0:
				ordered[i].after = ordered[i+1]
				ordered[i].before = nil
			case i == len(ordered)-1:
				ordered[i].before = ordered[i-1]
				ordered[i].after = nil
			default:
				ordered[i].after = ordered[i+1]
				ordered[i].before = ordered[i-1]
			}
		}
	}
}

func (x *sn) LeftMost(level int) *sn { //only if nested more than four!

	if x.left.isLeaf && x.right.isLeaf && x.Height() >= level {
		//log.Println("LeftMost", x)
		return x
	}

	if !x.left.isLeaf {
		checkLeft := x.left.LeftMost(level)
		if checkLeft != nil {
			return checkLeft
		}
	}
	if !x.right.isLeaf {
		checkRight := x.right.LeftMost(level)
		if checkRight != nil {
			return checkRight
		}
	}
	return nil
}

func (x *sn) Root() *sn {
	if x.parent == nil {
		return x
	}
	return x.parent.Root()
}

func (x *sn) Start() *sn {
	return x.ordered[0]
}

func (x *sn) Reduce() {
	x.LeftMost(4).Explode()
	if x.LeftMost(4) != nil {
		x.Reduce()
	}
	before := fmt.Sprint(x)
	x.Split()
	after := fmt.Sprint(x)
	if before != after {
		x.Reduce()
	}

}

func (x *sn) Explode() {
	if x == nil {
		return
	}
	x.Root().Decorate(true)

	if x.left.before != nil {
		x.left.before.val += x.left.val
	}
	if x.right.after != nil {
		x.right.after.val += x.right.val
	}

	x.val = 0

	x.left.parent = nil
	x.left = nil

	x.right.parent = nil
	x.right = nil

	x.isLeaf = true
}

func (x *sn) Split() bool {
	x.Root().Decorate(true)
	node := x.Start()

	for node != nil {

		if node.val > 9 {
			newLNode := &sn{isLeaf: true}
			newLNode.val = (node.val) / 2
			newLNode.parent = node

			node.left = newLNode

			newRNode := &sn{isLeaf: true}
			newRNode.val = (node.val + 1) / 2
			newRNode.parent = node
			node.right = newRNode

			node.val = 0
			node.isLeaf = false
			return true
		}
		node = node.after

	}

	return false
}
func (x *sn) Magnitude() int64 {
	if x.isLeaf {
		return int64(x.val)
	}
	return (3 * x.left.Magnitude()) + (2 * x.right.Magnitude())
}

func main() {

	//fname := "tiny.txt"
	fname := "input.txt"
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)

	snailfish := &sn{}
	first := true
	log.Println("starting")
	for scanner.Scan() {
		line := scanner.Text()
		if first {
			NewSn(line, snailfish)
			first = false
			continue
		}
		next := &sn{}
		NewSn(line, next)

		sum := Add(snailfish, next)
		sum.Reduce()

		snailfish = sum
	}
	fmt.Println(snailfish.Root().Magnitude())

}

func NewSn(str string, node *sn) {
	if len(str) < 2 {
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
