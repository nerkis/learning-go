package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"flag"
	"strings"
)

func main() {
	var count, correct, incorrect int
	var file_name string

	flag.StringVar(&file_name, "file", "problems.csv", "name of csv file containing the quiz")
	flag.Parse()
	csv_file, _ := os.Open(file_name)
	csv_reader := csv.NewReader(bufio.NewReader(csv_file))
	input_reader := bufio.NewReader(os.Stdin)

	for {
		count++
		line, error := csv_reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		fmt.Printf("Problem #%d: %s = ", count, line[0])
		user_ans, _ := input_reader.ReadString('\n')
		if line[1] == strings.TrimRight(user_ans, "\n") {
			correct++
		} else {
			incorrect++
		}
	}

	fmt.Printf("You got %d questions correct and %d incorrect.\n", correct, incorrect)
}
