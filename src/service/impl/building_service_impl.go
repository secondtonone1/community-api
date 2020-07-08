package impl

import (
	"community-api/dao/impl"
	"community-api/service"
	"sync"
)

var buildingServiceOnce sync.Once
var buildingService *buildingServiceImpl

type buildingServiceImpl struct{}

func init() {
	parkServiceOnce.Do(func() {
		buildingService = new(buildingServiceImpl)
	})
}

func NewBuildingService() service.BuildingService {
	return buildingService
}

/**
根据园区ID查询楼栋数量
businessId
	商户ID
parkId
	园区ID
status
	状态
*/
func (b *buildingServiceImpl) FindBuildingNum(businessId, parkId int64, status int) (int64, error) {
	return impl.NewBuildingDao().FindBuildingNum(businessId, parkId, status)
}
