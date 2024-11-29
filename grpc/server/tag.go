package server

import (
	"context"
	"encoding/json"
	"fmt"
	"grpcproj/pkg/bapi"
	"grpcproj/proto"
)

type TagServer struct {
	proto.UnimplementedTagServiceServer
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *proto.GetTagListRequest) (*proto.GetTagListResponse, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(r.GetName())
	if err != nil {
		return nil, err
	}
	tagList := proto.GetTagListResponse{}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, err
	}
	return &tagList, nil
}
