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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d`)

	var calibration_values []string
	for scanner.Scan() {
		line := scanner.Text()
		digit_indexes := re.FindAllStringIndex(line, -1)
		firstNumber := string(line[digit_indexes[0][0]])
		lastNumber := string(line[digit_indexes[len(digit_indexes)-1][0]])
		calibration_values = append(calibration_values, firstNumber+lastNumber)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	for _, value := range calibration_values {
		int_value, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		sum += int_value
	}

	fmt.Println(sum)
}
