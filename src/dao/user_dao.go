package dao

type UserDao interface {

	/**
	根据room信息查询user
		roomId
			房间id
	*/
	QueryUsersByRoom(roomId int64) ([]string, error)
}
