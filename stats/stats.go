// Package stats 提供统计相关函数
package stats

import "sort"

// Average 计算浮点数切片的平均值
func Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// Median 计算中位数（依赖sort包）
func Median(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}
