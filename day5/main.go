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

var grid = map[image.Point]int{}

type Vent struct {
	left  image.Point
	right image.Point
}

var vents = []Vent{}

func main() {
	log.Println("Day 5")
	//f, _ := os.Open("sample.txt")
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		xs := strings.Fields(line)
		left := strings.Split(xs[0], ",")
		right := strings.Split(xs[2], ",")
		x1, y1 := atoi(left[0]), atoi(left[1])
		x2, y2 := atoi(right[0]), atoi(right[1])
		vent := Vent{
			left:  image.Point{x1, y1},
			right: image.Point{x2, y2},
		}
		vents = append(vents, vent)
	}
	for _, vent := range vents {

		if vent.left.X == vent.right.X {
			i, j := minmax(vent.left.Y, vent.right.Y)

			for ; i < j+1; i++ {
				//fmt.Println(vent, "  ", vent.left.X, i)
				grid[image.Point{vent.left.X, i}]++
			}
			continue
		}
		if vent.left.Y == vent.right.Y {
			i, j := minmax(vent.left.X, vent.right.X)

			for ; i < j+1; i++ {
				grid[image.Point{i, vent.left.Y}]++
			}
			continue
		}
		if (vent.left.X < vent.right.X) && (vent.left.Y < vent.right.Y) {
			for i, j := vent.left.X, vent.left.Y; i < vent.right.X+1; i++ {
				grid[image.Point{i, j}]++
				j++
				continue
			}
		}
		if (vent.left.X < vent.right.X) && (vent.left.Y > vent.right.Y) {
			for i, j := vent.left.X, vent.left.Y; i < vent.right.X+1; i++ {
				grid[image.Point{i, j}]++
				j--
				continue
			}
		}
		if (vent.left.X > vent.right.X) && (vent.left.Y < vent.right.Y) {
			for i, j := vent.left.X, vent.left.Y; i > vent.right.X-1; i-- {
				grid[image.Point{i, j}]++
				j++
				continue
			}
		}
		if (vent.left.X > vent.right.X) && (vent.left.Y > vent.right.Y) {
			for i, j := vent.left.X, vent.left.Y; i > vent.right.X-1; i-- {
				grid[image.Point{i, j}]++
				j--
				continue
			}
		}

	}

	//display(grid)
	total := 0
	for _, p := range grid {
		if p > 1 {
			total++
		}
	}
	log.Println("overlapping vents", total)
}

func atoi(str string) (x int) {
	x, _ = strconv.Atoi(str)
	return
}

func minmax(x int, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}

func display(grid map[image.Point]int) {
	fmt.Println("**********")
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Print(grid[image.Point{x, y}])
		}
		fmt.Println()
	}
	fmt.Println("**********")
}
