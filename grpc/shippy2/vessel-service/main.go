package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"log"
)
import "goToy/grpc/shippy2/vessel-service/proto/vessel"

type Repository interface {
	FindAvailable(in *vessel.Specification) (*vessel.Vessel, error)
}

type VesselRepo struct {
	vessels []*vessel.Vessel
}

func (v VesselRepo) FindAvailable(in *vessel.Specification) (*vessel.Vessel, error) {
	for _, item := range v.vessels {
		if in.Capacity <= item.Capacity && in.MaxHeight <= item.Capacity {
			return item, nil
		}
	}
	return nil, errors.New("vessel not found!")
}

type Service struct {
	repo Repository
}

func (s *Service) FindAvailable(ctx context.Context, in *vessel.Specification, out *vessel.Response) error {
	if result, err := s.repo.FindAvailable(in); err != nil {
		return err
	} else {
		out.Vessel = result
		return nil
	}
}

func main() {
	vessels := []*vessel.Vessel{{
		Id:        "Vessels01",
		Capacity:  500,
		MaxWeight: 200000,
		Name:      "monsterVessels",
		Available: true,
		OwnerId:   "11",
	}}

	repo := VesselRepo{vessels: vessels}
	service := micro.NewService(
		micro.Name("vessel"),
		micro.Version("latest"),
	)
	service.Init()

	vessel.RegisterVesselServiceHandler(service.Server(), &Service{repo: repo})

	if err := service.Run(); err != nil {
		log.Fatalf("run error : %v", err)
	}
}
