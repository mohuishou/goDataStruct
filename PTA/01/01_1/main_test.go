package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var (
	testDatas []int //测试的序列
	testLen   int   //测试的长度
)

//初始化测试，根据输入的序列长度，随机生成一个序列
func init() {
	testLen = 10 * 10 * 10 * 10 * 10 * 10 * 10 * 10
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 0
	for i := 0; i < testLen; i++ {
		m := r.Intn(100)
		if i%3 == 0 {
			m = -m
			count++
		}
		testDatas = append(testDatas, m)
	}

	fmt.Println("初始化成功，一共生成" + strconv.Itoa(testLen) + "随机数" + "包括" + strconv.Itoa(count) + "个负数")
}

func Test_1(t *testing.T) {
	if testLen > 100000 {
		fmt.Println("恕我不想等了")
		return
	}
	max := maxSubSeq1(testDatas, testLen)
	fmt.Println(max)
}

func Test_2(t *testing.T) {
	fmt.Println(maxSubSeq2(testDatas, testLen))
}

func Test_3(t *testing.T) {
	fmt.Println(maxSubSeq3(testDatas, testLen))
}
