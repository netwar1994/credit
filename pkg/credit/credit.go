package credit

import (
	"math"
)

func Calculate(amount, period, percent int64) (monthly int64, overpayment int64, total int64) {
	monthlyCreditRate := float64(percent) / 12 / 100
	allPeriodCreditRate := math.Pow(1 + monthlyCreditRate, 36)
	annuityRatio := monthlyCreditRate * allPeriodCreditRate / (allPeriodCreditRate - 1)
	monthly = int64(annuityRatio * float64(amount))
	overpayment = int64(annuityRatio * float64(amount) * float64(period)) - amount
	total = int64(annuityRatio * float64(amount) * float64(period))
	return
}
