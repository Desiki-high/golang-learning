package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个带有缓冲区的 channel，缓冲区大小为 2
	ch := make(chan int, 2)

	// 启动两个 goroutine 并发向 channel 中发送数据
	go func() {
		ch <- 1
		fmt.Println("Sent 1 to channel", time.Now().Second())
		ch <- 2
		fmt.Println("Sent 2 to channel", time.Now().Second())
		ch <- 3 // 当缓冲区已满时，该操作会被阻塞直到可以写入
		fmt.Println("Sent 3 to channel", time.Now().Second())
		time.Sleep(5 * time.Second)
		ch <- 4
		fmt.Println("Sent 4 to channel", time.Now().Second())
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Received", <-ch, time.Now().Second())
		time.Sleep(1 * time.Second)
		fmt.Println("Received", <-ch, time.Now().Second())
		fmt.Println("Received", <-ch, time.Now().Second())
		fmt.Println("Received", <-ch, time.Now().Second()) // 当缓冲区为空时，该操作会被阻塞 直到有数据写入
	}()

	time.Sleep(7 * time.Second)
}
