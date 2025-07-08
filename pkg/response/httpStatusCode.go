package response

const (
	ErrorCodeSuccess      = 20001 // success
	ErrorCodeParamInvalid = 20002 // email invalid
	ErrorCodeTokenInvalid = 30001
)

// Error messages corresponding to each code
var Msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorCodeTokenInvalid: "Invalid Token Bro",
}

func GetMsg(code int) string {
	if msg, ok := Msg[code]; ok {
		return msg
	}
	return "Unknown error"
}
