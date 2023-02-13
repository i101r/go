package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	pb "github.com/i101r/go/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		fmt.Printf("Listening failed!")
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterStorageServer(grpcServer, &server{})

	grpcServer.Serve(listener)

	fmt.Printf("hello")

}

func Set() {

}

func (s *server) Get (c context.Context, request *pb.Request)	(c context.Context, request *pb.Request){

		output := "Get response"

		response = &pb.Response{
			Message: output,
		}
	return response, nil   
}

func Delete() {

}
