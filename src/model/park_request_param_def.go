package model

//添加园区 请求
type RequestAddPark struct {
	Name string `json:"park_name"`   //园区名称
}
//添加园区 响应
type ResponseAddPark struct {
	ParkId int64 `json:"park_id"`   //园区名称
}


//编辑园区 请求
type RequestEditPark struct {
	Name string `json:"park_name"`   //园区名称
}
