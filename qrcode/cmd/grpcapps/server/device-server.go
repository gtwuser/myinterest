package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kapjoshi/myinterest/qrcode/cmd/grpcapps/pob"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type deviceServer struct{}

func (*deviceServer) ManagedDevice(ctx context.Context, req *pob.DeviceRequest) (*pob.DeviceResponse, error) {
	fmt.Printf("Device info requested %v", req)
	d := &pob.DeviceDetails{
		Name: "MyDevice",
		Type: "Simple",
	}
	return &pob.DeviceResponse{DevDetails: d}, nil
}

func main() {
	//l, err := net.Listen("tcp", ":8081")
	//if err != nil {
	//	log.Fatal("Error in creating listener", err)
	//}

	//s := grpc.NewServer()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pob.RegisterManageDeviceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		log.Fatal("Error in registering service to server", err)
	}
	if err = http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal("Error in starting server", err)
	}
}
