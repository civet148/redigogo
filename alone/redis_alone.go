package alone

import (
	"fmt"
	"github.com/civet148/redigogo"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	REDIS_SCHEME_PREFIX = "redis://"
)

type RedisAlone struct {
	pool *redis.Pool //Redis连接池
}

func init() {

	if err := redigogo.Register(redigogo.AdapterType_Alone, newCache); err != nil {
		panic(err.Error())
	}
}

func newCache(c *redigogo.Config) (cache redigogo.Cache) {

	pool := &redis.Pool{

		MaxIdle:     1,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			r, err := redis.DialURL(REDIS_SCHEME_PREFIX + c.MasterHost) // "redis://127.0.0.1:6379"
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}

			if c.Password != "" {
				//redis auth password
				if _, err := r.Do(redigogo.REDIS_CMD_AUTH, c.Password); err != nil {

					return nil, fmt.Errorf("redis auth password error [%s]", err.Error())
				}
			}

			return r, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do(redigogo.REDIS_CMD_PING)
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
	return &RedisAlone{
		pool: pool, //redis connection pool
	}
}

func (r *RedisAlone) Close() {
	_ = r.pool.Close()
}

func (r *RedisAlone) Do(cmd string, args ...interface{}) (v interface{}, e error) {
	c := r.pool.Get()
	return c.Do(cmd, args...)
}

func (r *RedisAlone) Int(reply interface{}, err error) (v int, e error) {

	return redis.Int(reply, err)
}

func (r *RedisAlone) Int64(reply interface{}, err error) (v int64, e error) {
	return redis.Int64(reply, err)
}

func (r *RedisAlone) Float64(reply interface{}, err error) (v float64, e error) {
	return redis.Float64(reply, err)
}

func (r *RedisAlone) String(reply interface{}, err error) (v string, e error) {
	return redis.String(reply, err)
}

func (r *RedisAlone) Bytes(reply interface{}, err error) (v []byte, e error) {
	return redis.Bytes(reply, err)
}

func (r *RedisAlone) Bool(reply interface{}, err error) (v bool, e error) {
	return redis.Bool(reply, err)
}

func (r *RedisAlone) Values(reply interface{}, err error) (v []interface{}, e error) {
	return redis.Values(reply, err)
}

func (r *RedisAlone) Ints(reply interface{}, err error) (v []int, e error) {
	return redis.Ints(reply, err)
}

func (r *RedisAlone) Strings(reply interface{}, err error) (v []string, e error) {
	return redis.Strings(reply, err)
}

func (r *RedisAlone) StringMap(result interface{}, err error) (v map[string]string, e error) {
	return redis.StringMap(result, err)
}

func (r *RedisAlone) Scan(src []interface{}, dst ...interface{}) (v []interface{}, e error) {
	return redis.Scan(src, dst...)
}
