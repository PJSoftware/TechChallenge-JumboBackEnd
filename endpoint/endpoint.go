package endpoint

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"../api"
	"../apiresponse"
	"../dummydb"
)

// Request holds the results of the call to Parser()
type Request struct {
	endpoint string
	data     map[string]string
	suburl   string
	method   string
	w        http.ResponseWriter
	req      *http.Request
	validKey bool
}

// Parse determines which endpoint is being communicated with
func (r *Request) Parse(w http.ResponseWriter, req *http.Request) {
	r.reset()
	r.w = w
	r.req = req

	r.validKey = api.IsKeyValid(r.req.Header)
	r.w.Header().Set("Content-Type", "application/json")

	extractEP := regexp.MustCompile("^/([a-z0-9]*)(/[^?]+)?([?].+)?$")

	matches := extractEP.FindStringSubmatch(req.RequestURI)
	if len(matches) > 0 {
		r.endpoint = matches[1]
		r.suburl = matches[2]
		if matches[3] != "" {
			txt := strings.TrimPrefix(matches[3], "?")
			for _, pair := range strings.Split(txt, "&") {
				kv := strings.Split(pair, "=")
				r.data[kv[0]] = kv[1]
			}
		}
		r.method = req.Method
	} else {
		r.status(http.StatusInternalServerError, "error parsing url")
		return
	}

	switch r.endpoint {
	case "reset4test":
		dummydb.InitDB()
		r.status(http.StatusOK, "")
	case "pet":
		r.handlePets()
	case "store":
		r.handleStore()
	case "user":
		r.handleUser()
	default:
		r.status(http.StatusBadRequest, "bad endpoint requested")
	}
}

func (r *Request) reset() {
	r.endpoint = ""
	r.suburl = ""
	r.method = ""
	r.data = make(map[string]string)
}

func (r *Request) status(code int32, desc string) {
	resp := apiresponse.ByCode(code, desc)
	j, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		r.w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("error '%s' occurred while processing status code %d, '%s'", err.Error(), resp.Code, resp.ResponseType)
		js := fmt.Sprintf(`{"code":%d,"type":"Internal Server Error","message":"%s"}`, http.StatusInternalServerError, msg)
		r.w.Write([]byte(js))
		return
	}
	r.w.WriteHeader(int(code))
	r.w.Write(j)
}

func (r *Request) jsonOut(v interface{}) []byte {
	j, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		r.w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("error '%s' occurred while processing output '%v'", err.Error(), v)
		js := fmt.Sprintf(`{"code":%d,"type":"Internal Server Error","message":"%s"}`, http.StatusInternalServerError, msg)
		r.w.Write([]byte(js))
		return []byte("")
	}
	return j
}

func (r *Request) extractJSON() *json.Decoder {
	if r.req.Header.Get("Content-Type") == "application/json" {
		return json.NewDecoder(r.req.Body)
	}
	return nil
}
