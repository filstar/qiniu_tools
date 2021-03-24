package main

import (
	"log"
	"os"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
	"github.com/qiniupd/qiniu-go-sdk/x/log.v7"
)

//使用方法：QINIU="../cfg.toml" go run rename_file.go

func main() {
	uploader := operation.NewUploaderV2()
	if uploader == nil {
		log.Println("load config file", os.Getenv("QINIU"), "failed")
	}

	fromFile := os.Args[1]
	toKey := fromFile[1:]
	log.Println("fromFile: ", fromFile, "  toKey: ", toKey)

	err := uploader.Upload(fromFile, toKey)

	if err != nil {
		log.Println("upload from", fromFile, "to", toKey, "fail:", err)
	} else {
		log.Println("upload from", fromFile, "to", toKey, "success")
	}
}
