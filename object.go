package fluentjson

import (
	"fmt"
	"reflect"
	"strconv"
)

type Object map[string]interface{}

func (obj Object) Copy() Object {
	ret := Object{}
	for k, v := range obj {
		ret[k] = v
	}
	return ret
}

func (obj Object) PutValue(k string, v interface{}) Object {
	obj[k] = v
	return obj
}

func (obj Object) Remove(k string) Object {
	delete(obj, k)
	return obj
}

func (obj Object) getValue(k string) (interface{}, error) {
	v, ok := obj[k]
	if !ok {
		return nil, ErrKeyNotFound{k}
	}
	return v, nil
}

func (obj Object) getString(k string) (string, error) {
	v, err := obj.getValue(k)
	if err != nil {
		return "", err
	}
	switch v.(type) {
	case string:
		return v.(string), nil
	case fmt.Stringer:
		return (v.(fmt.Stringer)).String(), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}

func (obj Object) getInt(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getFloat(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getDouble(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getBool(k string) (bool, error) {
	v, err := obj.getValue(k)
	if err != nil {
		return false, err
	}
	switch v.(type) {
	case bool:
		return v.(bool), nil
	case string:
		return strconv.ParseBool(v.(string))
	default:
		return false, ErrValueIsNotBool
	}
}

func (obj Object) getTime(k string) (interface{}, error) {
	return obj, nil
}

func (obj Object) getObject(k string) (Object, error) {
	v, ok := obj[k]
	if !ok {
		return Object{}, nil
	}
	switch v.(type) {
	case Object:
		return v.(Object), nil
	case map[string]interface{}:
		return v.(map[string]interface{}), nil
	default:
		return nil, ErrValueConv{reflect.TypeOf(v).String(), "fluentjson.Object"}
	}
}

func (obj Object) getArray(k string) (Array, error) {
	v, ok := obj[k]
	if !ok {
		return Array{}, nil
	}
	switch v.(type) {
	case Array:
		return v.(Array), nil
	case []interface{}:
		return v.([]interface{}), nil
	}
	return nil, ErrValueConv{reflect.TypeOf(v).String(), "fluentjson.Array"}
}

func (obj Object) Len() int {
	return len(obj)
}

func (obj Object) IsEmpty() bool {
	return obj.Len() == 0
}

func (obj Object) Contains(k string) bool {
	_, ok := obj[k]
	return ok
}
