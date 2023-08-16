package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hello_grpc "grpc_class/pb"
	"log"
	"time"
)

func main(){
	// 连接到server端，此处禁用安全传输
	coon,err := grpc.Dial("localhost:8972",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer coon.Close()

	client := hello_grpc.NewHelloGRPCClient(coon)	//new 一个client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)	//创建一个有过期时间的context，返回取消函数cancel()
	defer cancel()
	req,err := client.SayHi(ctx,&hello_grpc.Req{Message: "我从客户端来"})	//调用方法
	fmt.Println(req.GetMessage())	//获取返回值

}
