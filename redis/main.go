package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:32768", // you redis server address
		Password: "redispw",
		DB:       0, // use default DB
	})

	err := rdb.Set(ctx, "dean", "dut", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "dean").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("dean:", val)

	val2, err := rdb.Get(ctx, "key").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val2)
	}
	//dean: dut
	//key does not exist
}
