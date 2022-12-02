package resp_code

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorNotCompare     = 10007

	ErrorExistTeam    = 20002
	ErrorNotExistTeam = 20003

	ErrorAuthCheckTokenFail    = 30001
	ErrorAuthCheckTokenTimeout = 30002
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorDatabase              = 40001
)
