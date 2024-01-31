package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cowsCount int
	message   string
}

func (ice *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", ice.cowsCount, ice.message)
}

// DivideFood evenly divides the available fodder between a given number of cows
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	totalFodder, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (totalFodder * fatteningFactor / float64(cows)), nil
}

// ValidateInputAndDivideFood adds more robust error handling around the inputs to DivideFood
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows < 0 || cows == 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, cows)
}

// ValidateNumberOfCows returns an error if there are zero or negative cows
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cowsCount: cows,
			message:   "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			cowsCount: cows,
			message:   "no cows don't need food",
		}
	}

	return nil
}
