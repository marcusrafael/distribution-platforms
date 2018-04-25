package distribution

import "github.com/fatih/color"

type Requestor struct {}

func (* Requestor) Request(invocation Invocation) Termination {

    var marshaller Marshaller
    var termination Termination

    clientRequestHandler := ClientRequestHandler{invocation.Host, invocation.Port}

    requestHeader := RequestHeader{"", 0, true, 0, invocation.Operation}
    requestBody := RequestBody{invocation.Parameters}
    replyHeader := ReplyHeader{"", 0, 0}
    replyBody := ReplyBody{make(map[string]string)}
    messageHeader := MessageHeader{"MIOP", 0, false, 0, 0}

    messageBody := MessageBody{requestHeader, requestBody, replyHeader, replyBody}

    message := Message{messageHeader, messageBody}

    msgMarshalled, _ := marshaller.Marshal(message)

    color.Blue("calling send method on requestor [clientRequestHandler]")
    msgToBeUnmarshalled := clientRequestHandler.Send(msgMarshalled)
    color.Blue("CALLED send method on requestor [clientRequestHandler]")

    var msgUnmarshalled Message

    marshaller.Unmarshal(msgToBeUnmarshalled, &msgUnmarshalled)

    //log.Println("*********************************")
    //log.Println(msgUnmarshalled.Body.ReplyBody.Result)
    //log.Println("*********************************")
    termination.Result = msgUnmarshalled.Body.ReplyBody.Result

    return termination
}
