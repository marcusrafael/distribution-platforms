package main

import "github.com/marcusrafael/distribution-platforms/middleware/distribution"


func main() {

    namingService := distribution.NamingProxy{"localhost", "1234"}
    proxy := distribution.CalculatorProxy{"localhost", "4321", "miop", "invoker", "object"}
    namingService.Bind("calculator", proxy)

    server := distribution.ServerRequestHandler{"localhost", "4321"}
    server.Start()
}
