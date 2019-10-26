package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	proto "goToy/grpc/shippy2/consignment-service/proto/consignment"
	"io/ioutil"
	"log"
)

func ParseFile(filePath string) (*proto.Consignment, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {

		log.Fatal("file error")
	}
	var consignment *proto.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("parse file error")
	}
	return consignment, nil
}

func main() {
	//e := cmd.Init()
	//fmt.Println(e)
	service := micro.NewService(micro.Name("monster"))
	service.Init()

	client := proto.NewShippingServiceClient("monster", service.Client())
	file := "grpc/shippy2/consignment-cli/consignment.json"
	consignment, err := ParseFile(file)
	if err != nil {
		log.Fatal("error file")
	}
	resp, e := client.CreateConsignment(context.Background(), consignment)
	fmt.Println(e)
	log.Println(resp)
	resp, e = client.GetAllConsignments(context.Background(), &proto.GetAllReq{})
	fmt.Println(e)
	log.Println(resp)
}
