package main

import (
	pb "evenkey/mygrpc" // 导入生成的 protobuf 代码
	routeChat "evenkey/route"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	c := &routeChat.MyServiceServer{}

	// 定义命令行参数
	flag.IntVar(&c.Port, "port", 50051, "gRPC server port")
	flag.StringVar(&c.Addr, "addr", "localhost", "gRPC server address")
	flag.StringVar(&c.Username, "username", "root", "username")
	flag.StringVar(&c.Password, "password", "root", "password")
	flag.IntVar(&c.Keydelay, "delay", 10, "key press delay")

	// 解析命令行参数
	flag.Parse()

	// 打印解析出来的参数值
	fmt.Printf("port=%d, addr=%s, Username=%s, Password=%s\n",
		c.Port, c.Addr, c.Username, c.Password)

	// 创建 gRPC 服务器
	lis, err := net.Listen("tcp", ":50051") // 指定服务器监听的端口
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterMyServiceServer(s, c)

	// 启动服务器
	log.Println("Starting gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
