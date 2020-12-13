package algo

import (
	"math/rand"
	"time"
)

//两倍均值算法生成红包值

const min = 1

func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	// calculate the max amount
	max := amount - min*count
	// calculate the max average can use
	avg := max / count
	// 两倍的基础加上最小金额 防止出现0
	avg2 := 2*avg + min
	// 随机红包金额序列元素，把二倍均值作为随机数的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min
	return x
}
