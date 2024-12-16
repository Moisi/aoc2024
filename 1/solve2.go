package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var filePath = "inputs.csv"

func main() {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanWords)

	s1 := []int{}
	m2 := map[int]int{}

	for fileScanner.Scan() {
		num1, _ := strconv.Atoi(fileScanner.Text())
		s1 = append(s1, num1)
		fileScanner.Scan()
		num2, _ := strconv.Atoi(fileScanner.Text())
		m2[num2] = m2[num2] + 1
	}
	readFile.Close()
	sum := 0

	for _, val := range s1 {
		sum += val * m2[val]
	}

	fmt.Println("ans: ", sum)
}
