package service

type ParkService  interface {
	/**
	添加园区
		businessId
			商户ID
		userId
			操作人用户ID
		parkName
			园区名称
	*/
	AddPark(businessId, userId int64, parkName string) string

	/**
	  修改园区状态
	  	businessId
	  		商户ID
	  	userId
	  		操作人用户ID
	  	parkId
	  		园区ID
	  	status
	  		园区状态
	*/
	UpdateParkStatus(businessId, userId, parkId int64, status int) (bool, error)

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
	EditPark(businessId, userId, parkId int64, parkName string) (bool, error)
}

