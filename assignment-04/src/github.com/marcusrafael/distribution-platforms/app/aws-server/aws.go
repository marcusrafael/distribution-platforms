package main

import "github.com/marcusrafael/distribution-platforms/middleware"

func main() { middleware.Invoker(middleware.AWSService, middleware.UnknownAddress) }
