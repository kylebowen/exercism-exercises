package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	const percentConversion float64 = 100

	return float64(productionRate) * successRate / percentConversion
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	const minutesPerHour float64 = 60

	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(workingCarsPerHour / minutesPerHour)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const singleCarCost uint = 10000
	const tenCarCost uint = 95000

	groupsOfTenCarsCount := uint(carsCount / 10)
	fullPriceCarsCount := uint(carsCount % 10)

	return groupsOfTenCarsCount*tenCarCost + fullPriceCarsCount*singleCarCost
}
