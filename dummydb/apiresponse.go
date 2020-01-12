package dummydb

// APIResponse encodes our response types
type APIResponse struct {
	Code         int32  `json:"code"`
	ResponseType string `json:"type"`
	Message      string `json:"message"`
}

var apiResponseTBL = []APIResponse{
	APIResponse{200, "OK", "OK"},
	APIResponse{201, "Created", "Created"},
	APIResponse{202, "Accepted", "Accepted"},
	APIResponse{204, "No Content", "No Content"},
	APIResponse{400, "Bad Request", "Bad Request"},
	APIResponse{401, "Unauthorized", "Unauthorized"},
	APIResponse{403, "Forbidden", "Forbidden"},
	APIResponse{404, "Not Found", "Not Found"},
	APIResponse{405, "Method Not Allowed", "Method Not Allowed"},
	APIResponse{406, "Not Acceptable", "Not Acceptable"},
	APIResponse{412, "Precondition Failed", "Precondition Failed"},
	APIResponse{415, "Unsupported Media Type", "Unsupported Media Type"},
	APIResponse{500, "Internal Server Error", "Internal Server Error"},
	APIResponse{501, "Not Implemented", "Not Implemented"},
}
