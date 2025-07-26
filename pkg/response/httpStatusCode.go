package response

const (
	ErrorCodeSuccess       = 20001 // success
	ErrorCodeParamInvalid  = 20002 // email invalid
	ErrorCodeTokenInvalid  = 30001
	ErrorCodeUserHasExists = 50001
	ErrorCodeOtpStillValid = 50002
	ErrorCodeRedisError    = 50003
	ErrorCodeEmailSend     = 50004
	ErrorCodeNotFound      = 50005
	ErrorCodeInternal      = 50006
)

// Error messages corresponding to each code
var Msg = map[int]string{
	ErrorCodeSuccess:       "success",
	ErrorCodeParamInvalid:  "Email is invalid",
	ErrorCodeTokenInvalid:  "Invalid Token Bro",
	ErrorCodeUserHasExists: "User has already register",
	ErrorCodeOtpStillValid: "OTP already sent and still valid",
	ErrorCodeRedisError:    "Redis error occurred",
	ErrorCodeEmailSend:     "Email send fail bro",
	ErrorCodeNotFound:      "Email can't find",
	ErrorCodeInternal:      "Can't marsall",
}

func GetMsg(code int) string {
	if msg, ok := Msg[code]; ok {
		return msg
	}
	return "Unknown error"
}
