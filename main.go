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
func generateRandomElements(size int) ([]int, error) {
	// ваш код здесь
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if size <= 0 {
		return nil, fmt.Errorf("given size, zero or less")
	}

	numbers := make([]int, 0, size)
	for i := 0; i < size; i++ {
		numbers = append(numbers, int(rnd.Int63()))
	}
	return numbers, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	// ваш код здесь
	if len(data) == 0 {
		return 0, fmt.Errorf("len of given slice is zero")
	}
	if len(data) == 1 {
		return data[0], nil
	}

	// да, решил использовать сортировку sort.Ints(), а не простой перебор элементов с условием max < v
	// а потом суммарно 8 часов идиотии, что бы понять, почему нижележайший массив меняется и в maxChunks передается отсортированный массив (по началу показалось, что вообще другой)
	//sort.Ints(data)
	//max := data[len(data)-1]
	// поэтому for range data ...
	// и минус 8 часов и массу нервов
	max := data[0]
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	// ваш код здесь
	// check data parameters
	if len(data) == 0 {
		return 0, fmt.Errorf("len of given slice is zero")
	}

	if len(data) == 1 {
		return data[0], nil
	}

	if CHUNKS <= 0 {
		return 0, fmt.Errorf("given number of chunks is zero or less")
	}

	// if len(data) < 8, then chunks count must be lower than default CHUNKS
	numChunks := CHUNKS
	if len(data) < 8 {
		numChunks = len(data)
	}

	// done channel
	done := make(chan struct{})
	defer close(done)
	// channel for goroutines, that send result
	getMaxNum := make(chan int, numChunks)
	// slice for results from goroutines
	var chunksResults []int
	go func() {
		for v := range getMaxNum {
			chunksResults = append(chunksResults, v)
		}
		done <- struct{}{}
	}()

	var wg sync.WaitGroup
	// ok, lets chunk data and run goroutines
	for i := 0; i < numChunks; i++ {
		start := i * len(data) / numChunks
		end := ((i + 1) * len(data)) / numChunks
		wg.Add(1)
		go func() {
			defer wg.Done()
			chunk := data[start:end]
			max, _ := maximum(chunk)
			getMaxNum <- max
		}()
	}
	wg.Wait()
	//all gouroutines did their task, close the channel
	close(getMaxNum)
	//wait when goroutine that collecting results, complete
	<-done
	return maximum(chunksResults)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	rndNumbers, err := generateRandomElements(SIZE)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max, err := maximum(rndNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max, err = maxChunks(rndNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
