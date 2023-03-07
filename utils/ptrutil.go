package utils

import (
	"unsafe"
)

func ByteBufPtr(str []byte) uintptr {
	return uintptr(unsafe.Pointer(&str[0]))
}
func IntPtr(in int) uintptr {
	return uintptr(in)
}
func StringVal(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	var vbyte []byte

	for i := 0; ; i++ {
		sbyte := *((*byte)(unsafe.Pointer(ptr)))
		if sbyte == 0 {
			break
		}
		vbyte = append(vbyte, sbyte&0xff)
		ptr += 1
	}
	return string(vbyte)
}

func BytesVal(ptr uintptr, len int) []byte {
	var vbyte []byte

	for i := 0; i < len; i++ {
		sbyte := *((*byte)(unsafe.Pointer(ptr)))
		if sbyte == 0 {
			break
		}
		vbyte = append(vbyte, sbyte&0xff)
		ptr += 1
	}
	return vbyte
	//return *(*[]byte)(unsafe.Pointer(ptr))
}
func IntVal(ptr uintptr) int {
	return *(*int)(unsafe.Pointer(ptr))
}
func UIntVal(ptr uintptr) uint {
	return *(*uint)(unsafe.Pointer(ptr))
}

func Int32Val(ptr uintptr) int32 {
	return *(*int32)(unsafe.Pointer(ptr))
}

func UInt32Val(ptr uintptr) uint32 {
	return *(*uint32)(unsafe.Pointer(ptr))
}

func BoolVal(ptr uintptr) bool {
	return *(*bool)(unsafe.Pointer(ptr))
}
