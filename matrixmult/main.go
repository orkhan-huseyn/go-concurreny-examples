package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	matrixSize = 250
)

var (
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result  = [matrixSize][matrixSize]int{}
	rwLock  = sync.RWMutex{}
	cond    = sync.NewCond(rwLock.RLocker())
	wg      = sync.WaitGroup{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := range matrixSize {
		for col := range matrixSize {
			matrix[row][col] += rand.IntN(10) - 5
		}
	}
}

func workoutRow(row int) {
	rwLock.RLock()
	for {
		wg.Done()
		cond.Wait()
		for col := range matrixSize {
			for i := range matrixSize {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
	}
}

func main() {
	wg.Add(matrixSize)
	for row := range matrixSize {
		go workoutRow(row)
	}

	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Wait()
		rwLock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)

		wg.Add(matrixSize)
		rwLock.Unlock()
		cond.Broadcast()
	}

	elapsed := time.Since(start)
	fmt.Printf("Processing took %s\n", elapsed)
}
