package distribution
import "log"
import "strconv"
import "github.com/fatih/color"


type NamingProxy struct {
    Host string
    Port string
}

func (np NamingProxy) Lookup(service string) CalculatorProxy {

    var requestor Requestor

    parameters := make(map[string]string)
    parameters["Service"] = service
    invocation := Invocation{np.Host, np.Port, "Lookup", parameters}
    termination := requestor.Request(invocation)
    log.Println(termination.Result)
    color.Red("started")
    var calculatorProxy CalculatorProxy

    calculatorProxy.Host = termination.Result["Host"]
    calculatorProxy.Port = termination.Result["Port"]
    calculatorProxy.InvokerId = termination.Result["InvokerId"]
    calculatorProxy.Protocol = termination.Result["Protocol"]
    calculatorProxy.ObjectId = termination.Result["ObjectId"]

    log.Println(calculatorProxy.Host)
    log.Println(calculatorProxy.Port)
    log.Println(calculatorProxy.InvokerId)
    log.Println(calculatorProxy.Protocol)
    log.Println(calculatorProxy.ObjectId)

    return calculatorProxy

}

func (np NamingProxy) Bind(serviceName string, calculatorProxy CalculatorProxy) {

    var requestor Requestor

    parameters := make(map[string]string)
    parameters["Service"] = serviceName
    parameters["Host"] = calculatorProxy.Host
    parameters["Port"] = calculatorProxy.Port
    parameters["InvokerId"] = calculatorProxy.InvokerId
    parameters["Protocol"] = calculatorProxy.Protocol
    parameters["ObjectId"] = calculatorProxy.ObjectId
    invocation := Invocation{np.Host, np.Port, "Bind", parameters}
    log.Println(invocation)
    requestor.Request(invocation)

}


type CalculatorProxy struct {
    Host string
    Port string
    Protocol string
    InvokerId string
    ObjectId string
}

func (cp CalculatorProxy) operation(operation string, x float64, y float64) float64 {

    var invocation Invocation
    var termination Termination
    var requestor Requestor

    parameters := make(map[string]string)
    xResultAsString := strconv.FormatFloat(x, 'f', 2, 64)
    yResultAsString := strconv.FormatFloat(y, 'f', 2, 64)
    parameters["first"] = xResultAsString
    parameters["second"] = yResultAsString
    invocation.Host = cp.Host
    invocation.Port = cp.Port
    invocation.Operation = operation
    invocation.Parameters = parameters

    termination = requestor.Request(invocation)
    result, _ := strconv.ParseFloat(termination.Result["result"], 64)

    return result

}


func (cp CalculatorProxy) Sum(x float64, y float64) float64 {
    return cp.operation("Sum", x, y)
}

func (cp CalculatorProxy) Sub(x float64, y float64) float64 {
    return cp.operation("Sub", x, y)
}

func (cp CalculatorProxy) Div(x float64, y float64) float64 {
    return cp.operation("Div", x, y)
}

func (cp CalculatorProxy) Mul(x float64, y float64) float64 {
    return cp.operation("Mul", x, y)
}
