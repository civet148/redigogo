package main

import (
	"github.com/civet148/gotools/log"
	"github.com/civet148/redigogo"
	_ "github.com/civet148/redigogo/alone"
	_ "github.com/civet148/redigogo/cluster"
)

func main() {
	c := redigogo.Config{
		Password:       "",
		Index:          0, //default db index on redis
		MasterHost:     "127.0.0.1:6379",
		ReplicateHosts: nil,
		ConnTimeout:    500,
		ReadTimeout:    500,
		WriteTimeout:   500,
		KeepAlive:      30,
		AliveTime:      60,
	}
	cache := redigogo.NewCache(&c)
	_ = cache
	cache.Do(redigogo.REDIS_CMD_SET, "mykey", "123456")

	if v, err := cache.String(cache.Do(redigogo.REDIS_CMD_GET, "mykey")); err != nil {
		log.Debugf("get err %s", err)
	} else {
		log.Debugf("get value %s", v)
	}
}
