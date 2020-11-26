package credit

import (
	"fmt"
	"math"
)

func Calculate(amount, period, percent int64) (monthly int64, overpayment int64, total int64) {
	monthlyCreditRate := float64(percent) / 12 / 100
	fmt.Println(monthlyCreditRate)
	allPeriodCreditRate := math.Pow(1 + monthlyCreditRate, 36)
	fmt.Println(allPeriodCreditRate)
	annuityRatio := monthlyCreditRate * allPeriodCreditRate / (allPeriodCreditRate - 1)
	fmt.Println(annuityRatio)
	monthly = int64(annuityRatio * float64(amount))
	overpayment = monthly * period - amount
	total = monthly * period
	fmt.Println(amount * 100)
	return
}
