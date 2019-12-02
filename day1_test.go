package main

import (
	"testing"
)

func TestDay1P1(t *testing.T) {
	testData := `12
14
1969
100756`

	ints := toIntArray(testData)
	if aggregate(ints, getFuel) != 34241 {
		t.Fatalf("Got incorrect value: %d", aggregate(ints, getFuel))
	}

	realInts := toIntArray(loadFileAsString("day1.data"))

	if aggregate(realInts, getFuel) != 3323874 {
		t.Fatalf("Got incorrect value: %d", aggregate(realInts, getFuel))
	}

}

func TestDay1P2(t *testing.T) {
	if getFuelAdjustedFuel(14) != 2 {
		t.Fatalf("Didn't get 2 for fuel. Got %d", getFuelAdjustedFuel(14))
	}

	if getFuelAdjustedFuel(1969) != 966 {
		t.Fatalf("Didn't get 2 for fuel. Got %d", getFuelAdjustedFuel(1969))
	}

	realInts := toIntArray(loadFileAsString("day1.data"))

	if aggregate(realInts, getFuelAdjustedFuel) != 4982961 {
		t.Fatalf("Got incorrect value: %d", aggregate(realInts, getFuelAdjustedFuel))
	}
}

func getFuelAdjustedFuel(mass int) int {
	currentFuelRequirement := getFuel(mass)
	tally := currentFuelRequirement
	currentFuelRequirement = getFuel(currentFuelRequirement)
	for currentFuelRequirement > 0 {
		tally += currentFuelRequirement
		currentFuelRequirement = getFuel(currentFuelRequirement)
	}

	return tally
}
func aggregate(masses []int, fuelCalcFn func(int) int) int {
	fuelReq := 0
	for _, mass := range masses {
		fuelReq += fuelCalcFn(mass)
	}
	return fuelReq
}

func getFuel(input int) int {
	return int(input/3) - 2
}
