package main

import (
	"log"
	"os"

	"github.com/qiniupd/qiniu-go-sdk/syncdata/operation"
)

//使用方法：QINIU="../cfg.toml" go run upload_genesis-sectors.go

func main() {
	var err error

	//上传
	uploader := operation.NewUploaderV2()
	if uploader == nil {
		log.Println("load config file", os.Getenv("QINIU"), "failed")
	} else {
		err = uploader.Upload("/root/.genesis-sectors/sealed/s-t01000-0", "root/.genesis-sectors/sealed/s-t01000-0")
		err = uploader.Upload("/root/.genesis-sectors/cache/s-t01000-0/p_aux", "root/.genesis-sectors/cache/s-t01000-0/p_aux")
		err = uploader.Upload("/root/.genesis-sectors/cache/s-t01000-0/sc-02-data-tree-r-last.dat", "root/.genesis-sectors/cache/s-t01000-0/sc-02-data-tree-r-last.dat")
		err = uploader.Upload("/root/.genesis-sectors/cache/s-t01000-0/t_aux", "root/.genesis-sectors/cache/s-t01000-0/t_aux")

		err = uploader.Upload("/root/.genesis-sectors/sealed/s-t01000-1", "root/.genesis-sectors/sealed/s-t01000-1")
		err = uploader.Upload("/root/.genesis-sectors/cache/s-t01000-1/p_aux", "root/.genesis-sectors/cache/s-t01000-1/p_aux")
		err = uploader.Upload("/root/.genesis-sectors/cache/s-t01000-1/sc-02-data-tree-r-last.dat", "root/.genesis-sectors/cache/s-t01000-1/sc-02-data-tree-r-last.dat")
		err = uploader.Upload("/root/.genesis-sectors/cache/s-t01000-1/t_aux", "root/.genesis-sectors/cache/s-t01000-1/t_aux")
		if err != nil {
			log.Println("上传失败：", err)
		}
	}

	////查看文件
	//lister := operation.NewListerV2()
	//if lister == nil {
	//	panic(errors.New("列表初始化失败"))
	//}
	//filestats := lister.ListStat([]string{"root/a.txt", "root/b.txt"})
	//for _, v := range filestats {
	//	log.Println("filestat:", v)
	//}
	//    //下载文件
	//    downloader := operation.NewDownloaderV2()
	//    f, err := downloader.DownloadFile("key", "/tmp/path")
	//    if err != nil {
	//       log.Println(fmt.Sprintf("下载%s到%s失败：%s", "key", "/tmp/path", err))
	//    }
	//defer f.Close()

}
