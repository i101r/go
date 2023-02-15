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

type strg interface{
	Connect() bool
	Get(key string) (value []byte, flags int, err error)
	Set(key string, value []byte, flags int, exptime int64) (err error)
	Delete(key string) (err error) 
}

var st strg = &storage.Memcache{}
// var st strg = &storage.Cache{}

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		fmt.Println("Listening :5300")
		return 
	}

	fmt.Println("Listening :5300")

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterStorageServer(grpcServer, &server{})

	if !st.Connect(){
		st = &storage.Cache{}
	}

	grpcServer.Serve(listener)

}

func (s *server) Set (c context.Context, request *pb.SetRequest) ( response *pb.Response,err error){
	
	st.Set(request.Name, []byte(request.Value), 0, 86400)
	
	output := "Ok"

	response = &pb.Response{
		Message: output,
	}
	return response, nil   
}

func (s *server) Get (c context.Context, request *pb.GetRequest) ( response *pb.Response,err error){

		output, _, _:=st.Get(request.Name)

		response = &pb.Response{
			Message: string(output),
		}

	return response, nil   
}

func (s *server) Delete (c context.Context, request *pb.GetRequest) ( response *pb.Response,err error){
	
	st.Delete(request.Name)

	output := "Delete"

	response = &pb.Response{
		Message: output,
	}

	return response, nil   
}


