package service

type BuildingService  interface {
	/**
	根据园区ID查询楼栋数量
	businessId
		商户ID
	parkId
		园区ID
	status
		状态
	 */
	FindBuildingNum(businessId, parkId int64, status int) (int64, error)
}

