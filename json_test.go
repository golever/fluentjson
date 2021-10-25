package fluentjson

import "testing"

func TestStringToArray(t *testing.T) {
	var s1 = "[]"
	var s2 = "[\"Sunday\", \"Monday\", \"Tuesday\", \"Wednesday\", \"Thursday\", \"Friday\", \"Saturday\"]"

	arr1, err := StringToArray(s1)
	if err != nil {
		t.Errorf("wrong")
	}
	if !arr1.IsEmpty() {
		t.Errorf("wrong len of arr1:%v", arr1)
	}

	arr2, err := StringToArray(s2)
	if err != nil {
		t.Errorf("wrong")
	}
	if arr2.Len() != 7 {
		t.Errorf("wrong len of arr2:%v", arr1)
	}
}

func TestStringToObject(t *testing.T) {
	var s1 = "{}"
	var s2 = "{\n  \"name\": \"foo\",\n  \"email\": \"bar@gmail.com\",\n  \"age\": 18\n}"

	obj1, err := StringToObject(s1)
	if err != nil {
		t.Errorf("wrong")
	}
	if !obj1.IsEmpty() {
		t.Errorf("wrong len of obj1:%v", obj1)
	}

	obj2, err := StringToObject(s2)
	if err != nil {
		t.Errorf("wrong")
	}
	v, ok := obj2["age"]
	if !ok {
		t.Errorf("must contais key age")
	}
	age, ok := v.(float64)
	if age != 18 {
		t.Errorf("age must be 18")
	}
}
