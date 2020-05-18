package commands

// Response is a generic response interface
type Response interface {
	// IsSuccess returns true if the sent command was executed successfully
	IsSuccess() bool
	// GetResult returns the response body
	GetResult() interface{}
	// GetErrorType returns the type of error if it was not succesful
	GetErrorType() string
	// GetErrorMessage returns the error message if it was not successful
	GetErrorMessage() string
}

// BaseResponse contains all possible response fields
type BaseResponse struct {
	Success      bool
	Result       interface{}
	ErrorType    string
	ErrorMessage string
}

// IsSuccess returns true if the sent command was executed successfully
func (br *BaseResponse) IsSuccess() bool {
	return br.Success
}

// GetResult returns the response body
func (br *BaseResponse) GetResult() interface{} {
	return br.Result
}

// GetErrorType returns the type of error if it was not succesful
func (br *BaseResponse) GetErrorType() string {
	return br.ErrorType
}

// GetErrorMessage returns the error message if it was not successful
func (br *BaseResponse) GetErrorMessage() string {
	return br.ErrorMessage
}
