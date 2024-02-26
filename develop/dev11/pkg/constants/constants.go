package constants

type ctxKey string

const (
	KeyUserID      ctxKey = "user_id"
	KeyLogger      ctxKey = "logger"
	KeyRequestId   ctxKey = "request_id"
	KeyRequestInfo ctxKey = "request_info"
)
