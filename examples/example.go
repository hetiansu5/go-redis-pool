package main

import (
	"github.com/hetiansu5/go-redis-pool"
	"fmt"
)

func main() {
	cf := go_redis_pool.RedisConfig{
		Host: "127.0.0.1",
		Port: 6379,
	}
	conf := &go_redis_pool.ReplicaConfig{
		Master: cf,
		//Slaves: []go_redis_pool.RedisConfig{cf},
	}
	pool := go_redis_pool.NewReplicaPool(conf)
	key := "my:test"
	res, err := pool.Set(key, "aaa")
	if err != nil {
		fmt.Println(err)
		return
	} else if res != true {
		fmt.Println("return error")
		return
	}

	reply, err := pool.Get(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	if reply == "" {
		fmt.Println("it is empty")
		return
	}

	fmt.Print(reply)
}
