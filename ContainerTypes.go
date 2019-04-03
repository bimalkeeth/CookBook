package main

import "unsafe"

type MySlice struct {
	elems unsafe.Pointer
	len   int
	cap   int
}

const e = 3

func main() {
	var maping = make(map[int]int)
	maping[e] = e
}
