package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "grpc_class/pb"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}


func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res,err error){
	fmt.Println(req.GetMessage())
return &hello_grpc.Res{Message: "我是从服务端返回的grpc内容"},nil
}

func main(){
	l,err := net.Listen("tcp",":8972")
	if err != nil{
		fmt.Printf("failed to listen:%v",err)
		return
	}
	s := grpc.NewServer()	//创建gRPC服务器
	hello_grpc.RegisterHelloGRPCServer(s,&server{})		//在gRpc服务端注册服务

	err = s.Serve(l)	//启动服务
	if err != nil{
		fmt.Printf("failed to serve:%v",err)
		return
	}

}

