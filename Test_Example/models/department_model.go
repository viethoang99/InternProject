package models

import (
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

const BS_DEPART = "DEPARTMENT"

func (this *Department) String() string {
	return this.Department_id
}

func (this *Department) GetBsKey() generic.TStringKey{
	return generic.TStringKey(fmt.Sprint(BS_DEPART))
}

func (this *Department) GetAll() []Department {
	if totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey()); err == nil {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(), 0, int32(totalCount))
		if err != nil {
			return make([]Department, 0)
		}
		transaction, err := this.UnMarshalArrayTItem(slice)
		return transaction
	}
	return make([]Department, 0)
}

func (this *Department) AddDepartment() error{
	bDepartment, id, err := MarshalBytes(this)
	if err != nil {
		return err
	}
	bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   id,
		Value: bDepartment,
	})
	return nil
}
func (this *Department) Get() (interface{}, error) {

	bytes, err := this.GetItemBytes()

	if err != nil {
		return nil, err
	}

	return UnMarshalBytes(bytes)
}
func (this *Department) GetU(mapb string) (interface{}, error) {

	bytes, err := this.GetItemBytesD(mapb)

	if err != nil {
		return nil, err
	}

	return UnMarshalBytes(bytes)
}
func (this *Department) PutItem() error {
	bDepart, key, err := MarshalBytes(this)
	if err != nil {
		return err
	}
	_, err = this.Get()
	if err != nil {
		return err
	}
	return bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bDepart,
	})
}
func (this *Department) GetItemBytes() ([]byte, error) {
	var p *User
	tDepartment, err := bigsetIf.BsGetItem(p.GetBsKey(), generic.TItemKey(this.String()))
	if err != nil {
		return nil, err
	}
	return tDepartment.GetValue(), nil
}
func (this *Department) GetItemBytesD(mapb string) ([]byte, error) {
	var p *User
	tDepartment, err := bigsetIf.BsGetItem(p.GetBsKey(), generic.TItemKey(mapb))
	if err != nil {
		return nil, err
	}
	return tDepartment.GetValue(), nil
}
func (this *Department) UnMarshalArrayTItem(objects []*generic.TItem) ([]Department, error) {
	objs := make([]Department, 0)
	for _, object := range objects {
		obj := Department{}
		err := json.Unmarshal(object.GetValue(), &obj)

		if err != nil {
			return make([]Department, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}
func (this *Department) Delete() error {
	t:=this.Department_id
	return bigsetIf.BsRemoveItem(this.GetBsKey(), generic.TItemKey(t))
}

func (this *Department) UpdateUser(u User) string{
	user,_:=bigsetIf.BsGetItem(u.GetBsKey(), generic.TItemKey(u.User_id))
	obj:=User{}
	err:=json.Unmarshal(user.GetValue(),&obj)
	if err!=nil{
		return "INTERNAL SERVER ERROR"
	}
	obj.Department_id=u.Department_id
	bNewUser, key, err := MarshalBytes1(obj.User_id,obj)
	if err != nil {
		return "Loi"
	}
	if err != nil {
		return "Loi"
	}
	err1:= bigsetIf.BsPutItem(u.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bNewUser,
	})
	if err1!=nil{
		return "INTERNAL SERVER ERROR"
	}else{
		return "SUCCESS"
	}
}

