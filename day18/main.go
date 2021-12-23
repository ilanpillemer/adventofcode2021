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
	//log.Println("adding", left, right)
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
		//log.Println("ordered", x.ordered)
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

func (x *sn) Reduce() {
	before := fmt.Sprint(x)
	log.Println(x.LeftMost(4))
	x.LeftMost(4).Explode()
	after := fmt.Sprint(x)
	if x.LeftMost(4) != nil {
		log.Println("       : ", before)
		log.Println("explode: ", after)
		x.Reduce()
	}
	before = fmt.Sprint(x)
	x.Split()
	after = fmt.Sprint(x)
	if before != after {
		log.Println("       : ", before)
		log.Println("split  : ", after)
		x.Reduce()
	}
	//fmt.Println(after)

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
	//log.Println("split", x)

	if x.left != nil && x.left.isLeaf && x.left.val > 9 {
		newLNode := &sn{isLeaf: true}
		newLNode.val = (x.left.val) / 2
		newLNode.parent = x.left
		x.left.left = newLNode

		newRNode := &sn{isLeaf: true}
		newRNode.val = (x.left.val + 1) / 2
		newRNode.parent = x.left
		x.left.right = newRNode

		x.left.val = 0
		x.left.isLeaf = false
		return true
	}

	if x.right != nil && x.right.isLeaf && x.right.val > 9 {
		newLNode := &sn{isLeaf: true}
		newLNode.val = (x.right.val) / 2
		newLNode.parent = x.right

		x.right.left = newLNode

		newRNode := &sn{isLeaf: true}
		newRNode.val = (x.right.val + 1) / 2
		newRNode.parent = x.right
		x.right.right = newRNode

		x.right.val = 0
		x.right.isLeaf = false
		return true
	}

	if x.left != nil {
		checkLeft := x.left.Split()
		if checkLeft {
			return checkLeft
		}
	}

	if x.right != nil {
		checkRight := x.right.Split()
		if checkRight {
			return checkRight
		}
	}
	return false
}

//var ltest = "[[[[5,11],[13,0]],[[8,[7,7]],[[7,9],[5,0]]]],[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]]"

func main() {

	//runTests()
	//runSeqTest()
	//runReduceTest()
	//runReduceL1()
	//runReduceL2()
	runReduceLN()
	os.Exit(0)
	fname := "tiny1.txt"
	fname = "tiny2.txt"
	fname = "tiny3.txt"
	fname = "tiny.txt"
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
		log.Println(" ", snailfish)
		log.Println("+", next)
		sum := Add(snailfish, next)

		sum.Reduce()
		log.Println("=", sum)
		fmt.Println("########################")
		fmt.Println("")
		//fmt.Println(sum)
		snailfish = sum
	}

}

func runReduceL1() {
	l := "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]"
	r := "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"
	left := &sn{}
	right := &sn{}
	NewSn(l, left)
	NewSn(r, right)
	fmt.Println("  ", l)
	fmt.Println("+ ", r)

	combined := Add(left, right)
	fmt.Println("=", combined)
	fmt.Println("Left Most", combined.LeftMost(4))
	combined.Reduce()
	fmt.Println("=", combined)
	fmt.Println(combined)
}

func runReduceL2() {
	l := "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"
	r := "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]"
	left := &sn{}
	right := &sn{}
	NewSn(l, left)
	NewSn(r, right)
	fmt.Println("  ", l)
	fmt.Println("+ ", r)

	combined := Add(left, right)
	fmt.Println("=", combined)
	fmt.Println("Left Most", combined.LeftMost(4))
	combined.Reduce()
	fmt.Println("=", combined)
	fmt.Println(combined)

}

func runReduceLN() {

	l := "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]"
	r := "[7,[5,[[3,8],[1,4]]]]"
	left := &sn{}
	right := &sn{}
	NewSn(l, left)
	NewSn(r, right)
	fmt.Println("  ", l)
	fmt.Println("+ ", r)

	combined := Add(left, right)
	fmt.Println("=", combined)
	fmt.Println("Left Most", combined.LeftMost(4))
	combined.Reduce()
	fmt.Println("=", combined)
	fmt.Println(combined)
	want := "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]"
	if want != combined.String() {
		fmt.Printf("want [%s]\ngot  [%s]\n", want, combined)
	}

}

func runReduceTest() {
	l := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	r := "[1,1]"
	left := &sn{}
	right := &sn{}
	NewSn(l, left)
	NewSn(r, right)
	combined := Add(left, right)
	combined.Reduce()
	fmt.Println(combined)
}

func runSeqTest() {
	l := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	r := "[1,1]"
	left := &sn{}
	right := &sn{}
	NewSn(l, left)
	NewSn(r, right)

	combined := Add(left, right)
	fmt.Println("after addition", combined)
	combined.LeftMost(4).Explode()
	fmt.Println("after explode", combined)
	combined.LeftMost(4).Explode()
	fmt.Println("after explode", combined)
	combined.Split()
	fmt.Println("after split", combined)
	combined.Split()
	fmt.Println("after split", combined)
	combined.LeftMost(4).Explode()
	fmt.Println("after explode", combined)
}

func runTests() {

	//runTest("test1", "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]", "")

	runTest("test2", "[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]")
	runTest("test3", "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]") // hmmm
	runTest("test4", "[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]")
	runTest("test5", "[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]")
	runTest("test6", "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	runTest("test7", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]")
	runTest("one", "[1,2]", "[1,2]")
	runTest("two", "[[1,2],3]", "[[1,2],3]")
	//	runSplitTest("one", "[[[[0,7],4],[15,[0,13]]],[1,1]]")
}

func runTest(name string, test string, want string) {
	fmt.Println("**********")
	fmt.Println(name, test)
	fmt.Println("**********")

	root := &sn{}
	NewSn(test, root)
	result := fmt.Sprint(root)

	if result != test {
		panic("tree failed to build")
	}
	log.Printf("%s -> %s\n", test, root)
	fmt.Println("**********")
	fmt.Println("left most", root.LeftMost(4))
	fmt.Println("**********")
	root.LeftMost(4).Explode()
	fmt.Println("after left mode explode", root)
	if fmt.Sprint(root) != want {
		panic(fmt.Sprintf("want: %s, got :%s", want, root))
	}
	fmt.Println(test, "-->", root)
	fmt.Println("**********")

}

func runSplitTest(name string, test string) {
	fmt.Println("**********")
	fmt.Println(name)
	fmt.Println("**********")

	root := &sn{}
	NewSn(test, root)
	result := fmt.Sprint(root)

	if result != test {
		log.Println("oops", result)
		panic("tree failed to build")
	}
	log.Printf("%s -> %s\n", test, root)
	fmt.Println("**********")
	fmt.Println("**********")
	root.Split()
	fmt.Println("after Split explode", root)

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
