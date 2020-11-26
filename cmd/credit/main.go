package main

import (
	"fmt"
	"github.com/netwar1994/credit/pkg/credit"
)



func main() {
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
}