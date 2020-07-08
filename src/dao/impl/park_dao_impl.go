package impl

import (
	"community-api/constants"
	"community-api/dao"
	"community-api/db"
	"sync"
)

var parkDaoOnce sync.Once
var parkDao *parkDaoImpl

type parkDaoImpl struct{}

func init() {
	parkDaoOnce.Do(func() {
		parkDao = new(parkDaoImpl)
	})
}

func NewParkDao() dao.ParkDao {
	return parkDao
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
func (o *parkDaoImpl) AddPark(businessId, userId int64, parkName string) (int64, error) {
	inserDef := &db.InsertDef{
		Table:      constants.DB_PSCC_PARK_TABLE,
		InsertCols: []string{"name", "business_id", "opertor_id"},
		InsertVals: []interface{}{parkName, businessId, userId},
	}
	return db.GetDbGroupHandler(constants.COMMUNITY_CLOUD_DB_NAME).InsertReturnPrimaryKeyValue(inserDef)
}

/**
修改园区状态
	businessId
		商户ID
	userId
		操作人用户ID
*/
func (o *parkDaoImpl) UpdateParkStatus(businessId, userId, parkId int64, status int) (bool, error) {
	updateDef := &db.UpdateDef{
		Table:      constants.DB_PSCC_PARK_TABLE,
		Conds:      []string{"park_id=?", "business_id=?"},
		UpdateCols: []string{"status=?", "opertor_id=?"},
		Vals:       []interface{}{status, userId, parkId, businessId},
	}
	affectedNum, e := db.GetDbGroupHandler(constants.COMMUNITY_CLOUD_DB_NAME).UpdateForRowAffected(updateDef)
	if e != nil {
		return false, e
	}
	//如果修改条数大于0，则修改成功，否则修改失败
	return affectedNum > 0, nil
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
func (o *parkDaoImpl) EditPark(businessId, userId, parkId int64, status int, parkName string) (bool, error) {
	updateDef := &db.UpdateDef{
		Table:      constants.DB_PSCC_PARK_TABLE,
		Conds:      []string{"park_id=?", "business_id=?", "status=?"},
		UpdateCols: []string{"name=?", "opertor_id=?"},
		Vals:       []interface{}{parkName, userId, parkId, businessId, status},
	}
	_, e := db.GetDbGroupHandler(constants.COMMUNITY_CLOUD_DB_NAME).UpdateForRowAffected(updateDef)
	if e != nil {
		return false, e
	}
	//如果修改条数大于0，则修改成功，否则修改失败
	//return affectedNum > 0, nil
	return true, nil
}
