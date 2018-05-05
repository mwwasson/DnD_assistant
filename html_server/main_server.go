package html_server

import (
    "log"
    "net"
    "net/http"
    "fmt"
)

func Start_server(port_no int) {
    port_str := fmt.Sprintf(":%d",port_no);

    currentIP := get_server_ip()
    http.HandleFunc("/",sayHelloWorld)

    fmt.Printf("Starting DnD Assistant Server:\n")
    fmt.Printf("IP Address: %s\n", currentIP)
    fmt.Printf("Port: %d\n",port_no)
    http.ListenAndServe(port_str,nil)
    
}

func get_server_ip() net.IP {
    conn, err := net.Dial("udp", "1.1.1.1:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    
    return localAddr.IP
    
}
