package main

import (
	"fmt"

	"github.com/nlepage/golang-wasm/js/bind"
)

type Location struct {
	Hash     func() string `js:"hash"`
	Host     func() string `js:"host"`
	Hostname func() string `js:"hostname"`
	Href     func() string `js:"href"`
	// ...
}
type Element struct {
	TagName func() string `js:"tagName`
}
type Document struct {
	GetElementById func(string) Element `js:"getElementById`
}

type Window struct {
	Location func() Location `js:"location"`
}
type Global struct {
	Document func() Document `js:"document"`
}

func main() {
	window := &Window{}
	if err := bind.BindGlobals(window); err != nil {
		panic(err)
	}

	global := &Global{}
	if err := bind.BindGlobals(global); err != nil {
		panic(err)
	}
	fmt.Println(window.Location().Href())
	fmt.Println(global.Document().GetElementById("hello-world"))
	//.GetElementById("hello-world"))
}
