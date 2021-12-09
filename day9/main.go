package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"sort"
	"strconv"
)

var width = 0
var height = 0 //2147483647

func main() {

	grid := map[image.Point]int{}

	fmt.Println("day9")
	//f, _ := os.Open("sample.txt")
	f, _ := os.Open("input.txt")
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

	px := lowPoints(grid, width, height)
	sum := 0
	for _, k := range px {
		sum += k
		sum += 1
	}

	fmt.Println("***** part 1 *****")
	fmt.Println(sum)
	cx := lowPointsCoords(grid, width, height)
	all := []int{}
	for _, lp := range cx {
		basin := map[image.Point]bool{}
		basinSize(grid, lp, basin)
		all = append(all, len(basin))

	}
	sort.Ints(all)

	scores := all[len(all)-3:]

	final := int64(1)
	for _, score := range scores {
		final = final * int64(score)
	}
	fmt.Println("***** part 2 *****")
	fmt.Println(final)
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

func display(grid map[image.Point]int, x int, y int) {
	for i := 0; i < y+1; i++ {
		for j := 0; j < x+1; j++ {
			fmt.Print(grid[image.Point{j, i}])
		}
		fmt.Println()
	}
}

func basinSize(grid map[image.Point]int, px image.Point, acc map[image.Point]bool) {
	if _, ok := acc[px]; ok {
		return
	}
	if px.X < 0 || px.X > width {
		return
	}
	if px.Y < 0 || px.Y > height {
		return
	}
	if grid[px] == 9 {
		return
	}
	acc[px] = true
	basinSize(grid, image.Point{px.X - 1, px.Y}, acc)
	basinSize(grid, image.Point{px.X + 1, px.Y}, acc)
	basinSize(grid, image.Point{px.X, px.Y + 1}, acc)
	basinSize(grid, image.Point{px.X, px.Y - 1}, acc)

}

func lowPointsCoords(grid map[image.Point]int, x int, y int) []image.Point {
	px := []image.Point{}
	for i := 0; i < y+1; i++ {
		for j := 0; j < x+1; j++ {
			if ok, _ := isLow(grid, j, i); ok {
				px = append(px, image.Point{j, i})
			}
		}
	}
	return px
}

func lowPoints(grid map[image.Point]int, x int, y int) []int {
	px := []int{}
	for i := 0; i < y+1; i++ {
		for j := 0; j < x+1; j++ {
			if ok, value := isLow(grid, j, i); ok {
				px = append(px, value)
			}
		}
	}
	return px
}

func atoi(str string) int {

	i, _ := strconv.Atoi(str)
	return i
}

func isLow(grid map[image.Point]int, x, y int) (ok bool, value int) {

	p := image.Point{x, y}

	if x == 0 {
		right := image.Point{x + 1, y}
		if grid[right] <= grid[p] {
			return false, -1
		}
	}
	if x == width {
		left := image.Point{x - 1, y}
		if grid[left] <= grid[p] {
			return false, -1
		}
	}
	if x > 0 && x < width {

		left := image.Point{x - 1, y}
		right := image.Point{x + 1, y}

		if grid[left] <= grid[p] || grid[right] <= grid[p] {
			return false, -1
		}
	}

	if y == 0 {
		down := image.Point{x, y + 1}
		if grid[down] <= grid[p] {
			return false, -1
		}

	}
	if y == height {
		up := image.Point{x, y - 1}
		if grid[up] <= grid[p] {
			return false, -1
		}
	}
	if y > 0 && y < height {
		up := image.Point{x, y - 1}
		down := image.Point{x, y + 1}
		if grid[up] <= grid[p] || grid[down] <= grid[p] {
			return false, -1
		}
	}
	return true, grid[p]
}
