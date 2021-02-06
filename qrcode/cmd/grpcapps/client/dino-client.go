package main

//
//import (
//	"fmt"
//	"github.com/gtwuser/myinterest/qrcode/cmd/d1/grpcapps/pb"
//	"google.golang.org/grpc"
//)
//
//
//func main(){
//	lis, _ := grpc.Dial("localhost:50051", grpc.WithInsecure)
//	defer lis.Close()
//
//	dc := dinopb.NewDinoServiceClient(lis)
//	in := &dinopb.Greet{
//		Firstname: "Gatotkach",
//		Lastname: "Bheem"
//	}
//	res, _ := dc.Greeting(context.Background(), in)
//	fmt.Println("Result", res.Greeting)
//
//}
