package testing

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"log"
	"testing"
)

func TestApiGetDep(t *testing.T) {
	result:=httplib.Get("http://localhost:8080/v1/department/")
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

func TestApiPostDep(t *testing.T) {
	req := httplib.Post("http://localhost:8080/v1/department/")
	params :=map[string]string{"Department_id": "5","Name":"Nhan su"}
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
//
func TestApiPutDep(t *testing.T){
	req := httplib.Post("http://localhost:8080/v1/department/")
	params :=map[string]string{"Department_id": "5","Name": "Ky Thuat"}
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

func TestApiDeleteDep(t *testing.T){
	result:=httplib.Delete("http://localhost:8080/v1/department/")
	result.Param("Departmentid","5")
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

func TestApiGetSameDep(t *testing.T){
	result:=httplib.Get("http://localhost:8080/v1/department/2")
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


func TestApiGetUserSameDep(t *testing.T){
	result:=httplib.Get("http://localhost:8080/v1/department/1")
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

