package apierror

//APIError is a structure of single error of api
type APIError struct {
	code     string
	resource string
	message  string
}

//APIErrors is a structure of errors of api
type APIErrors struct {
	errors []APIError
}

//errorCode is a map of error code
var errorCode = map[int]string{
	//User error
	1: "RequestParameterInvalid", //request parameter is not formatted enough to make SQL query
	2: "BadRequestParameter",     //request parameter is well-formatted ,but meanless
	3: "TimeConflicted",          //time of request parameter is conflicted with record in DataBase
	//Internal Server Error
	11: "FailedExecQuery", //query string is correct,but execing query failed in some reason
}

//MakeError is a function to make new APIError
func MakeError(code int, resource string, message string) APIError {
	return APIError{code: errorCode[code], resource: resource, message: message}
}
