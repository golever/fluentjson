package fluentjson

import (
	"testing"
)

func TestArray_SetNil(t *testing.T) {
	arr := Array{"1", 2, "3"}
	arr.SetNil(1)

	v, _ := arr.getValue(1)
	if v != nil {
		t.Error("value must be nil")
	}
}
