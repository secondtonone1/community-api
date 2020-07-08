package model

import "community-api/constants"

type ResponseCode struct {
	Code constants.ResponseCodeType `json:"code"`
	Desc string                     `json:"desc"`
}

type ResponseStruct struct {
	ResponseCode
	Data interface{} `json:"data"`
}
