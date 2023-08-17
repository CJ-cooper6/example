package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		//当你使用 Docker Compose 启动多个容器时，这些容器默认会处于同一个网络中。
		//因此它们可以通过服务名称来相互访问。这使得内部的服务可以通过它们的服务名称进行通信，而不需要使用具体的 IP 地址或端口。
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result() // Check if Redis is accessible
	if err != nil {
		fmt.Println(11111)
		panic(err)
	}
	r := gin.Default()
	r.GET("/", index)
	r.Run(":8080")

}

func index(c *gin.Context) {
	ctx := context.Background()
	val, err := rdb.Incr(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	response := fmt.Sprintf("页面总访问次数: %d", val)
	c.String(http.StatusOK, response)
}
