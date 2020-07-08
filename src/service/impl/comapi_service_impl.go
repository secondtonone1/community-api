package impl

import (
	"community-api/dao/impl"
	"community-api/logging"
	"community-api/protobuffer_def"
	serivce "community-api/service"
	"fmt"
	"strconv"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

var (
	csi     *comApiServiceImpl
	csiOnce = &sync.Once{}
)

type comApiServiceImpl struct{}

func init() {
	fmt.Println("init comapi service")
}

func NewComApiService() serivce.ComApiService {
	csiOnce.Do(func() {
		csi = &comApiServiceImpl{}
	})
	return csi
}

func (s *comApiServiceImpl) preDeal(baseRequest *protobuffer_def.BaseRequest,
	baseResponse *protobuffer_def.BaseResponse, request proto.Message) bool {
	//解析body
	if baseRequest.GetBody() == nil {
		baseResponse.Code = protobuffer_def.ReturnCode_BODY_IS_NULL
		baseResponse.Desc = "body is null"
		return false
	}
	//反序列化
	err := ptypes.UnmarshalAny(baseRequest.GetBody(), request)
	if err != nil {
		baseResponse.Code = protobuffer_def.ReturnCode_DESERIALIZATION_ERROR
		baseResponse.Desc = "body deserialization error"
		return false
	}
	return true
}

func (s *comApiServiceImpl) QueryUsersByRoom(baseReq *protobuffer_def.BaseRequest,
	baseRsp *protobuffer_def.BaseResponse) error {
	baseRsp.Code = protobuffer_def.ReturnCode_SUCCESS
	//解析请求参数
	request := &protobuffer_def.QueryUsersByRoomReq{}
	if !s.preDeal(baseReq, baseRsp, request) {
		return nil
	}
	roomId, err := strconv.ParseInt(request.RoomId, 10, 64)
	if err != nil {
		logging.Logger.Error("roomid parse int failed, err is ", err)
		baseRsp.Code = protobuffer_def.ReturnCode_UNKOWN_ERROR
		baseRsp.Desc = "QueryUsersByRoom err"
		return err
	}
	users, err := impl.NewUserDao().QueryUsersByRoom(roomId)
	if err != nil {
		logging.Logger.Error("QueryUsersByRoom failed, err is ", err)
		baseRsp.Code = protobuffer_def.ReturnCode_UNKOWN_ERROR
		baseRsp.Desc = "QueryUsersByRoom err"
		return err
	}

	userrsp := &protobuffer_def.QueryUserByRoomRsp{}
	for _, id := range users {
		userinfo := &protobuffer_def.QueryUserByRoomRsp_UserInfo{
			UserId: id,
		}
		userrsp.UserInfos = append(userrsp.UserInfos, userinfo)
	}
	logging.Logger.Debug("userrsp is ", userrsp)
	body, err := ptypes.MarshalAny(userrsp)
	if err != nil {
		baseRsp.Code = protobuffer_def.ReturnCode_SERIALIZATION_ERROR
		baseRsp.Desc = "MarshalAny error"
		return err
	}

	baseRsp.Code = protobuffer_def.ReturnCode_SUCCESS
	baseRsp.Desc = "success"

	baseRsp.Body = body

	return nil
}
