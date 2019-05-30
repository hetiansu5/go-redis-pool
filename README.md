# Redis Pool

Redis pool supports master-slave access model in redis protol.

## How to use

See the [example.go](./example/example.go)

### Redis master-slave


```
conf := &go_redis_pool.ReplicaConfig{
	Master: "127.0.0.1:6379",
    Slaves: []string{"127.0.0.1:6379"},
}
pool := go_redis_pool.NewReplicaPool(conf, nil)
pool.Do("GET", "a")
```
