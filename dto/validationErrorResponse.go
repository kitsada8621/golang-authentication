package dto

type ValidationErrorResponse struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
