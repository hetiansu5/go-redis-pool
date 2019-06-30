package go_redis_pool

import (
	"math/rand"
	"strings"
	"time"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	ErrNil = redis.ErrNil
)

type ReplicaPool struct {
	master         *redis.Pool
	slaves         []*redis.Pool
	readFromMaster bool
}

type ReplicaConfig struct {
	Master RedisConfig
	Slaves []RedisConfig
	Opts   Options
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func (r *RedisConfig) GetPortOrDefault() int {
	if r.Port == 0 {
		return 6379
	} else {
		return r.Port
	}
}

var readCommandTab map[string]bool

func NewReplicaPool(conf *ReplicaConfig) *ReplicaPool {
	opts := FillOptions(conf.Opts)
	pool := NewReplicaPoolWithDial(conf, func(r *RedisConfig) (redis.Conn, error) {
		addr := fmt.Sprintf("%s:%d", r.Host, r.GetPortOrDefault())
		conn, err := redis.DialTimeout("tcp", addr, opts.ConnectTimeout, opts.ReadTimeout, opts.WriteTimeout)
		if err != nil {
			return conn, err
		}
		if r.Password != "" {
			if _, err := conn.Do("AUTH", r.Password); err != nil {
				conn.Close()
				return nil, err
			}
		}
		if r.DB != 0 {
			if _, err := conn.Do("SELECT", r.DB); err != nil {
				conn.Close()
				return nil, err
			}
		}
		return conn, nil
	})
	pool.SetWait(opts.Wait)
	pool.SetMaxIdle(opts.MaxIdle)
	pool.SetMaxActive(opts.MaxActive)
	pool.SetIdleTimeout(opts.IdleTimeout)
	pool.readFromMaster = false
	return pool
}

func NewReplicaPoolWithDial(conf *ReplicaConfig, newFn func(redisConf *RedisConfig) (redis.Conn, error)) *ReplicaPool {
	if conf.Master.Host == "" {
		return nil
	}

	rp := new(ReplicaPool)
	rp.master = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return newFn(&conf.Master)
		},
	}

	rp.slaves = make([]*redis.Pool, len(conf.Slaves))
	if conf.Slaves != nil {
		for i, slave := range conf.Slaves {
			rp.slaves[i] = &redis.Pool{
				Dial: func() (redis.Conn, error) {
					return newFn(&slave)
				},
			}
		}
	}
	return rp
}

func (rp *ReplicaPool) SetTestOnBorrow(tb func(c redis.Conn, t time.Time) error) {
	rp.master.TestOnBorrow = tb
	for _, slave := range rp.slaves {
		slave.TestOnBorrow = tb
	}
}

func (rp *ReplicaPool) SetMaxIdle(maxIdle int) {
	rp.master.MaxIdle = maxIdle
	for _, slave := range rp.slaves {
		slave.MaxIdle = maxIdle
	}
}

func (rp *ReplicaPool) SetMaxActive(maxActive int) {
	rp.master.MaxActive = maxActive
	for _, slave := range rp.slaves {
		slave.MaxActive = maxActive
	}
}

func (rp *ReplicaPool) SetWait(wait bool) {
	rp.master.Wait = wait
	for _, slave := range rp.slaves {
		slave.Wait = wait
	}
}

func (rp *ReplicaPool) SetIdleTimeout(idleTimeout time.Duration) {
	rp.master.IdleTimeout = idleTimeout
	for _, slave := range rp.slaves {
		slave.IdleTimeout = idleTimeout
	}
}

func (rp *ReplicaPool) pickPool(command string) *redis.Pool {
	if rp.readFromMaster {
		return rp.master
	}
	if _, ok := readCommandTab[strings.ToUpper(command)]; !ok {
		// write command, just retrive master pool
		return rp.master
	}
	if len(rp.slaves) == 0 {
		return rp.master
	}
	idx := rand.Intn(len(rp.slaves))
	// FIXME: check liveness and retry
	return rp.slaves[idx%len(rp.slaves)]
}

func (rp *ReplicaPool) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := rp.pickPool(commandName).Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}

func (rp *ReplicaPool) WithMaster() *ReplicaPool {
	newPool := new(ReplicaPool)
	// copy master pointer from old pool
	newPool.master = rp.master
	newPool.readFromMaster = true
	return newPool
}

func (rp *ReplicaPool) GetConn(master bool) redis.Conn {
	if master {
		return rp.master.Get()
	}
	idx := rand.Intn(len(rp.slaves))
	// FIXME: check liveness and retry
	return rp.slaves[idx%len(rp.slaves)].Get()
}

func (rp *ReplicaPool) Close() {
	rp.master.Close()
	for _, slave := range rp.slaves {
		slave.Close()
	}
}

func init() {
	readCommands := []string{
		"GET", "TTL", "MGET", "PING", "TYPE", "DUMP", "STRLEN", "EXISTS", "GETRANGE",
		// bit
		"GETBIT", "BITCOUNT",
		// hash
		"HGET", "HLEN", "HVALS", "HKEYS", "HSCAN", "HGETALL", "HEXISTS",
		// list
		"LLEN", "LINDEX", "LRANGE",
		// hll
		"PFCOUNT",
		// zset
		"ZCARD", "ZRANK", "ZSCAN", "ZCOUNT", "ZRANGE", "ZSCORE", "ZREVRANK", "ZRANGEBYLEX",
		"ZRANGEBYSCORE",
		// set
		"SCARD", "SSCAN", "SDIFF", "SINTER", "SUNION", "SMEMBERS", "SISMEMBER", "SRANDMEMBER",
	}
	readCommandTab = make(map[string]bool, 0)
	for _, command := range readCommands {
		readCommandTab[command] = true
	}
}
