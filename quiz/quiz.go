package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"strings"
)

func main() {
	var count, correct int

	file_name := flag.String("file", "problems.csv", "name of csv file in the format of 'question,answer'")
	flag.Parse()

	csv_file, err := os.Open(*file_name)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s\n", *file_name))
	}
	csv_reader := csv.NewReader(csv_file)
	input_reader := bufio.NewReader(os.Stdin)

	lines, err := csv_reader.ReadAll()
	if err != nil {
		exit("Failed to parse file")
	}

	for _, line := range lines {
		count++
		fmt.Printf("Problem #%d: %s = ", count, line[0])
		user_ans, _ := input_reader.ReadString('\n')
		trimmed_answer := strings.TrimRight(user_ans, "\n")
		if line[1] == strings.TrimSpace(trimmed_answer) {
			correct++
		}
	}
	fmt.Printf("You got %d questions correct and %d incorrect.\n", correct, (count - correct))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
