package mytransform

import (
	"bytes"
	"encoding/binary"
	"math"
	"unsafe"
)

// IntToByte 数字转Byte
func IntToByte(data int) (ret []byte) {
	len := unsafe.Sizeof(data)
	ret = make([]byte, len)
	tmp := 0xff
	var index uint
	for index = 0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	}
	return ret
}

// ByteToInt Byte转int
func ByteToInt(data []byte) int {
	ret := 0
	len := len(data)
	var i uint
	for i = 0; i < uint(len); i++ {
		ret = ret | (int(data[i]) << (i * 8))
	}
	return ret
}

// ByteToFloat64 Byte转Float
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

// ByteToFloat32 Byte转Float32
func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

// BytesToString Bytes转字符串
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
