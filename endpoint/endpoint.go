package endpoint

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"../apiresponse"
)

// Request holds the results of the call to Parser()
type Request struct {
	endpoint string
	suburl   string
	method   string
	w        http.ResponseWriter
	req      *http.Request
}

// Parser determines which endpoint is being communicated with
func (r *Request) Parser(w http.ResponseWriter, req *http.Request) {
	r.reset()
	r.w = w
	r.req = req
	r.w.Header().Set("Content-Type", "application/json")

	extractEP := regexp.MustCompile("^/([a-z]+)(/.*)?$")
	matches := extractEP.FindStringSubmatch(req.RequestURI)
	if len(matches) > 0 {
		r.endpoint = matches[1]
		r.suburl = matches[2]
		r.method = req.Method
	}

	switch r.endpoint {
	case "pets":
		r.handlePets()
	case "store":
		r.handleStore()
	case "user":
		r.handleUser()
	default:
		r.status(400, "bad endpoint requested")
	}
}

func (r *Request) reset() {
	r.endpoint = ""
	r.suburl = ""
	r.method = ""
}

func (r *Request) status(code int32, desc string) {
	resp := apiresponse.ByCode(code, desc)
	j, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		r.w.WriteHeader(500)
		msg := fmt.Sprintf("unknown error occurred while processing status code %d, '%s'", resp.Code, resp.ResponseType)
		js := fmt.Sprintf(`{"code":500,"type":"Internal Server Error","message":"%s"}`, msg)
		r.w.Write([]byte(js))
		return
	}
	r.w.WriteHeader(int(code))
	r.w.Write(j)
}
