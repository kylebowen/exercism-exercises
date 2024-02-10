package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) (percentageRate float32) {
	switch {
	case balance < 0.0:
		percentageRate = 3.213
	case 0.0 <= balance && balance < 1000.0:
		percentageRate = 0.5
	case 1000.0 <= balance && balance < 5000.0:
		percentageRate = 1.621
	case 5000.0 <= balance:
		percentageRate = 2.475
	}

	return percentageRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * (float64(InterestRate(balance)) / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int

	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
	}
	return years
}
