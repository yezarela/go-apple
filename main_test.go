package main

import (
	"fmt"
	"testing"
)

func TestCountBoxes(t *testing.T) {
	testCases := []struct {
		Cakes    int
		Apples   int
		Expected int
	}{
		{20, 25, 5},
		{20, 12, 4},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Count Box %d", i), func(t *testing.T) {

			b := CountBoxes(tc.Cakes, tc.Apples)

			if b != tc.Expected {
				t.Errorf("Expected %d but found %d", tc.Expected, b)
			}
		})
	}
}

func TestCountItemAmount(t *testing.T) {
	testCases := []struct {
		Cakes         int
		Apples        int
		Boxes         int
		ExpectedCake  int
		ExpectedApple int
	}{
		{20, 25, 5, 4, 5},
		{20, 12, 4, 5, 3},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Count Item Amount %d", i), func(t *testing.T) {

			x, y := CountItemAmount(tc.Cakes, tc.Apples, tc.Boxes)

			if x != tc.ExpectedCake {
				t.Errorf("Expected %d but found %d", tc.ExpectedCake, x)
			}
			if y != tc.ExpectedApple {
				t.Errorf("Expected %d but found %d", tc.ExpectedApple, y)
			}
		})
	}
}
