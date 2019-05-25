package apierror

//APIError is a structure of single error of api
type APIError struct {
	Code     string `json:"code"`
	Resource string `json:"resource"`
	Message  string `json:"message"`
}

//errorCode is a map of error code
var errorCode = map[int]string{
	//User error
	1: "RequestParameterInvalid", //request parameter is not formatted enough to make SQL query
	2: "BadRequestParameter",     //request parameter is well-formatted ,but meanless
	3: "TimeConflicted",          //time of request parameter is conflicted with record in DataBase
	4: "ResourceAlreadyExist",    //unique resource is already exist
	5: "AuthenticationFailed",    //authentication of user failed with some reason

	//Internal Server Error
	11: "FailedExecQuery", //query string is correct,but execing query failed in some reason
}

//Error is a function to make new APIError
func Error(code int, resource string, message string) APIError {
	return APIError{Code: errorCode[code], Resource: resource, Message: message}
}
