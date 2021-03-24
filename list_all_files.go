package main

import (
	"fmt"
	"log"
	"os"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

//使用方法：QINIU="../cfg.toml" go run list_public_cloud_files.go

func main() {

	lister := operation.NewListerV2()
	if lister == nil {
		log.Println("load config file", os.Getenv("QINIU"), "failed")
	} else {
		a := lister.ListPrefix("")
		for _, v := range a {
			fmt.Println(v)
		}
	}
}
