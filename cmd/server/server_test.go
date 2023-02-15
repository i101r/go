 package main

 import(
	"testing"
	pb "github.com/i101r/go/grpc"
	context "context"
 )

 var setRequest = &pb.SetRequest{}
 var getRequest = &pb.GetRequest{}
 var srv = &server{}

 func TestSet(t *testing.T){
	st.Connect()
	
	setRequest.Name ="key1"
	setRequest.Value ="incididunt"

	ctx := context.Background()

	res, err := srv.Set(ctx, setRequest)
	
	if err != nil  {
        t.Errorf("\nTestSet %s", getRequest.Name)
    }


 }

 func TestGet(t *testing.T){
	
	getRequest.Name ="key1"

	ctx := context.Background()

	res, err := srv.Get(ctx, getRequest)
	
	if err != nil &&  res.Message!="incididunt" {
        t.Errorf("\nTestGet %s", getRequest.Name)
    }


 }

 func TestDelete(t *testing.T){
	
	getRequest.Name ="key1"

	ctx := context.Background()

	res, err := srv.Get(ctx, getRequest)
	
	if err != nil {
        t.Errorf("\nTestDelete %s", getRequest.Name)
    }


 }