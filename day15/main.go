package main

import (
	"bufio"
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

// 870 863
func main() {
	//load("sample.txt")
	load("input.txt")
	//display(distances)
	//display(unvisited)
	assign()
	//display(weights)
	log.Println(weights[image.Pt(width, height)])
}

func load(fname string) {
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	x := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		x = 0
		for _, j := range line {
			distances[image.Point{x, y}] = atoi(string(j))
			weights[image.Point{x, y}] = 2147483647
			unvisited[image.Point{x, y}] = 1
			width = max(width, x)
			x++
		}
		height = max(height, y)
		y++
	}
}

func assign() {
	start := image.Pt(0, 0)

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

	weight := 2147483647
	var pt image.Point
	//display(unvisited)
	for k, v := range weights {
		if _, ok := unvisited[k]; ok {
			if v < weight {
				weight = v
				pt = k
			}
		}
	}

	return pt
}

func finished() bool {
	return len(unvisited) == 0
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
			}
		}
	}
	delete(unvisited, curr)
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
			//ch := "."
			//if pt == 1 {
			//	ch = "#"
			//}
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
