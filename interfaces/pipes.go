package interfaces

import (
	"github.com/lunny/log"
	"io"
	"os"
)

func PipeExample()error{
	r,w:=io.Pipe()
	go func(){
		_,err:=w.Write([]byte("test\n"))
		if err!=nil{
			log.Error(err)
		}
		err=w.Close()
	}()
	if _,err:=io.Copy(os.Stdout,r);err!=nil{
		return err
	}
	return nil
}