package distribution

import (
    "log"
    "reflect"
    "strconv"
    "strings"
)

func invokerFailOnError(err error, msg string) {
     if err != nil {
          log.Fatalln("%s: %s", msg, err)
          return
     }
}

type Invoker struct {
    aor map[string]map[string]string
}

func (invoker *Invoker) Invoke(data []byte) []byte {

    var message Message
    var marshaller Marshaller

    marshaller.Unmarshal(data, &message)

    operation := message.Body.RequestHeader.Operation

    method := strings.Title(operation)

    switch method {

        case "Lookup":
            message = invoker.Lookup(message)
        case "Bind":
            message = invoker.Bind(message)
        default:
            message = invoker.Calculator(message)

    }

    msgMarshalled, _ := marshaller.Marshal(message)

    return msgMarshalled

}


func (invoker *Invoker) Calculator(message Message) Message {

    log.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
    first, second := before(message)

    operation := message.Body.RequestHeader.Operation

    method := strings.Title(operation)

    calculator := Calculator{first, second}

    result := reflect.ValueOf(&calculator).MethodByName(method).Call([]reflect.Value{})

    final := result[0].Float()

    return after(message, final)
}

func (invoker *Invoker) Bind(message Message) Message {
    //Protocol:miop Service:calculator Host:localhost InvokerId:invoker ObjectId:object Port:4321
    parameters := message.Body.RequestBody.Parameters
    serviceName := parameters["Service"]
    invoker.aor[serviceName] = parameters
    log.Println("!@#$%^&*()", invoker.aor)
    return message
}

func (invoker *Invoker) Lookup(message Message) Message {
    service := message.Body.RequestBody.Parameters["Service"]
    neww := invoker.aor[service]
    message.Body.ReplyBody.Result = neww
    log.Println("###################################################")
    log.Println(message.Body.ReplyBody.Result)
    log.Println("###################################################")
    return message
}

func before(message Message) (float64, float64) {

    first := message.Body.RequestBody.Parameters["first"]
    second := message.Body.RequestBody.Parameters["second"]
    number1, err1 := strconv.ParseFloat(first, 64)
    invokerFailOnError(err1, "parser error")
    number2, err2 := strconv.ParseFloat(second, 64)
    invokerFailOnError(err2, "parser error")
    return number1, number2
}

func after(message Message, result float64) Message {

    resultAsString := strconv.FormatFloat(result, 'f', 2, 64)
    message.Body.ReplyBody.Result["result"] = resultAsString
    return message

}
