package main

import "fmt"

func main() {

	var (
		k     int
		data  int
		sum   int
		max   int
		start int
		end   int
	)

	//获取输入
	fmt.Scan(&k)
	datas := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&data)
		datas[i] = data
	}

	start = datas[0]
	start1 := datas[0]
	for i := 0; i < k; i++ {
		sum = sum + datas[i]
		if sum > max {
			end = datas[i]
			max = sum
			start = start1
		}
		if max == 0 {
			start = start1
		}
		if sum < 0 {
			if i < k-1 && datas[i+1] >= 0 {
				start1 = datas[i+1]
			}
			sum = 0
		}
	}
	if max == 0 && start < 0 {
		end = datas[k-1]
	}
	fmt.Println(max, start, end)
}
