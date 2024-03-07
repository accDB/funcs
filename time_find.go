package funcs

import (
	"errors"
	"fmt"
	"unsafe"
)

type TimeFindIn struct {
	Log string
}

func (t *TimeFindIn) From(args []string) error {
	if len(args) != 1 {
		return errors.New(fmt.Sprintf("error args len should be 1, but have:%d", len(args)))
	}
	t.Log = args[0]
	return nil
}

func (t *TimeFindIn) To() ([]string, error) {
	if len(t.Log) < 1 {
		return nil, errors.New("error in len should be more 0")
	}
	return []string{t.Log}, nil
}

type TimeFindOut struct {
	LogTime string
	Log     string
}

func (o *TimeFindOut) To() ([]byte, error) {
	return append(getHeader(o), getData(o)...), nil
}

func (o *TimeFindOut) From(d []byte) error {
	*o = *(*TimeFindOut)(unsafe.Pointer(&d[0]))
	len1 := len(o.LogTime)
	o.LogTime = unsafe.String(&d[32], len1)
	len1 += 32
	o.Log = unsafe.String(&d[len1], len(o.Log))
	return nil
}

func (o *TimeFindOut) ReUse() {
	o.LogTime = ""
	o.Log = ""
}

func getData(out *TimeFindOut) []byte {
	data := unsafe.Slice(unsafe.StringData(out.LogTime), len(out.LogTime))
	return append(data, unsafe.Slice(unsafe.StringData(out.Log), len(out.Log))...)
}

func getHeader(out *TimeFindOut) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(out)), 32)
}
