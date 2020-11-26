package credit_test

import (
	"github.com/netwar1994/credit/pkg/credit"
	"fmt"
)

func ExampleCalculate() {
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
	fmt.Println(credit.Calculate(10_000_00, 36, 20))
	// Output:
	// 3716358 33788900 133788900
	// 37163 337889 1337889

}
