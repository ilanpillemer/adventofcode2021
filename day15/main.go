package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
)

//https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
var width = 0
var height = 0
var distances = map[image.Point]int{}
var weights = map[image.Point]int{}
var unvisited = map[image.Point]int{}
var count = 0
var pq PriorityQueue

func main() {
	//	load("sample.txt")
	load("input.txt")
	assign()
	log.Println(weights[image.Pt(width, height)])
}

func wrap(val int, mod int) int {
	if val < mod {
		return val
	}
	return (val % mod) + 1
}

func load(fname string) {
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	x := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		x = 0
		width = len(line)
		for _, j := range line {
			for k := 0; k < width-1; k++ {
				distances[image.Pt(x+(k*width), y)] = wrap((atoi(string(j)) + k), 10)
				distances[image.Pt(x+(k*width), y+width)] = wrap((atoi(string(j)) + (k + 1)), 10)
				distances[image.Pt(x+(k*width), y+(2*width))] = wrap((atoi(string(j)) + (k + 2)), 10)
				distances[image.Pt(x+(k*width), y+(3*width))] = wrap((atoi(string(j)) + (k + 3)), 10)
				distances[image.Pt(x+(k*width), y+(4*width))] = wrap((atoi(string(j)) + (k + 4)), 10)
			}
			x++
		}
		y++
	}
	for k := range distances {
		weights[k] = 2147483647
		unvisited[k] = 1
	}
	log.Println("width", width)
	width = (width * 5) - 1
	height = width
	log.Println(width, height)
	pq = PriorityQueue{}
	log.Println(pq)
	item := &Item{
		value:    image.Pt(0, 0),
		priority: 0,
	}

	heap.Init(&pq)
	heap.Push(&pq, item)
}

func assign() {
	start := image.Pt(0, 0)
	item := &Item{
		value:    start,
		priority: 0,
	}
	heap.Push(&pq, item)
	weights[start] = 0
	process(start)

	for {
		if finished() {
			break
		} else {
			n := getNextNode()
			process(n)
		}
	}
}

func getNextNode() image.Point {
	item := heap.Pop(&pq).(*Item)
	return item.value

}

func finished() bool {
	if _, ok := unvisited[image.Pt(width, height)]; !ok {
		return true
	}
	return false
}

func process(curr image.Point) {
	ns := neighbours(curr)
	myWeight := weights[curr]
	for _, n := range ns {
		if _, ok := unvisited[n]; ok {
			weight := weights[n]
			proposed := myWeight + distances[n]
			if proposed < weight {
				weights[n] = proposed
				item := &Item{
					value:    n,
					priority: proposed,
				}
				heap.Push(&pq, item)
			}
		}
	}
	delete(unvisited, curr)
	count++
}

func neighbours(pt image.Point) []image.Point {
	nx := []image.Point{}
	up := pt.Add(image.Pt(0, -1))
	down := pt.Add(image.Pt(0, 1))
	left := pt.Add(image.Pt(-1, 0))
	right := pt.Add(image.Pt(1, 0))
	if up.Y >= 0 {
		nx = append(nx, up)
	}
	if down.Y <= height {
		nx = append(nx, down)
	}
	if left.X >= 0 {
		nx = append(nx, left)
	}
	if right.X <= width {
		nx = append(nx, right)
	}
	return nx

}

func apply(grid map[image.Point]int, f func(image.Point), g func()) map[image.Point]int {
	x := width
	y := height
	for i := 0; i < y+1; i++ {
		for j := 0; j < x+1; j++ {
			f(image.Point{j, i})
		}
		g()
	}
	return grid
}

func display(grid map[image.Point]int) {
	apply(grid,
		func(x image.Point) {
			pt := grid[x]
			fmt.Print(pt)
		},
		func() { fmt.Println() },
	)
	fmt.Println()
}

func atoi(str string) int {
	x, _ := strconv.Atoi(str)
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func deepcopy(in map[image.Point]int) map[image.Point]int {
	out := map[image.Point]int{}
	for k, v := range in {
		out[k] = v
	}
	return out
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    image.Point // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value image.Point, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
