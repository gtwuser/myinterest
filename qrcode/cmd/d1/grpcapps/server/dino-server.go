package main

import (
	"fmt"
)

type server struct {
	firstName string
}

func main() {
	var sa []*server
	s := &server{firstName: `Amit`}
	k := &server{firstName: `Kishore`}
	sa = append(sa, s)
	sa = append(sa, k)

	for _, v := range sa {
		fmt.Println(*v)
	}

}

//func (s *server) Greeting(ctx context.Context, grt *dinopb.Greet) (*dinopb.GreetResponse, error) {
//	rs := `Hello ` + grt.GetFirstname()
//	grs := &dinopb.GreetResponse{
//		Greeting: rs,
//	}
//	return grs, nil
//}
