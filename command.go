package go_redis_pool

import "time"

/*
* Replica only command
	RandomKey() string
	Rename(key, newkey string) string
	RenameNX(key, newkey string) bool
	Scan(cursor uint64, match string, count int64) *ScanCmd
	BLPop(timeout time.Duration, keys ...string) []string
	BRPop(timeout time.Duration, keys ...string) []string
*/
type ZSetObject struct {
	Score  float64
	Member string
}

type command interface {
	GetSet(key string, value interface{}) (string, error)
	Del(keys ...string) (int, error)
	Unlink(keys ...string) (int, error)
	Exists(keys ...string) (int, error)
	Expire(key string, expiration time.Duration) (bool, error)
	ExpireAt(key string, tm time.Time) (bool, error)
	TTL(key string) (int, error)
	Type(key string) (string, error)

	/* Int */
	Incr(key string) (int, error)
	IncrBy(key string, value int64) (int, error)
	IncrByFloat(key string, value float64) (string, error)
	Decr(key string) (int, error)
	DecrBy(key string, decrement int64) (int, error)

	/* String */
	Append(key, value string) (int, error)
	Get(key string) (string, error)
	GetRange(key string, start, end int64) (string, error)
	MGet(keys ...string) ([]interface{}, error)
	MSet(pairs ...interface{}) (bool, error)
	MSetNX(pairs ...interface{}) (bool, error)
	Set(key string, value interface{}, expiration time.Duration) (bool, error)
	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)
	SetXX(key string, value interface{}, expiration time.Duration) (bool, error)
	SetRange(key string, offset int64, value string) (int, error)
	StrLen(key string) (int, error)

	/* Bit */
	GetBit(key string, offset int64) int
	SetBit(key string, offset int64, value int) int
	BitCount(key string, start, end int64) int
	BitOpAnd(destKey string, keys ...string) int
	BitOpOr(destKey string, keys ...string) int
	BitOpXor(destKey string, keys ...string) int
	BitOpNot(destKey string, key string) int
	BitPos(key string, bit int64, pos ...int64) int

	/* Hash */
	HDel(key string, fields ...string) int
	HExists(key, field string) bool
	HGet(key, field string) string
	HGetAll(key string) map[string]string
	HIncrBy(key, field string, incr int64) int
	HIncrByFloat(key, field string, incr float64) float64
	HKeys(key string) []string
	HLen(key string) int
	HMGet(key string, fields ...string) []interface{}
	HMSet(key string, fields map[string]interface{}) string
	HSet(key, field string, value interface{}) bool
	HSetNX(key, field string, value interface{}) bool
	HVals(key string) []string

	/* List */
	BRPopLPush(source, destination string, timeout time.Duration) string
	LIndex(key string, index int64) string
	LInsert(key, op string, pivot, value interface{}) int
	LInsertBefore(key string, pivot, value interface{}) int
	LInsertAfter(key string, pivot, value interface{}) int
	LLen(key string) int
	LPop(key string) string
	BLPop(args ...interface{}) ([]string, error)
	LPush(key string, values ...interface{}) int
	LPushX(key string, value interface{}) int
	LRange(key string, start, stop int64) []string
	LRem(key string, count int64, value interface{}) int
	LSet(key string, index int64, value interface{}) string
	LTrim(key string, start, stop int64) string
	RPop(key string) string
	RPopLPush(source, destination string) string
	RPush(key string, values ...interface{}) int
	RPushX(key string, value interface{}) int

	/* Set */
	SAdd(key string, members ...interface{}) int
	SCard(key string) int
	SDiff(keys ...string) []string
	SDiffStore(destination string, keys ...string) int
	SInter(keys ...string) []string
	SInterStore(destination string, keys ...string) int
	SIsMember(key string, member interface{}) bool
	SMembers(key string) []string
	SMembersMap(key string) map[string]struct{}
	SMove(source, destination string, member interface{}) bool
	SPop(key string) string
	SPopN(key string, count int64) []string
	SRandMember(key string) string
	SRandMemberN(key string, count int64) []string
	SRem(key string, members ...interface{}) int
	SUnion(keys ...string) []string
	SUnionStore(destination string, keys ...string) int

	/* ZSet */
	ZAdd(key string, members ...ZSetObject) int
	ZCard(key string) int
	ZCount(key, min, max string) int
	ZIncrBy(key string, increment float64, member string) float64
	ZInterStore(destination string, params ...interface{}) int
	ZLexCount(key, min, max string) int
	ZRange(key string, start, stop int64) []string
	ZRangeByLex(key string, offset, count int64, min, max string) []string
	ZRangeWithScores(key string, start, stop int64) []ZSetObject
	ZRangeByScore(key string, offset, count int64, min, max string) []string
	ZRangeByScoreWithScores(key string, offset, count int64, min, max string) []ZSetObject
	ZRank(key, member string) int
	ZRem(key string, members ...interface{}) int
	ZRemRangeByRank(key string, start, stop int64) int
	ZRemRangeByScore(key, min, max string) int
	ZRemRangeByLex(key, min, max string) int
	ZRevRange(key string, start, stop int64) []string
	ZRevRangeWithScores(key string, start, stop int64) []ZSetObject
	ZRevRangeByScore(key string, offset, count int64, min, max string) []string
	ZRevRangeByLex(key string, offset, count int64, min, max string) []string
	ZRevRangeByScoreWithScores(key string, offset, count int64, min, max string) []ZSetObject
	ZRevRank(key, member string) int
	ZScore(key, member string) float64
	ZUnionStore(dest string, store []ZSetObject, keys ...string) int

	/* Hyperloglog */
	PFAdd(key string, els ...interface{}) int
	PFCount(keys ...string) int
	PFMerge(dest string, keys ...string) string

	/* GEO */
	/*
		GeoAdd(key string, geoLocation ...*GeoLocation) int
		GeoPos(key string, members ...string) *GeoPosCmd
		GeoRadius(key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
		GeoRadiusRO(key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
		GeoRadiusByMember(key, member string, query *GeoRadiusQuery) *GeoLocationCmd
		GeoRadiusByMemberRO(key, member string, query *GeoRadiusQuery) *GeoLocationCmd
		GeoDist(key string, member1, member2, unit string) float64
		GeoHash(key string, members ...string) []string

		Pipeline() Pipeliner
		Pipelined(fn func(Pipeliner) error) ([]Cmder, error)
		TxPipelined(fn func(Pipeliner) error) ([]Cmder, error)
		TxPipeline() Pipeliner
	*/

	Dump(key string) string
}
