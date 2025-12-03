package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run b.go <filename>")
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	totalSum := big.NewInt(0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		k := len(line) - 12
		var row []byte

		for i := 0; i < len(line); i++ {
			digit := line[i]

			for k > 0 && len(row) > 0 && row[len(row)-1] < digit {
				row = row[:len(row)-1]
				k--
			}

			row = append(row, digit)
		}

		if k > 0 && k <= len(row) {
			row = row[:len(row)-k]
		} else if k > len(row) {
			row = []byte{}
		}

		numStr := string(row)
		num := new(big.Int)
		num.SetString(numStr, 10)

		totalSum.Add(totalSum, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	fmt.Printf("Total: %d\n", totalSum)
	fmt.Printf("Time: %s\n", elapsed)
}
