package main

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/i101r/go/grpc"
	storage "github.com/i101r/go/storage"
)

type server struct{
	pb.UnimplementedStorageServer
}

var memc = &storage.Memcache{}

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		fmt.Printf("\nListening :5300")
		return 
	}

	fmt.Printf("\nListening :5300")

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterStorageServer(grpcServer, &server{})

	memc.Connect()

	grpcServer.Serve(listener)

}

func (s *server) Set (c context.Context, request *pb.SetRequest) ( response *pb.Response,err error){
	
	memc.Set(request.Name, []byte(request.Value), 0, 86400)
	
	output := "Ok"

	response = &pb.Response{
		Message: output,
	}
	return response, nil   
}

func (s *server) Get (c context.Context, request *pb.GetRequest) ( response *pb.Response,err error){

		output, _, _:=memc.Get(request.Name)

		response = &pb.Response{
			Message: string(output),
		}

	return response, nil   
}


func (s *server) Delete (c context.Context, request *pb.GetRequest) ( response *pb.Response,err error){
	
	memc.Delete(request.Name)

	output := "Delete"

	response = &pb.Response{
		Message: output,
	}

	return response, nil   
}


