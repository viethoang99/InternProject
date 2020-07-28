package models

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

const BS_Data = "NHANVIEN"
type User struct {
	User_id string `json:"user_id" xml:"user_id"`
	Username string `json:"username" xml:"username"`
	Department_id string
}

type Department struct {
	Department_id string
	Name string
}

func (this *User) String() string {
	return this.User_id
}

func (this *User) GetBsKey() generic.TStringKey{
	return generic.TStringKey(fmt.Sprint(BS_Data))
}

func (this *User) GetAll() []User {
	if total, err := bigsetIf.GetTotalCount(this.GetBsKey()); err == nil {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(), 0, int32(total))
		if err != nil {
			return make([]User, 0)
		}
		transaction, err := this.UnMarshalArrayTItem(slice)
		if err != nil {
			println(err.Error(), "akjflaksjdf")
		}
		return transaction
	}
	return make([]User, 0)
}
var pagenumber =1
func (this *User) GetAll1(pagesize int) []User {
	if count, err := bigsetIf.GetTotalCount(this.GetBsKey()); err == nil {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(), int32((pagenumber-1)*pagesize), int32(pagenumber*pagesize))
		if err != nil {
			return make([]User, 0)
		}
		module:=int(count) % pagesize
		page:=int(count) / pagesize
		if module!=0{
			page=page+1
		}
		if pagenumber==page{
			pagenumber=1
		}else{
			pagenumber++
		}
		transaction, err := this.UnMarshalArrayTItem(slice)
		if err != nil {
			println(err.Error(), "akjflaksjdf")
		}
		return transaction
	}

	return make([]User, 0)
}
var dep Department
func (this *User) AddUser() string{

	depart:= dep.GetAll()
	num:=0
	for i:=0;i<len(depart);i++{
		if depart[i].Department_id==this.Department_id {
			num = 1
			break
		}
	}
	if num==0{
		return "Khong ton tai phong ban"
	}
	bUser, id, err := MarshalBytes(this)
	if err != nil {
		return "INTERNAL SERVER ERROR"
	}
	err =bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   id,
		Value: bUser,
	})
	if err==nil{
		return "Them moi thanh cong"
	}else{
		return "INTERNAL SERVER ERROR"
	}
}
func (this *User) Get() (interface{}, error) {

	bytes, err := this.GetItemBytes()

	if err != nil {
		return nil, err
	}

	return UnMarshalBytes(bytes)
}
func (this *User) PutItem() error {
	bUser, key, err := MarshalBytes(this)
	if err != nil {
		return err
	}
	_, err = this.Get()
	if err != nil {
		return err
	}
	return bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bUser,
	})
}
func (this *User) GetItemBytes() ([]byte, error) {
	tUser, err := bigsetIf.BsGetItem(this.GetBsKey(), generic.TItemKey(this.String()))
	if err != nil {
		return nil, err
	}
	return tUser.GetValue(), nil
}
func UnMarshalBytes(bytes []byte) (interface{}, error) {
	var obj interface{}

	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
func (this *User) UnMarshalArrayTItem(objects []*generic.TItem) ([]User, error) {
	objs := make([]User, 0)
	for _, object := range objects {
		obj := User{}
		err := json.Unmarshal(object.GetValue(), &obj)

		if err != nil {
			return make([]User, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}
func (this *User) Delete() error {
	return bigsetIf.BsRemoveItem(this.GetBsKey(), generic.TItemKey(this.User_id))
}


