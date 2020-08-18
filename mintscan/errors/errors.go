package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorCode represents custom error code used in this application.
type ErrorCode uint32

// ErrorMsg represents error message that will be returned to client for any error.
type ErrorMsg string

// WrapError wraps both error code and error message.
type WrapError struct {
	ErrorCode ErrorCode `json:"error_code"`
	ErrorMsg  ErrorMsg  `json:"error_msg"`
}

const (
	InternalServer ErrorCode = 101

	DuplicateAccount   ErrorCode = 201
	InvalidFormat      ErrorCode = 202
	NotExist           ErrorCode = 203
	NotAllowed         ErrorCode = 204
	FailedConversion   ErrorCode = 205
	InvalidMessageType ErrorCode = 207

	OverMaxLimit                      ErrorCode = 301
	FailedUnmarshalJSON               ErrorCode = 302
	FailedMarshalBinaryLengthPrefixed ErrorCode = 303

	RequiredParam ErrorCode = 601
	InvalidParam  ErrorCode = 602
)

// ErrorCodeToErrorMsg returns error message from error code
func ErrorCodeToErrorMsg(code ErrorCode) ErrorMsg {
	switch code {
	case InternalServer:
		return "Internal server error"
	case DuplicateAccount:
		return "Duplicate account"
	case InvalidFormat:
		return "Invalid format"
	case InvalidMessageType:
		return "Invalid Message Type"
	case NotExist:
		return "NotExist"
	case NotAllowed:
		return "NotAllowed"
	case FailedConversion:
		return "FailedConversion"
	case OverMaxLimit:
		return "OverMaxLimit"
	case FailedUnmarshalJSON:
		return "FailedUnmarshalJSON"
	case FailedMarshalBinaryLengthPrefixed:
		return "FailedMarshalBinaryLengthPrefixed"
	default:
		return "Unknown"
	}
}

// ErrorCodeToErrorMsgs returns error message concatenating with custom message from error code
func ErrorCodeToErrorMsgs(code ErrorCode, msg string) ErrorMsg {
	switch code {
	case RequiredParam:
		return ErrorMsg("Required Parameter: " + msg)
	case InvalidParam:
		return ErrorMsg("Invalid Parameter: " + msg)
	default:
		return "Unknown"
	}
}

// --------------------
// Error Types
// --------------------

func ErrInternalServer(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: InternalServer,
		ErrorMsg:  ErrorCodeToErrorMsg(InternalServer),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrDuplicateAccount(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: DuplicateAccount,
		ErrorMsg:  ErrorCodeToErrorMsg(DuplicateAccount),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrInvalidFormat(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: InvalidFormat,
		ErrorMsg:  ErrorCodeToErrorMsg(InvalidFormat),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrInvalidMessageType(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: InvalidMessageType,
		ErrorMsg:  ErrorCodeToErrorMsg(InvalidMessageType),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrNotExist(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: NotExist,
		ErrorMsg:  ErrorCodeToErrorMsg(NotExist),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrNotAllowed(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: NotAllowed,
		ErrorMsg:  ErrorCodeToErrorMsg(NotAllowed),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrFailedConversion(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: FailedConversion,
		ErrorMsg:  ErrorCodeToErrorMsg(FailedConversion),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrOverMaxLimit(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: OverMaxLimit,
		ErrorMsg:  ErrorCodeToErrorMsg(OverMaxLimit),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrFailedUnmarshalJSON(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: FailedUnmarshalJSON,
		ErrorMsg:  ErrorCodeToErrorMsg(FailedUnmarshalJSON),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrFailedMarshalBinaryLengthPrefixed(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: FailedMarshalBinaryLengthPrefixed,
		ErrorMsg:  ErrorCodeToErrorMsg(FailedMarshalBinaryLengthPrefixed),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrRequiredParam(w http.ResponseWriter, statusCode int, msg string) {
	wrapError := WrapError{
		ErrorCode: RequiredParam,
		ErrorMsg:  ErrorCodeToErrorMsgs(RequiredParam, msg),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrInvalidParam(w http.ResponseWriter, statusCode int, msg string) {
	wrapError := WrapError{
		ErrorCode: InvalidParam,
		ErrorMsg:  ErrorCodeToErrorMsgs(InvalidParam, msg),
	}
	PrintException(w, statusCode, wrapError)
}

// --------------------
// PrintException
// --------------------

// PrintException prints out the exception result
func PrintException(w http.ResponseWriter, statusCode int, err WrapError) {
	w.Header().Add("Content-Type", "application/json")

	// Write HTTP status code
	w.WriteHeader(statusCode)

	result, _ := json.Marshal(err)

	fmt.Fprint(w, string(result))
}
