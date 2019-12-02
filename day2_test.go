package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDay2P1(t *testing.T) {
	testFn := func(testProg string, expectedOutput string) {
		progOutput := processGravityAssist(testProg)
		if progOutput != expectedOutput {
			t.Fatalf("Prog %s failed.\n Expected %s\n got %s", testProg, expectedOutput, progOutput)
		}
	}

	testFn(`1,0,0,0,99`, `2,0,0,0,99`)
	testFn(`1,0,0,3,99`, `1,0,0,2,99`)
	testFn(`2,3,0,3,99`, `2,3,0,6,99`)
	testFn(`2,4,4,5,99,0`, `2,4,4,5,99,9801`)
	testFn(`1,1,1,4,99,5,6,0,99`, `30,1,1,4,2,5,6,0,99`)
	realProgInput := csvStringToIntArray(loadFileAsString("day2.data"))
	realProgInput[1] = 12
	realProgInput[2] = 2
	testFn(intsToString(realProgInput), `3058646,12,2,2,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,24,2,19,6,48,1,23,5,49,1,9,27,52,1,31,10,56,2,35,9,168,1,5,39,169,2,43,9,507,1,5,47,508,2,51,13,2540,1,55,10,2544,1,59,10,2548,2,9,63,7644,1,67,5,7645,2,13,71,38225,1,75,10,38229,1,79,6,38231,2,13,83,191155,1,87,6,191157,1,6,91,191159,1,10,95,191163,2,99,6,382326,1,103,5,382327,2,6,107,764654,1,10,111,764658,1,115,5,764659,2,6,119,1529318,1,123,5,1529319,2,127,6,3058638,1,131,5,3058639,1,2,135,3058641,1,139,13,0,99,2,0,14,0`)

}

func TestDay2P2(t *testing.T) {
	programString := loadFileAsString("day2.data")

	realProgInput := csvStringToIntArray(programString)
	realProgInput[1] = 12
	realProgInput[2] = 2

	output := csvStringToIntArray(processGravityAssist(intsToString(realProgInput)))
	if output[0] != 3058646 {
		t.Fatalf("Expected output 3058646, got %d", output[0])
	}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			realProgInput = csvStringToIntArray(programString)
			realProgInput[1] = noun
			realProgInput[2] = verb
			output = csvStringToIntArray(processGravityAssist(intsToString(realProgInput)))
			if output[0] == 19690720 {
				fmt.Printf("Found noun %d and verb %d makes 19690720. Answer is %d\n", noun, verb, (100*noun)+verb)
			}

		}
	}

}

func processGravityAssist(inputString string) string {
	inputInts := csvStringToIntArray(inputString)

	cursor := 0
	inputInts = update(inputInts, cursor)

	for inputInts[cursor] != 99 {
		cursor += 4
		inputInts = update(inputInts, cursor)
	}

	return intsToString(inputInts)

}

func intsToString(ints []int) string {
	b := []byte{}
	for _, n := range ints {
		b = strconv.AppendInt(b, int64(n), 10)
		b = append(b, ',')
	}
	b = b[:len(b)-1]
	return string(b)
}

func update(ints []int, cursor int) []int {
	if ints[cursor] == 99 {
		return ints
	}
	if ints[cursor] == 1 {
		int1 := ints[cursor+1]
		int2 := ints[cursor+2]
		storePos := ints[cursor+3]
		ints[storePos] = ints[int1] + ints[int2]
		return ints
	}
	if ints[cursor] == 2 {
		int1 := ints[cursor+1]
		int2 := ints[cursor+2]
		storePos := ints[cursor+3]
		ints[storePos] = ints[int1] * ints[int2]
		return ints
	}

	return ints
}
