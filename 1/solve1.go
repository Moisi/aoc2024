package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	pq "github.com/Jcowwell/go-algorithm-club/PriorityQueue"
)

var filePath = "inputs.csv"

func lessThan(m1, m2 int) bool {
	return m1 < m2
}

func main() {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanWords)

	queue1 := pq.PriorityQueueInit[int](lessThan)
	queue2 := pq.PriorityQueueInit[int](lessThan)

	for fileScanner.Scan() {
		num1, _ := strconv.Atoi(fileScanner.Text())
		queue1.Enqueue(num1)
		fileScanner.Scan()
		num2, _ := strconv.Atoi(fileScanner.Text())
		queue2.Enqueue(num2)
	}
	readFile.Close()

	sum := 0
	if queue1.Count() != queue2.Count() {
		panic("queues not of the same sizes")
	}

	for !queue1.IsEmpty() {
		num1, _ := queue1.Dequeue()
		num2, _ := queue2.Dequeue()
		val := num1 - num2
		if val > 0 {
			sum += val
		} else {
			sum -= val
		}
	}

	fmt.Println("ans: ", sum)
}
