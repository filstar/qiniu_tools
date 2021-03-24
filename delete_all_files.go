package main

import (
	"log"
	"os"
	"time"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

//使用方法：QINIU="../cfg.toml" go run delete_all_files.go

func main() {
	lister := operation.NewListerV2()
	if lister == nil {
		log.Println("load config file", os.Getenv("QINIU"), "failed")
	}
	a := lister.ListPrefix("")
	for _, v := range a {
		go func(v string) {
			err := lister.Delete(v)
			log.Println(v, err)
		}(v)
	}
	time.Sleep(time.Second * 10)
}
