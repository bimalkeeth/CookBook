package main

import (
	int "CookBook/interfaces"
	"bytes"
	"fmt"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := int.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer =", out.String())
	fmt.Print("stdout on PipeExample = ")
	if err := int.PipeExample(); err != nil {
		panic(err)
	}
}
