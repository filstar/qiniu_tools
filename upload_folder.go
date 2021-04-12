package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

//使用方法：QINIU="../cfg.toml" go run upload_folder.go

func main() {
	fromFolder := os.Args[1]
	log.Println("upload from ", fromFolder)
	upload(fromFolder)
	println("end")
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
		a := regexp.MustCompile("cache|sealed")
		var str = a.FindAllString(v, -1)
		if len(str) > 0 {
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
}
