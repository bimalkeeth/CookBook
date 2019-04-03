package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type LowerCaseReader struct{
	text string
}
func NewLowerCaseReader(text string ) *LowerCaseReader {
	return &LowerCaseReader{text: text}
}

func(r *LowerCaseReader)Read(p []byte)(int,error){
	buf:=make([]byte,len(r.text))
	for i:=0;i<len(buf);i++{
		buf[i]=r.text[i]| 0x20
	}
	n:=copy(p,buf)
	return n,io.EOF
}


func main() {

	r:=NewLowerCaseReader("ALL CAPITALS")
	resp,err:=ioutil.ReadAll(r)
	if err!=nil{
		log.Fatal("error")
	}
	fmt.Println(string(resp))
}
