package response

//any = interface{}
// in modern Go (1.18 and later), it's now preferred to use the alias any instead of interface{} when you simply mean “any type.”

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessResponse(data any) ResponseData {
	return ResponseData{
		Code:    ErrorCodeSuccess,
		Message: GetMsg(ErrorCodeSuccess),
		Data:    data,
	}
}

func ErrorResponse(code int, data any) ResponseData {
	return ResponseData{
		Code:    code,
		Message: GetMsg(code),
		Data:    data,
	}
}
