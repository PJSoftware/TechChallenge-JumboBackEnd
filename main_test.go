package main_test

import (
	"fmt"
	"os/exec"
	"testing"
)

const curlEXE string = "c:\\Windows\\System32\\curl.exe"
const server string = "http://localhost:8080"
const apiKey string = "special-key"

type testData struct {
	name   string
	url    string
	method string
	usekey bool
	form   string
	json   string
	want   string
}

var testTable = []testData{
	testData{"Reset Database", "/reset4test", "GET", true, "", "",
		`{"code":200,"type":"OK","message":"OK"}`},

	testData{"Bad Endpoint", "/", "GET", true, "", "",
		`{"code":400,"type":"Bad Request","message":"bad endpoint requested"}`},

	testData{"No API-key", "/pet/1", "GET", false, "", "",
		`{"code":400,"type":"Bad Request","message":"Invalid ID supplied"}`},

	testData{"Retrieve Pet 1", "/pet/1", "GET", true, "", "",
		`{"id":1,"category":{"id":1,"name":"Dog"},"name":"Fido","photoUrls":["fido.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":2,"name":"Fluffy"},{"id":3,"name":"Friendly"},{"id":4,"name":"Big"}],"status":"sold"}`},
	testData{"Retrieve Pet 3", "/pet/3", "GET", true, "", "",
		`{"id":3,"category":{"id":2,"name":"Cat"},"name":"Moggy","photoUrls":["moggy1.jpg","moggy2.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":2,"name":"Fluffy"},{"id":5,"name":"Small"}],"status":"available"}`},
	testData{"Retrieve Pet 5", "/pet/5", "GET", true, "", "",
		`{"id":5,"category":{"id":1,"name":"Dog"},"name":"Rover","photoUrls":["fido.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":2,"name":"Fluffy"},{"id":3,"name":"Friendly"}],"status":"available"}`},
	testData{"No Pet 0", "/pet/0", "GET", true, "", "",
		`{"code":404,"type":"Not Found","message":"Pet not found"}`},
	testData{"No Pet 6", "/pet/6", "GET", true, "", "",
		`{"code":404,"type":"Not Found","message":"Pet not found"}`},

	testData{"Change Name", "/pet/2", "POST", true, "name=Slimy", "",
		`{"id":2,"category":{"id":4,"name":"Snake"},"name":"Slimy","photoUrls":["severus.jpg"],"tags":[{"id":3,"name":"Friendly"},{"id":4,"name":"Big"}],"status":"pending"}`},
	testData{"Change Status", "/pet/2", "POST", true, "status=sold", "",
		`{"id":2,"category":{"id":4,"name":"Snake"},"name":"Slimy","photoUrls":["severus.jpg"],"tags":[{"id":3,"name":"Friendly"},{"id":4,"name":"Big"}],"status":"sold"}`},
	testData{"Unknown Status", "/pet/2", "POST", true, "status=xxx", "",
		`{"code":405,"type":"Method Not Allowed","message":"Invalid input: bad status value"}`},
	testData{"Change Both", "/pet/2", "POST", true, "name=Severus&status=pending", "",
		`{"id":2,"category":{"id":4,"name":"Snake"},"name":"Severus","photoUrls":["severus.jpg"],"tags":[{"id":3,"name":"Friendly"},{"id":4,"name":"Big"}],"status":"pending"}`},

	testData{"Retrieve Pet 4", "/pet/4", "GET", true, "", "",
		`{"id":4,"category":{"id":3,"name":"Gerbil"},"name":"Gerry","photoUrls":["gerry.jpg"],"tags":[{"id":2,"name":"Fluffy"},{"id":5,"name":"Small"}],"status":"available"}`},
	testData{"Delete Pet 4", "/pet/4", "DELETE", true, "", "",
		`{"code":200,"type":"OK","message":"pet deleted successfully"}`},
	testData{"No More 4", "/pet/4", "GET", true, "", "",
		`{"code":404,"type":"Not Found","message":"Pet not found"}`},

	testData{"Upload Via GET Fails", "/pet/2/uploadImage", "GET", true, "file=snake1.jpg", "",
		`{"code":405,"type":"Method Not Allowed","message":"only POST valid for image upload"}`},
	testData{"Upload Via PUT", "/pet/2/uploadImage", "POST", true, "file=snake1.jpg", "",
		`{"code":200,"type":"OK","message":"image uploaded successfully"}`},
	testData{"Confirm Upload", "/pet/2", "GET", true, "", "",
		`{"id":2,"category":{"id":4,"name":"Snake"},"name":"Severus","photoUrls":["severus.jpg","snake1.jpg"],"tags":[{"id":3,"name":"Friendly"},{"id":4,"name":"Big"}],"status":"pending"}`},

	testData{"Invalid Method", "/pet", "GET", true, "", "",
		`{"code":405,"type":"Method Not Allowed","message":"unrecognised method for processing pets"}`},

	testData{"New Pet No Data", "/pet", "POST", true, "", "",
		`{"code":405,"type":"Method Not Allowed","message":"invalid input"}`},
	testData{"New Pet With JSON", "/pet", "POST", true, "", `{"id":1,"category":{"id":1,"name":"Dog"},"name":"Good Boy","photoUrls":["doggy.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":2,"name":"Fluffy"},{"id":3,"name":"Friendly"}],"status":"available"}`,
		`{"code":200,"type":"OK","message":"pet added successfully"}`},
	testData{"Retrieve New Pet", "/pet/6", "GET", true, "", "",
		`{"id":6,"category":{"id":1,"name":"Dog"},"name":"Good Boy","photoUrls":["doggy.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":2,"name":"Fluffy"},{"id":3,"name":"Friendly"}],"status":"available"}`},

	testData{"Update Pet No Data", "/pet", "PUT", true, "", "",
		`{"code":405,"type":"Method Not Allowed","message":"invalid input"}`},
	testData{"Update Missing Pet With JSON", "/pet", "PUT", true, "", `{"id":7,"category":{"id":2,"name":"Cat"},"name":"Tigger","photoUrls":["tigger.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":5,"name":"Small"}],"status":"available"}`,
		`{"code":404,"type":"Not Found","message":"Pet not found"}`},
	testData{"Update Pet With JSON", "/pet", "PUT", true, "", `{"id":6,"category":{"id":2,"name":"Cat"},"name":"Tigger","photoUrls":["tigger.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":5,"name":"Small"}],"status":"available"}`,
		`{"code":200,"type":"OK","message":"pet updated successfully"}`},
	testData{"Retrieve Updated Pet", "/pet/6", "GET", true, "", "",
		`{"id":6,"category":{"id":2,"name":"Cat"},"name":"Tigger","photoUrls":["tigger.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":5,"name":"Small"}],"status":"available"}`},

	testData{"Search By Status Bad Method", "/pet/findByStatus", "PUT", true, "", "",
		`{"code":405,"type":"Method Not Allowed","message":"unrecognised method for processing pets"}`},
	testData{"Search By Status No Data", "/pet/findByStatus", "GET", true, "", "",
		`null`},
	testData{"Search By Status", "/pet/findByStatus", "GET", true, "status=sold", "",
		`[{"id":1,"category":{"id":1,"name":"Dog"},"name":"Fido","photoUrls":["fido.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":2,"name":"Fluffy"},{"id":3,"name":"Friendly"},{"id":4,"name":"Big"}],"status":"sold"}]`},
	testData{"Search By Status Mult", "/pet/findByStatus", "GET", true, "status=sold,pending", "",
		`[{"id":1,"category":{"id":1,"name":"Dog"},"name":"Fido","photoUrls":["fido.jpg"],"tags":[{"id":1,"name":"Cute"},{"id":2,"name":"Fluffy"},{"id":3,"name":"Friendly"},{"id":4,"name":"Big"}],"status":"sold"},{"id":2,"category":{"id":4,"name":"Snake"},"name":"Severus","photoUrls":["severus.jpg","snake1.jpg"],"tags":[{"id":3,"name":"Friendly"},{"id":4,"name":"Big"}],"status":"pending"}]`},
}

func TestPetEndPoint(t *testing.T) {
	for idx, tst := range testTable {
		result := runCurl(tst)
		if result != "" {
			t.Logf("Pet Test %d (%s):\n%s", idx, tst.name, result)
			t.Fail()
		}
	}
}

func runCurl(td testData) string {
	mth := td.method
	key := `api-key: ` + apiKey
	url := server + td.url
	data := td.form
	if data == "" {
		data = td.json
	}

	var cmd *exec.Cmd
	if td.usekey {
		if data != "" {
			if td.method == "GET" && td.form != "" {
				url = url + "?" + data
				cmd = exec.Command(curlEXE, "-X", mth, "-H", key, url)
			} else {
				if td.form == "" {
					cmd = exec.Command(curlEXE, "-X", mth, "-H", key, "-H", "Content-Type: application/json", "-d", data, url)
				} else {
					cmd = exec.Command(curlEXE, "-X", mth, "-H", key, "-d", data, url)
				}
			}
		} else {
			cmd = exec.Command(curlEXE, "-X", mth, "-H", key, url)
		}
	} else {
		cmd = exec.Command(curlEXE, "-X", mth, url)
	}

	bgot, err := cmd.Output()
	got := string(bgot)
	if err != nil {
		return "error running cURL:\n" + got + "\n" + err.Error()
	}

	if got != td.want {
		return fmt.Sprintf("expected: %s;\n     got: %s", td.want, got)
	}

	return ""
}
