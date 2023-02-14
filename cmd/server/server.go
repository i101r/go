package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/i101r/go/grpc"
)

type server struct{
	pb.UnimplementedStorageServer
}

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

func (s *server) Set (c context.Context, request *pb.Uid) ( response *pb.Response,err error){

	output := "Set response"

	response = &pb.Response{
		Message: output,
	}
	return response, nil   
}

func (s *server) Get (c context.Context, request *pb.Uid) ( response *pb.Response,err error){

		output := "Get response  final"

		response = &pb.Response{
			Message: output,
		}
	return response, nil   
}


func (s *server) Delete (c context.Context, request *pb.Uid) ( response *pb.Response,err error){

	output := "Delete response"

	response = &pb.Response{
		Message: output,
	}

	return response, nil   
}


