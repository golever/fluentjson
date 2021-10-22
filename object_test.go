package fluentjson

import "testing"

func TestObject_Size(t *testing.T) {
	obj := Object{
		"first": "first",
	}
	if obj.Len() != 1 {
		t.Errorf("except size %d", obj.Len())
	}

	obj.PutValue("second", "second")
	if obj.Len() != 2 {
		t.Errorf("except size %d", obj.Len())
	}

	obj.PutValue("second", "second-1")
	if obj.Len() != 2 {
		t.Errorf("except size %d", obj.Len())
	}

	obj.Remove("second")
	if obj.Len() != 1 {
		t.Errorf("except size %d", obj.Len())
	}

}
