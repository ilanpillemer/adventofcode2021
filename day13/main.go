package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
)

var width = 0
var height = 0
var grid = map[image.Point]int{}

//sample folds
var yfold = []int{7}
var xfold = []int{5}

func main() {
	load("input.txt")
	//load("sample.txt")
	//fold along x=655
	//fold along y=447
	//fold along x=327
	//fold along y=223
	//fold along x=163
	//fold along y=111
	//fold along x=81
	//fold along y=55
	//fold along x=40
	//fold along y=27
	//fold along y=13
	//fold along y=6

	foldx(655)
	foldy(447)
	foldx(327)
	foldy(223)
	foldx(163)
	foldy(111)
	foldx(81)
	foldy(55)
	foldx(40)
	foldy(27)
	foldy(13)
	foldy(6)
	width = 40
	height = 6
	log.Println(len(grid))
	display(grid)
}

func load(fname string) {
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			xs := strings.Split(line, ",")
			x, y := atoi(xs[0]), atoi(xs[1])
			width = max(width, x)
			height = max(height, y)
			grid[image.Pt(x, y)] = 1
		}
	}
}

func foldy(fy int) {
	apply(grid,
		func(pt image.Point) {
			if _, ok := grid[pt]; ok {
				if pt.Y > fy {
					x, y := pt.X, pt.Y
					y = y - fy
					grid[image.Pt(x, fy-y)] = 1
					delete(grid, pt)
				}
			}
		},
		func() {},
	)
}

func foldx(fx int) {
	apply(grid,
		func(pt image.Point) {
			if _, ok := grid[pt]; ok {
				if pt.X > fx {
					x, y := pt.X, pt.Y
					x = x - fx
					grid[image.Pt(fx-x, y)] = 1
					delete(grid, pt)
				}
			}
		},
		func() {},
	)
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
			ch := "."
			if pt == 1 {
				ch = "#"
			}
			fmt.Print(ch)
		},
		func() { fmt.Println() },
	)
	fmt.Println()
}
