package redigogo

import (
	"fmt"
)

const (
	REDIS_CMD_AUTH   = "AUTH"
	REDIS_CMD_SET    = "SET"
	REDIS_CMD_GET    = "GET"
	REDIS_CMD_DEL    = "DEL"
	REDIS_CMD_EVAL   = "EVAL"
	REDIS_CMD_EXPIRE = "EXPIRE"
	REDIS_CMD_PING   = "PING"
	REDIS_CMD_SELECT = "SELECT"
)

const (
	REDIS_CONN_TIMEOUT    = 500
	REDIS_READ_TIMEOUT    = 500
	REDIS_WRITE_TIMEOUT   = 500
	REDIS_KEEP_ALIVE_TIME = 30
	REDIS_ALIVE_TIME      = 60
)

type Config struct {
	Password       string   `json:"password"`        //auth password
	Index          int      `json:"db_index"`        //db index
	MasterHost     string   `json:"master_host"`     //master node
	ReplicateHosts []string `json:"replicate_hosts"` //replicate node
	ConnTimeout    int      `json:"conn_timeout"`    //millisecond, default 500 ms
	ReadTimeout    int      `json:"read_timeout"`    //millisecond, default 500 ms
	WriteTimeout   int      `json:"write_timeout"`   //millisecond, default 500 ms
	KeepAlive      int      `json:"keep_alive"`      //second, default 30s
	AliveTime      int      `json:"alive_time"`      //second, default 60s
}

type AdaperType int

const (
	AdapterType_Alone   = 1
	AdapterType_Cluster = 2
)

func (a AdaperType) GoString() string {
	return a.String()
}

func (a AdaperType) String() string {
	switch a {
	case AdapterType_Alone:
		return "AdapterType_Alone"
	case AdapterType_Cluster:
		return "AdapterType_Cluster"
	}
	return "AdapterType_Unknown"
}

type Cache interface {
	Close()
	Do(cmd string, args ...interface{}) (interface{}, error)
	Int(reply interface{}, err error) (int, error)
	Int64(reply interface{}, err error) (int64, error)
	Float64(reply interface{}, err error) (float64, error)
	String(reply interface{}, err error) (string, error)
	Bytes(reply interface{}, err error) ([]byte, error)
	Bool(reply interface{}, err error) (bool, error)
	Values(reply interface{}, err error) ([]interface{}, error)
	Ints(reply interface{}, err error) ([]int, error)
	Strings(reply interface{}, err error) ([]string, error)
	StringMap(result interface{}, err error) (map[string]string, error)
	Scan(src []interface{}, dst ...interface{}) ([]interface{}, error)
	Unmarshal(dst interface{}, reply interface{}, err error) error
	IsNilReply(err error) bool
}

type Instance func(c *Config) Cache

var AdapterMap = make(map[AdaperType]Instance)

func Register(adapter AdaperType, inst Instance) (err error) {
	if _, ok := AdapterMap[adapter]; !ok {

		AdapterMap[adapter] = inst
		return
	}
	err = fmt.Errorf("adapter [%v] instance already exists", adapter)
	return
}

func NewCache(c *Config) (cache Cache) {

	if c.ConnTimeout <= 0 {
		c.ConnTimeout = REDIS_CONN_TIMEOUT
	}
	if c.ReadTimeout <= 0 {
		c.ReadTimeout = REDIS_READ_TIMEOUT
	}
	if c.WriteTimeout <= 0 {
		c.WriteTimeout = REDIS_WRITE_TIMEOUT
	}
	if c.KeepAlive <= 0 {
		c.KeepAlive = REDIS_KEEP_ALIVE_TIME
	}
	if c.AliveTime <= 0 {
		c.AliveTime = REDIS_ALIVE_TIME
	}

	if len(c.ReplicateHosts) == 0 {
		if fn, ok := AdapterMap[AdapterType_Alone]; !ok {
			panic("redis adapter alone instance not registered")
		} else {
			return fn(c)
		}
	} else {
		if fn, ok := AdapterMap[AdapterType_Cluster]; !ok {
			panic("redis adapter cluster instance not registered")
		} else {
			return fn(c)
		}
	}
	return
}
