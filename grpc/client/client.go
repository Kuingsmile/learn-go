package main

import (
	"context"
	"fmt"
	"grpcproj/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, _ := GetClientConn("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer clientConn.Close()
	tagServiceClient := proto.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(context.Background(), &proto.GetTagListRequest{Name: "RRRR"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

func GetClientConn(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.NewClient(target, opts...)
}
