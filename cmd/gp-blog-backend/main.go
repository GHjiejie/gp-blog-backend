package main

import (
	"gp-blog-backend/server"

	"google.golang.org/grpc"
)

func main() {
	// 创建grpc服务器
	grpcServer := grpc.NewServer()
	// 注册UserService服务
	pb.RegisterUserServiceServer(grpcServer, &server.UserServiceServer{})

}
