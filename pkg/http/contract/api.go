package contract

type APIResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Error   *Error      `json:"error,omitempty"`
	Success bool        `json:"success"`
}

type Error struct {
	Description string `json:"description,omitempty"`
}

func NewSuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
	}
}

func NewFailureResponse(description string) APIResponse {
	return APIResponse{
		Success: false,
		Error: &Error{
			Description: description,
		},
	}
}
