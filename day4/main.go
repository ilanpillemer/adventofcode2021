package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type row []int
type col []int

type board struct {
	id          int
	bestTurn    int
	bestRow     int
	bestRowTurn int
	bestCol     int
	bestColTurn int

	rows     []row
	srcRows  []row
	cols     []col
	srcCols  []col
	unmarked []int

	pos map[int]int
	num map[int]int
}

func (b board) String() string {
	builder := strings.Builder{}
	builder.WriteString("board: " + fmt.Sprint(b.id) + "\n")
	builder.WriteString(fmt.Sprintln("best for",
		b.id, "-", b.unmarked, sum(b.unmarked),
		"with piece", b.num[b.bestTurn],
		"with score", b.num[b.bestTurn]*sum(b.unmarked),
		"at turn", b.bestTurn))

	return builder.String()
}

func (b board) CreateCols() board {
	//the best row is the minimum highest value per row
	cols := []col{}
	for i := 0; i < 5; i++ {
		cols = append(cols, make([]int, 5))
	}

	for i, row := range b.rows {
		for j, square := range row {
			cols[j][i] = square
		}
	}
	b.cols = cols
	return b
}

func (b board) UpdateBestRow() board {
	//the best row is the minimum highest value per row
	best := 2147483647
	bestRow := best
	for i, row := range b.rows {
		//highest value
		highest := max(row)
		if highest < best {
			best = highest
			bestRow = i
		}
	}
	b.bestRowTurn = best
	b.bestRow = bestRow
	return b
}

func (b board) UpdateTurn() board {
	b.bestTurn = b.bestColTurn
	if b.bestTurn > b.bestRowTurn {
		b.bestTurn = b.bestRowTurn
	}

	//	b.worstTurn = b.worstColTurn
	//	if b.worstTurn > b.worstRowTurn {
	//		b.worstTurn = b.worstRowTurn
	//	}

	return b
}

func (b board) UpdateBestCol() board {
	//the best row is the minimum highest value per row
	best := 2147483647
	bestCol := best
	for i, col := range b.cols {
		//highest value
		highest := max(col)
		if highest < best {
			best = highest
			bestCol = i
		}
	}
	b.bestColTurn = best
	b.bestCol = bestCol
	return b
}

func (b board) UpdateUnMarked() board {
	xs := []int{}

	for i, row := range b.rows {
		for j, turn := range row {
			if turn > b.bestTurn {
				xs = append(xs, b.srcRows[i][j])
			}
		}
	}
	b.unmarked = xs

	return b
}

func main() {
	log.Println("Day4")
	//items := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	items := []int{87, 7, 82, 21, 47, 88, 12, 71, 24, 35, 10, 90, 4, 97, 30, 55, 36, 74, 19, 50, 23, 46, 13, 44, 69, 27, 2, 0, 37, 33, 99, 49, 77, 15, 89, 98, 31, 51, 22, 96, 73, 94, 95, 18, 52, 78, 32, 83, 85, 54, 75, 84, 59, 25, 76, 45, 20, 48, 9, 28, 39, 70, 63, 56, 5, 68, 61, 26, 58, 92, 67, 53, 43, 62, 17, 81, 80, 66, 91, 93, 41, 64, 14, 8, 57, 38, 34, 16, 42, 11, 86, 72, 40, 65, 79, 6, 3, 29, 60, 1}
	pos := map[int]int{}
	num := map[int]int{}
	boards := map[int]board{}
	for i, item := range items {
		pos[item] = i
		num[i] = item
	}
	//	file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()
	n := 0
	scanner := bufio.NewScanner(file)
	current := board{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards[n] = current
			n++
			boards[n] = board{id: n, pos: pos, num: num}
			continue
		}
		current = boards[n]
		current.rows = append(current.rows, reverse(pos, line))
		current.srcRows = append(current.srcRows, toInts(line))
		boards[n] = current
	}

	best := board{}
	turn := 0
	bestRow := 0
	value := 0
	_, _, _, _ = best, turn, bestRow, value
	for k, v := range boards {
		b := board{}
		b = v.CreateCols()
		b = b.UpdateBestRow()
		b = b.UpdateBestCol()
		b = b.UpdateTurn()
		b = b.UpdateUnMarked()
		boards[k] = b
	}
	log.Println("WINNER!")
	winner := board{}
	winTurn := 2147483647
	for _, b := range boards {
		if b.bestTurn < winTurn {
			winner = b
			winTurn = b.bestTurn
		}
	}
	log.Println(winner)
	log.Println("Loser!")
	winTurn = -1
	for _, b := range boards {
		if b.bestTurn > winTurn {
			winner = b
			winTurn = b.bestTurn
		}
	}
	log.Println(winner)
}

func reverse(pos map[int]int, line string) []int {
	ints := []int{}
	for k, v := range pos {
		ints = toInts(line)
		for i, v := range ints {
			ints[i] = pos[v]
		}
		_, _ = k, v
	}
	return ints
}

func toInts(line string) []int {
	ints := []int{}
	xs := strings.Fields(line)
	for _, x := range xs {
		ints = append(ints, atoi(x))
	}
	return ints
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func min(ints []int) int {
	if len(ints) == 0 {
		return -1
	}
	x := ints[0]
	for _, v := range ints {
		if v < x {
			x = v
		}
	}
	return x
}

func max(ints []int) int {
	if len(ints) == 0 {
		return -1
	}
	x := ints[0]
	for _, v := range ints {
		if v > x {
			x = v
		}
	}
	return x
}

func sum(ints []int) int {
	tot := 0
	for _, v := range ints {
		tot = tot + v
	}
	return tot
}
