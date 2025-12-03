package aml

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Response represents an API response.
type Response struct {
	*http.Response
}

// newResponse creates a new Response.
func newResponse(httpResp *http.Response) *Response {
	return &Response{Response: httpResp}
}

// RawBody represents a raw body.
type RawBody struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
	Error  json.RawMessage `json:"errors"`
}

// ErrorDetail represents the detail of an error.
type ErrorDetail struct {
	Row          int    `json:"row"`
	PrimaryKey   string `json:"primaryKey"`
	ErrorMessage string `json:"errorMessage"`
}

// UploadBatchResponseData represents the data of an upload batch response.
type UploadBatchResponseData struct {
	BatchNo string `json:"batchNo"`
}

// NameCheckResponseData represents the data of a name check response.
type NameCheckResponseData struct {
	RCScore int `json:"RCScore"`
}

// QueryOverResponseData represents the data of a query over response.
type QueryOverResponseData struct {
	Remain int `json:"remain"`
	Over   int `json:"over"`
}

// UploadHandlerResponseData represents the data of an upload handler response.
type UploadHandlerResponseData struct {
	Message string `json:"message"`
	Rows    int    `json:"rows"`
}

// ErrorResponse represents a tapd error response.
type ErrorResponse struct {
	response *http.Response
	rawBody  *RawBody
	err      error
}

func (e *ErrorResponse) Error() string {
	if e.rawBody != nil {
		return fmt.Sprintf("status: %s, error: %s", e.rawBody.Status, string(e.rawBody.Error))
	}

	if e.response != nil {
		return fmt.Sprintf("status code: %d, err: %v", e.response.StatusCode, e.err)
	}

	return e.err.Error()
}

func (e *ErrorResponse) Unwrap() error {
	return e.err
}

func IsErrorResponse(err error) bool {
	var e *ErrorResponse
	return errors.As(err, &e)
}
