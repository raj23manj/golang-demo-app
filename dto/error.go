package dto

type DtoErrorResponse struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

func (e *DtoErrorResponse) GetMessage() string {
	return e.Message
}
