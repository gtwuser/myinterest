package main

import (
	"context"
	"fmt"
	"github.com/kapjoshi/myinterest/qrcode/cmd/grpcapps/pob"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type greetServer struct{}

func (*greetServer) Greet(ctx context.Context, req *pob.GreetRequest) (*pob.GreetResponse, error) {
	return &pob.GreetResponse{Greet: `Hi ` + req.GetFname()}, nil
}

func (*greetServer) Recipes(ctx context.Context, req *pob.RecipeRequest) (*pob.RecipeResponse, error) {
	return &pob.RecipeResponse{Taste: req.GetName() + " Mast bana hai !!"}, nil
}

func main() {
	fmt.Println(`Serving now .... `)
	l, err := net.Listen(`tcp`, `:8081`)
	if err != nil {
		log.Fatal("creating listener failed ", err)
	}

	s := grpc.NewServer()

	//srv := &greetServer{}
	//pob.RegisterGreetServiceServer(s, srv)
	if err := s.Serve(l); err != nil {
		log.Fatal(`Issue in serving greet service`)
	}
}
