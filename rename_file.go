package main

import (
	"log"
	"os"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

//使用方法：QINIU="../cfg.toml" go run rename_file.go

func main() {
	lister := operation.NewListerV2()
	if lister == nil {
		log.Println("load config file", os.Getenv("QINIU"), "failed")
	} else {
		fromKey, toKey := "fromKey", "toKey"
		err := lister.Rename(fromKey, toKey)
		if err != nil {
			log.Println("rename from", fromKey, "to", toKey, "fail:", err)
		} else {
			log.Println("rename from", fromKey, "to", toKey, "success")
		}
	}
}
