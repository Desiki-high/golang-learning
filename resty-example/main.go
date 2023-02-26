package main

import "golang-learning/resty-example/client"

func main() {
	client.RestyGet("http://127.0.0.1:8848/index/get")
}
