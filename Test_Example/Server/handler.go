package main

import (
	"Test_Example/conf/gen-go/tutorial"
	"Test_Example/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
)

type ServiceHandler struct{
	log map[int]*tutorial.SharedStruct
}


func (p *ServiceHandler) Ping(ctx context.Context) (err error) {
	fmt.Print("ping()\n")
	return nil
}

func (p *ServiceHandler) GetAll(ctx context.Context) (r []*tutorial.User, err error) {
	if totalCount, err := models.GetBigsetIf().GetTotalCount(p.GetBsKey()); err == nil {
		slice, err := models.GetBigsetIf().BsGetSliceR(p.GetBsKey(), 0, int32(totalCount))
		if err != nil {
			return make([]*tutorial.User, 0),err
		}
		transaction, err := p.UnMarshalArrayTItem(slice)
		return transaction,nil
	}
	return make([]*tutorial.User, 0),nil
}

func (p *ServiceHandler) AddUser(ctx context.Context, u *tutorial.User) (r string, err error) {
	bUser, id, err := MarshalBytes(u.UserID,u)
	if err != nil {
		return "INTERNAL SERVER ERROR",err
	}
	err = models.GetBigsetIf().BsPutItem(p.GetBsKey(), &generic.TItem{
		Key:   id,
		Value: bUser,
	})
	if err!=nil{
		return "NOTIFICATION: ",err
	}else{
		return "Adding Success!",nil
	}
}

func (p *ServiceHandler) Put(ctx context.Context, u *tutorial.User) (r string, err error) {
	bProd, key, err := MarshalBytes(u.UserID,u)
	err1:=models.GetBigsetIf().BsPutItem(p.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bProd,
	})
	if err1!=nil{
		return "NOTIFICATION: ",err1
	}else{
		return "Update Success!",nil
	}
}

func (p *ServiceHandler) Delete(ctx context.Context, id string) (r string, err error) {
	err=models.GetBigsetIf().BsRemoveItem(p.GetBsKey(), generic.TItemKey(id))
	return "DELETE",err
}

func NewServiceHandler() *ServiceHandler{
	return &ServiceHandler{log:make(map[int]*tutorial.SharedStruct)}
}

func (p *ServiceHandler) GetBsKey() generic.TStringKey{
	return generic.TStringKey(fmt.Sprint(tutorial.BS_Data))
}

func (p *ServiceHandler) Get(id string) (interface{}, error) {

	bytes, err := p.GetItemBytes(id)

	if err != nil {
		return nil, err
	}

	return UnMarshalBytes(bytes)
}
func (p *ServiceHandler) UnMarshalArrayTItem(objects []*generic.TItem) ([]*tutorial.User, error) {
	objs := make([]*tutorial.User, 0)
	for _, object := range objects {
		obj := &tutorial.User{}
		err := json.Unmarshal(object.GetValue(), &obj)
		fmt.Printf("%v\n", string(object.GetKey()))
		if err != nil {
			return make([]*tutorial.User, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}
func MarshalBytes(id string,object interface{}) ([]byte, []byte, error) {
	var key []byte
	if object == nil {
		return nil, nil, errors.New("object must be not nil")
	}

	obj, err := json.Marshal(object)
	if err != nil {
		return nil, nil, err
	}

	key = []byte(id)
	test := fmt.Sprintf("%v", object)
	log.Println(test, "-- test")

	return obj, key, nil
}
func UnMarshalBytes(bytes []byte) (interface{}, error) {
	var obj interface{}

	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
func (p *ServiceHandler) GetItemBytes(id string) ([]byte, error) {
	tUser, err := models.GetBigsetIf().BsGetItem(p.GetBsKey(), generic.TItemKey(id))
	if err != nil {
		return nil, err
	}
	return tUser.GetValue(), nil
}
