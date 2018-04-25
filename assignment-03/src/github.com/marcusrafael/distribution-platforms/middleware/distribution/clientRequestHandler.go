package distribution

import(
    "log"
    "net"
    "strings"
)

func clientFailOnError(err error, msg string) {
     if err != nil {
          log.Fatalln("%s: %s", msg, err)
          return
     }
}


type ClientRequestHandler struct {
    Host string
    Port string
}

func (crh *ClientRequestHandler) Send(message []byte) []byte {

    address := strings.Join([]string{crh.Host, crh.Port},":")
    conn, err1 := net.Dial("tcp", address)
    clientFailOnError(err1, "error while creating socket")
    defer conn.Close()
    conn.Write(message)
    log.Printf("message sent to server")
    buffer := make([]byte, 1024)
    data, err2 := conn.Read(buffer)
    clientFailOnError(err2, "error while reading socket")
    log.Printf("message receive from server")
    return buffer[:data]

}
