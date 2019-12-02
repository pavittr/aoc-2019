package main

import "testing"

func TestStringConverter(t *testing.T) {
	testString := `12
14
1969
100756`

	intArray := toIntArray(testString)

	if len(intArray) != 4 {
		t.Fatalf("Expected an array with 4 elemnts, got %+v", intArray)
	}

	if intArray[0] != int(12) {
		t.Fatalf("First value should be 12, was %d", intArray[0])
	}

	if intArray[1] != int(14) {
		t.Fatalf("First value should be 14, was %d", intArray[0])
	}
	if intArray[2] != int(1969) {
		t.Fatalf("First value should be 1969, was %d", intArray[0])
	}
	if intArray[3] != int(100756) {
		t.Fatalf("First value should be 100756, was %d", intArray[0])
	}

}

func TestDataLoad(t *testing.T) {
	dataLoaded := loadFileAsString("util.data")
	if dataLoaded != `12
14
1969
100756` {
		t.Fatalf("Failed to load Data, got '%s'", dataLoaded)
	}
}

func TestCSVStringToInt(t *testing.T) {
	testString := "1,2,3"
	resultInts := csvStringToIntArray(testString)
	if len(resultInts) != 3 {
		t.Fatalf("Expected 3 elements, got %+v", resultInts)
	}
	if resultInts[0] != 1 {
		t.Fatalf("Expected 1, got %d", resultInts[0])
	}
	if resultInts[1] != 2 {
		t.Fatalf("Expected 2, got %d", resultInts[1])
	}
	if resultInts[2] != 3 {
		t.Fatalf("Expected 3, got %d", resultInts[2])
	}

}
