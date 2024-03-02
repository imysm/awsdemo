package handler

import (
	"awsdemo/pkg/connector"
	"awsdemo/pkg/secret"
	"fmt"
	"github.com/gin-gonic/gin"
	goredis "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

func Cache(c *gin.Context) {
	r := secret.GetRedisOption(viper.GetString("secretRedisName"), viper.GetString("region"))

	optsRedis := &connector.RedisOptions{Addr: r.Host + r.Port, Username: r.Username, Password: r.Password}

	//log.Println(optsRedis)
	rdb := connector.NewRedis(optsRedis)
	key := c.Query("name")
	rdb.Set(c, key, c.Query("message"), 0)

	val, err := rdb.Get(c, key).Result()
	if err == goredis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		fmt.Printf("err : %s", err.Error())
	} else {
		SendResponse(c, nil, fmt.Sprintf("%s  %s = %s \n", time.Now().Format("2006-01-02 15:04:05"), key, val))
	}
	//stats := rdb.PoolStats()
	//SendResponse(c, nil,
	//	fmt.Sprintf("Hits=%d Misses=%d Timeouts=%d TotalConns=%d IdleConns=%d StaleConns=%d\n",
	//		stats.Hits, stats.Misses, stats.Timeouts, stats.TotalConns, stats.IdleConns, stats.StaleConns))
}
