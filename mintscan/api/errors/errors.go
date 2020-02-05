package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error code and msg
type ErrorCode uint32
type ErrorMsg string

// WrapError parses the error into an object-like struct for exporting
type WrapError struct {
	ErrorCode ErrorCode `json:"error_code"`
	ErrorMsg  ErrorMsg  `json:"error_msg"`
}

const (
	InternalServer ErrorCode = 101

	DuplicateAccount  ErrorCode = 201
	InvalidFormat     ErrorCode = 202
	NotExist          ErrorCode = 203
	FailedConversion  ErrorCode = 204
	NotExistValidator ErrorCode = 205

	OverMaxLimit                      ErrorCode = 301
	FailedUnmarshalJSON               ErrorCode = 302
	FailedMarshalBinaryLengthPrefixed ErrorCode = 303
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
	case NotExist:
		return "NotExist"
	case NotExistValidator:
		return "NotExistValidator"
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

/*
	----------------------------------------------- Error Types
*/

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

func ErrNotExist(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: NotExist,
		ErrorMsg:  ErrorCodeToErrorMsg(NotExist),
	}
	PrintException(w, statusCode, wrapError)
}

func ErrNotExistValidator(w http.ResponseWriter, statusCode int) {
	wrapError := WrapError{
		ErrorCode: NotExistValidator,
		ErrorMsg:  ErrorCodeToErrorMsg(NotExistValidator),
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

/*
	----------------------------------------------- PrintException
*/

// PrintException prints out the exception result
func PrintException(w http.ResponseWriter, statusCode int, err WrapError) {
	w.Header().Add("Content-Type", "application/json")

	// Write HTTP status code
	w.WriteHeader(statusCode)

	result, _ := json.Marshal(err)

	fmt.Fprint(w, string(result))
}
