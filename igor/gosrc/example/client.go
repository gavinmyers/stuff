// You can edit this code!
// Click here and start typing.
package main

import (
  "fmt"
  "net"
  "os"
  )

func main() {
    conn, err := net.Dial("tcp4", "localhost:9988")
    if err != nil {
        fmt.Println(err)
  os.Exit(1)
    }
    
    defer conn.Close()                                                          
                                                                                
    var readbuf [512]byte                                                       
    var writebuf [512]byte                                                      
                                                                                
    for {                                                                       
                                                                                
        // Read connection -> Write stdout
        n, _ := conn.Read(readbuf[0:])                                        
  os.Stdout.Write(readbuf[0:n])                                  
        
  // Read stdin -> Write connection
        n, _ = os.Stdin.Read(writebuf[0:])                                    
  conn.Write(writebuf[0:n])                                      
    } 
}
