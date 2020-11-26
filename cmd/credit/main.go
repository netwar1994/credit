package main

import (
	"fmt"
	"github.com/netwar1994/credit/pkg/credit"
)

func main() {
	monthlyPayment, creditOverpayment, totalPayment := credit.Calculate(1_000_000_00, 36, 20)
	fmt.Println("Ежемесячный платеж:", monthlyPayment)
	fmt.Println("Переплата по кредиту:", creditOverpayment)
	fmt.Println("Общая выплата:", totalPayment)
}