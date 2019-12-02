package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func toIntArray(data string) []int {
	return stringToIntArray(data, "\n")

}

func stringToIntArray(input string, sep string) []int {

	intArray := make([]int, 0)

	for _, str := range strings.Split(input, sep) {
		newInt, err := strconv.Atoi(str)

		if err != nil {
			return nil
		}
		intArray = append(intArray, newInt)
	}

	return intArray
}

func loadFileAsString(fileName string) string {
	fileReader, err := os.Open(fileName)
	if err != nil {
		return fmt.Sprintf("FIle error: %+v", err)

	}
	bytes, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return fmt.Sprintf("Reader err: %+v", err)
	}
	return strings.TrimSpace(string(bytes))
}

func csvStringToIntArray(input string) []int {
	return stringToIntArray(input, ",")
}
