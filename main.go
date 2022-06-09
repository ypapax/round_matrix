package main

import (
	"fmt"
	"log"
)

func main() {
	round(3)
}

func round(N int) {
	m := make([][]int, N)
	var counter = 1
	roundM(m, 0, N, counter)

	printMatrix(m)
}

func printMatrix(m [][]int) {
	log.Println("matrix")
	for _, l := range m {
		var line string
		for _, v := range l {
			line += fmt.Sprintf(" %+v", v)
		}
		log.Println(line)
	}
}

func roundM(m [][]int, topRow int, N int, counter int) {
	if topRow >= N {
		return
	}
	counterBefore := counter
	var rightColumn int
	var leftColumn = -1
	for i := 0; i < N; i++ {
		if m[i] == nil {
			m[i] = make([]int, N)
		}
		if m[topRow][i] > 0 {
			continue
		}
		if leftColumn < 0 {
			leftColumn = i
		}
		rightColumn = i
		m[topRow][i] = counter
		counter++
	}
	log.Printf("rightColumn: %+v\n", rightColumn)
	var bottomRow int
	for i := 0; i < N; i++ {
		log.Printf("i: %+v, rightColumn: %+v", i, rightColumn)
		printMatrix(m)

		if m[i][rightColumn] > 0 {
			continue
		}
		bottomRow = i
		m[i][rightColumn] = counter
		counter++
	}

	for i := rightColumn; i >= leftColumn && i>=0; i-- {
		log.Println("bottomRow", bottomRow)
		log.Println("i", i)
		if m[bottomRow][i] > 0 {
			continue
		}
		m[bottomRow][i] = counter
		counter++
	}

	for i := bottomRow; i > topRow; i-- {
		log.Println("i", i)
		log.Println("leftColumn", leftColumn)
		if m[i][leftColumn] > 0 {
			continue
		}
		m[i][leftColumn] = counter
		counter++
	}
	if counter != counterBefore {
		roundM(m, topRow+1, N, counter)
	}
}