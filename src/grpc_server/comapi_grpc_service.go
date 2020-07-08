package grpc_server

import (
	"community-api/config"
	"community-api/protobuffer_def"
	"community-api/service/impl"
	"context"
	"sync"
	"time"

	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
)

var (
	comApiService service.Service
	comApiOnce    = &sync.Once{}
	comApiServer  server.Server
)

//初始化grpc service服务
func StartComApiService(config *config.Config) {
	comApiOnce.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// create GRPC service
		service := grpc.NewService(
			service.Address(config.Base.GRPCAddr),
			service.Name(config.Base.ServiceName),
			service.Registry(config.RegisterCenter.GetRegisterCenter()),
			service.RegisterTTL(time.Second*30),
			service.RegisterInterval(time.Second*20),
			service.Context(ctx),
		)
		service.Init()

		comApiServer = service.Server()
		// register test handler
		protobuffer_def.RegisterComApiServerHandler(service.Server(), &comApiGrpcServiceImpl{})

		//启动服务
		if err := service.Run(); err != nil {
			panic(err)
		}

	})
}

func StopComApiService() {
	comApiServer.Stop()
}

type comApiGrpcServiceImpl struct {
}

func (s *comApiGrpcServiceImpl) BaseInterface(context context.Context, baseRequest *protobuffer_def.BaseRequest, baseResponse *protobuffer_def.BaseResponse) error {
	baseResponse.C = baseRequest.GetC()
	baseResponse.RequestId = baseRequest.GetRequestId()
	baseResponse.Code = protobuffer_def.ReturnCode_SUCCESS

	switch baseRequest.GetC() {
	case protobuffer_def.CMD_QUERY_USERS_BY_ROOM: //根据roomid查询用户
		return impl.NewComApiService().QueryUsersByRoom(baseRequest, baseResponse)
	default:
		baseResponse.Desc = "unkown cmd"
		baseResponse.Code = protobuffer_def.ReturnCode_UNKOWN_CMD //示知的指令
	}
	return nil
}
