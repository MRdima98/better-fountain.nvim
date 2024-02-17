package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
    socket, err := net.Listen("unix", "/tmp/gotem.sock")
    if err != nil {
        log.Fatal(err)
    }

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        socket.Close()
        os.Exit(1)
    }()

    for {
        conn, err := socket.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go func(conn net.Conn) {
            defer conn.Close()
            // Create a buffer for incoming data.
            buf := make([]byte, 4096)

            // Read data from the connection.
            n, err := conn.Read(buf)
            if err != nil {
                log.Fatal(err)
            }

            // Echo the data back to the connection.
            _, err = conn.Write(buf[:n])
            if err != nil {
                log.Fatal(err)
            }
        }(conn)
    }
}
