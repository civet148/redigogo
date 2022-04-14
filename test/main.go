package main

import (
	"github.com/civet148/log"
	"github.com/civet148/redigogo"
	_ "github.com/civet148/redigogo/alone"
	_ "github.com/civet148/redigogo/cluster"
	"github.com/civet148/redigogo/cmd"
)

func main() {
	c := redigogo.Config{
		Password:       "",
		Index:          0, //default db index on redis
		MasterHost:     "192.168.20.105:6379",
		ReplicateHosts: nil,
		ConnTimeout:    500,
		ReadTimeout:    500,
		WriteTimeout:   500,
		KeepAlive:      30,
		AliveTime:      60,
	}
	cache := redigogo.NewCache(&c)
	_ = cache
	cache.Do(cmd.RedisCmdSet, "mykey", "123456")

	if v, err := cache.String(cache.Do(cmd.RedisCmdGet, "mykey")); err != nil {
		log.Debugf("get err %s", err)
	} else {
		log.Debugf("get value %+v", v)
	}

	if v, err := cache.Strings(cache.Do(cmd.RedisCmdConfig, "get", "dbfilename")); err != nil {
		log.Debugf("get err %s", err)
	} else {
		log.Debugf("get value %+v", v)
	}
}
