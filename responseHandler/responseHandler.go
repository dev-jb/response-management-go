package responsehandler

import "encoding/json"

//ErrorMessage Struct
type ErrorMessage struct {
	ErrorCode    int32  `json:"errorCode,omitempty"`
	ErrorMessage string `json:"erroMessage,omitempty"`
}

//SuccessMessage Struct
type SuccessMessage struct {
	SuccessCode    int32       `json:"successCode,omitempty"`
	SuccessMessage string      `json:"successMessage,omitempty"`
	Result         interface{} `json:"result,omitempty"`
}

// Response Struct
type Response struct {
	StatusCode int32       `json:"statusCode,omitempty"`
	Message    string      `json:"message,omitempty"`
	Result     interface{} `json:"result,omitempty"`
}

// HandleResponse function to handle all the errors and success responses.
func HandleResponse(message string, statusCode int32, args interface{}) Response {
	// response := Response{}
	// response.StatusCode = statusCode
	// response.Message = message
	// response.Result = args
//	responseAsBytes, _ := json.Marshal(args)
	// if statusCode != status.OK {
	// 	if statusCode == status.INTERNALSERVERERROR {
	// 		// notify admin with emails or others...
	// 	}
	// 	return errorResponse(message, statusCode)
	// 	// return shim.Error(string(responseAsBytes))
	// 	// return defaultResponse.Error(message, statusCode)
	// }

	return response(message, args, statusCode)
	// return shim.Success(responseAsBytes)
	// return defaultResponse.Success(message, payload, statusCode)
}

func response(message string, payload interface{}, statusCode int32) Response {
	response := Response{}
	response.StatusCode = statusCode
	response.Message = message
	response.Result = payload
	return response
}
