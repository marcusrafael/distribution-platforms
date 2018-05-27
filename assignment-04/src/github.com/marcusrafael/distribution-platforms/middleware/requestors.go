package middleware

import (
	"encoding/gob"
	"net"
)

func Request(message Message) (Message, error) {
	address := message.Parameters["address"]
	if message.Parameters["dns"] != "" {
		address = KnownAddress
	}
	conn, err := net.Dial("tcp", address)
	FailOnError(err, "fail connecting to server socket")
	encoder := gob.NewEncoder(conn)
	encoder.Encode(message)
	response := &Message{}
	decoder := gob.NewDecoder(conn)
	decoder.Decode(response)
	conn.Close()
	return *response, nil
}
