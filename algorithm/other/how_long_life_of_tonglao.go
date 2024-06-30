package main

import "fmt"

//每30年返老还童一次，每次恢复的时间逐渐变长。36岁需要30天，66岁需要60天，96岁需要90天。
//直到某一次需要的恢复时间太长，我们认为天山童姥到达生命尽头。她能活多少岁呢？

const (
	daysOfOneYear       = 365 // 简单认为一年有365天，不区分闰年/平年
	yearsOfOneCycle     = 30  // 每30年进行一次神功循环
	initialRecycleYear  = 36  // 天山童姥开始还童的年龄
	initialPracticeYear = 6   // 天山童姥开始练功的时间
)

func main() {
	var yearOfTongLao = initialRecycleYear
	for {
		if yearOfTongLao-initialPracticeYear <= yearsOfOneCycle*daysOfOneYear {
			yearOfTongLao += yearsOfOneCycle
		} else {
			yearOfTongLao += yearsOfOneCycle - 1
			break
		}
	}

	fmt.Printf("天山童姥能活%d岁\n", yearOfTongLao)
}
