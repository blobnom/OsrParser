package OsrParser

import (
	"encoding/binary"
	"errors"
	"io"
)

// RBString reads an Osu! Encoded string of the Given io.Reader returns an String else Error
func rBString(data io.Reader) (s string, err error) {
	bufferSlice := make([]byte, 1)
	data.Read(bufferSlice)
	if bufferSlice[0] != 11 {
		return "", nil
	}
	length := ulebUnmarshalReader(data)
	bufferSlice = make([]byte, length)
	b, err := data.Read(bufferSlice)
	if b < length {
		err = errors.New("Unexpected end of string")
	}
	s = string(bufferSlice)
	return
}

// RInt Reads an Binary encoded int with the given io.Reader returns int else error
func rInt(data io.Reader) (i int, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RUInt Reads an Binary encoded unsigned Int and returns a uint else an error
func rUInt(data io.Reader) (i uint, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RInt8 Reads an Binary encoded int8 with the given io.Reader returns int8 else error
func rInt8(data io.Reader) (i int8, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RUInt8 Reads an Binary encoded unsigned Int8 and returns a uint8 else an error
func rUInt8(data io.Reader) (i uint8, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RInt16 Reads an Binary encoded int16 with the given io.Reader returns int16 else error
func rInt16(data io.Reader) (i int16, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RUInt16 Reads an Binary encoded unsigned Int16 and returns a uint16 else an error
func rUInt16(data io.Reader) (i uint16, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RInt32 Reads an Binary encoded int32 with the given io.Reader returns int32 else error
func rInt32(data io.Reader) (i int32, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RUInt32 Reads an Binary encoded unsigned Int32 and returns a uint32 else an error
func rUInt32(data io.Reader) (i uint32, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RInt64 Reads an Binary encoded int64 with the given io.Reader returns int64 else error
func rInt64(data io.Reader) (i int64, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RUInt64 Reads an Binary encoded unsigned Int64 and returns a uint64 else an error
func rUInt64(data io.Reader) (i uint64, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RFloat32 Reads an float32 of the given io.Reader, returns float32 else error
func rFloat32(data io.Reader) (i float32, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RFloat64 Reads an float64 of the given io.Reader, returns float64 else error
func rFloat64(data io.Reader) (i float64, err error) {
	err = binary.Read(data, binary.LittleEndian, &i)
	return
}

// RBool reads a Binary encoded boolean using int8, returns bool else error
func rBool(data io.Reader) (i bool, err error) {
	var m int8
	err = binary.Read(data, binary.LittleEndian, &m)
	i = m > 0
	return
}

func rSlice(data io.Reader, length int32) (s []byte, err error) {
	s = make([]byte, length)
	_, err = data.Read(s)
	return
}
