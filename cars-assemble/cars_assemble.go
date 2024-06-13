package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * float64((successRate / 100)))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64((productionRate / 60)) * (successRate / 100))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	/*
		var individualCarCost uint = 10000
		var discountedCarCost uint = 9500

		if carsCount%10 == 0 {
			return uint(carsCount) * discountedCarCost
		} else {
			return ((uint(carsCount) / 10) * discountedCarCost * 10) + (uint(carsCount)%10)*individualCarCost
		}
	*/
	var individualCarCost uint = 10000
	var discountedCarCost uint = 9500

	totalCost := (uint(carsCount) / 10) * discountedCarCost * 10
	remainingCars := uint(carsCount) % 10
	totalCost += remainingCars * individualCarCost

	return totalCost

}
