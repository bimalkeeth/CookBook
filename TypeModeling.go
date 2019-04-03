package main

import "fmt"

type Eraser interface {
	Erase() bool
}

type Cleaner interface {
	Clean() bool
}

type Destoyer interface {
	Cleaner
	Eraser
}

type PPS=**string
type Webcontroller struct{}

func(wc *Webcontroller)GetName()string{
	return "Web Controller"
}

type Indexer interface {
	Index()
}

type AppController struct{
	*Webcontroller
	Indexer
}

type IndexString string

func(hs IndexString) Index(){
	fmt.Println("Index Page")
}


func main() {
	ac:=new(AppController)
	fmt.Println(ac.Webcontroller.GetName())
	ac=&AppController{new(Webcontroller),IndexString("")}
	ac.Index()
}
