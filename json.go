package fluentjson

import "encoding/json"

func StringToObject(s string) (Object, error) {
	return BytesToObject([]byte(s))
}

func BytesToObject(b []byte) (Object, error) {
	var obj = Object{}
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func StringToArray(s string) (Array, error) {
	return BytesToArray([]byte(s))
}

func BytesToArray(b []byte) (Array, error) {
	var arr = Array{}
	err := json.Unmarshal(b, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
