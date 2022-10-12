package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Если ответ существует, верните список из двух элементов
// Если нет - то верните пустой список
func twoSum(array []int, targetSum int) []int {
	// Ваше решение
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == targetSum {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}

func twoSumV2(array []int, targetSum int) []int {
	sort.Ints(array)
	left := 0
	right := len(array) - 1
	for left < right {
		c_sum := array[left] + array[right]
		if c_sum < targetSum {
			left++
		} else if c_sum > targetSum {
			right--
		} else {
			return []int{array[left], array[right]}
		}
	}
	return []int{}
}

func twoSumV3(array []int, targetSum int) []int {
	previous := make(map[int]int)
	for i := 0; i < len(array); i++ {
		x := targetSum - array[i]
		if _, ok := previous[x]; ok {
			return []int{array[i], x}
		} else {
			previous[array[i]] = i
		}
	}
	return []int{}
}
func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	targetSum := readInt(scanner)
	result := twoSum(array, targetSum)
	if len(result) == 0 {
		fmt.Println("None")
	} else {
		fmt.Print(result[0])
		fmt.Print(" ")
		fmt.Print(result[1])
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
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
