package web

import (
	"community-api/constants"
	"community-api/logging"
	"community-api/model"
	"community-api/service/impl"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
添加园区
*/
func addParkV1(c *gin.Context) {
	stringBusinessId := c.Param("business_id")
	stringUserId := c.GetHeader("user_id")
	body, err := readBody(c, true)
	if err != nil {
		logging.Logger.Errorf("web.addParkV1() read body erro=[%s]", err)
		return
	}
	if len(body) > 0 {

		request := &model.RequestAddPark{}
		//解析请求参数
		err := parseParamV1(body, request, c, true)
		if err != nil {
			logging.Logger.Error("web.addParkV1 parseParamV1 body=[%s],error=[%s]", body, err)
			return
		}
		//验证参数
		if request.Name == "" {
			c.JSON(200,model.ResponseCode{
				Code: constants.ResponseCodeLessParamError,
				Desc: constants.ResponseCodeLessParamError.String(),
			})
			logging.Logger.Error("web.addParkV1() stringBusinessId=%s, stringUserId=%s park_name is null",
				stringBusinessId, stringUserId)
			return
		}

		//数据库添加插入园区记录
		int64UserId, _:= strconv.ParseInt(stringUserId,10,64)
		int64BusinessId, _:= strconv.ParseInt(stringBusinessId,10,64)
		resultJson := impl.NewParkService().AddPark(int64BusinessId, int64UserId, request.Name)
		//返回结果
		c.Writer.WriteString(resultJson)
	}
}

/**
修改园区
 */
func editParkV1(c *gin.Context) {
	stringBusinessId := c.Param("business_id")
	stringParkId := c.Param("park_id")
	stringUserId := c.GetHeader("user_id")
	body, err := readBody(c, true)
	if err != nil {
		logging.Logger.Errorf("web.editParkV1() read body erro=[%s]", err)
		return
	}
	if len(body) > 0 {

		request := &model.RequestEditPark{}
		//解析请求参数
		err := parseParamV1(body, request, c, true)
		if err != nil {
			logging.Logger.Error("web.editParkV1 parseParamV1 body=[%s],error=[%s]", body, err)
			return
		}
		//验证参数
		if stringParkId == "" || request.Name == "" {
			c.JSON(200,model.ResponseCode{
				Code: constants.ResponseCodeLessParamError,
				Desc: constants.ResponseCodeLessParamError.String(),
			})
			logging.Logger.Error("web.addParkV1() stringBusinessId=%s, stringUserId=%s , park_name=[%s],stringParkId=[%d] error",
				stringBusinessId, stringUserId, stringParkId)
			return
		}
		var int64ParkId int64 = 0
		if int64ParkId, err = strconv.ParseInt(stringParkId,10,64); err != nil {
			c.JSON(200,model.ResponseCode{
				Code: constants.ResponseCodeParamError,
				Desc: constants.ResponseCodeParamError.String(),
			})
			logging.Logger.Error("web.addParkV1() stringBusinessId=%s, stringUserId=%s ,stringParkId=[%d], error[=[%s]",
				stringBusinessId, stringUserId, stringParkId, err)
			return
		}

		//数据库修改园区记录
		int64UserId, _:= strconv.ParseInt(stringUserId,10,64)
		int64BusinessId, _:= strconv.ParseInt(stringBusinessId,10,64)
		success, err := impl.NewParkService().EditPark(int64BusinessId, int64UserId, int64ParkId, request.Name)
		//判断是否处理异常
		if err != nil {
			c.JSON(200,model.ResponseCode{
				Code: constants.ResponseCodeServerError,
				Desc: constants.ResponseCodeServerError.String(),
			})
			logging.Logger.Error()
			return
		}
		//判断是否处理失败
		if !success {
			c.JSON(200,model.ResponseCode{
				Code: constants.ResponseFail,
				Desc: constants.ResponseFail.String(),
			})
			return
		}

		//返回成功结果
		c.JSON(200,model.ResponseCode{
			Code: constants.ResponseCodeSuccess,
			Desc: constants.ResponseCodeSuccess.String(),
		})
	}

}

//删除园区
func delParkV1(c *gin.Context) {
	stringBusinessId := c.Param("business_id")
	stringParkId := c.Param("park_id")
	stringUserId := c.GetHeader("user_id")

	//验证参数
	if stringParkId == "" {
		c.JSON(200, model.ResponseCode{
			Code: constants.ResponseCodeLessParamError,
			Desc: constants.ResponseCodeLessParamError.String(),
		})
		logging.Logger.Errorf("web.delParkV1() stringBusinessId=%s, stringUserId=%s , park_name=[%s],stringParkId=[%d] error",
			stringBusinessId, stringUserId, stringParkId)
		return
	}

	var err error
	var int64ParkId int64 = 0
	if int64ParkId, err = strconv.ParseInt(stringParkId, 10, 64); err != nil {
		c.JSON(200, model.ResponseCode{
			Code: constants.ResponseCodeParamError,
			Desc: constants.ResponseCodeParamError.String(),
		})
		logging.Logger.Errorf("web.delParkV1() stringBusinessId=%s, stringUserId=%s ,stringParkId=[%d], error[=[%s]",
			stringBusinessId, stringUserId, stringParkId, err)
		return
	}

	//数据库园区记录状态
	int64UserId, _ := strconv.ParseInt(stringUserId, 10, 64)
	int64BusinessId, _ := strconv.ParseInt(stringBusinessId, 10, 64)
	status := constants.DB_PSCC_PARK_TABLE_FILED_STAUT_1
	success, err := impl.NewParkService().UpdateParkStatus(int64BusinessId, int64UserId, int64ParkId, status)

	//判断是否处理异常
	if err != nil {
		//判断是否还有下级数据
		if err.Error() == constants.ResponseHaveDataNoAllowDel.String() {
			c.JSON(200, model.ResponseCode{
				Code: constants.ResponseHaveDataNoAllowDel,
				Desc: constants.ResponseHaveDataNoAllowDel.String(),
			})
		} else {
			c.JSON(200, model.ResponseCode{
				Code: constants.ResponseCodeServerError,
				Desc: constants.ResponseCodeServerError.String(),
			})
		}

		logging.Logger.Errorf("web.delParkV1() impl.NewParkService().UpdateParkStatus(%d,%d,%d,%d), error[=[%s]",
			int64BusinessId, int64UserId, int64ParkId, status, err)
		return
	}

	//判断是否处理失败
	if !success {
		c.JSON(200, model.ResponseCode{
			Code: constants.ResponseFail,
			Desc: constants.ResponseFail.String(),
		})
		return
	}

	//返回成功结果
	c.JSON(200, model.ResponseCode{
		Code: constants.ResponseCodeSuccess,
		Desc: constants.ResponseCodeSuccess.String(),
	})
}