package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file")
	flag.Parse()

	problems := parseLines(readFile(*csvFilename))

	correct := 0
	for i, p := range problems {
		readAnswer(i, p, &correct)
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func readFile(csvFilename string) [][]string {
	file, err := os.Open(csvFilename)

	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s", csvFilename)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the CSV file.")
	}

	return lines
}

func readAnswer(i int, p problem, correctCounter *int) {
	fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
	var answer string
	fmt.Scanf("%s\n", &answer)
	if answer == p.a {
		*correctCounter++
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
