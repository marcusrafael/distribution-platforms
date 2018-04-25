package main

import "fmt"
import "github.com/marcusrafael/distribution-platforms/middleware/distribution"


func main() {

   namingProxy := distribution.NamingProxy{"localhost", "1234"}
   calculatorProxy := namingProxy.Lookup("calculator")

   fmt.Println(calculatorProxy.Sum(21, 21))
   fmt.Println(calculatorProxy.Sub(21, 21))
   fmt.Println(calculatorProxy.Div(21, 21))
   fmt.Println(calculatorProxy.Mul(21, 21))

}
