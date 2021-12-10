package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	larger_count := 0

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading all bytes", err)
	}

	lines := strings.Split(string(bytes), "\n")

	var measurements []int
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Error converting to int")
		}
		measurements = append(measurements, i)
	}

	previous := measurements[0]

	for _, measurement := range measurements {
		if measurement > previous {
			larger_count += 1
		}
		previous = measurement
	}

	fmt.Printf("measurements larger: %d", larger_count)
}
