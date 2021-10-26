package fluentjson

import (
	"errors"
	"fmt"
)

var (
	ErrIndexOutOfRange  = errors.New("index out of range")
	ErrValueIsNotNumber = errors.New("value is not number")
	ErrValueIsNotBool   = errors.New("value is not bool")
	ErrValueIsNotTime   = errors.New("value is not time")
)

type ErrKeyNotFound struct {
	Key string
}

type ErrValueConv struct {
	Sou string
	Tar string
}

func (e ErrKeyNotFound) Error() string {
	return fmt.Sprintf("key[%s] not existed", e.Key)
}

func (e ErrValueConv) Error() string {
	return fmt.Sprintf("cannot transform %s into %s", e.Sou, e.Tar)
}
