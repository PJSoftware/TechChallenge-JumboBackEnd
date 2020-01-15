package apiresponse

import "net/http"

// APIResponse encodes our response types
type APIResponse struct {
	Code         int32  `json:"code"`
	ResponseType string `json:"type"`
	Message      string `json:"message"`
}

var apiResponseTBL []APIResponse
var responseTableInitialized bool

// initResponses is called internally when required
// It sets up json structures for responses we're likely to use
// Only runs once, so no checking required for existing codes
func initResponses() {
	if !responseTableInitialized {
		codes := []int32{
			http.StatusOK,                  // 200
			http.StatusBadRequest,          // 400
			http.StatusNotFound,            // 404
			http.StatusMethodNotAllowed,    // 405
			http.StatusInternalServerError, // 500
			http.StatusNotImplemented,      // 501
		}

		for _, code := range codes {
			initResponse(code)
		}
	}
	responseTableInitialized = true
}

func initResponse(code int32) {
	resp := new(APIResponse)
	resp.Code = code
	resp.ResponseType = http.StatusText(int(code))
	resp.Message = resp.ResponseType
	apiResponseTBL = append(apiResponseTBL, *resp)
}

// ByCode returns pointer to Pet with specified ID
func ByCode(code int32, msg string) *APIResponse {
	initResponses()
	for _, r := range apiResponseTBL {
		if r.Code == code {
			if msg != "" {
				r.Message = msg
			} else {
				r.Message = r.ResponseType
			}
			return &r
		}
	}
	return nil
}
