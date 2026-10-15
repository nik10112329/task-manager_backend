package apperror

import "encoding/json"

var (
	ErrorNotFound = &AppError{
		Err:              nil,
		Code:             "US-000003",
		Message:          "not found",
		DeveloperMessage: "",
	}
)

type AppError struct {
	Err              error  `json:"-"`
	Code             string `json:"code,omitempty"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
}

func (e *AppError) NewAppError(err error, code, message, developerMessage string) *AppError {
	return &AppError{
		Err:              err,
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) MarshalJSON() ([]byte, error) {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
