package middleware

import (
	"io/ioutil"
	"log"
	"net"
)

func Bind(socket net.Listener, service string) error {
	params := make(map[string]string)
	params["operation"] = "Bind"
	params["service"] = service
	params["address"] = socket.Addr().String()
	params["dns"] = KnownAddress
	message := Message{params, []byte{}}
	message, err := Request(message)
	return err
}

func Lookup(service string) (Proxy, error) {
	params := make(map[string]string)
	params["operation"] = "Lookup"
	params["service"] = service
	params["dns"] = KnownAddress
	message := Message{params, []byte{}}
	message, err := Request(message)
	rparams := make(map[string]map[string]string)
	tmp := make(map[string]string)
	tmp["service"] = message.Parameters["service"]
	tmp["address"] = message.Parameters["address"]
	rparams["result"] = tmp
	name := Name{rparams}
	proxy := Proxy{name}
	return proxy, err
}

type Proxy struct {
	Name Name
}

func (proxy *Proxy) Upload(filepath string) error {
	log.Println("uploading", filepath)
	params := make(map[string]string)
	params["operation"] = "Upload"
	params["filepath"] = filepath
	params["service"] = proxy.Name.Binds["result"]["service"]
	params["address"] = proxy.Name.Binds["result"]["address"]
	data, err := ioutil.ReadFile(params["filepath"])
	message := Message{params, data}
	_, err = Request(message)
	return err
}

func (proxy Proxy) Download(filepath string) error {
	log.Println("downloading", filepath)
	params := make(map[string]string)
	params["operation"] = "Download"
	params["filepath"] = filepath
	params["service"] = proxy.Name.Binds["result"]["service"]
	params["address"] = proxy.Name.Binds["result"]["address"]
	message := Message{params, []byte{}}
	message, err := Request(message)
	err = ioutil.WriteFile(params["filepath"], message.Data, 0644)
	FailOnError(err, "error writing file")
	return err
}
