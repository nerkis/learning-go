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

	fileName := flag.String("file", "problems.csv", "name of csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit for the quiz, in seconds")
	flag.Parse()

	csvFile, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s\n", *fileName))
	}
	csvReader := csv.NewReader(csvFile)
	inputReader := bufio.NewReader(os.Stdin)

	lines, err := csvReader.ReadAll()
	if err != nil {
		exit("Failed to parse file")
	}

	fmt.Println("Ready to start? Press enter to begin")
	inputReader.ReadString('\n')

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	go func() {
		<-timer.C
		exit(fmt.Sprintf("\nTimer expired! You got %d questions correct and %d incorrect.\n", correct, (count - correct)))
	}()

	for _, line := range lines {
		count++
		fmt.Printf("Problem #%d: %s = ", count, line[0])
		ans, _ := inputReader.ReadString('\n')
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
