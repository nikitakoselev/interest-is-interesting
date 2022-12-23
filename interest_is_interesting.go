package interest

import (
	"math"
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	result := float64(0)

	switch b := balance; {
	case b < 0:
		result = 3.213
	case b >= 0 && b < 1000:
		result = 0.5
	case b >= 1000 && b < 5000:
		result = 1.621
	default:
		result = 2.475
	}

	return float32(result)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) / 100 * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + balance/100*float64(InterestRate(balance))
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	balanceInterest := InterestRate(balance) / 100
	result := math.Round(math.Log2(targetBalance/balance) / math.Log2(float64(1+balanceInterest)))
	return int(result)
}
