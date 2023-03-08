package errorcode

import "google.golang.org/grpc/codes"

type Code struct {
	StatusCode  codes.Code
	ErrCode     int64
	ErrMesssage string
}

var (
	// 基础
	ArgumentInvalid      = Code{StatusCode: codes.InvalidArgument, ErrCode: 400, ErrMesssage: "Argument invalid"}
	AuthorizationInvalid = Code{StatusCode: codes.Unauthenticated, ErrCode: 401, ErrMesssage: "Authorization invalid"}
	PermissionDenied     = Code{StatusCode: codes.PermissionDenied, ErrCode: 403, ErrMesssage: "Permission denied"}

	TokenCreateFailed   = Code{StatusCode: codes.Unknown, ErrCode: 401001, ErrMesssage: "Token creation failed"}
	SessionCreateFailed = Code{StatusCode: codes.Unknown, ErrCode: 401002, ErrMesssage: "Session creation failed"}
	SessionInvalid      = Code{StatusCode: codes.Unknown, ErrCode: 403001, ErrMesssage: "Session invalid"}

	MissingRequiredParams = Code{StatusCode: codes.InvalidArgument, ErrCode: 422001, ErrMesssage: "Missing required parameters"}
	IDRequired            = Code{StatusCode: codes.InvalidArgument, ErrCode: 422002, ErrMesssage: "Field `id` is required"}
	NameRequired          = Code{StatusCode: codes.InvalidArgument, ErrCode: 422003, ErrMesssage: "Field `name` is required"}
	Unknown               = Code{StatusCode: codes.Unknown, ErrCode: 500, ErrMesssage: "Unknown error"}
	RecordNotFound        = Code{StatusCode: codes.Unknown, ErrCode: 500001, ErrMesssage: "Record was not found"}
	RecordCreationFailed  = Code{StatusCode: codes.Unknown, ErrCode: 500002, ErrMesssage: "Record creation failed"}
	RecordDeletionFailed  = Code{StatusCode: codes.Unknown, ErrCode: 500003, ErrMesssage: "Record deletion failed"}
	RecordUpdateFailed    = Code{StatusCode: codes.Unknown, ErrCode: 500004, ErrMesssage: "Record update failed"}
	RecordCountFailed     = Code{StatusCode: codes.Unknown, ErrCode: 500005, ErrMesssage: "Record count failed"}
	RecordsNotFound       = Code{StatusCode: codes.Unknown, ErrCode: 500006, ErrMesssage: "Records was not found"}
	// administrator 服务 编号6xx00x
	NameAndPasswordRequired = Code{StatusCode: codes.InvalidArgument, ErrCode: 641001, ErrMesssage: "Field `name` and `password` are required"}
	AdministratorIDRequired = Code{StatusCode: codes.InvalidArgument, ErrCode: 641002, ErrMesssage: "Field `administrator_id` is required"}
	MethodRequired          = Code{StatusCode: codes.InvalidArgument, ErrCode: 642003, ErrMesssage: "Field `method` is required"}
	AdministratorExisted    = Code{StatusCode: codes.Unknown, ErrCode: 651001, ErrMesssage: "Administrator existed"}
	PasswordError           = Code{StatusCode: codes.Unknown, ErrCode: 651002, ErrMesssage: "Field `password` error"}
	NameAndPasswordError    = Code{StatusCode: codes.Unknown, ErrCode: 651003, ErrMesssage: "Field `name` and `password` error"}
	RoleTypeExisted         = Code{StatusCode: codes.Unknown, ErrCode: 652001, ErrMesssage: "Role type existed"}
	RolesBindFailed         = Code{StatusCode: codes.Unknown, ErrCode: 652002, ErrMesssage: "Roles bind failed"}
	RolesUnbindFailed       = Code{StatusCode: codes.Unknown, ErrCode: 652003, ErrMesssage: "Roles unbind failed"}
	PermissionsBindFailed   = Code{StatusCode: codes.Unknown, ErrCode: 653001, ErrMesssage: "Permissions bind failed"}
	PermissionsUnbindFailed = Code{StatusCode: codes.Unknown, ErrCode: 653002, ErrMesssage: "Permissions unbind failed"}
	// article 服务 编号7xx00x
)
