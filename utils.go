package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func toIntArray(data string) ([]int, error) {

	intArray := make([]int, 0)

	for _, str := range strings.Split(data, "\n") {
		newInt, err := strconv.Atoi(str)

		if err != nil {
			return nil, err
		}
		intArray = append(intArray, newInt)
	}

	return intArray, nil

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
