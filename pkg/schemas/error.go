package schemas

// ErrorResponse represents a human-readable error response for API endpoints.
type ErrorResponse struct {
	// Message provides a brief, user-friendly description of the error.
	Message string `json:"message"`
	// HelpTexts contains additional context or suggestions to help the user
	// understand or resolve the error.
	HelpTexts []string `json:"helpTexts,omitempty"`
}

func NewErrorResponse(message string, helpTexts []string) ErrorResponse {
	return ErrorResponse{Message: message, HelpTexts: helpTexts}
}
