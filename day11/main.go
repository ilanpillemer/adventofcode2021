package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
)

var width = 0
var height = 0
var grid = map[image.Point]int{}

//var fname = "simple.txt"
//var fname = "sample.txt"

var fname = "input.txt"

func main() {
	fmt.Println("Day 11")
	load()
	total := 0
	for i := 0; i < 100; i++ {
		var n int
		grid, n = step(grid)
		total = total + n
	}

	log.Println("**** part 1 *****")
	display(grid)
	fmt.Println(total)
	fmt.Println("simultaneous?", isSimultaneous(grid))
	log.Println("**** part 1 *****")
	log.Println("**** part 2 *****")
	load()
	total = 0
	i := 0
	for {
		var n int
		grid, n = step(grid)
		total = total + n
		i++
		if isSimultaneous(grid) {
			break
		}
	}
	display(grid)

	fmt.Println(total)
	fmt.Println("simultaneous?", isSimultaneous(grid), i)
	log.Println("**** part 2 *****")
}

func isSimultaneous(in map[image.Point]int) bool {
	for _, v := range in {
		if v != 0 {
			return false
		}
	}

	return true

}

func step(in map[image.Point]int) (map[image.Point]int, int) {
	y := incNeighbours(in)
	next, i := flash(y, map[image.Point]bool{})
	return next, i
}

func neighbours(p image.Point) []image.Point {
	return []image.Point{
		p.Add(image.Pt(1, 0)),
		p.Add(image.Pt(0, 1)),
		p.Add(image.Pt(-1, 0)),
		p.Add(image.Pt(0, -1)),
		p.Add(image.Pt(1, 1)),
		p.Add(image.Pt(-1, 1)),
		p.Add(image.Pt(1, -1)),
		p.Add(image.Pt(-1, -1)),
	}
}

func load() {
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	x := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		x = 0
		for _, j := range line {
			grid[image.Point{x, y}] = atoi(string(j))
			width = max(width, x)
			x++
		}
		height = max(height, y)
		y++

	}

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
		func(x image.Point) { fmt.Print(grid[x]) },
		func() { fmt.Println() },
	)
	fmt.Println()
}

func incNeighbours(grid map[image.Point]int) map[image.Point]int {
	x := width
	y := height
	variant := map[image.Point]int{}
	// inc all by one
	for i := 0; i < y+1; i++ {
		for j := 0; j < x+1; j++ {
			ns := neighbours(image.Point{j, i})
			for _, nx := range ns {
				if value, ok := grid[nx]; ok {
					variant[nx] = (value + 1)
				}
			}
		}
	}
	return variant
}

func flash(grid map[image.Point]int, flashed map[image.Point]bool) (map[image.Point]int, int) {
	x, y := width, height
	// next iteration of grid
	variant := map[image.Point]int{}
	for k, v := range grid { // initial state of variant
		variant[k] = v
	}
	isDark := true
	for i := 0; i < y+1; i++ {
		for j := 0; j < x+1; j++ {
			// for each point in grid
			// should it maybe flash?
			// it should flash if its value is zero and it has not flashed before
			flash1 := grid[image.Pt(j, i)] > 9          // it could flash
			if _, ok := flashed[image.Pt(j, i)]; !ok && // has not flashed before
				flash1 {
				isDark = false
				// a flash has occured
				ns := neighbours(image.Point{j, i})
				// go through all neighbours of the flashed point
				for _, nx := range ns {
					if value, ok := variant[nx]; ok { // if neighbour exists in grid
						variant[nx] = value + 1
					}
				}
				flashed[image.Pt(j, i)] = true
			}
		}
	}
	if isDark {
		// all fish that flashed are now very tired
		for pt := range flashed {
			variant[pt] = 0
		}
		return variant, len(flashed)
	}
	return flash(variant, flashed)
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func conv(in map[image.Point]int, flashed map[image.Point]bool) map[image.Point]int {
	out := map[image.Point]int{}
	for k := range in {
		value := 0
		if flashed[k] {
			value = 1
		}
		out[k] = value
	}
	return out
}
