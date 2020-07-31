package main

import (
	"context"
	"fmt"
	"github.com/kapjoshi/myinterest/qrcode/cmd/grpcapps/pob"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(`:8081`, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Issue in connecting to grpc server", err)
	}

	cc := pob.NewGreetServiceClient(conn)

	resp, err := cc.Greet(context.Background(), &pob.GreetRequest{
		Fname: "Kapil",
		Lname: "Joshi",
	})

	if err != nil {
		log.Fatal(`Error in invoking server code`, err)
	}
	fmt.Println(resp.GetGreet())

	recipeResponse, err := cc.Recipes(context.Background(), &pob.RecipeRequest{
		Name:           "chenna puda",
		DurationInHour: 2,
	})
	if err != nil {
		log.Fatal(`Error in invoking server code`, err)
	}

	fmt.Println("Recipe response", recipeResponse)
}
