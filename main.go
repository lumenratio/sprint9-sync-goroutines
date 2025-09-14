package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]int, 0, size)
	for i := 0; i < size; i++ {
		numbers = append(numbers, rnd.Int())
	}
	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	max := data[0]
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	// check data parameters
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	// slice for results from goroutines
	chunksResults := make([]int, CHUNKS)
	var wg sync.WaitGroup
	// ok, lets chunk data and run goroutines
	for i := 0; i < CHUNKS; i++ {
		start := i * len(data) / CHUNKS
		end := ((i + 1) * len(data)) / CHUNKS
		chunk := data[start:end]
		wg.Add(1)
		go func(chunk []int, i int) {
			defer wg.Done()
			max := maximum(chunk)
			chunksResults[i] = max
		}(chunk, i)
	}
	wg.Wait()
	//wait when goroutine that collecting results, complete
	return maximum(chunksResults)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	rndNumbers := generateRandomElements(SIZE)
	if rndNumbers == nil {
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(rndNumbers)
	if max == 0 {
		return
	}
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(rndNumbers)
	if max == 0 {
		return
	}
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
