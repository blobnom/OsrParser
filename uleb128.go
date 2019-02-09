package OsrParser

import (
	"io"
)

//UlebMarshal convert int to Uleb byte array
func ulebMarshal(i int) (r []byte) {
	var len int
	if i == 0 {
		r = []byte{0}
		return
	}

	for i > 0 {
		r = append(r, 0)
		r[len] = byte(i & 0x7F)
		i >>= 7
		if i != 0 {
			r[len] |= 0x80
		}
		len++
	}
	return
}

//UlebUnmarshal convert Uleb byte array to int
func ulebUnmarshal(r []byte) (total int, len int) {
	var shift uint
	for {
		b := r[len]
		len++
		total |= (int(b&0x7F) << shift)
		if (b & 0x80) == 0 {
			break
		}
		shift += 7
	}
	return
}

//ulebUnmarshalReader converts io.reader to int
func ulebUnmarshalReader(r io.Reader) (total int) {
	var shift uint
	var lastByte byte

	for {
		b := make([]byte, 1)
		r.Read(b)
		lastByte = b[0]
		total |= (int(lastByte&0x7F) << shift)
		if (lastByte & 0x80) == 0 {
			break
		}
		shift += 7
	}
	return
}
