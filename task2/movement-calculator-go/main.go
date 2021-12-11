package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command string

const (
	Forward Command = "forward"
	Up      Command = "up"
	Down    Command = "down"
)

func main() {
	fmt.Printf("Task2 Part1: %d\n", part1())
	fmt.Printf("Task2 Part2: %d\n", part2())
}

func part1() int {
	file, _ := os.Open("input")
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading array", err)
	}

	distance := 0
	height := 0

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")

		amount, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("Error converting to int")
		}

		move := split[0]

		switch move {
		case "forward":
			distance += amount
		case "up":
			height -= amount
		case "down":
			height += amount
		default:
			log.Fatalf("New move ay? %s", move)
		}
	}

	return height * distance
}

func part2() int {
	file, _ := os.Open("input")
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading array", err)
	}

	distance := 0
	height := 0
	aim := 0

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		split := strings.Split(line, " ")

		amount, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("Error converting to int")
		}

		move := split[0]

		switch move {
		case "forward":
			distance += amount
			height += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		default:
			log.Fatalf("New move ay? %s", move)
		}
	}

	return height * distance
}
