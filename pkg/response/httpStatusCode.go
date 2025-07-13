package response

const (
	ErrorCodeSuccess       = 20001 // success
	ErrorCodeParamInvalid  = 20002 // email invalid
	ErrorCodeTokenInvalid  = 30001
	ErrorCodeUserHasExists = 50001
)

// Error messages corresponding to each code
var Msg = map[int]string{
	ErrorCodeSuccess:       "success",
	ErrorCodeParamInvalid:  "Email is invalid",
	ErrorCodeTokenInvalid:  "Invalid Token Bro",
	ErrorCodeUserHasExists: "User has already register",
}

func GetMsg(code int) string {
	if msg, ok := Msg[code]; ok {
		return msg
	}
	return "Unknown error"
}
