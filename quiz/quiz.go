package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"strings"
	"time"
)

func main() {
	var count, correct int

	file_name := flag.String("file_name", "problems.csv", "name of csv file in the format of 'question,answer'")
	time_limit := flag.Int("time_limit", 30, "time limit for the quiz")
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

	fmt.Println("Ready to start? Press enter to begin")
	input_reader.ReadString('\n')

	timer := time.NewTimer(time.Duration(*time_limit) * time.Second)
	go func() {
		<-timer.C
		exit(fmt.Sprintf("\nTimer expired! You got %d questions correct and %d incorrect.\n", correct, (count - correct)))
	}()

	for _, line := range lines {
		count++
		fmt.Printf("Problem #%d: %s = ", count, line[0])
		ans, _ := input_reader.ReadString('\n')
		if IsCorrect(line[1], ans) {
			correct++
		}
	}
	fmt.Printf("You got %d questions correct and %d incorrect.\n", correct, (count - correct))
}

func IsCorrect(expected, ans string) bool {
	ans = strings.ToLower(strings.TrimSpace(strings.TrimRight(ans, "\n")))
	return strings.ToLower(expected) == ans
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
