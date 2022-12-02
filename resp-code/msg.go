package resp_code

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "Invalid params",

	ErrorAuthCheckTokenFail:    "Token auth failed",
	ErrorAuthCheckTokenTimeout: "Token expired",
	ErrorAuthToken:             "Token generation failed",
	ErrorAuth:                  "Token error",
	ErrorNotCompare:            "Password not match",
	ErrorDatabase:              "Database error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
