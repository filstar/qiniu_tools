package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

//使用方法：QINIU="../cfg.toml" go run delete_all_files.go

func main() {

	strPassword := ""
	// log.Println("please input password:")
	_, err1 := fmt.Scanln(&strPassword)
	if nil == err1 {
		// log.Println("strPassword:", strPassword)

		if strings.Compare("buyaoluanlai", strPassword) == 0 {
			log.Println("")
			lister := operation.NewListerV2()
			if lister == nil {
				log.Println("load config file", os.Getenv("QINIU"), "failed")
			} else {
				a := lister.ListPrefix("")
				for _, v := range a {
					go func(v string) {
						err := lister.Delete(v)
						log.Println(v, err)
					}(v)
				}
				time.Sleep(time.Second * 10)
			}
		} else {
			fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "\n请联系管理员\n", 0x1B)
		}
	} else {
		fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "\n请联系管理员\n", 0x1B)
	}

}
