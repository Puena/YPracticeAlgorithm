package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func movingAverage(array []int, windowSize int) []float64 {
	// Ваше решение
	result := make([]float64, len(array) - (windowSize - 1))
	for i := 0; i <= len(array) - windowSize; i++ {
		sum := 0
		for j := i; j < i + windowSize; j++ {
			sum += array[j]
		}
		// result = append(result, float64(sum) / float64(windowSize))
		result[i] = float64(sum) / float64(windowSize)
	}
	return result
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	windowSize := readInt(scanner)
	printArray(movingAverage(array, windowSize))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []float64) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.FormatFloat(arr[i], 'f', 8, 64))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
