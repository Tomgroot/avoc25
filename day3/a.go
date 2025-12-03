package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run a.go <filename>")
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var totalSum int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		max := -1
		var row []int
		for i, char := range line {
			digit, _ := strconv.Atoi(string(char))
			row = append(row, digit)

			if digit > max && i != len(line)-1 {
				max = digit
				row = nil
			}
		}


		max2 := -1
		for _, num := range row {
			if num > max2 {
				max2 = num
			}
		}

		combinedNum := (max * 10) + max2
		totalSum += combinedNum
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	
	fmt.Printf("Total: %d\n", totalSum)
	fmt.Printf("Time: %s\n", elapsed)
}