package main

import "github.com/marcusrafael/distribution-platforms/middleware/distribution"


func main() {

    server := distribution.ServerRequestHandler{"localhost", "1234"}
    server.Start()

}
