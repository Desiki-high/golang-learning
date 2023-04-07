package main

import (
	"context"
	"fmt"
	"time"
)

/*
当你在编写一个并发程序时，你可能需要控制一组goroutine的生命周期，
例如在一个HTTP请求处理中启动多个goroutine来处理请求，但是如果请求被取消或超时，你可能需要停止所有的goroutine以释放资源。
Context就提供了这样的功能，当你创建一个Context时，可以指定一个“父”Context，当父Context被取消时，所有的子Context都会被取消，这样就可以方便地控制一组goroutine的生命周期。
*/

// 控制goroutine的生命周期
func controlRoutineGroup() {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)

	go worker(ctx, "A")
	go worker(ctx, "B")

	// 模拟请求取消操作
	time.Sleep(5 * time.Second)
	//上层上下文关闭，协程会select ctx.Done() channel 打印stopping return 协程结束
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Parent function is stopping")
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %s is stopping\n", name)
			return
		default:
			fmt.Printf("worker %s is working\n", name)
			time.Sleep(time.Second)
		}
	}
}

/*
当你处理一个请求时，可能需要在多个goroutine之间传递一些请求相关的元数据，如请求的截止时间、请求的用户身份信息等等。
Context就可以用来传递这些元数据，每个goroutine都可以从Context中获取这些元数据，而不需要通过函数参数传递。
*/

// 跨goroutine传递请求相关的元数据
func transferMsgBetweenRoutine() {
	parentCtx := context.Background()

	// 创建一个带有用户ID的子Context
	userID := 123
	ctx := context.WithValue(parentCtx, "userID", userID)

	// 处理请求
	req := &Request{ID: 1, UserID: userID, Created: time.Now()}
	handleRequestMsg(ctx, req)
}

type Request struct {
	ID      int
	UserID  int
	Created time.Time
}

func handleRequestMsg(ctx context.Context, req *Request) {
	// 从Context中获取用户ID
	userID := ctx.Value("userID").(int)

	fmt.Printf("processing request %d for user %d\n", req.ID, userID)
}

/*
当你处理一个请求时，可能需要支持请求的取消操作，例如当用户在等待请求响应时，可以通过取消请求来提高响应速度。
Context就提供了这样的功能，当你创建一个Context时，可以指定一个“取消函数”，当需要取消请求时，调用这个取消函数就可以取消所有与这个Context相关的goroutine。
*/

// 支持请求的取消操作
func cancelRequest() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 9*time.Second)
	defer cancel()

	go handleRequest(ctx)

	// 等待处理完成或者超时
	select {
	case <-ctx.Done():
		fmt.Println("Parent function is cancelled")
	}
}

func handleRequest(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("request is cancelled")
	case <-time.After(10 * time.Second):
		fmt.Println("request is processed")
	}
}

func main() {
	controlRoutineGroup()
	transferMsgBetweenRoutine()
	cancelRequest()
	//等待handleRequest打印
	time.Sleep(1)
}
