package main

import (
	"github.com/hetiansu5/go-redis-pool"
	"fmt"
)

func main() {
	conf := &go_redis_pool.ReplicaConfig{
		Master: "127.0.0.1:6379",
		Slaves: []string{"127.0.0.1:6379"},
	}
	pool := go_redis_pool.NewReplicaPool(conf, nil)
	reply, err := pool.Do("GET", "a")
	if err != nil {
		fmt.Println(err)
		return;
	}
	if reply == nil {
		fmt.Println("it is nil")
		return;
	}
	by := reply.([]byte)
	fmt.Print(string(by))
}
