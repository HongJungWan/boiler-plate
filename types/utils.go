package types

type ApiResponse struct {
	Result      int64  `json:"result"`
	Description string `json:"description"`
}

func NewApiResponse(description string, result int64) *ApiResponse {
	return &ApiResponse{
		Description: description,
		Result:      result,
	}
}
