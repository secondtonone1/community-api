package grpc_client

import (
	"community-api/protobuffer_def"
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/zookeeper/v2"
)

func Start() {
	r := zookeeper.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2181"}
		op.Context = context.Background()
		op.Timeout = time.Second * 5
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create GRPC service
	service := grpc.NewService(
		service.Name("test.client2"),
		service.Registry(r),
		service.Context(ctx),
	)

	err := service.Client().Init(client.Retries(3), client.PoolSize(200), client.PoolTTL(time.Second*20), client.RequestTimeout(time.Second*5))
	if err != nil {
		fmt.Println("service client init failed, err is ", err)
		return
	}

	test := protobuffer_def.NewComApiServerService("api-service", service.Client())
	fmt.Println("test is ", test)
	baseReq := &protobuffer_def.BaseRequest{}
	baseReq.RequestId = "query_user_by_room"
	baseReq.C = protobuffer_def.CMD_QUERY_USERS_BY_ROOM
	queryReq := &protobuffer_def.QueryUsersByRoomReq{RoomId: "101"}

	body, err := ptypes.MarshalAny(queryReq)
	if err != nil {
		fmt.Println("proto marshal failed")
		return
	}
	baseReq.Body = body

	fmt.Println("baseReq is ", baseReq)
	baseRsp, err := test.BaseInterface(context.Background(), baseReq)
	if err != nil {
		fmt.Println("query failed, err is ", err)
		return
	}

	if baseRsp.Code != protobuffer_def.ReturnCode_SUCCESS {
		fmt.Println("query failed, code is ", baseRsp.Code)
		return
	}

	queryRsp := &protobuffer_def.QueryUserByRoomRsp{}
	err = ptypes.UnmarshalAny(baseRsp.GetBody(), queryRsp)
	if err != nil {
		fmt.Println("proto unmarsh failed")
		return
	}
	fmt.Println("query rsp is ", queryRsp)
}
