package main

import (
	"encoding/binary"
	"io"
	"reflect"
	"time"

	"github.com/ugorji/go/codec"
)

type FBitDecoder struct {
	handle *codec.MsgpackHandle
	mpdec  *codec.Decoder
}

type FLBTime struct {
	time.Time
}

func (f FLBTime) WriteExt(interface{}) []byte {
	panic("unsupported")
}

func (f FLBTime) ReadExt(i interface{}, b []byte) {
	out := i.(*FLBTime)
	sec := binary.BigEndian.Uint32(b)
	usec := binary.BigEndian.Uint32(b[4:])
	out.Time = time.Unix(int64(sec), int64(usec))
}

func (f FLBTime) ConvertExt(v interface{}) interface{} {
	return nil
}

func (f FLBTime) UpdateExt(dest interface{}, v interface{}) {
	panic("unsupported")
}

func NewDecoder(r io.Reader) *FBitDecoder {
	dec := new(FBitDecoder)
	dec.handle = new(codec.MsgpackHandle)
	dec.handle.RawToString = true
	dec.handle.SetExt(reflect.TypeOf(FLBTime{}), 0, &FLBTime{})
	dec.mpdec = codec.NewDecoder(r, dec.handle)

	return dec
}

func GetRecord(dec *FBitDecoder) (ret int, ts interface{}, rec map[interface{}]interface{}) {
	var m interface{}

	err := dec.mpdec.Decode(&m)
	if err != nil {
		return -1, 0, nil
	}

	slice := reflect.ValueOf(m)
	t := slice.Index(0).Interface()
	data := slice.Index(1)

	mapdata := data.Interface().(map[interface{}]interface{})

	return 0, t, mapdata
}
