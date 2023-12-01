package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
    file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var calibration_values []string
	for scanner.Scan() {
		line := scanner.Text()
		var firstNumber, lastNumber string
		for i, ch := range line {
			if unicode.IsDigit(ch) {
				calibration_value := string(ch)
				if firstNumber == "" {
					firstNumber = calibration_value
				}
				lastNumber = calibration_value
			}
			
			if i == len(line) - 1 {
				if lastNumber != "" {
					calibration_values = append(calibration_values, firstNumber+lastNumber)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
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