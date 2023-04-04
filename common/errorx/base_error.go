package errorx

const defaultCode = 1001

type CodeError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Flag    bool        `json:"flag"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewErrorData(code int, msg string, data interface{}) error {
	return &CodeError{Code: code, Message: msg, Data: data}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Message
}

type CodeErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Flag    bool        `json:"flag"`
}

func (e *CodeError) Info() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
		Data:    e.Data,
		Flag:    e.Flag,
	}

}
