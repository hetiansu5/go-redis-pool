package go_redis_pool

import (
	"github.com/garyburd/redigo/redis"
)

func (rp *ReplicaPool) Get(key string) (string, error) {
	return redis.String(rp.Do("GET", key))
}

func (rp *ReplicaPool) GetSet(key string, value interface{}) (string, error) {
	return redis.String(rp.Do("GETSET", key, value))
}

func (rp *ReplicaPool) Del(keys ...interface{}) (int, error) {
	return redis.Int(rp.Do("DEL", keys...))
}

func (rp *ReplicaPool) Exists(key ...interface{}) (int, error) {
	return redis.Int(rp.Do("EXISTS", key...))
}

func (rp *ReplicaPool) Expire(key string, expiration int64) (bool, error) {
	return redis.Bool(rp.Do("EXPIRE", key, expiration))
}

func (rp *ReplicaPool) ExpireAt(key string, tm int64) (bool, error) {
	return redis.Bool(rp.Do("EXPIREAT", key, tm))
}

func (rp *ReplicaPool) TTL(key string) (int, error) {
	return redis.Int(rp.Do("TTL", key))
}

func (rp *ReplicaPool) PTTL(key string) (int, error) {
	return redis.Int(rp.Do("PTTL", key))
}

func (rp *ReplicaPool) Type(key string) (string, error) {
	return redis.String(rp.Do("TYPE", key))
}

func (rp *ReplicaPool) Incr(key string) (int, error) {
	return rp.IncrBy(key, 1)
}

func (rp *ReplicaPool) IncrBy(key string, value int64) (int, error) {
	return redis.Int(rp.Do("INCRBY", key, value))
}

func (rp *ReplicaPool) IncrByFloat(key string, value float64) (float64, error) {
	return redis.Float64(rp.Do("INCRBYFLOAT", key, value))
}

func (rp *ReplicaPool) Decr(key string) (int, error) {
	return rp.DecrBy(key, 1)
}

func (rp *ReplicaPool) DecrBy(key string, value int64) (int, error) {
	return redis.Int(rp.Do("DECRBY", key, value))
}

func (rp *ReplicaPool) Append(key string, value string) (int, error) {
	return redis.Int(rp.Do("APPEND", key, value))
}

func (rp *ReplicaPool) GetRange(key string, start, end int64) (string, error) {
	return redis.String(rp.Do("GETRANGE", key, start, end))
}

func (rp *ReplicaPool) SetRange(key string, offset int64, value string) (int, error) {
	return redis.Int(rp.Do("SETRANGE", key, offset, value))
}

func (rp *ReplicaPool) MGet(keys ...interface{}) ([]string, error) {
	return redis.Strings(rp.Do("MGET", keys...))
}

func (rp *ReplicaPool) Set(key string, value interface{}) (bool, error) {
	return isOKString(redis.String(rp.Do("SET", key, value)))
}

func (rp *ReplicaPool) SetEX(key string, value interface{}, seconds int64) (bool, error) {
	return isOKString(redis.String(rp.Do("SET", key, value, "EX", seconds)))
}

func (rp *ReplicaPool) SetPX(key string, value interface{}, milliseconds int64) (bool, error) {
	return isOKString(redis.String(rp.Do("SET", key, value, "PX", milliseconds)))
}

func (rp *ReplicaPool) SetNX(key string, value interface{}) (bool, error) {
	return redis.Bool(rp.Do("SETNX", key, value))
}

func (rp *ReplicaPool) SetXX(key string, value interface{}) (bool, error) {
	return isOKString(redis.String(rp.Do("SET", key, value, "XX")))
}

func (rp *ReplicaPool) MSet(pairs ...interface{}) (bool, error) {
	return isOKString(redis.String(rp.Do("MSET", pairs...)))
}

func (rp *ReplicaPool) MSetNX(pairs ...interface{}) (bool, error) {
	return redis.Bool(rp.Do("MSETNX", pairs...))
}

func (rp *ReplicaPool) StrLen(key string) (int, error) {
	return redis.Int(rp.Do("STRLEN", key))
}

func (rp *ReplicaPool) GetBit(key string, offset int64) (int, error) {
	return redis.Int(rp.Do("GETBIT", key, offset))
}

func (rp *ReplicaPool) SetBit(key string, offset int64, value int) (int, error) {
	return redis.Int(rp.Do("SETBIT", key, offset, value))
}

func (rp *ReplicaPool) BitCount(key string, offsets ...int64) (int, error) {
	switch len(offsets) {
	case 0:
		return redis.Int(rp.Do("BITCOUNT", key))
	case 2:
		return redis.Int(rp.Do("BITCOUNT", key, offsets[0], offsets[1]))
	default:
		return 0, errWrongArguments
	}
}

func (rp *ReplicaPool) BitOpAnd(destKey string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, "AND", destKey)
	args = append(args, keys...)
	return redis.Int(rp.Do("BITOP", args...))
}

func (rp *ReplicaPool) BitOpOr(destKey string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, "OR", destKey)
	args = append(args, keys...)
	return redis.Int(rp.Do("BITOP", args...))
}

func (rp *ReplicaPool) BitOpXor(destKey string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, "XOR", destKey)
	args = append(args, keys...)
	return redis.Int(rp.Do("BITOP", args...))
}

func (rp *ReplicaPool) BitOpNot(destKey, key string) (int, error) {
	return redis.Int(rp.Do("BITOP", "NOT", destKey, key))
}

func (rp *ReplicaPool) BitPos(key string, bit int64, offsets ...int64) (int, error) {
	switch len(offsets) {
	case 0:
		return redis.Int(rp.Do("BITPOS", key, bit))
	case 1:
		return redis.Int(rp.Do("BITPOS", key, bit, offsets[0]))
	case 2:
		return redis.Int(rp.Do("BITPOS", key, bit, offsets[0], offsets[1]))
	default:
		return 0, errWrongArguments
	}
}

func (rp *ReplicaPool) HDel(key string, fields ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, fields...)
	return redis.Int(rp.Do("HDEL", args...))
}

func (rp *ReplicaPool) HExists(key, field string) (bool, error) {
	return redis.Bool(rp.Do("HEXISTS", key, field))
}

func (rp *ReplicaPool) HGet(key, field string) (string, error) {
	return redis.String(rp.Do("HGET", key, field))
}

func (rp *ReplicaPool) HGetAll(key string) (map[string]string, error) {
	return redis.StringMap(rp.Do("HGETALL", key))
}

func (rp *ReplicaPool) HIncrBy(key, field string, value int) (int, error) {
	return redis.Int(rp.Do("HINCRBY", key, field, value))
}

func (rp *ReplicaPool) HIncrByFloat(key, field string, value float64) (float64, error) {
	return redis.Float64(rp.Do("HINCRBYFLOAT", key, field, value))
}

func (rp *ReplicaPool) HMGet(key string, fields ...interface{}) ([]interface{}, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, fields...)
	return redis.Values(rp.Do("HMGET", args...))
}

func (rp *ReplicaPool) HMSet(key string, pairs ...interface{}) (bool, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, pairs...)
	return isOKString(redis.String(rp.Do("HMSET", args...)))
}

func (rp *ReplicaPool) HSet(key string, field string, value interface{}) (bool, error) {
	return redis.Bool(rp.Do("HSET", key, field, value))
}

func (rp *ReplicaPool) HSetNX(key string, field string, value interface{}) (bool, error) {
	return redis.Bool(rp.Do("HSETNX", key, field, value))
}

func (rp *ReplicaPool) HVals(key string) ([]string, error) {
	return redis.Strings(rp.Do("HVALS", key))
}

func (rp *ReplicaPool) HLen(key string) (int, error) {
	return redis.Int(rp.Do("HLEN", key))
}

func (rp *ReplicaPool) BRPopLPush(source, destination string, timeout uint64) (string, error) {
	return redis.String(rp.Do("BRPOPLPUSH", source, destination, timeout))
}

func (rp *ReplicaPool) LIndex(key string, index int64) (string, error) {
	return redis.String(rp.Do("LINDEX", key, index))
}

func (rp *ReplicaPool) LInsert(key string, op string, pivot, value interface{}) (int, error) {
	return redis.Int(rp.Do("LINSERT", key, op, pivot, value))
}

func (rp *ReplicaPool) LInsertBefore(key string, pivot, value interface{}) (int, error) {
	return redis.Int(rp.Do("LINSERT", key, "BEFORE", pivot, value))
}

func (rp *ReplicaPool) LInsertAfter(key string, pivot, value interface{}) (int, error) {
	return redis.Int(rp.Do("LINSERT", key, "AFTER", pivot, value))
}

func (rp *ReplicaPool) LLen(key string) (int, error) {
	return redis.Int(rp.Do("LLEN", key))
}

func (rp *ReplicaPool) LPop(key string) (string, error) {
	return redis.String(rp.Do("LPOP", key))
}

func (rp *ReplicaPool) BLPop(args ...interface{}) ([]string, error) {
	return redis.Strings(rp.Do("BLPOP", args...))
}

func (rp *ReplicaPool) LPush(key string, values ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, values...)
	return redis.Int(rp.Do("LPUSH", args...))
}

func (rp *ReplicaPool) LPushX(key string, value interface{}) (int, error) {
	return redis.Int(rp.Do("LPUSHX", key, value))
}

func (rp *ReplicaPool) LRange(key string, start, stop int64) ([]string, error) {
	return redis.Strings(rp.Do("LRANGE", key, start, stop))
}

func (rp *ReplicaPool) LRem(key string, count int64, value interface{}) (int, error) {
	return redis.Int(rp.Do("LREM", key, count, value))
}

func (rp *ReplicaPool) LSet(key string, index int64, value interface{}) (bool, error) {
	return isOKString(redis.String(rp.Do("LSET", key, index, value)))
}

func (rp *ReplicaPool) LTrim(key string, start, stop int64) (bool, error) {
	return isOKString(redis.String(rp.Do("LTRIM", key, start, stop)))
}

func (rp *ReplicaPool) RPop(key string) (string, error) {
	return redis.String(rp.Do("RPOP", key))
}

func (rp *ReplicaPool) RPopLPush(source, destination string) (string, error) {
	return redis.String(rp.Do("RPOPLPUSH", source, destination))
}

func (rp *ReplicaPool) RPush(key string, values ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, values...)
	return redis.Int(rp.Do("RPUSH", args...))
}

func (rp *ReplicaPool) RPushX(key string, value interface{}) (int, error) {
	return redis.Int(rp.Do("RPUSHX", key, value))
}

func (rp *ReplicaPool) SAdd(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, members...)
	return redis.Int(rp.Do("SADD", args...))
}

func (rp *ReplicaPool) SCard(key string) (int, error) {
	return redis.Int(rp.Do("SCARD", key))
}

func (rp *ReplicaPool) SDiff(keys ...interface{}) ([]string, error) {
	return redis.Strings(rp.Do("SDIFF", keys...))
}

func (rp *ReplicaPool) SDiffStore(destination string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destination)
	args = append(args, keys...)
	return redis.Int(rp.Do("SDIFFSTORE", args...))
}

func (rp *ReplicaPool) SInter(keys ...interface{}) ([]string, error) {
	return redis.Strings(rp.Do("SINTER", keys...))
}

func (rp *ReplicaPool) SInterStore(destination string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destination)
	args = append(args, keys...)
	return redis.Int(rp.Do("SINTERSTORE", args...))
}

func (rp *ReplicaPool) SIsMember(key string, member interface{}) (bool, error) {
	return redis.Bool(rp.Do("SISMEMBER", key, member))
}

func (rp *ReplicaPool) SMove(source, destination string, member interface{}) (bool, error) {
	return redis.Bool(rp.Do("SMOVE", source, destination, member))
}

func (rp *ReplicaPool) SMembers(key string) ([]string, error) {
	return redis.Strings(rp.Do("SMEMBERS", key))
}

func (rp *ReplicaPool) SPop(key string) (string, error) {
	return redis.String(rp.Do("SPOP", key))
}

func (rp *ReplicaPool) SPopN(key string, count int64) ([]string, error) {
	return redis.Strings(rp.Do("SPOP", key, count))
}

func (rp *ReplicaPool) SRandMember(key string) (string, error) {
	return redis.String(rp.Do("SRANDMEMBER", key))
}

func (rp *ReplicaPool) SRandMemberN(key string, count int64) ([]string, error) {
	return redis.Strings(rp.Do("SRANDMEMBER", key, count))
}

func (rp *ReplicaPool) SRem(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, members...)
	return redis.Int(rp.Do("SREM", args...))
}

func (rp *ReplicaPool) SUnion(keys ...interface{}) ([]string, error) {
	return redis.Strings(rp.Do("SUNION", keys...))
}

func (rp *ReplicaPool) SUnionStore(destionation string, keys ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destionation)
	args = append(args, keys...)
	return redis.Int(rp.Do("SUNIONSTORE", args...))
}

func (rp *ReplicaPool) ZAdd(key string, pairs ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, pairs...)
	return redis.Int(rp.Do("ZADD", args...))
}

func (rp *ReplicaPool) ZCard(key string) (int, error) {
	return redis.Int(rp.Do("ZCARD", key))
}

func (rp *ReplicaPool) ZCount(key string, min, max interface{}) (int, error) {
	return redis.Int(rp.Do("ZCOUNT", key, min, max))
}

func (rp *ReplicaPool) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return redis.Float64(rp.Do("ZINCRBY", key, increment, member))
}

func (rp *ReplicaPool) ZInterStore(destination string, nkeys int, params ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destination)
	args = append(args, nkeys)
	args = append(args, params...)
	return redis.Int(rp.Do("ZINTERSTORE", args...))
}

func (rp *ReplicaPool) ZLexCount(key, min, max string) (int, error) {
	return redis.Int(rp.Do("ZLEXCOUNT", key, min, max))
}

func (rp *ReplicaPool) ZRange(key string, start, stop int64) ([]string, error) {
	return redis.Strings(rp.Do("ZRANGE", key, start, stop))
}

func (rp *ReplicaPool) ZRangeWithScores(key string, start, stop int64) (map[string]string, error) {
	return redis.StringMap(rp.Do("ZRANGE", key, start, stop, "WITHSCORES"))
}

func (rp *ReplicaPool) ZRangeByLex(key, min, max string, offset, count int64) ([]string, error) {
	return redis.Strings(rp.Do("ZRANGEBYLEX", key, min, max, "LIMIT", offset, count))
}

func (rp *ReplicaPool) ZRangeByScore(key, min, max interface{}, offset, count int64) ([]string, error) {
	return redis.Strings(rp.Do("ZRANGEBYSCORE", key, min, max, "LIMIT", offset, count))
}

func (rp *ReplicaPool) ZRank(key, member string) (int, error) {
	return redis.Int(rp.Do("ZRANK", key, member))
}

func (rp *ReplicaPool) ZRem(key string, members ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, members...)
	return redis.Int(rp.Do("ZREM", args...))
}

func (rp *ReplicaPool) ZRemRangeByLex(key, min, max string) (int, error) {
	return redis.Int(rp.Do("ZREMRANGEBYLEX", key, min, max))
}

func (rp *ReplicaPool) ZRemRangeByRank(key string, start, stop int64) (int, error) {
	return redis.Int(rp.Do("ZREMRANGEByRANK", key, start, stop))
}

func (rp *ReplicaPool) ZRemRangeByScore(key, min, max interface{}) (int, error) {
	return redis.Int(rp.Do("ZREMRANGEBYSCORE", key, min, max))
}

func (rp *ReplicaPool) ZRevRange(key string, start, stop int64) ([]string, error) {
	return redis.Strings(rp.Do("ZREVRANGE", key, start, stop))
}

func (rp *ReplicaPool) ZRevRangeWithScores(key string, start, stop int64) (map[string]string, error) {
	return redis.StringMap(rp.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
}

func (rp *ReplicaPool) ZRevRangeByLex(key, max, min string, offset, count int64) ([]string, error) {
	return redis.Strings(rp.Do("ZREVRANGEBYLEX", key, max, min, "LIMIT", offset, count))
}

func (rp *ReplicaPool) ZRevRangeByScore(key, max, min interface{}, offset, count int64) ([]string, error) {
	return redis.Strings(rp.Do("ZREVRANGEBYSCORE", key, max, min, "LIMIT", offset, count))
}

func (rp *ReplicaPool) ZRevRangeByScoreWithScores(key, max, min interface{}, offset, count int64) (map[string]string, error) {
	return redis.StringMap(rp.Do("ZREVRANGEBYSCORE", key, max, min, "WITHSCORES", "LIMIT", offset, count))
}

func (rp *ReplicaPool) ZRevRank(key, member string) (int, error) {
	return redis.Int(rp.Do("ZREVRANK", key, member))
}

func (rp *ReplicaPool) ZScore(key, member string) (float64, error) {
	return redis.Float64(rp.Do("ZSCORE", key, member))
}

func (rp *ReplicaPool) ZUnionStore(destination string, nkeys int, params ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, destination)
	args = append(args, nkeys)
	args = append(args, params...)
	return redis.Int(rp.Do("ZUNIONSTORE", args...))
}

func (rp *ReplicaPool) PFAdd(key string, els ...interface{}) (int, error) {
	args := make([]interface{}, 0)
	args = append(args, key)
	args = append(args, els...)
	return redis.Int(rp.Do("PFADD", args...))
}

func (rp *ReplicaPool) PFCount(keys ...interface{}) (int, error) {
	return redis.Int(rp.Do("PFCOUNT", keys...))
}

func (rp *ReplicaPool) PFMerge(dest string, keys ...interface{}) (bool, error) {
	args := make([]interface{}, 0)
	args = append(args, dest)
	args = append(args, keys...)
	return isOKString(redis.String(rp.Do("PFMERGE", args...)))
}

func (rp *ReplicaPool) Publish(channel, msg string) (int, error) {
	return redis.Int(rp.Do("PUBLISH", channel, msg))
}

func (rp *ReplicaPool) RawCommand(command string, args ...interface{}) (interface{}, error) {
	return rp.Do(command, args...)
}
