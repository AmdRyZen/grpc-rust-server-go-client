package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	pb "mix-grpc/protos"

	"github.com/mix-go/dotenv"
	"google.golang.org/grpc"
)

type GrpcClientCommand struct {
}

type Page struct {
	Total       int          `json:"total"`
	Pages       int          `json:"pages"`
	PageNo      int          `json:"page_no"`
	PageSize    int          `json:"page_size"`
	SearchCount int          `json:"search_count"`
	Records     []MemberInfo `json:"records"`
}
type MemberInfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

//  go build -o bin/go_build_main_go.exe main.go
// bin/go_build_main_go grpc:client
func (t *GrpcClientCommand) Main() {
	addr := dotenv.Getenv("GIN_ADDR").String("localhost:50051")
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("----------------")
		fmt.Println(err)
	}
	fmt.Println("-------------------------------")
	fmt.Println(conn.GetState())

	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewGreeterClient(conn)
	fmt.Println(cli)
	req := pb.HelloRequest{
		Name: "testjm001",
	}
	resp, err := cli.SayHello(ctx, &req)
	if err != nil {
		fmt.Println("111111111111111111111111111111")
		fmt.Println(err)
	}
	/*fmt.Println(resp)*/

	var jsonValue = Page{}
	json.Unmarshal([]byte(resp.Message), &jsonValue)

	for i := range jsonValue.Records {
		fmt.Println(jsonValue.Records[i].Name)
	}

	fmt.Println("--------------------------------")
	fmt.Println(jsonValue)
}
