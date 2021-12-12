package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var edges = map[string][]string{}
var total = [][]string{}

func main() {
	//	load("sample.txt")
	load("input.txt")

	seen := map[string]int{}
	seen["start"] = 1

	acc := []string{"start"}
	paths("start", seen, acc)

	log.Println("finished")
	log.Printf("Found %d paths \n", len(total))
}

func paths(x string, seen map[string]int, acc []string) {

	if x == "end" {
		//	log.Println(acc)
		total = append(total, acc)

	}
	for _, edge := range edges[x] {

		_, ok := seen[edge]
		if ok {
			ok = override(edge, seen)
		}

		if !ok {
			nseen := deepcopy(seen)
			nseen = seeEdge(nseen, edge)

			acc = append(acc, edge)
			paths(edge, nseen, acc)
		}
	}
}

func override(edge string, seen map[string]int) bool {
	if edge == "start" || edge == "end" {
		return true
	}

	for _, value := range seen {
		if value == 2 {
			return true
		}
	}
	return false
}

func seeEdge(seen map[string]int, edge string) map[string]int {
	if edge == strings.ToLower(edge) {
		seen[edge]++
	}
	return seen
}

func load(fname string) {
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		xs := strings.Split(line, "-")
		x := edges[xs[0]]
		if x == nil {
			x = []string{}
		}
		x = append(x, xs[1])
		edges[xs[0]] = x

		x = edges[xs[1]]
		if x == nil {
			x = []string{}
		}
		x = append(x, xs[0])
		edges[xs[1]] = x
	}
	fmt.Println(edges)
}

func deepcopy(x map[string]int) map[string]int {
	y := map[string]int{}
	for k, v := range x {
		y[k] = v
	}
	return y
}
