package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
func partOne() int {
	numbers := loadNumbers("../input.txt")
	var current, previous, counter int
	i := 1
	for i < len(numbers) {
		current = numbers[i]
		previous = numbers[i-1]
		diff := previous - current
		if diff < 0 {
			counter++
		}
		i++
	}
	return counter
}

func partTwo() int {
	numbers := loadNumbers("../input_pt2.txt")
	var counter, i int

	for i < len(numbers)-3 {
		prevWindow := sumArr(numbers[i : i+3])
		currWindow := sumArr(numbers[i+1 : i+4])
		diff := prevWindow - currWindow
		if diff < 0 {
			counter++
		}
		i++
	}
	return counter
}

func sumArr(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}

func loadNumbers(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		value := scanner.Text()
		v, _ := strconv.Atoi(value)
		numbers = append(numbers, v)
	}
	return numbers
}
