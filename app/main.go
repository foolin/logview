package main

import (
	"os"
	"log"
	"github.com/foolin/logview"
)

func main()  {
	if len(os.Args) <= 1{
		log.Fatalf("Please use command and args. such: logview tail -f demo.log")
		return
	}
	err := logview.RunCmd(os.Args[1], os.Args[2:]...)
	if err != nil {
		log.Fatalf("run cmd:%v, error: %v", os.Args[1:], err)
	}
}


