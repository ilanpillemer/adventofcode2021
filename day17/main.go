package main

import (
	"fmt"
	"image"
	"log"
)

//type velocity image.Point
//type point image.Point
var lowesty int

func main() {
	log.Println("day 17")
	// sample
	// target area: x=20..30, y=-10..-5

	//lowesty = -10
	//	target := image.Rect(20, -10, 30, -5)
	//target := image.Rect(20, -10, 31, -4)
	// input
	//target area: x=236..262, y=-78..-58
	lowesty = -78
	//	target := image.Rect(236, -78, 262, -58)
	target := image.Rect(236, -78, 263, -57)
	//	maxy := abs(lowesty) * 100
	maxx := abs(lowesty) * 100
	highest := 0
	count := 0
	for x := 0; x < maxx; x++ {
		for y := -300; y < 300; y++ {
			high, ok := fire(image.Pt(0, 0), image.Pt(x, y), target)
			if ok {
				//	log.Println(image.Pt(x, y))
				count++
				if high > highest {
					highest = high
				}
			}
		}

	}
	fmt.Println("highest:", highest)
	fmt.Println("count:", count)
	log.Println("done")

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func step(point image.Point, velocity image.Point) (image.Point, image.Point) {
	point = point.Add(velocity)
	switch {
	case velocity.X > 0:
		velocity.X = velocity.X - 1
	case velocity.X < 0:
		velocity.X = velocity.X + 1
	}
	velocity.Y = velocity.Y - 1
	//fmt.Println(point, velocity)
	return point, velocity
}

// target area: x=20..30, y=-10..-5
//target area: x=236..262, y=-78..-58

func fire(point image.Point, velocity image.Point, target image.Rectangle) (int, bool) {
	highest := -10
	for {
		point, velocity = step(point, velocity)
		if point.Y > highest {
			highest = point.Y
		}
		if point.In(target) {
			//	log.Println("hit target", velocity)
			return highest, true
		}
		if point.Y < lowesty || point.X == 0 {
			break
		}
	}
	return highest, false
}
