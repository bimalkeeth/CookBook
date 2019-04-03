package main

import "fmt"

type Helper=interface {
	Help() string
}

type HelpString string

func(h HelpString) Help() string{
	return string(h)
}

type UnHelpString struct{}

func(uhs * UnHelpString)Help()string{
	return "I cannot help you"
}

var _=Helper(HelpString(""))

func main() {
	var h Helper=HelpString("Help me")
	fmt.Println(h.Help())
	var explicit=interface{Help()string}.Help(h)
	fmt.Println(explicit)
}
