package fluentjson

import (
	"testing"
)

func Test_json(t *testing.T) {
	var s = "{'name':'z', 'age': 18}"

	obj, _ := StringToObject(s)
	t.Log(obj)

	obj2, _ := StringToArray(s)
	t.Log(obj2)
}
