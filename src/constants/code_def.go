package constants

type ResponseCodeType int

const (
	ResponseCodeSuccess        ResponseCodeType = 1
	ResponseCodeBodyIsNull     ResponseCodeType = 2
	ResponseCodeJsonParsError  ResponseCodeType = 3
	ResponseCodeLessParamError ResponseCodeType = 4
	ResponseCodeServerError    ResponseCodeType = 5
	ResponseCodeParamError     ResponseCodeType = 6
	ResponseFail               ResponseCodeType = 7
	ResponseHaveDataNoAllowDel ResponseCodeType = 8
	ResponseCodeAuthError      ResponseCodeType = 403
)

func (p ResponseCodeType) String() string {
	switch (p) {
	case ResponseCodeSuccess:
		return "success"
	case ResponseCodeBodyIsNull:
		return "body is null"
	case ResponseCodeJsonParsError:
		return "json parse error"
	case ResponseCodeLessParamError:
		return "less param"
	case ResponseCodeServerError:
		return "server error"
	case ResponseCodeParamError:
		return "param error"
	case ResponseFail:
		return "fail"
	case ResponseHaveDataNoAllowDel:
		return "have data, not allow delete"
	case ResponseCodeAuthError:
		return "check auth fail"
	default:
		return "unknown code desc"
	}
}
