package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Successfully Opened users.json")

	scanner := bufio.NewScanner(f)

	countValidOld := 0
	countValidNew := 0

	for scanner.Scan() { // internally, it advances token based on sperator
		line := scanner.Text()
		hifen := strings.Index(line, "-")
		colon := strings.Index(line, ":")
		space := strings.Index(line, " ")
		min, _ := strconv.Atoi(line[:hifen])
		max, _ := strconv.Atoi(line[hifen+1 : space])
		char := line[space+1 : colon]
		password := line[colon+2:]
		// fmt.Printf("Min: %d, Max: %d, Char: '%s', Pass: '%s', Line: %s\n", min, max, char, password, line)

		countValidOld += validOldPassword(password, char, min, max)
		countValidNew += validNewPassword(password, char, min, max)
	}

	fmt.Println(countValidOld, countValidNew)
}

func validOldPassword(pass string, char string, min int, max int) int {
	count := strings.Count(pass, char)

	if count >= min && count <= max {
		return 1
	}
	return 0
}

func validNewPassword(pass string, char string, min int, max int) int {
	charMin, charMax := string(pass[min-1]), string(pass[max-1])
	if (charMin == char && charMax == char) || (charMin != char && charMax != char) {
		return 0
	}
	return 1
}
