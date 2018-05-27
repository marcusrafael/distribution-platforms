package middleware

import (
	"encoding/gob"
	"log"
	"net"
)

func bind(binds map[string]map[string]string, message *Message) {
	err := NameBind(binds, message)
	FailOnError(err, "fail binding")
	log.Println("bind :)")
}

func lookup(binds map[string]map[string]string, message *Message) {
	err := NameLookup(binds, message)
	FailOnError(err, "fail looking up")
	log.Println("lookup :)")
}

func upload(message *Message) {
	log.Println(message.Parameters["service"])
	service := message.Parameters["service"]
	if service == AWSService {
		log.Println("uploading to aws :)")
		err := AWSUpload(message)
		FailOnError(err, "fail aws uploading")
	}
	if service == GCPService {
		log.Println("uploading to gcp :)")
		err := GCPUpload(message)
		FailOnError(err, "fail gcp uploading")
	}
}

func download(message *Message) {
	service := message.Parameters["service"]
	if service == AWSService {
		log.Println("downloading to aws :)")
		err := AWSDownload(message)
		FailOnError(err, "fail aws downloading")
	}
	if service == GCPService {
		log.Println("downloading to gcp :)")
		err := GCPDownload(message)
		FailOnError(err, "fail gcp downloading")
	}
}

func Invoker(service string, address string) {

	listener, err := net.Listen("tcp", address)
	FailOnError(err, "fail start server socket")
	binds := make(map[string]map[string]string)

	if address != KnownAddress {
		log.Println("binding service:", service)
		log.Println("binding address:", listener.Addr().String())
		Bind(listener, service)
	}

	for {
		conn, err := listener.Accept()
		FailOnError(err, "fail accepting connection")
		message := &Message{}
		decoder := gob.NewDecoder(conn)
		decoder.Decode(message)
		operation := message.Parameters["operation"]

		log.Println("invoking:", operation)

		switch operation {
		case "Bind":
			bind(binds, message)
		case "Lookup":
			lookup(binds, message)
		case "Upload":
			upload(message)
		case "Download":
			download(message)
		}
		encoder := gob.NewEncoder(conn)
		encoder.Encode(message)
	}
}
