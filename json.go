package fluentjson

import "encoding/json"

func StringToObject(s string) (Object, error) {
	err := json.Unmarshal([]byte(s), Object{})
	return Object{}, err
}

func BytesToObject(b []byte) (Object, error) {
	err := json.Unmarshal(b, Object{})
	return Object{}, err
}

func StringToArray(s string) (Array, error) {
	err := json.Unmarshal([]byte(s), Object{})
	return Array{}.Add(Object{}), err
}

func BytesToArray(b []byte) (Array, error) {
	err := json.Unmarshal(b, Object{})
	return Array{}.Add(Object{}), err
}
