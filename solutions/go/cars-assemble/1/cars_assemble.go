package cars

const (
	// CarCost is the cost of a single car.
	CarCost = 10000

	// CarsInGroup is the number of cars in a group for a discounted price.
	GroupCost = 95000

	// CarsInGroup is the number of cars in a group for a discounted price.
	CarsInGroup = 10
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var fullGroups int = carsCount / CarsInGroup
	var indivualCars int = carsCount % CarsInGroup
	return uint(fullGroups*GroupCost + indivualCars*CarCost)
}
