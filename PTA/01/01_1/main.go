package main

import "fmt"

func main() {

	var (
		k    int
		data int
	)

	//获取输入
	fmt.Scan(&k)
	datas := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&data)
		datas[i] = data
	}
	max := maxSubSeq3(datas, k)
	fmt.Println(max)
}

//最大子列问题，解法一 时间复杂度O(n^2)
func maxSubSeq1(datas []int, n int) int {
	max := 0
	for i := 0; i <= n; i++ { //子列起始位置
		sum := 0
		for j := i; j < n; j++ {
			sum = sum + datas[j]
			if sum > max {
				max = sum
			}
			if sum < 0 {
				sum = 0
			}
		}
	}
	return max
}

//分治算法 时间复杂度O(nlogn)
func maxSubSeq2(datas []int, n int) int {
	return divideAndConquer(datas, 0, n-1)
}

func divideAndConquer(datas []int, left, right int) int {
	var (
		maxLeft, maxRight             int //左右两边的最大值
		maxLeftBorder, maxRightBorder int //跨边界线的左右两边的最大值
		sumLeftBorder, sumRightBorder int //跨边界线的左右两边的值
		center                        int //中点位置
	)

	//左右两边相等，说明已经分到了最小，只有最后一个数字，不能再分
	//如果该数为负数，直接返回0
	if left == right {
		if datas[left] > 0 {
			return datas[left]
		}
		return 0
	}

	//计算中点
	center = (left + right) / 2

	//递归算出两边子列的最大的和
	maxLeft = divideAndConquer(datas, left, center)
	maxRight = divideAndConquer(datas, center+1, right)

	//求跨分界线的最大子列和
	//从分界线分别开始从左右扫描得出结果
	//
	//从分界线向左扫描
	for i := center; i >= left; i-- {
		sumLeftBorder = sumLeftBorder + datas[i]
		if sumLeftBorder > maxLeftBorder {
			maxLeftBorder = sumLeftBorder
		}
	}

	//从分界线向右扫描
	for i := center + 1; i <= right; i++ {
		sumRightBorder = sumRightBorder + datas[i]
		if sumRightBorder > maxRightBorder {
			maxRightBorder = sumRightBorder
		}
	}

	//比较左右两边以及跨分界线的三个子列和
	return max(maxLeft, maxRight, maxLeftBorder+maxRightBorder)
}

//返回最大值
func max(a ...int) int {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			a[i+1] = a[i]
		}
	}
	return a[len(a)-1]
}

//动态规划，在线处理，O(n)
func maxSubSeq3(datas []int, n int) int {
	var (
		sum int
		max int
	)
	for i := 0; i < n; i++ {
		sum = sum + datas[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
