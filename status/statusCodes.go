package status

const (
	// OK constant - status code less than 400, endorser will endorse it.
	// OK means init or invoke successfully.
	OK = 200

	// ERRORTHRESHOLD constant - status code greater than or equal to 400 will be considered an error and rejected by endorser.
	// BADREQUEST ERROR
	ERRORTHRESHOLD = 400

	// UNAUTHORIZED constant - error specifically for use when authentication is required and has failed or has not yet been provided
	UNAUTHORIZED = 401

	// FORBIDDEN constant - error when the request contained valid data and was understood by the server, but the server is refusing action.
	FORBIDDEN = 403

	// NOTFOUND constant - error when requested resource could not be found but may be available in the future
	NOTFOUND = 404

	// INTERNALSERVERERROR constant - a generic error message, given when an unexpected condition was encountered and no more specific message is suitable
	INTERNALSERVERERROR = 500

	// NOTIMPLEMENTED constant - error when it does not recognize the request method/function, or it lacks the ability to fulfil the request
	NOTIMPLEMENTED = 501

	// CONFLICT constant - error when the request could not be processed because of conflict in the current state of the resource.
	CONFLICT = 409
)
