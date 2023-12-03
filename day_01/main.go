package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Print(calibration("day_01/input.txt"))
}

func calibration(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var calibration int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()
		regex := regexp.MustCompile(`\d`)
		indexOfNumbers := regex.FindAllStringIndex(raw, -1)

		var number string
		number = addFirstNumber(raw, indexOfNumbers)
		number = number + addLastNumber(raw, indexOfNumbers)

		internalNumber, _ := strconv.Atoi(number)
		calibration += internalNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return calibration
}

func addLastNumber(raw string, indexOfNumbers [][]int) string {
	lastIndex := len(indexOfNumbers) - 1
	return raw[indexOfNumbers[lastIndex][0]:indexOfNumbers[lastIndex][1]]
}

func addFirstNumber(raw string, indexOfNumbers [][]int) string {
	return raw[indexOfNumbers[0][0]:indexOfNumbers[0][1]]
}
