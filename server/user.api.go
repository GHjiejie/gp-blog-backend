package server

import (
	"context"
	pb "gp-blog-backend/repository/userpb"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer // 确保向后兼容
}

func (s *UserServiceServer) RegisterUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.RegisterUserResponse, error) {

}
