package impl

import (
	"community-api/constants"
	"community-api/dao"
	"community-api/db"
	"sync"
)

var buildingDaoOnce sync.Once
var buildingDao *buildingDaoImpl

type buildingDaoImpl struct{}

func init() {
	buildingDaoOnce.Do(func() {
		buildingDao = new(buildingDaoImpl)
	})
}

func NewBuildingDao() dao.BuildingDao {
	return buildingDao
}

/**
添加楼栋
	businessId
		商户ID
	userId
		操作人用户ID
	parkId
		园区ID
	parkName
		园区名称
*/
func (o *buildingDaoImpl) AddBuilding(businessId, userId, parkId int64, buildingName string) (int64, error) {
	return 0, nil
}

/**
修改楼栋状态
	businessId
		商户ID
	userId
		操作人用户ID
	buildingId
		楼栋ID
*/
func (o *buildingDaoImpl) UpdateBuilding(businessId, userId, buildingId int64, status int) (bool, error) {
	return true, nil
}

/**
编辑楼栋
	businessId
		商户ID
	userId
		操作人用户ID
	buildingId
		楼栋ID
	buildingName
		楼栋名称
*/
func (o *buildingDaoImpl) EditBuilding(businessId, userId, buildingId int64, status int, buildingName string) (bool, error) {
	return true, nil
}

/**
查询园区楼栋数量
*/
func (o *buildingDaoImpl) FindBuildingNum(businessId, parkId int64, status int) (int64, error) {
	querDef := &db.QueryCountDef{
		Table: constants.DB_PSCC_BUILDING_TABLE,
		Conds: []string{"park_id=?", "business_id=?", "status=?"},
		Vals:  []interface{}{parkId, businessId, status},
	}
	return db.GetDbGroupHandler(constants.COMMUNITY_CLOUD_DB_NAME).Count(querDef)
}
