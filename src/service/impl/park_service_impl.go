package impl

import (
	"community-api/constants"
	"community-api/dao/impl"
	"community-api/logging"
	"community-api/model"
	"community-api/service"
	"encoding/json"
	"errors"
	"sync"
)

var parkServiceOnce sync.Once
var parkService *parkServiceImpl

type parkServiceImpl struct{}

func init() {
	parkServiceOnce.Do(func() {
		parkService = new(parkServiceImpl)
	})
}

func NewParkService() service.ParkService {
	return parkService
}

/**
添加园区
	businessId
		商户ID
	userId
		操作人用户ID
	parkName
		园区名称
*/
func (o *parkServiceImpl) AddPark(businessId, userId int64, parkName string) string {
	parkId, err := impl.NewParkDao().AddPark(businessId, userId, parkName)
	if err != nil {
		data, _ := json.Marshal(model.ResponseCode{Code: constants.ResponseCodeServerError, Desc: constants.ResponseCodeServerError.String()})
		logging.Logger.Error("service.AddPark(%d, %d, %s) error=[%s]", businessId, userId, parkName, err)
		return string(data)
	}
	responseStruct := &model.ResponseStruct{Data: &model.ResponseAddPark{ParkId: parkId}}
	responseStruct.Code = constants.ResponseCodeSuccess
	responseStruct.Desc = constants.ResponseCodeSuccess.String()
	data, _ := json.Marshal(responseStruct)
	return string(data)
}

/**
  修改园区状态
  	businessId
  		商户ID
  	userId
  		操作人用户ID
  	parkId
  		园区ID
*/
func (o *parkServiceImpl) UpdateParkStatus(businessId, userId, parkId int64, status int) (bool, error) {
	//验证该园区下面是否有楼栋
	buildingNum, err := NewBuildingService().FindBuildingNum(businessId, parkId, constants.DB_PSCC_BUILDING_TABLE_FILED_STAUT_0)
	if err != nil {
		logging.Logger.Error("service.UpdateParkStatus(%d, %d, %s) error=[%s]", businessId, parkId, status, err)
		return false, err
	}
	if buildingNum == 0 {
		return impl.NewParkDao().UpdateParkStatus(businessId, userId, parkId, status)
	}
	return false, errors.New(constants.ResponseHaveDataNoAllowDel.String())
}

/**
编辑园区
	businessId
		商户ID
	userId
		操作人用户ID
	parkId
		园区ID
	parkName
		园区名称
*/
func (o *parkServiceImpl) EditPark(businessId, userId, parkId int64, parkName string) (bool, error) {
	return impl.NewParkDao().EditPark(businessId, userId, parkId, constants.DB_PSCC_PARK_TABLE_FILED_STAUT_0, parkName)
}
