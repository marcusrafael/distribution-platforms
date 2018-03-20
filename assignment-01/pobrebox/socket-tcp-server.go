package main

import (
    "fmt"
    "io"
    "net"
    "os"
    "strconv"
    "io/ioutil"
)


const BUFFERSIZE = 1024




func main() {

    server, err := net.Listen("tcp", "0.0.0.0:1337")
    if err != nil {
        fmt.Println("Error listetning: ", err)
        os.Exit(1)
    }

    defer server.Close()
    fmt.Println("Server started! Waiting for connections...")
    for {
        connection, err := server.Accept()
        if err != nil {
            fmt.Println("Error: ", err)
            os.Exit(1)
        }
        fmt.Println("Client connected")
        sendFileToClient(connection)
    }
}


func sendFileToClient(connection net.Conn) {
    fmt.Println("Someone has connected!")
    files,_ := ioutil.ReadDir("/home/ubuntu/pobrebox/")

    var theFileName string
    for _, f := range files {
            theFileName = f.Name()
            theFileName = "/home/ubuntu/pobrebox/" + theFileName
            file, err := os.Open(theFileName)
            if err != nil {
                fmt.Println(err)
                return
            }
            fileInfo, err := file.Stat()
            if err != nil {
                fmt.Println(err)
                return
            }
            fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
            fileName := fillString(fileInfo.Name(), 64)
            fmt.Println("Sending filename and filesize!")
            connection.Write([]byte(fileSize))
            connection.Write([]byte(fileName))
            sendBuffer := make([]byte, BUFFERSIZE)
            fmt.Println("Start sending file!")
            for {
                _, err = file.Read(sendBuffer)
                if err == io.EOF {
                    break
                }
                connection.Write(sendBuffer)
            }
            fmt.Println("File has been sent, closing connection!")
            os.Remove(theFileName)
            return
    }
    connection.Close()
}


func fillString(retunString string, toLength int) string {
    for {
        lengtString := len(retunString)
        if lengtString < toLength {
            retunString = retunString + ":"
            continue
        }
        break
    }
    return retunString
}
