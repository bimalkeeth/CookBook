package main

import (
	"github.com/gopherjs/jsbuiltin"
	"honnef.co/go/js/dom"
)

func main() {

}

func builtindemo(element dom.Element) {
	if jsbuiltin.TypeOf(element) == "object" {
		println("using type of operator")
	}
}
