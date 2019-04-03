package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type callBackChan chan struct{}

func CheckEvery(ctx context.Context, d time.Duration,cb callBackChan){
	for{
		select {
		   case <-ctx.Done():
			   return
		   case <-time.After(d):
		   	if cb!=nil{
		   		cb <- struct {

				}{}
			}
		}
	}
}

func PrintProcesslist(){
	psCommand:=exec.Command("ps","a")
	resp,err:=psCommand.CombinedOutput()
	if err!=nil{
		log.Fatal("error: ps command failed")
	}
	out:=string(resp)
	lines:=strings.Split(out,"\n")
	for _,line:=range lines{
		if line!=""{
			fmt.Println(line)
		}
	}
}

func main() {
	ctx:=context.Background()
	PrintProcesslist()
	callback:=make(callBackChan)
	go CheckEvery(ctx,3 * time.Second,callback)
	go func(){
		for{
			select {
			   case <- callback:
			   	PrintProcesslist()
			}
		}
	}()
	for{
		time.Sleep(10* time.Microsecond)
	}
}
