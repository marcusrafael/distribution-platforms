package main

import "github.com/marcusrafael/distribution-platforms/middleware"

//import "time"
import "log"

func main() {

	// $ DO NOT RUN PERFORMANCE EVALUATION WITHOUT COMMENT UPLOAD/DOWNLOAD ACTIONS :)
	//
	// Actions:
	// 1) comment download
	// 2) perform upload
	// 3) comment upload
	// 4) perform download

	proxy, err := middleware.Lookup(middleware.StorageService)
	middleware.FailOnError(err, "fail looking up service")

	//err = proxy.Upload("app/storage-client/files/10MB.zip")
	//middleware.FailOnError(err, "fail uploading file")
	log.Println(proxy)
	err = proxy.Upload("app/storage-client/files/XAVIERMB.zip")
	middleware.FailOnError(err, "fail uploading file")
	//proxy, err := middleware.Lookup(middleware.GCPService)
	//proxy, err := middleware.Lookup(middleware.AWSService)
	//middleware.FailOnError(err, "fail looking up service")
	//log.Println(proxy)
	//start := time.Now()
	//log.Println(start)
	//for i := 0; i < 10000; i++ {
	//_, err := proxy.Upload("app/scenarios/files/10MB.zip")
	//middleware.FailOnError(err, "fail uploading file")
	//time.Sleep(10 * time.Millisecond)
	//}
	//elapsed := time.Since(start)
	//log.Println(elapsed)
	//proxy.Download("app/scenarios/files/image.jpg")
	//middleware.FailOnError(err, "fail downloading file")
}
