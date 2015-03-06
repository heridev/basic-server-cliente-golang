package main

import (
  "net"
  "fmt"
  "time"
  "os"
)

func main() {

  port := "127.0.0.1:" + os.Getenv("PORT")

  udpAddress, err := net.ResolveUDPAddr("udp4",port)

  if err != nil {
    fmt.Println("error resolving UDP address on ", port)
    fmt.Println(err)
    return
  }

  conn ,err := net.ListenUDP("udp", udpAddress)

  if err != nil {
    fmt.Println("error listening on UDP port ", port)
    fmt.Println(err)
    return
  }

  defer conn.Close()

  var buf []byte = make([]byte, 1500)

  for {

    time.Sleep(100 * time.Millisecond)

    n,address, err := conn.ReadFromUDP(buf)

    if err != nil {
      fmt.Println("error reading data from connection")
      fmt.Println(err)
      return
    }

    if address != nil {

      fmt.Println("got message from ", address, " with n = ", n)

      if n > 0 {
        fmt.Println("from address", address, "got message:", string(buf[0:n]), n)
      }

      ip := address.String()

      clientconn, err := net.Dial("udp", ip)

      if err != nil {
        fmt.Println("Could not resolve udp address or connect to it  on " , ip)
        fmt.Println(err)
        return
      }
      defer clientconn.Close()

      a, err := conn.Write([]byte("SOS ... \n"))
      fmt.Println(a)
      if err != nil {
        fmt.Println("error writing data to server", ip)
        fmt.Println(err)
        return
      }


      fmt.Println("Responding to server at ", ip)

    }
  }

}
