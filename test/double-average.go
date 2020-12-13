package main

import (
	"HongBao/algo"
	"fmt"
)

func main() {
	count, amount := int64(10), int64(10000) // 最小单位为分
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := algo.DoubleAverage(count -i, remain)
		remain -= x
		sum += x
		fmt.Printf("第%d个红包，金额为%f, 剩余金额为：%f\n", i+1,
			float64(x)/float64(100), float64(remain)/float64(100))
	}
	fmt.Println("sum =", sum)
}
