package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn, file *os.File) {
    defer conn.Close()

    file.Write([]byte("Handling connection\n"))

    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            file.Write([]byte("Error reading\n"))
            break
        }

        message := string(buf[0:n])
        fmt.Println("Received message:", message)
        file.Write([] byte("Received message: " + message + "\n"))
        if err != nil {
            fmt.Println("Error sending response:", err.Error())
            break
        }
    }
}


func main() {
    os.Remove("/tmp/echo.sock")
    os.Remove("/tmp/gotem.log")
    os.WriteFile("/tmp/gotem.log", []byte("\n\n\nServer launch\n"), 0644)
    file, _ := os.OpenFile("/tmp/gotem.log", os.O_APPEND|os.O_WRONLY, 0644)

    l, err := net.Listen("unix", "/tmp/echo.sock")
    if err != nil {
        errLog := "Listen error: " + err.Error()
        file.Write([]byte(errLog))
    }
    defer l.Close()


    file.Write([]byte("Listening on /tmp/unixsock ... \n"))
    for {
        file.Write([]byte("Server loop\n"))
        conn, err := l.Accept()
        file.Write([]byte("Connection Accepted\n"))
        if err != nil {
            errLog := "Error accepting connection: " + err.Error()
            file.Write([]byte(errLog))
        }

        file.Write([]byte("Accepted new connection\n"))

        go handleConnection(conn, file)
    }
}
