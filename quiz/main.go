package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func loadCSV() map[string]string {

	problemsArray := make(map[string]string)

	csvFile := "problems.csv"

	// Open file
	f, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Read file content into a variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		problemsArray[line[0]] = line[1]

	}

	return problemsArray

}

func executeExercise() {

	var correctGuesses int
	problemsArray := loadCSV()

	reader := bufio.NewReader(os.Stdin)

	for key, value := range problemsArray {
		fmt.Println("What " + key + ", sir?")
		userInput, _ := reader.ReadString('\n')

		if strings.TrimRight(userInput, "\n") == value {
			correctGuesses++
		}
	}

	fmt.Println("Answered " + strconv.Itoa(correctGuesses) + " out of " + strconv.Itoa(len(problemsArray)) + " questions correctly.")

}

func main() {

	go executeExercise()

	time.Sleep(15 * time.Second)

}
