package response

import (
	responseHandler "github.com/dev-jb/response-management-go/responseHandler"
	"github.com/dev-jb/response-management-go/status"
)

const (

	// SomethingWentWrongMsg Message
	SomethingWentWrongMsg = "Something went wrong"

	// InvalidFunctionNameMsg message
	InvalidFunctionNameMsg = "Invalid Smart Contract function name"

	// IncorrectNoOfArgsMsg message
	IncorrectNoOfArgsMsg = "Incorrect number of arguments. Expecting"

	// NotFoundMsg message
	NotFoundMsg = "Not Found"

	// AlreadyExistsMsg message
	AlreadyExistsMsg = "Already Exists"

	// UnauthorizedToAccessMsg message
	UnauthorizedToAccessMsg = "Unauthorized To Access"

	// FailedToGetCurrentDateTimeMsg message
	FailedToGetCurrentDateTimeMsg = "Failed to get current date and time"

	// RegistrationFailedMsg message
	RegistrationFailedMsg = "Registration Failed"

	// UpdateFailedMsg message
	UpdateFailedMsg = "Update Failed"

	// DeleteFailedMsg message
	DeleteFailedMsg = "Delete Failed"

	// GetQueryResultFailedMsg message
	GetQueryResultFailedMsg = "Unable to retrieve data"

	// RegistrationSuccessMsg message
	RegistrationSuccessMsg = "Registration Successful"

	// UpdateSuccessMsg message
	UpdateSuccessMsg = "Update Successful"

	// DeleteSuccessMsg message
	DeleteSuccessMsg = "Delete Successful"

	// DataRetrieveSuccessMsg message
	DataRetrieveSuccessMsg = "Data Retrieved Successful"

	// SuccessMsg message
	SuccessMsg = "Success"

	// EmptyMsg message
	EmptyMsg = "Empty"

	// InitializeSuccessMsg message
	InitializeSuccessMsg = "Initialized Successfully"
)

/* ==============================Success Response============================== */

// Success Function: call when there are no any described success functions.
func Success(message string, statusCode int32, args interface{}) responseHandler.Response {
	return responseHandler.HandleResponse(message, statusCode, args)
}

// RegistrationSuccess Function: calls when asset registered successfully. Returns message = Registration Successful: 'Your Message' with Status Code 200(OK) and args as result
func RegistrationSuccess(message string, args interface{}) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(RegistrationSuccessMsg, message), status.OK, args)
}

// UpdateSuccess Function: calls when asset updated successfully. Returns message = Update Successful: 'Your Message' with Status Code 200(OK) and args as result
func UpdateSuccess(message string, args interface{}) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(UpdateSuccessMsg, message), status.OK, args)
}

// DeleteSuccess Function: calls when asset deleted successfully. Returns message = Delete Successful: 'Your Message' with Status Code 200(OK) and args as result
func DeleteSuccess(message string, args interface{}) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(DeleteSuccessMsg, message), status.OK, args)
}

// DataReceivedSuccess Function: calls when data detched successfully. Returns message = Data Retrieved Successful: 'Your Message' with Status Code 200(OK) and args as result
func DataReceivedSuccess(message string, args interface{}) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(DataRetrieveSuccessMsg, message), status.OK, args)
}

// Empty Function: call when list is empty. Returns message = Empty: 'Your Message' with status code 200 (OK)
func Empty(message string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(EmptyMsg, message), status.OK, nil)
}

// InitializeSuccess Function: call when list is empty. Returns message= Initialized Successfully: 'Your Message' with status code 200(OK) and args as result
func InitializeSuccess(message string, args interface{}) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(InitializeSuccessMsg, message), status.OK, args)
}

/* ==============================Success Response============================== */

/* ==============================Custom Error============================== */

// Error function: call when there is no desired error defined.
func Error(message string, statusCode int32) responseHandler.Response {
	return responseHandler.HandleResponse(message, statusCode, nil)
}

/* ==============================Custom Error============================== */

/* ==============================Internal Server Errors============================== */

// JSONMarshalFailed Function: calls when encode to JSON fails. Returns Somehing Went Wrong Message With Status Code 500(Internal Server Error)
func JSONMarshalFailed() responseHandler.Response {
	return responseHandler.HandleResponse(SomethingWentWrongMsg, status.INTERNALSERVERERROR, nil)
}

// JSONUnMarshalFailed Function: calls when decode of JSON fails. Returns Somehing Went Wrong Message With Status Code 500(Internal Server Error)
func JSONUnMarshalFailed() responseHandler.Response {
	return responseHandler.HandleResponse(SomethingWentWrongMsg, status.INTERNALSERVERERROR, nil)
}

// FailedToGetCurrentDateTime Function: calls when there is error on getting current date and time. Returns Somehing Went Wrong Message With Status Code 500(Internal Server Error)
func FailedToGetCurrentDateTime() responseHandler.Response {
	return responseHandler.HandleResponse(SomethingWentWrongMsg, status.INTERNALSERVERERROR, nil)
}

// GenerateRandomNoFailed Function: calls when random no generation failed. Returns Somehing Went Wrong Message With Status Code 500(Internal Server Error)
func GenerateRandomNoFailed() responseHandler.Response {
	return responseHandler.HandleResponse(SomethingWentWrongMsg, status.INTERNALSERVERERROR, nil)
}

/* ==============================Internal Server Errors============================== */

// RegistrationFailed Function: calls when there is error on writing the ledger. Returns message = Registration Failed: 'Your Message' With Status Code 500(Internal Server Error)
func RegistrationFailed(message string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(RegistrationFailedMsg, message), status.INTERNALSERVERERROR, nil)
}

// UpdateFailed Function: calls when there is error on udpdate Function. Returns message = Update Failed: 'Your Message' With Status Code 500(Internal Server Error)
func UpdateFailed(message string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(UpdateFailedMsg, message), status.INTERNALSERVERERROR, nil)
}

// DeleteFailed Function: calls when there is error on writing the ledger. Returns message = Delete Failed: 'Your Message' With Status Code 500(Internal Server Error)
func DeleteFailed(message string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(DeleteFailedMsg, message), status.INTERNALSERVERERROR, nil)
}

// GetQueryResultFailed Function: call when there is error on getting the query result from ledger. Returns message = Unable to retrieve data: 'Your Message' With Status Code 500(Internal Server Error)
func GetQueryResultFailed(message string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(GetQueryResultFailedMsg, message), status.INTERNALSERVERERROR, nil)
}

/* ==============================Not Found Errors============================== */

// NotFound Function: calls when asset is not found. Returns message= Not Found: 'Your Message' with status code 404(Not Found)
func NotFound(message string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(NotFoundMsg, message), status.NOTFOUND, nil)
}

/* ==============================Not Found Errors============================== */

/* ==============================Forbidden Errors============================== */

// IncorrectNoOfArgs Function: call when no of args does not match the expected. Returns message= Incorrect No of Args. Expecting: 'Expected Value' with status code 403(Forbidden)
func IncorrectNoOfArgs(expectedValue string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(IncorrectNoOfArgsMsg, expectedValue), status.FORBIDDEN, nil)
}

/* ==============================Forbidden Errors============================== */

/* ==============================Not Implemented Errors============================== */

// InvalidFunctionName function: call when the expected function name does not exists. Returns message= Invalid Function Name: 'Function Name' with status code 501(Not Implemented)
func InvalidFunctionName(functioName string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(InvalidFunctionNameMsg, functioName), status.NOTIMPLEMENTED, nil)
}

/* ==============================Not Implemented Errors============================== */

/* ==============================Conflict Errors============================== */

// AlreadyExist Function: call when asset already exists on the ledger. Returns message= Already Exists: 'Your Message' with status code 409(Conflict)
func AlreadyExist(message string) responseHandler.Response {
	return responseHandler.HandleResponse(WrapMessage(AlreadyExistsMsg, message), status.CONFLICT, nil)
}

/* ==============================Conflict Errors============================== */

// UnauthorizedToAccess Function: call when user not authorized to access the information. Returns message= Unauthorized To Access: 'Your Message' with status code 401(Unauthorized)
func UnauthorizedToAccess() responseHandler.Response {
	return responseHandler.HandleResponse(UnauthorizedToAccessMsg, status.UNAUTHORIZED, nil)
}

// WrapMessage function
func WrapMessage(defaultMessage string, customMessage string) string {
	return defaultMessage + ":" + customMessage
}
