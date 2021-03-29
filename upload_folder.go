package main

import (
	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
	"strings"

	"log"
	"os"
	"path/filepath"
)

//使用方法：QINIU="../cfg.toml" go run rename_file.go

func main() {
	fromFolder := os.Args[1]
	log.Println("upload from ", fromFolder)
	upload(fromFolder)
	log.Println("end")
}

func upload(fullPath string) {
	var filePathWalkDir = func(root string) ([]string, error) {
		var files []string
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		return files, err
	}
	fileList, err := filePathWalkDir(fullPath)

	for _, v := range fileList {
		if strings.Contains(v, "fetching") {
			continue
		}
		uploader := operation.NewUploaderV2()
		if uploader == nil {
			log.Println("init uploader failed", err)
			continue
		}
		if err := uploader.Upload(v, v); err != nil {
			log.Println("upload failed", v, err)
			continue
		}
		//os.Remove(v)
	}
}
