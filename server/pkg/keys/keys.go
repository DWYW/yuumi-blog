package keys

type CustomKey string

const (
	RequestIDKey             = CustomKey("RequestID")
	RequestFullMethodKey     = CustomKey("RequestFullMethod")
	RequestStartTimeKey      = CustomKey("RequestStartTime")
	ListPageMaxNumKey        = CustomKey("ListPageMaxNum")
	AdministratorAuthDataKey = CustomKey("AdministratorAuthData")
	VisitorAuthDataKey       = CustomKey("VisitorAuthData")
)

type AdministratorAuthData struct {
	ID   int64
	Name string
}

type VisitorAuthData struct {
	ID   int64
	Name string
}
