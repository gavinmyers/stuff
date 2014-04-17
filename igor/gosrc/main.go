package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
)

type Request struct {
    Type int64
}

type Response struct {
    Message string
}
func main() {
    fmt.Println("start client");
    conn, err := net.Dial("tcp", "localhost:6667")
    if err != nil {
        log.Fatal("Connection error", err)
    }

    encoder := gob.NewEncoder(conn)
    p := &Request{1}
    encoder.Encode(p)

    dec := gob.NewDecoder(conn)
    v := &Response{}
    dec.DecodeValue(v)
    fmt.Printf("Received : %+v", v)

    conn.Close()
    fmt.Println("done");
}
