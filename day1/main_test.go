package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {
	arr := []int{1721, 979, 366, 299, 675, 1456}

	t.Run("Passing asset arr to FindSumPair should return 514579", func(t *testing.T) {
		expectedResult := 514579

		result, _ := FindSumPair(arr)

		if result != expectedResult {
			fmt.Printf("Expected %s\nbut got %s\n", strconv.Itoa(expectedResult), strconv.Itoa(result))
			t.Fail()
		}
	})

	t.Run("Passing asset arr to FindSumTrio shoud return 241861950", func(t *testing.T) {
		expectedResult := 241861950

		result, _ := FindSumTrio(arr)

		if result != expectedResult {
			fmt.Printf("Expected %s\nbut got %s\n", strconv.Itoa(expectedResult), strconv.Itoa(result))
			t.Fail()
		}
	})
}
