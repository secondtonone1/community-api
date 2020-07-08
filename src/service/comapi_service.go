package service

import (
	"community-api/protobuffer_def"
)

type ComApiService interface {
	QueryUsersByRoom(baseRequest *protobuffer_def.BaseRequest, baseResponse *protobuffer_def.BaseResponse) error
}
