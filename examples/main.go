package main

import (
	"fmt"

	"github.com/jerry6654/mathutil/arithmetic"
	"github.com/jerry6654/mathutil/stats"
)

func main() {
	sum := arithmetic.Add(3, 5)
	avg := stats.Average([]float64{1, 2, 3, 4, 5})
	fmt.Printf("3 + 5 = %.1f\n", sum)
	fmt.Printf("Average: %.1f\n", avg)
}
