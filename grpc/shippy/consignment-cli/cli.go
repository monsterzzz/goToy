package main

import (
	"context"
	"encoding/json"
	"errors"
	proto "goToy/grpc/shippy/consignment-service/proto/consignment"
	"google.golang.org/grpc"
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
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal("conn err")
	}
	defer conn.Close()

	client := proto.NewShippingServiceClient(conn)

	info := "grpc/shippy/consignment-cli/consignment.json"
	consignment, err := ParseFile(info)
	if err != nil {
		log.Fatal("111")
	}
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatal("error req")
	}

	log.Println(resp.Created)

	resp, err = client.GetAllConsignments(context.Background(), &proto.GetAllReq{})
	if err != nil {
		log.Fatal("error req")
	}
	log.Println(resp.Consignments)
	log.Println(resp)
}
