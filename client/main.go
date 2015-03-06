package main

import (
  "fmt"
  "net"
  "time"
  "os"
)

func main() {
  fmt.Println(len(os.Args))
  if len(os.Args) != 2{
    fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
    os.Exit(1)
  }

  service := os.Args[1]

  fmt.Println("Connecting to server at ", service)

  conn1, err := net.Dial("udp",service)
  fmt.Println("Connecting to server at ", service)

  if err != nil {
    fmt.Println("Could not resolve udp address or connect to it  on " , service)
    fmt.Println(err)
    return
  }

  fmt.Println("Connected to server at ", service)

  defer conn1.Close()

  for {

    time.Sleep(1000*time.Millisecond)

    n, err := conn1.Write([]byte("SOS ... \n"))
    if err != nil {
      fmt.Println("error writing data to server", service)
      fmt.Println(err)
      return
    }

    if n > 0 {
      fmt.Println("Wrote ",n, " bytes to server at ", service)
    }
  }

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

    n, address, err := conn.ReadFromUDP(buf)

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

    }
  }

}


 //members := [3]string{"1200", "1201", "1202"}

//func main() {

  //fmt.Println(len(os.Args))
  //if len(os.Args) != 2{
    //fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
    //os.Exit(1)
  //}

  //service := os.Args[1]

  //fmt.Println("Connecting to server at ", service)

  //conn, err := net.Dial("udp",service)
  //fmt.Println("Connecting to server at ", service)

  //if err != nil {
    //fmt.Println("Could not resolve udp address or connect to it  on " , service)
    //fmt.Println(err)
    //return
  //}

  //fmt.Println("Connected to server at ", service)

  //defer conn.Close()

  //fmt.Println("About to write to connection")

  //for {

    //n, err := conn.Write([]byte("SOS ... \n"))
    //if err != nil {
      //fmt.Println("error writing data to server", service)
      //fmt.Println(err)
      //return
    //}

    //if n > 0 {
      //fmt.Println("Wrote ",n, " bytes to server at ", service)
    //}
  //}

//}
