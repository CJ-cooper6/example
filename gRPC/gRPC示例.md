# gRPC示例



### 编写proto代码

关于 `Protocol Buffers` 的教程可以查看 [Protocol Buffers V3中文指南](https://www.liwenzhou.com/posts/Go/Protobuf3-language-guide-zh/)

创建`hello_grpc.proto`  文件

```
syntax = "proto3";  // 版本声明

option go_package = "./;hello_grpc";
package hello_grpc;

//请求消息
message Req{
  string message = 1;
}

//响应消息
message Res{
  string message = 1;
}

//定义服务
service HelloGRPC{
  rpc SayHi(Req)returns(Res);
}


```



在项目根目录下执行以下命令，根据`hello.proto`生成 go 源码文件。

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/hello.proto		//最后这个生成位置可以修改
```



### 编写Server端Go代码

1. **取出server**

   ```
   type server struct {
   	hello_grpc.UnimplementedHelloGRPCServer
   }
   ```

   

2. **挂载方法**

   ```go
   func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res,err error){
      fmt.Println(req.GetMessage())
   return &hello_grpc.Res{Message: "我是从服务端返回的grpc内容"},nil
   }
   ```

3. **注册服务**

   ```go
   l,err := net.Listen("tcp",":8972")
   if err != nil{
      fmt.Printf("failed to listen:%v",err)
      return
   }
   s := grpc.NewServer()  //创建gRPC服务器
   hello_grpc.RegisterHelloGRPCServer(s,&server{})   //在gRpc服务端注册服务
   ```

   

4. **创建监听**

   ```go
   err = s.Serve(l)    //启动服务
   if err != nil{
      fmt.Printf("failed to serve:%v",err)
      return
   }
   ```

   最后编译并执行 



### 编写Client端Go代码

```go
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
```

编译并执行，得到以下输出结果，说明RPC调用成功。

```go
//客户端输出：
我是从服务端返回的grpc内容
```

```go
//服务端输出:
我从客户端来
```

