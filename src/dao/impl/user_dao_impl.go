package impl

import (
	"community-api/constants"
	"community-api/dao"
	"community-api/db"
	"community-api/logging"
	"fmt"
	"sync"
)

var userDaoOnce sync.Once
var userDao *userDaoImpl

type userDaoImpl struct{}

func init() {
	userDaoOnce.Do(func() {
		parkDao = new(parkDaoImpl)
	})
}

func NewUserDao() dao.UserDao {
	return userDao
}

/**
根据room信息查询user
	roomId
		房间id
*/
func (o *userDaoImpl) QueryUsersByRoom(roomId int64) ([]string, error) {
	var condStr = fmt.Sprintf("room_id = %d", roomId)
	queryDef := &db.QueryDef{
		Table: constants.DB_PSCC_USER_TABLE,
		Cols:  []string{"user_id"},
		Conds: []string{condStr},
	}

	res, err := db.GetDbGroupHandler(constants.COMMUNITY_CLOUD_DB_NAME).Get(queryDef)
	if err != nil {
		logging.Logger.Error("db get users by room failed, err is ", err)
		return []string{}, err
	}

	users := []string{}
	for _, val := range res {
		for _, mval := range val {
			users = append(users, mval)
		}

	}
	return users, nil
}
