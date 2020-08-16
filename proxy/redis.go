package proxy

import (
	"SpiderTool/conf"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

//type RedisClient struct {
//
//}

func RedisClient() (cli *redis.Client, con *conf.Config) {
	var c conf.Config
	config := c.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisConf.Address,
		Password: config.RedisConf.Password,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Printf("redis连接失败：%s", err)
		return
	}
	fmt.Printf("redis连接成功：%s", pong)

	return client, config
}

func AddProxy(client *redis.Client, config *conf.Config, proxy string) {
	score, err := strconv.ParseFloat(config.RedisConf.InitialScore, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	//if client.Get(config.RedisConf.Key) == nil {
	//	client.set
	//}
	//if client.ZScore(config.RedisConf.Key, proxy) == nil {
		client.ZAdd(config.RedisConf.Key, redis.Z{Score: score, Member: proxy})
	//}
	//client.ZAdd()
}
