package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

//func main() {
//	log.Println("Day 8, part1")
//
//	uniques := map[string]int{}
//	f, _ := os.Open("simple.txt")
//	//f, _ := os.Open("sample.txt")
//	//f, _ := os.Open("input.txt")
//	scanner := bufio.NewScanner(f)
//	for scanner.Scan() {
//		line := scanner.Text()
//
//		p1 := strings.Split(line, "|")
//		outputs := p1[1]
//		words := strings.Fields(outputs)
//
//		for _, word := range words {
//			word = SortString(word)
//			uniques[word]++
//		}
//	}
//	//fmt.Println(uniques)
//	sum := 0
//	for k, x := range uniques {
//		if len(k) == 2 || // numeral 1
//			len(k) == 4 || // numeral 4
//			len(k) == 3 || // numeral 7
//			len(k) == 7 { // numeral 8
//			sum += x
//		}
//	}
//	println("part1", sum)
//}

func main() {
	log.Println("Day 8, part2")
	log.Println("Calculating brute force style...")
	uniques := map[string]int{}
	//f, _ := os.Open("simple.txt")
	//f, _ := os.Open("sample.txt")
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	total := 0
	fmt.Println("0")
	for scanner.Scan() {
		line := scanner.Text()
		p1 := strings.Split(line, "|")
		outputs := p1[1]
		inputs := p1[0]
		words := strings.Fields(outputs)
		for i, w := range words {
			words[i] = SortString(w)
		}

		wirings := strings.Fields(inputs)

		normalised := []string{}
		for _, wiring := range wirings {
			wiring = SortString(wiring)
			normalised = append(normalised, wiring)
			uniques[wiring]++
		}
		sort.Strings(normalised)
		hash := fmt.Sprint(normalised)
		Perm([]rune("abcdefg"), func(a []rune) {
			wiring, wHash := wire(string(a))
			if hash == wHash {
				value := wiring[words[0]]*1000 +
					wiring[words[1]]*100 +
					wiring[words[2]]*10 +
					wiring[words[3]]
				println("+", value)
				total = total + value
			}

		})

	}
	fmt.Println("================")
	fmt.Println(total)
}

func wire(perm string) (map[string]int, string) {
	// abcdefg
	// 0123456
	a := string(perm[0])
	b := string(perm[1])
	c := string(perm[2])
	d := string(perm[3])
	e := string(perm[4])
	f := string(perm[5])
	g := string(perm[6])
	wiring := map[int]string{}
	display := map[string]int{}
	wiring[0] = SortString(a + b + c + e + f + g)
	wiring[1] = SortString(c + f)
	wiring[2] = SortString(a + c + d + e + g)
	wiring[3] = SortString(a + c + d + f + g)
	wiring[4] = SortString(b + c + d + f)
	wiring[5] = SortString(a + b + d + f + g)
	wiring[6] = SortString(a + b + d + e + f + g)
	wiring[7] = SortString(a + c + f)
	wiring[8] = SortString(a + b + c + d + e + f + g)
	wiring[9] = SortString(a + b + c + d + f + g)

	for k, v := range wiring {
		display[v] = k
	}

	normalised := []string{}
	for _, v := range wiring {
		normalised = append(normalised, v)
	}
	sort.Strings(normalised)
	hash := fmt.Sprint(normalised)

	return display, hash
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func contains(word string, words []string) bool {
	for _, v := range words {
		if word == v {
			fmt.Println("found", word)
			return true
		}
	}
	return false
}

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
