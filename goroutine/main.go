package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("Hello, World!")
	say("Hello, Goroutine!")
	handlerMultiTask()
	handlerMultiData()
}

// 使用Goroutine并发地获取多个网页的内容
func handlerMultiTask() {
	urls := []string{
		"https://www.baidu.com",
		"https://www.microsoft.com",
		"https://www.baidu.com",
		"https://www.dlut.edu.cn",
	}

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					panic("close error")
				}
			}(resp.Body)

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Printf("URL: %s, Length: %d\n", url, len(body))
		}(url)
	}
	time.Sleep(5 * time.Second)
}

// 使用Goroutine并发地对一个数组进行排序
func handlerMultiData() {
	rand.Seed(time.Now().UnixNano())

	//创建一个长度为1000000的整型切片nums
	nums := make([]int, 1000000)
	for i := 0; i < len(nums); i++ {
		// 赋值随机数,范围是0-999999
		nums[i] = rand.Intn(1000000)
	}

	start := time.Now()
	// 创建channel让协程能够返回排完序数组
	ch := make(chan []int)
	go mergeSort(nums, ch)

	// 阻塞直到 协程写入了排序完的数组
	result := <-ch
	elapsed := time.Since(start)

	fmt.Printf("Sorted %d numbers in %s\n", len(result), elapsed)
}

func mergeSort(nums []int, ch chan []int) {
	if len(nums) <= 1 {
		ch <- nums
		return
	}

	mid := len(nums) / 2
	leftCh := make(chan []int)
	rightCh := make(chan []int)

	go mergeSort(nums[:mid], leftCh)
	go mergeSort(nums[mid:], rightCh)

	left := <-leftCh
	right := <-rightCh

	close(leftCh)
	close(rightCh)

	result := merge(left, right)
	ch <- result
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
