package main

import (
    "fmt"
    "net"
    "encoding/gob"
)

type Request struct {
    Type int64
}
type Response struct {
    Message string
}
func handleConnection(conn net.Conn) *Response {
    dec := gob.NewDecoder(conn)
    p := &Request{}
    dec.Decode(p)
    fmt.Printf("Received : %+v", p);
    return &Response{"Yar"}
}

func main() {
    fmt.Println("start");
   ln, err := net.Listen("tcp", ":6667")
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept() // this blocks until connection or error
        if err != nil {
            // handle error
            continue
        }
        go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
    }
}
