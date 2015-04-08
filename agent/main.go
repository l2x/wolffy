package main

import (
	"flag"
	"fmt"
)

var (
	PrivateKey = ""
)

func init() {
	flag.StringVar(&PrivateKey, "pk", "", "help message for flagname")
	flag.Parse()
}

func main() {
	if PrivateKey == "" {
		fmt.Println("illegal privateKey")
		fmt.Println("usage: ./agent -pk=privateKey")
		return
	}

	router()
}
