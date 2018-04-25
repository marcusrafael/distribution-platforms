package distribution

import(
    "log"
    "net"
    "strings"
)

func serverFailOnError(err error, msg string) {
     if err != nil {
          log.Fatalln("%s: %s", msg, err)
          return
     }
}


type ServerRequestHandler struct {
    Host string
    Port string
}

func (srh *ServerRequestHandler) Start() {

    address := strings.Join([]string{srh.Host, srh.Port},":")
    listener, errListen := net.Listen("tcp", address)
    serverFailOnError(errListen, "error while listening socket")
    log.Println("server socket initilized")

    invoker := Invoker{make(map[string]map[string]string)}

    for {
        conn, err1 := listener.Accept()
        serverFailOnError(err1, "error while accepting socket")
        defer listener.Close()
        buffer := make([]byte, 1024)
        dataSize, err2 := conn.Read(buffer)
        serverFailOnError(err2, "error while reading buffer")
        data := buffer[:dataSize]
        result := invoker.Invoke(data)
        conn.Write(result)
    }
}
