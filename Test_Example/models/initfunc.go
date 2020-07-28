package models
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
)

func MarshalArrayBytes(objects interface{}) ([][]byte, [][]byte, error) {
	var objs, keys [][]byte
	if reflect.ValueOf(objects).Kind().String() != "slice" {
		return nil, nil, errors.New("kind of objects doesn't match slice type")
	}

	for _, object := range objects.([]interface{}) {
		marshal, err := json.Marshal(object)
		if err != nil {
			return nil, nil, err
		}

		objs = append(objs, marshal)
		keys = append(keys, []byte(fmt.Sprintf("%v", object)))
	}

	return objs, keys, nil
}

func GetTItems(objects interface{}) ([][]byte, [][]byte, error) {
	var objs, keys [][]byte
	if reflect.ValueOf(objects).Kind().String() != "slice" {
		return nil, nil, errors.New("kind of objects doesn't match slice type")
	}

	for _, object := range objects.([]interface{}) {
		marshal, err := json.Marshal(object)
		if err != nil {
			return nil, nil, err
		}

		objs = append(objs, marshal)
		keys = append(keys, []byte(fmt.Sprintf("%v", object)))
	}

	return objs, keys, nil
}

func MarshalBytes(object interface{}) ([]byte, []byte, error) {
	var key []byte
	if object == nil {
		return nil, nil, errors.New("object must be not nil")
	}

	obj, err := json.Marshal(object)
	if err != nil {
		return nil, nil, err
	}

	key = []byte(fmt.Sprintf("%v", object))
	test := fmt.Sprintf("%v", object)
	log.Println(test, "-- test")

	return obj, key, nil
}
func MarshalBytes1(id string,object interface{}) ([]byte, []byte, error) {
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
