# Go Client Redis Pool

Go client redis pool supports master-slave access model in redis protol.

## How to use

See the [example.go](./example/example.go)

### Redis master-slave


```
cf := go_redis_pool.RedisConfig{
	Host: "127.0.0.1",
	Port: 6379,
}
conf := &go_redis_pool.ReplicaConfig{
	Master: cf
}
pool := go_redis_pool.NewReplicaPool(conf, nil)
pool.Do("GET", "a")
```
