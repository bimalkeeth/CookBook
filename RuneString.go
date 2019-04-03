package main

import "unsafe"

func ReadMemory(ptr unsafe.Pointer,size uintptr)[]byte{
	out:=make([]byte,size)
	for i:=range out{
		out[i]=*((*byte)(unsafe.Pointer(uintptr(ptr)+uintptr(i))))
	}
	return out
}
type goString struct{
	elements[]byte
	len int
}

func main() {
	//s:=[]byte("Hello world")
	//var stringExample="Hello world"
	//var anotherStringExmple="Hello world 2"
	//var goString=goString{s,12}
}
