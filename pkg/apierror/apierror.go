package apierror

type APIError struct {
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Context    map[string]interface{} `json:"context,omitempty"`
	StatusCode int                    `json:"status_code"`
}

func (e *APIError) HTTPStatusCode() int {
	return e.StatusCode
}

func (e *APIError) Error() string {
	return e.Message
}
