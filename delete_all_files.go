package main

import (
	"log"
	"os"
	"time"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

func main() {
	var err error

	lister := operation.NewListerV2()
	if lister == nil {
		log.Println("load config file", os.Getenv("QINIU"), "failed")
	}
	a := operation.ListPrefix("")
	for _, v := range a {
		go func(v string) {
			err := operation.Delete(v)
			log.Println(v, err)
		}(v)
	}
	time.Sleep(time.Second * 10)

	/*
		//加载配置
		config, err := operation.Load("./cfg.toml")
		if err != nil {
			log.Println("加载配置文件失败：", err)
			return
		}
		op := operation.NewLister(config)

		a := op.ListPrefix("")
		for _, v := range a {
			go func(v string) {
				err := op.Delete(v)
				log.Println(v, err)
			}(v)
		}
		time.Sleep(time.Second * 10)
	*/
}
