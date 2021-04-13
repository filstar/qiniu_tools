package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

//使用方法：QINIU="../cfg.toml" go run download_folder.go

func main() {

	downloader := operation.NewDownloaderV2()
	if downloader == nil {
		log.Println("load config file", os.Getenv("QINIU"), "failed")
	} else {
		fromKey := os.Args[1]
		toFolder := os.Args[2]
		fmt.Println("fromKey: ", fromKey, "  toFolder: ", toFolder)

		exist, err := PathExists(toFolder)
		if err != nil {
			fmt.Printf("get dir error![%v]\n", err)
			return
		}
		if !exist {
			err := os.MkdirAll(toFolder, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			} else {
				fmt.Printf("mkdir success!\n")
			}
		}

		lister := operation.NewListerV2()
		if lister != nil {
			a := lister.ListPrefix("")
			for _, v := range a {
				if strings.Contains(v, fromKey) {
					tmpfromKey := strings.TrimSpace(v)
					tmptoFile := strings.Replace(tmpfromKey, fromKey, toFolder, 1)
					//fmt.Println("fromKey: ", tmpfromKey, "  toFile: ", tmptoFile)
					f, err := downloader.DownloadFile(tmpfromKey, tmptoFile)
					if err != nil {
						log.Println("download from", tmpfromKey, "to", tmptoFile, "fail:", err)
					} else {
						log.Println("download from", tmpfromKey, "to", tmptoFile, "success")
					}
					defer f.Close()
				}
			}
		}
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
