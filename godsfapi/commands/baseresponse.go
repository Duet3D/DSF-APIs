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

type BaseResponse struct {
	Success      bool
	Result       interface{}
	ErrorType    string
	ErrorMessage string
}

func (br *BaseResponse) IsSuccess() bool {
	return br.Success
}

func (br *BaseResponse) GetResult() interface{} {
	return br.Result
}

func (br *BaseResponse) GetErrorType() string {
	return br.ErrorType
}

func (br *BaseResponse) GetErrorMessage() string {
	return br.ErrorMessage
}
