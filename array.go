package fluentjson

import (
	"errors"
	"fmt"
	"reflect"
)

type Array []interface{}

func NewArray() Array {
	return Array{}
}

func (arr Array) Add(v interface{}) Array {
	return append(arr, v)
}

func (arr Array) AddNil() Array {
	var v interface{}
	return append(arr, v)
}

func (arr Array) Set(idx int, v interface{}) Array {
	if idx >= arr.Len() && idx < 0 {
		return arr
	}
	arr[idx] = v
	return arr
}

func (arr Array) SetNil(idx int) Array {
	var v interface{}
	return arr.Set(idx, v)
}

func (arr Array) Remove(idx int) Array {
	if idx >= arr.Len() && idx < 0 {
		return arr
	}
	return append(arr[:idx], arr[idx+1:])
}

func (arr Array) Copy() Array {
	ret := Array{}
	copy(ret, arr)
	return ret
}

func (arr Array) getValue(idx int) (interface{}, error) {
	if idx >= arr.Len() && idx < 0 {
		return nil, ErrIndexOutOfRange
	}
	return arr[idx], nil
}

func (arr Array) getString(idx int) (string, error) {
	v, err := arr.getValue(idx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", v), nil
}

func (arr Array) getInt(idx int) (interface{}, error) {
	return arr, nil
}

func (arr Array) getFloat(idx int) (interface{}, error) {
	return arr, nil
}

func (arr Array) getDouble(idx int) (interface{}, error) {
	return arr, nil
}

func (arr Array) getBool(idx int) (interface{}, error) {
	return arr, nil
}

func (arr Array) getTime(idx int) (interface{}, error) {
	return arr, nil
}

func (arr Array) getObject(idx int) (Object, error) {
	v, err := arr.getValue(idx)
	if err != nil {
		return nil, err
	}
	switch v.(type) {
	case Object:
		return v.(Object), nil
	case map[string]interface{}:
		return v.(map[string]interface{}), nil
	}
	return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not fluentjson.Object", reflect.TypeOf(v)))
}

func (arr Array) getArray(idx int) (Array, error) {
	v, err := arr.getValue(idx)
	if err != nil {
		return nil, err
	}
	switch v.(type) {
	case Array:
		return v.(Array), nil
	case []interface{}:
		return v.([]interface{}), nil
	}
	return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not fluentjson.Array", reflect.TypeOf(v)))
}

func (arr Array) Len() int {
	return len(arr)
}

func (arr Array) IsEmpty() bool {
	return arr.Len() == 0
}
