package cluster

import (
	"encoding/json"
	"fmt"
	"github.com/civet148/redigogo"
	"github.com/gitstliu/go-redis-cluster"
	"strings"
	"time"
)

const (
	REDIS_CLUSTER_NIL_REPLY = "nil reply"
)

type RedisCluster struct {
	cluster *redis.Cluster
}

func init() {

	if err := redigogo.Register(redigogo.AdapterType_Cluster, newCache); err != nil {
		panic(err.Error())
	}
}

func newCache(c *redigogo.Config) (cache redigogo.Cache) {

	var StartNodes []string
	StartNodes = append(StartNodes, c.MasterHost)
	StartNodes = append(StartNodes, c.ReplicateHosts...)
	cluster, err := redis.NewCluster(&redis.Options{
		StartNodes:   StartNodes,
		ConnTimeout:  time.Duration(c.ConnTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Millisecond,
		KeepAlive:    c.KeepAlive,
		AliveTime:    time.Duration(c.AliveTime) * time.Second,
	})

	if err != nil {
		panic(fmt.Sprintf("redis cluster node [%v] connect error [%v] with config [%+v]", StartNodes, err.Error(), c))
	}
	//redis auth password
	if _, err := cluster.Do(redigogo.REDIS_CMD_AUTH, c.Password); err != nil {

		panic(fmt.Sprintf("redis auth error [%s]", err))
	}
	//redis db select
	if c.Index != 0 {
		if _, err := cluster.Do(redigogo.REDIS_CMD_SELECT, c.Index); err != nil {

			panic(fmt.Sprintf("redis select db index error [%s]", err))
		}
	}
	return &RedisCluster{cluster: cluster}
}

func (c *RedisCluster) Close() {
	c.cluster.Close()
}

func (c *RedisCluster) Do(cmd string, args ...interface{}) (v interface{}, e error) {

	return c.cluster.Do(cmd, args...)
}

func (c *RedisCluster) Int(reply interface{}, err error) (v int, e error) {

	return redis.Int(reply, err)
}

func (c *RedisCluster) Int64(reply interface{}, err error) (v int64, e error) {
	return redis.Int64(reply, err)
}

func (c *RedisCluster) Float64(reply interface{}, err error) (v float64, e error) {
	return redis.Float64(reply, err)
}

func (c *RedisCluster) String(reply interface{}, err error) (v string, e error) {
	return redis.String(reply, err)
}

func (c *RedisCluster) Bytes(reply interface{}, err error) (v []byte, e error) {
	return redis.Bytes(reply, err)
}

func (c *RedisCluster) Bool(reply interface{}, err error) (v bool, e error) {
	return redis.Bool(reply, err)
}

func (c *RedisCluster) Values(reply interface{}, err error) (v []interface{}, e error) {
	return redis.Values(reply, err)
}

func (c *RedisCluster) Ints(reply interface{}, err error) (v []int, e error) {
	return redis.Ints(reply, err)
}

func (c *RedisCluster) Unmarshal(dst interface{}, reply interface{}, err error) (e error) {
	var v []byte
	if v, e = c.Bytes(reply, err); e != nil {
		return
	}

	if e = json.Unmarshal(v, dst); e != nil {
		return
	}
	return
}

func (c *RedisCluster) Strings(reply interface{}, err error) (v []string, e error) {
	return redis.Strings(reply, err)
}

func (c *RedisCluster) StringMap(result interface{}, err error) (v map[string]string, e error) {
	return redis.StringMap(result, err)
}

func (c *RedisCluster) Scan(src []interface{}, dst ...interface{}) (v []interface{}, e error) {
	return redis.Scan(src, dst...)
}

func (c *RedisCluster) IsNilReply(err error) bool {
	if strings.Contains(err.Error(), REDIS_CLUSTER_NIL_REPLY) {
		return true
	}
	return false
}
