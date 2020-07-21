package models

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
)

const BS_Data = "QUANLY"
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
	log.Println("GetAll()")
	if totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey()); err == nil {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(), 0, int32(totalCount))
		log.Println("---- models/user_m.go:33")
		if err != nil {
			log.Println("---- models/user_m.go:35")
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

func (this *User) AddUser() error{
	bUser, id, err := MarshalBytes(this)
	if err != nil {
		return err
	}
	bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   id,
		Value: bUser,
	})
	return nil
}
func (this *User) Get() (interface{}, error) {

	bytes, err := this.GetItemBytes()

	if err != nil {
		return nil, err
	}

	return UnMarshalBytes(bytes)
}
func (this *User) PutItem() error {
	bProd, key, err := MarshalBytes(this)
	if err != nil {
		return err
	}
	_, err = this.Get()
	if err != nil {
		return err
	}
	return bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bProd,
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


