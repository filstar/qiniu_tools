package main

import (
	"fmt"
	"log"
	"os"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

//使用方法：QINIU="../cfg.toml" go run download_file.go

func main() {
	downloader := operation.NewDownloaderV2()
	if downloader == nil {
		log.Println("load config file", os.Getenv("QINIU"), "failed")
	} else {
		fromKey := os.Args[1]
		toFile := os.Args[2]
		fmt.Println("fromKey: ", fromKey, "  toFile: ", toFile)

		f, err := downloader.DownloadFile(fromKey, toFile)
		if err != nil {
			log.Println("download from", fromKey, "to", toFile, "fail:", err)
		} else {
			log.Println("download from", fromKey, "to", toFile, "success")
		}
		defer f.Close()
	}
}
