package main

import (
	"errors"
	"flag"
	"fmt"
)

var (
	PrivateKey = ""
	Master     = ""
	Port       = ""
)

var (
	usege              = "for help: ./agent -help"
	ERR_PK             = errors.New("illegal privateKey \n" + usege)
	ERR_MASTER_EMPTY   = errors.New("illegal master \n" + usege)
	ERR_MASTER_CONNECT = errors.New("can not connect master. please make sure the master is correct")
)

func init() {
	flag.StringVar(&PrivateKey, "pk", "", "private key")
	flag.StringVar(&Master, "master", "", "master address. eg.127.0.0.1:9020")
	flag.StringVar(&Port, "port", ":9021", "port")
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
		fmt.Println(ERR_MASTER_CONNECT.Error() + "\n" + Master + "\n" + err.Error())
		return
	}

	router()
}
