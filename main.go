package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	if filename == "" {
		log.Fatal("Must provide a file")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("No such file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	titleText := ""
	if scanner.Scan() {
		titleText = scanner.Text()
	}

	result := fmt.Sprintf("# %s\n", titleText)

	for scanner.Scan() {
		currLine := scanner.Text()
		result += fmt.Sprintf("* [ ] %s\n", currLine)
	}
	fmt.Println(result)
}
