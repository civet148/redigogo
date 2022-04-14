package cmd

//reference https://www.redis.net.cn/order/
const (
	//键(key) 命令
	RedisCmdType      = "Type"      //返回 key 所储存的值的类型
	RedisCmdPexpireAt = "PEXPIREAT" //设置 key 过期时间的时间戳(unix timestamp) 以毫秒计
	RedisCmdRename    = "Rename"    //修改 key 的名称
	RedisCmdPersist   = "PERSIST"   //移除 key 的过期时间，key 将持久保持
	RedisCmdMove      = "Move"      //将当前数据库的 key 移动到给定的数据库 db 当中
	RedisCmdRandomKey = "RANDOMKEY" //从当前数据库中随机返回一个 key
	RedisCmdDump      = "Dump"      //序列化给定 key ，并返回被序列化的值
	RedisCmdTTL       = "TTL"       //以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)
	RedisCmdExpire    = "Expire"    //seconds 为给定 key 设置过期时间
	RedisCmdDel       = "DEL"       //该命令用于在 key 存在是删除 key
	RedisCmdPttl      = "Pttl"      //以毫秒为单位返回 key 的剩余的过期时间
	RedisCmdRenameNx  = "Renamenx"  //仅当 newkey 不存在时，将 key 改名为 newkey
	RedisCmdExists    = "EXISTS"    //检查给定 key 是否存在
	RedisCmdExpireAt  = "Expireat"  //EXPIREAT 的作用和 EXPIRE 类似，都用于为 key 设置过期时间 不同在于 EXPIREAT 命令接受的时间参数是 UNIX 时间戳(unix timestamp)
	RedisCmdKeys      = "Keys"      //查找所有符合给定模式( pattern)的 key

	//字符串(String) 命令
	RedisCmdSetNx       = "Setnx"       //只有在 key 不存在时设置 key 的值
	RedisCmdGetRange    = "Getrange"    //返回 key 中字符串值的子字符
	RedisCmdMset        = "Mset"        //同时设置一个或多个 key-value 对
	RedisCmdSetEx       = "Setex"       //将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)
	RedisCmdSet         = "SET"         //设置指定 key 的值
	RedisCmdGet         = "Get"         //获取指定 key 的值
	RedisCmdGetBit      = "Getbit"      //对 key 所储存的字符串值，获取指定偏移量上的位(bit)
	RedisCmdSetBit      = "Setbit"      //对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)
	RedisCmdDecr        = "Decr"        //将 key 中储存的数字值减一
	RedisCmdDecrBy      = "Decrby"      //key 所储存的值减去给定的减量值（decrement）
	RedisCmdStrLen      = "Strlen"      //返回 key 所储存的字符串值的长度
	RedisCmdMsetNx      = "Msetnx"      //同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在
	RedisCmdIncrBy      = "Incrby"      //将 key 所储存的值加上给定的增量值（increment）
	RedisCmdIncrByFloat = "Incrbyfloat" //将 key 所储存的值加上给定的浮点增量值（increment）
	RedisCmdSetRange    = "Setrange"    //用 value 参数覆写给定 key 所储存的字符串值，从偏移量 offset 开始
	RedisCmdPsetEx      = "Psetex"      //这个命令和 SETEX 命令相似，但它以毫秒为单位设置 key 的生存时间，而不是像 SETEX 命令那样，以秒为单位
	RedisCmdAppend      = "Append"      //如果 key 已经存在并且是一个字符串， APPEND 命令将 value 追加到 key 原来的值的末尾
	RedisCmdGetSet      = "Getset"      //将给定 key 的值设为 value ，并返回 key 的旧值(old value)
	RedisCmdMget        = "Mget"        //获取所有(一个或多个)给定 key 的值
	RedisCmdIncr        = "Incr"        //

	//列表(List) 命令
	RedisCmdLindex     = "Lindex"     //通过索引获取列表中的元素
	RedisCmdRpush      = "Rpush"      //在列表中添加一个或多个值
	RedisCmdLrange     = "Lrange"     //获取列表指定范围内的元素
	RedisCmdRpoplpush  = "Rpoplpush"  //移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	RedisCmdBlpop      = "Blpop"      //移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	RedisCmdBrpop      = "Brpop"      //移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	RedisCmdBrpoplpush = "Brpoplpush" //从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	RedisCmdLrem       = "Lrem"       //移除列表元素
	RedisCmdLlen       = "Llen"       //获取列表长度
	RedisCmdLtrim      = "Ltrim"      //对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除
	RedisCmdLpop       = "Lpop"       //移出并获取列表的第一个元素
	RedisCmdLpushx     = "Lpushx"     //将一个或多个值插入到已存在的列表头部
	RedisCmdLinsert    = "Linsert"    //在列表的元素前或者后插入元素
	RedisCmdRpop       = "Rpop"       //移除并获取列表最后一个元素
	RedisCmdLset       = "Lset"       //通过索引设置列表元素的值
	RedisCmdLpush      = "Lpush"      //将一个或多个值插入到列表头部
	RedisCmdRpushx     = "Rpushx"     //为已存在的列表添加值

	//集合(Set) 命令
	RedisCmdSunion      = "Sunion"      //返回所有给定集合的并集
	RedisCmdScard       = "Scard"       //获取集合的成员数
	RedisCmdSrandMember = "Srandmember" //返回集合中一个或多个随机数
	RedisCmdSmembers    = "Smembers"    //返回集合中的所有成员
	RedisCmdSinter      = "Sinter"      //返回给定所有集合的交集
	RedisCmdSrem        = "Srem"        //移除集合中一个或多个成员
	RedisCmdSmove       = "Smove"       //将 member 元素从 source 集合移动到 destination 集合
	RedisCmdSadd        = "Sadd"        //向集合添加一个或多个成员
	RedisCmdSisMember   = "Sismember"   //判断 member 元素是否是集合 key 的成员
	RedisCmdSdiffStore  = "Sdiffstore"  //返回给定所有集合的差集并存储在 destination 中
	RedisCmdSdiff       = "Sdiff"       //返回给定所有集合的差集
	RedisCmdSscan       = "Sscan"       //迭代集合中的元素
	RedisCmdSinterStore = "Sinterstore" //返回给定所有集合的交集并存储在 destination 中
	RedisCmdSunionStore = "Sunionstore" //所有给定集合的并集存储在 destination 集合中
	RedisCmdSpop        = "Spop"        //移除并返回集合中的一个随机元素

	//有序集合(sorted set) 命令
	RedisCmdZrevRank         = "Zrevrank"         //返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
	RedisCmdZlexCount        = "Zlexcount"        //在有序集合中计算指定字典区间内成员数量
	RedisCmdZunionStore      = "Zunionstore"      //计算给定的一个或多个有序集的并集，并存储在新的 key 中
	RedisCmdZremRangeByRank  = "Zremrangebyrank"  //移除有序集合中给定的排名区间的所有成员
	RedisCmdZcard            = "Zcard"            //获取有序集合的成员数
	RedisCmdZrem             = "Zrem"             //移除有序集合中的一个或多个成员
	RedisCmdZinterStore      = "Zinterstore"      //计算给定的一个或多个有序集的交集并将结果集存储在新的有序集合 key 中
	RedisCmdZrank            = "Zrank"            //返回有序集合中指定成员的索引
	RedisCmdZincrBy          = "Zincrby"          //有序集合中对指定成员的分数加上增量 increment
	RedisCmdZrangeByScore    = "Zrangebyscore"    //通过分数返回有序集合指定区间内的成员
	RedisCmdZrangeByLex      = "Zrangebylex"      //通过字典区间返回有序集合的成员
	RedisCmdZscore           = "Zscore"           //返回有序集中，成员的分数值
	RedisCmdZremRangeByScore = "Zremrangebyscore" //移除有序集合中给定的分数区间的所有成员
	RedisCmdZscan            = "Zscan"            //迭代有序集合中的元素（包括元素成员和元素分值）
	RedisCmdZrevRangeByScore = "Zrevrangebyscore" //返回有序集中指定分数区间内的成员，分数从高到低排序
	RedisCmdZremRangeByLex   = "Zremrangebylex"   //移除有序集合中给定的字典区间的所有成员
	RedisCmdZrevRange        = "Zrevrange"        //返回有序集中指定区间内的成员，通过索引，分数从高到底
	RedisCmdZrange           = "Zrange"           //通过索引区间返回有序集合成指定区间内的成员
	RedisCmdZcount           = "Zcount"           //计算在有序集合中指定区间分数的成员数
	RedisCmdZadd             = "Zadd"             //向有序集合添加一个或多个成员，或者更新已存在成员的分数

	//连接 命令
	RedisCmdEcho   = "Echo"   //打印字符串
	RedisCmdSelect = "Select" //切换到指定的数据库
	RedisCmdPing   = "Ping"   //查看服务是否运行
	RedisCmdQuit   = "Quit"   //关闭当前连接
	RedisCmdAuth   = "Auth"   //验证密码是否正确

	//服务器 命令
	RedisCmdClientPause  = "Client Pause" //在指定时间内终止运行来自客户端的命令
	RedisCmdDebugObject  = "Debug Object" //获取 key 的调试信息
	RedisCmdFlushdb      = "Flushdb"      //删除当前数据库的所有key
	RedisCmdSave         = "Save"         //异步保存数据到硬盘
	RedisCmdShow         = "Show"         //管理 RedisCmd的慢日志
	RedisCmdLastSave     = "Lastsave"     //返回最近一次 RedisCmd成功将数据保存到磁盘上的时间，以 UNIX 时间戳格式表示
	RedisCmdConfig       = "Config"       //获取指定配置参数的值
	RedisCmdCommand      = "Command"      //获取 RedisCmd命令详情数组
	RedisCmdSlaveOf      = "Slaveof"      //将当前服务器转变为指定服务器的从属服务器(slave server)
	RedisCmdDebug        = "Debug"        //让 RedisCmd服务崩溃
	RedisCmdFlushAll     = "Flushall"     //删除所有数据库的所有key
	RedisCmdDbsize       = "Dbsize"       //返回当前数据库的 key 的数量
	RedisCmdBgRewriteAOF = "Bgrewriteaof" //异步执行一个 AOF（AppendOnly File） 文件重写操作
	RedisCmdCluster      = "Cluster"      //获取集群节点的映射数组
	RedisCmdShutdown     = "Shutdown"     //异步保存数据到硬盘，并关闭服务器
	RedisCmdSync         = "Sync"         //用于复制功能(replication)的内部命令
	RedisCmdClient       = "Client"       //客户端
	RedisCmdRole         = "Role"         //返回主从实例所属的角色
	RedisCmdMonitor      = "Monitor"      //实时打印出 RedisCmd服务器接收到的命令，调试用
	RedisCmdTime         = "Time"         //返回当前服务器时间
	RedisCmdInfo         = "Info"         //获取 RedisCmd服务器的各种信息和统计数值
	RedisCmdBgsave       = "Bgsave"       //在后台异步保存当前数据库的数据到磁盘

	//脚本 命令
	RedisCmdScript  = "Script"  //运行Lua脚本
	RedisCmdEval    = "Eval"    //执行 Lua 脚本
	RedisCmdEvalsha = "Evalsha" //执行 Lua 脚本

	//事务 命令
	RedisCmdExec    = "Exec"    //执行所有事务块内的命令
	RedisCmdWatch   = "Watch"   //监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断
	RedisCmdDiscard = "Discard" //取消事务，放弃执行事务块内的所有命令
	RedisCmdUnwatch = "Unwatch" //取消 WATCH 命令对所有 key 的监视
	RedisCmdMulti   = "Multi"   //标记一个事务块的开始

	//HyperLogLog 命令
	RedisCmdPgmerge = "Pgmerge" //将多个 HyperLogLog 合并为一个 HyperLogLog
	RedisCmdPfadd   = "Pfadd"   //添加指定元素到 HyperLogLog 中
	RedisCmdPfcount = "Pfcount" //返回给定 HyperLogLog 的基数估算值

	//发布订阅 命令
	RedisCmdUnsubscribe  = "Unsubscribe"  //指退订给定的频道
	RedisCmdSubscribe    = "Subscribe"    //订阅给定的一个或多个频道的信息
	RedisCmdPubsub       = "Pubsub"       //查看订阅与发布系统状态
	RedisCmdPunsubscribe = "Punsubscribe" //退订所有给定模式的频道
	RedisCmdPublish      = "Publish"      //将信息发送到指定的频道
	RedisCmdPsubscribe   = "Psubscribe"   //订阅一个或多个符合给定模式的频道

	//地理位置(geo) 命令
	RedisCmdGEOHASH           = "GEOHASH"           //返回一个或多个位置元素的 Geohash 表示
	RedisCmdGEOPOS            = "GEOPOS"            //从key里返回所有给定位置元素的位置（经度和纬度）
	RedisCmdGEODIST           = "GEODIST"           //返回两个给定位置之间的距离
	RedisCmdGEORADIUS         = "GEORADIUS"         //以给定的经纬度为中心， 找出某一半径内的元素
	RedisCmdGEOADD            = "GEOADD"            //将指定的地理空间位置（纬度、经度、名称）添加到指定的key中
	RedisCmdGEORADIUSBYMEMBER = "GEORADIUSBYMEMBER" //找出位于指定范围内的元素，中心点是由给定的位置元素决定
)
