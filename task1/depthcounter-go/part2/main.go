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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Borked input", err)
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

	previous := measurements[0] + measurements[1] + measurements[2]

	for i := 0; i < len(lines)-2; i++ {
		current := measurements[i] + measurements[i+1] + measurements[i+2]

		if current > previous {
			larger_count += 1
		}
		previous = current
	}

	fmt.Printf("measurements larger: %d", larger_count)
}
