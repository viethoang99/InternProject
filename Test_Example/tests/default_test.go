package testing

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/httplib"
)

var port = ""
var baseURL = "http://localhost:"

// TestHTTPRequest beego test request client
type TestHTTPRequest struct {
	httplib.BeegoHTTPRequest
}

func getPort() string {
	if port == "" {
		config, err := config.NewConfig("ini", "/home/hoangviet/go/src/Test_Example/conf/app.conf")
		if err != nil {
			return "8081"
		}
		port = config.String("httpport")
		return port
	}
	return port
}

// Get returns test client in GET method
func Get(path string) *TestHTTPRequest {
	return &TestHTTPRequest{*httplib.Get(baseURL + getPort() + path)}
}

// Post returns test client in POST method
func Post(path string) *TestHTTPRequest {
	return &TestHTTPRequest{*httplib.Post(baseURL + getPort() + path)}
}

// Put returns test client in PUT method
func Put(path string) *TestHTTPRequest {
	return &TestHTTPRequest{*httplib.Put(baseURL + getPort() + path)}
}

// Delete returns test client in DELETE method
func Delete(path string) *TestHTTPRequest {
	return &TestHTTPRequest{*httplib.Delete(baseURL + getPort() + path)}
}
