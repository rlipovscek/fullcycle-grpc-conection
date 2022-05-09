package services

import (
	"context"
	"fmt"

	"github.com/rlipovscek/fullcycle-grpc-conection/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.Name,
		Email: req.Email,
	}, nil
}
