package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	arr := readFile()
	resultPair, errPair := FindSumPair(arr)

	if errPair != nil {
		fmt.Println("Erro trying to find the Pair ", errPair)
		return
	}

	fmt.Println("Result Pair finded: ", resultPair)

	resultTrio, errTrio := FindSumTrio(arr)

	if errTrio != nil {
		fmt.Println("Erro trying to find the Trio ", errTrio)
		return
	}

	fmt.Println("Result Trio finded: ", resultTrio)
}

func readFile() []int {
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

	var tempArr []int

	for scanner.Scan() { // internally, it advances token based on sperator
		test := scanner.Text()
		testInt, _ := strconv.Atoi(test)
		tempArr = append(tempArr, testInt)
	}

	sort.Ints(tempArr)

	return tempArr
}

// FindSumPair look to a array and find a pair that the sum result in 2020
// return the result of the multiplication of them
func FindSumPair(arr []int) (int, error) {
	tempResults := make(map[int]int)

	for _, s := range arr {
		_, found := tempResults[s]
		if found {
			return s * tempResults[s], nil
		}
		tempResults[2020-s] = s
	}

	return 0, errors.New("Nao foi possivel encontrar um par")
}

// FindSumTrio look to a array and find a trio that sum result in 2020
// return the result of themultiplication of them
func FindSumTrio(arr []int) (int, error) {
	endNum := len(arr) - 1
	for i := 0; i < len(arr)-1; i++ {
		innerStartNum := 0
		innerEndNum := endNum - 1
		for j := 0; j < (endNum - 1); j++ {
			tempSum := arr[innerStartNum] + arr[innerEndNum] + arr[endNum]
			if tempSum > 2020 {
				innerEndNum--
			}
			if tempSum < 2020 {
				innerStartNum++
			}
			if tempSum == 2020 {
				result := arr[innerStartNum] * arr[innerEndNum] * arr[endNum]
				return result, nil
			}
		}
		endNum--
	}

	return 0, errors.New("Nao foi possivel encontrar o trio")
}

func sumItems(arr []int, value int) int {
	temp := 0

	for _, item := range arr {
		temp += item
	}

	temp += value

	return temp
}

func multItems(arr []int, value int) int {
	temp := 1

	for _, item := range arr {
		temp = item * temp
	}

	temp = temp * value

	return temp
}
