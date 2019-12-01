package main

import "testing"

func TestStringConverter(t *testing.T) {
	testString := `12
14
1969
100756`

	intArray, err := toIntArray(testString)
	if err != nil {
		t.Fatalf("Got error processing test string. Error is %+v", err)
	}

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
