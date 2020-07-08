package dao

type BuildingDao  interface {

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
	AddBuilding(businessId, userId, parkId int64, buildingName string) (int64, error)

	/**
	修改楼栋状态
		businessId
			商户ID
		userId
			操作人用户ID
		buildingId
			楼栋ID
	 */
	UpdateBuilding(businessId, userId, buildingId int64, status int) (bool, error)

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
	EditBuilding(businessId, userId, buildingId int64,  status int, buildingName string) (bool, error)

	/**
	查询园区楼栋数量
	 */
	FindBuildingNum(businessId, parkId int64, status int) (int64, error)
}