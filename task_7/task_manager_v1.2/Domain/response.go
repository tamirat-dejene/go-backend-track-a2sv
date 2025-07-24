package domain

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorRespone struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
