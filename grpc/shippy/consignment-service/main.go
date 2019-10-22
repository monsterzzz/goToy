package main

import (
	"context"
	proto "goToy/grpc/shippy/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 仓库接口
type IRepository interface {
	Create(consignment *proto.Consignment) (*proto.Consignment, error)
	GetAll() []*proto.Consignment
}

type Repository struct {
	Consignments []*proto.Consignment
}

// 实现
func (repo *Repository) Create(consignment *proto.Consignment) (*proto.Consignment, error) {
	repo.Consignments = append(repo.Consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetALL() []*proto.Consignment {
	return repo.Consignments
}

// 服务类
type Service struct {
	repo Repository
}

// 对于proto声明的接口的实现
func (s *Service) CreateConsignment(ctx context.Context, req *proto.Consignment) (*proto.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		log.Fatal("req error!")
	}
	resp := &proto.Response{
		Created:     true,
		Consignment: consignment,
	}
	return resp, nil
}

func (s *Service) GetAllConsignments(context.Context, *proto.GetAllReq) (*proto.Response, error) {
	consignments := s.repo.GetALL()
	resp := &proto.Response{Consignments: consignments}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal("port error")
	}
	log.Println("listening on port:", 5001)

	server := grpc.NewServer()
	repo := Repository{}
	// grpc 注册服务
	proto.RegisterShippingServiceServer(server, &Service{repo: repo})

	// 运行服务
	if err := server.Serve(listener); err != nil {
		log.Fatal("failed server")
	}
}
