package main

import (
	"fmt"
	"sync"
)

func multiply(A, B [][]int, result chan<- [3]int, wg *sync.WaitGroup, row, col int) {
	defer wg.Done()
	sum := 0
	for i := 0; i < len(A[0]); i++ {
		sum += A[row][i] * B[i][col]
	}
	result <- [3]int{row, col, sum}
}

func main() {
	n, m, k := 1000, 1000, 1000

	A := make([][]int, n)
	B := make([][]int, m)
	for i := range A {
		A[i] = make([]int, m)
		for j := range A[i] {
			A[i][j] = i + j
		}
	}
	for i := range B {
		B[i] = make([]int, k)
		for j := range B[i] {
			B[i][j] = i + j
		}
	}

	result := make(chan [3]int)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			wg.Add(1)
			go multiply(A, B, result, &wg, i, j)
		}
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Printf("[%d,%d]=%d\n", res[0], res[1], res[2])
	}
}
