package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
        "path/filepath"
)


const BUFFERSIZE = 1024



func main() {

                connection, err := net.Dial("tcp", "ec2-XX-XXX-XXX-XX.compute-1.amazonaws.com:1337")
		if err != nil {
			panic(err)
		}
		defer connection.Close()
		fmt.Println("Connected to server, start receiving the file name and file size")
		bufferFileName := make([]byte, 64)
		bufferFileSize := make([]byte, 10)
		connection.Read(bufferFileSize)
		fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
		connection.Read(bufferFileName)
		fileName := strings.Trim(string(bufferFileName), ":")
                fmt.Println(fileName)
                fileName = filepath.Base(fileName)
                fileName = "/var/www/html/" + fileName
		newFile, err := os.Create(fileName)
		if err != nil {
	            panic(err)
		}
		defer newFile.Close()
		var receivedBytes int64
		for {
			if (fileSize - receivedBytes) < BUFFERSIZE {
				io.CopyN(newFile, connection, (fileSize - receivedBytes))
				connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
				break
			}
			io.CopyN(newFile, connection, BUFFERSIZE)
			receivedBytes += BUFFERSIZE
		}
		fmt.Println("Received file completely!")
}
