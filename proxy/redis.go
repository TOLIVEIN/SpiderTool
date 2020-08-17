package proxy

import (
	"SpiderTool/conf"
	"fmt"
	"github.com/go-redis/redis"
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
	score := config.RedisConf.InitialScore
	client.ZAdd(config.RedisConf.Key, redis.Z{Score: score, Member: proxy})

	//Decrease(client, config, proxy)
	//}
	//client.ZAdd()
}

func Decrease(client *redis.Client, config *conf.Config, proxy string) {
	score := client.ZScore(config.RedisConf.Key, proxy).Val()
	fmt.Println(score)
	if score > config.RedisConf.MinScore {
		fmt.Printf("代理：%s，现在的分数：%s", proxy, score)
		client.ZIncrBy(config.RedisConf.Key, -1, proxy)
	} else {
		fmt.Printf("代理：%s，现在的分数：%s，删除。。。", proxy, score)
		client.ZRem(config.RedisConf.Key, proxy)
	}
}

func Max(client *redis.Client, config *conf.Config, proxy string) {
	fmt.Printf("代理：%s可用，设置分数为：%s", proxy, config.RedisConf.MaxScore)
	client.ZAdd(config.RedisConf.Key, redis.Z{Score: config.RedisConf.MaxScore, Member: proxy})
}


func All(client *redis.Client, config *conf.Config) []string {
	min := fmt.Sprintf("%g", config.RedisConf.MinScore)
	max := fmt.Sprintf("%g", config.RedisConf.MaxScore)

	//return client.ZRange(config.RedisConf.Key, 0, 100).Val()

	return client.ZRangeByScore(config.RedisConf.Key, redis.ZRangeBy{Min: min, Max: max }).Val()
}