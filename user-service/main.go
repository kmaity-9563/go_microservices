package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/kmaity-9563/shopping-app/proto/userpb"
	"google.golang.org/grpc"

	"github.com/hashicorp/consul/api"
)

type server struct {
	pb.UnimplementedUserServiceServer
	mu sync.Mutex
	users map[string]*pb.User
}

func (s *server) createUser(ctx context.Context ,req *pb.createUser) (res *pb.CreateUserRes, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := fmt.Sprintf("%d", len(s.users)+1)
	req.User.id = id
	s.users[id] = req.User

	return &pb.CreateUserRes{UserId: id}, nil
}

func (s *server) getUser(ctx context.Context, req *pb.GetUserReq ) (res *pb.GetUserRes, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.UserId]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return &pb.GetUserRes{User: user}, nil
}

func (s *server) updateUser(ctx context.Context, req *pb.updateUserReq ) (res *pb.UpdateUserRes, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_ , exists := s.users[req.user.id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	s.users[req.user.id] = req.user

	return &pb.UpdateUserRes{User: req.user}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserRes, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	delete(s.users, req.Id)

	return &pb.DeleteUserRes{Id: req.Id}, nil
}

func main() {
	lis , err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal("failed to create consul client %v", err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = "shopping-app"
	registration.Name = "shopping-app"
	registration.Address = "localhost"
	registration.Port = 8080
	registration.Check = &api.AgentServiceCheck{
		CheckID: "shopping-app-check",
		Name:    "shopping-app-check",
		Notes:   "shopping-app health check",
		Status:  api.HealthPassing,
		GRPC:     "localhost:8080",
		Interval: "10s",
	}

	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("failed to register service %v", err)
	}
}