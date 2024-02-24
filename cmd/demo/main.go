package main

import (
	"awsdemo/internal/demo/router"
	mw "awsdemo/internal/demo/router/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {

	viper.SetDefault("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJjb3JwIjoiTWljdnMiLCJpYXQiOjE3MDg2N")
	viper.SetDefault("addr", ":80")
	viper.SetDefault("region", "ap-east-1")
	viper.SetDefault("secretMySQLName", "dev/demo/mysql")
	viper.SetDefault("secretRedisName", "dev/demo/redis")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not exits，using default config")
		} else {
			fmt.Printf("read io err: %v\n", err)
		}
	}

	g := gin.New()
	router.Load(
		g,
		// Middlewares...
		mw.DefaultLogger(),
		mw.Trace(),
		mw.RequestId(),
	)

	log.Printf("start to listening the incoming requests on %s", viper.GetString("addr"))
	log.Println(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
