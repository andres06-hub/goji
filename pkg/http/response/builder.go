package resp

import (
	"net/http"
	"time"
)

// APIResponse represents the structure of a typical API response.
// It includes fields for indicating success, a message, a status code,
// a timestamp, and any additional data.
type APIResponse struct {
	Success   bool        `json:"success"`   // Indicates if the API call was successful
	Message   string      `json:"message"`   // A message providing additional information about the response
	Code      int         `json:"code"`      // The HTTP status code of the response
	Timestamp int64       `json:"timestamp"` // The timestamp when the response was generated
	Data      interface{} `json:"data"`      // Any additional data included in the response
}

// NewAPISuccessResponse creates a new APIResponse indicating a successful API call.
// It sets the Success field to true, the Code field to HTTP 200 OK, and the Timestamp field to the current time.
func NewAPISuccessResponse() *APIResponse {
	return &APIResponse{
		Success:   true,
		Code:      http.StatusOK,
		Timestamp: time.Now().Unix(),
	}
}

// NewAPIErrorResponse creates a new APIResponse indicating a failed API call.
// It sets the Success field to false, the Code field to HTTP 500 Internal Server Error,
// the Timestamp field to the current time, and the Data field to nil.
func NewAPIErrorResponse() *APIResponse {
	return &APIResponse{
		Success:   false,
		Code:      http.StatusInternalServerError,
		Timestamp: time.Now().Unix(),
		Data:      nil,
	}
}

// SetData sets the Data field of the APIResponse and returns the updated APIResponse.
// This method allows for chaining multiple setters together.
func (a *APIResponse) SetData(data interface{}) *APIResponse {
	a.Data = data
	return a
}

// WithMessage sets the Message field of the APIResponse and returns the updated APIResponse.
// This method allows for chaining multiple setters together.
func (a *APIResponse) WithMessage(message string) *APIResponse {
	a.Message = message
	return a
}

// WithCode sets the Code field of the APIResponse and returns the updated APIResponse.
// This method allows for chaining multiple setters together.
func (a *APIResponse) WithCode(code int) *APIResponse {
	a.Code = code
	return a
}

// WithTimestamp sets the Timestamp field of the APIResponse and returns the updated APIResponse.
// This method allows for chaining multiple setters together.
func (a *APIResponse) WithTimestamp(timestamp int64) *APIResponse {
	a.Timestamp = timestamp
	return a
}

// WithSuccess sets the Success field of the APIResponse and returns the updated APIResponse.
// This method allows for chaining multiple setters together.
func (a *APIResponse) WithSuccess(success bool) *APIResponse {
	a.Success = success
	return a
}

// Build sends the APIResponse as a JSON response using the provided http.ResponseWriter.
// This method finalizes the response and writes it to the client.
func (a *APIResponse) Build(ctx http.ResponseWriter) {
	xJSONResponse(ctx, a)
}
