package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "goToy/grpc/shippy2/consignment-service/proto/consignment"
	"goToy/grpc/shippy2/vessel-service/proto/vessel"

	"log"
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
	repo         Repository
	vesselClient vessel.VesselServiceClient
}

// 对于proto声明的接口的实现
func (s *Service) CreateConsignment(ctx context.Context, req *proto.Consignment, resp *proto.Response) error {

	s.vesselClient.FindAvailable()

	consignment, err := s.repo.Create(req)
	if err != nil {
		log.Fatal("req error!")
	}
	fmt.Println(s.repo)
	resp.Created = true
	resp.Consignment = consignment
	//resp = &proto.Response{Created: true, Consignment: consignment}
	return nil
}

func (s *Service) GetAllConsignments(ctx context.Context, req *proto.GetAllReq, resp *proto.Response) error {
	fmt.Println("!!!")
	consignments := s.repo.GetALL()
	*resp = proto.Response{Consignments: consignments}
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("monster"),
		micro.Version("latest"),
	)

	server.Init()
	proto.RegisterShippingServiceHandler(server.Server(), &Service{repo: Repository{}})

	if err := server.Run(); err != nil {
		log.Fatal("new server error")
	}

}
