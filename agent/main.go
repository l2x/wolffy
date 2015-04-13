package main

import (
	"errors"
	"flag"
	"fmt"
)

var (
	PrivateKey = ""
	Master     = ""
)

var (
	usege              = "usage: ./agent -pk=privateKey -master=master"
	ERR_PK             = errors.New("illegal privateKey \n " + usege)
	ERR_MASTER_EMPTY   = errors.New("illegal master \n " + usege)
	ERR_MASTER_CONNECT = errors.New(fmt.Sprintf("can not connect master[%s] \n please make sure the master is correct", Master))
)

func init() {
	flag.StringVar(&PrivateKey, "pk", "", "private key")
	flag.StringVar(&Master, "master", "", "master ip")
	flag.Parse()
}

func main() {
	if PrivateKey == "" {
		fmt.Println(ERR_PK.Error())
		return
	}

	if Master == "" {
		fmt.Println(ERR_MASTER_EMPTY.Error())
		return
	}

	err := report()
	if err != nil {
		fmt.Println(ERR_MASTER_CONNECT.Error())
		return
	}

	router()
}
