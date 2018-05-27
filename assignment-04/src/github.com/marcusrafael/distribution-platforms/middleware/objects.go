package middleware

import (
	"bytes"
	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

type Name struct {
	Binds map[string]map[string]string
}

type Message struct {
	Parameters map[string]string
	Data       []byte
}

func NameBind(binds map[string]map[string]string, message *Message) error {
	name := make(map[string]string)
	name["service"] = message.Parameters["service"]
	name["address"] = message.Parameters["address"]
	binds[name["service"]] = name
	return nil
}

func NameLookup(binds map[string]map[string]string, message *Message) error {
	service := message.Parameters["service"]
	keys := make([]string, 0, len(binds))
	for k := range binds {
		keys = append(keys, k)
	}
	founds := []string{} // found service's name
	for _, v := range keys {
		if strings.Contains(v, service) {
			founds = append(founds, v)
		}
	}
	params := make(map[string]string)
	for i, _ := range founds {
		name := binds[founds[i]]
		timeout := time.Duration(1 * time.Second)
		log.Println("trying... ", name["address"])
		conn, err := net.DialTimeout("tcp", name["address"], timeout)
		time.Sleep(3 * time.Second)
		if err == nil {
			defer conn.Close()
			params["service"] = name["service"]
			params["address"] = name["address"]
		} else {
			log.Println("Skipping:", name)
		}
	}
	log.Println(params)
	*message = Message{params, []byte{}}
	return nil
}

func GCPUpload(message *Message) error {
	params := message.Parameters
	ctx := context.Background()
	client, err1 := storage.NewClient(ctx)
	FailOnError(err1, "error while creating gcp client")
	bh := client.Bucket("middleware-dp")
	obj := bh.Object(params["filepath"])
	w := obj.NewWriter(ctx)
	dat, err := ioutil.ReadFile(params["filepath"])
	FailOnError(err, "error while creating gcp client")
	_, err2 := w.Write(dat)
	FailOnError(err2, "error while creating gcp client")
	err3 := w.Close()
	FailOnError(err3, "error while creating gcp client")
	return nil
}

func GCPDownload(message *Message) error {
	params := message.Parameters
	ctx := context.Background()
	client, err1 := storage.NewClient(ctx)
	FailOnError(err1, "error while creating gcp reader")
	bh := client.Bucket("middleware-dp")
	obj := bh.Object(params["filepath"])
	r, err1 := obj.NewReader(ctx)
	FailOnError(err1, "error while creating gcp reader")
	data, err2 := ioutil.ReadAll(r)
	FailOnError(err2, "error while reading all")
	*message = Message{params, data}
	defer r.Close()
	return nil
}

func AWSDownload(message *Message) error {
	params := message.Parameters
	cfg, err := external.LoadDefaultAWSConfig()
	FailOnError(err, "error loading conf file")
	cfg.Region = endpoints.UsEast1RegionID
	downloader := s3manager.NewDownloader(cfg)
	buffer := &aws.WriteAtBuffer{}
	_, err2 := downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String("kumo-bucket"),
		Key:    aws.String(params["filepath"]),
	})
	FailOnError(err2, "error downloading from aws")
	*message = Message{message.Parameters, buffer.Bytes()}
	return nil
}

func AWSUpload(message *Message) error {
	params := message.Parameters
	cfg, err := external.LoadDefaultAWSConfig()
	FailOnError(err, "error loading conf file")
	cfg.Region = endpoints.UsEast1RegionID
	uploader := s3manager.NewUploader(cfg)
	uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("kumo-bucket"),
		Key:    aws.String(params["filepath"]),
		Body:   bytes.NewReader(message.Data),
	})
	return nil
}
