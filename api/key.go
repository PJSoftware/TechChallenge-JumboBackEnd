package api

import "net/http"

const validKey string = "special-key"

// IsKeyValid confirms that a valid api-key has been provided
// per the petstore.swagger.io page, "special-key" is what
// we are expecting
func IsKeyValid(h http.Header) bool {
	key := h.Get("api-key")
	return key == validKey
}
