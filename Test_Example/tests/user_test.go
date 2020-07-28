package testing

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"log"
	"testing"
)

func TestApiGet(t *testing.T) {
	result:=httplib.Get("http://localhost:8080/v1/user/")
	r,err:=result.Response()
	if err != nil {
		log.Fatal(err)
	}
	bData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(bData))
}

func TestApiPost(t *testing.T) {
	req := httplib.Post("http://localhost:8080/v1/user/")
	params :=map[string]string{"user_id": "17150212","username": "Hoang","Department_ID": "2"}
	bData, err := json.Marshal(params)
	req.Body(bData)
	resp, err := req.Response()
	if err == nil {
		bData, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(string(bData))
	}
}

func TestApiPut(t *testing.T){
	req := httplib.Post("http://localhost:8080/v1/user/")
	params :=map[string]string{"user_id": "17150212","username": "Hoang","Department_ID": "3"}
	bData, err := json.Marshal(params)
	req.Body(bData)
	resp, err := req.Response()
	if err == nil {
		bData, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(string(bData))
	}
}

func TestApiDelete(t *testing.T){
	result:=httplib.Delete("http://localhost:8080/v1/user/")
	result.Param("userid","17150212")
	r,err:=result.Response()
	if err != nil {
		log.Fatal(err)
	}
	bData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(bData))
}
