package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)
func main() {
	target := flag.String("target", "baidu.com:443", "The relay server (the connect-back address)")
	flag.Parse()
	locallistener, err :=net.Listen("tcp", "0.0.0.0:443")
	if err != nil {
		fmt.Errorf("Could not bind to port : %v\n", err)
		return
	}
	defer locallistener.Close()
	fmt.Printf("PortFwd to : %s\n", *target)
	for {
		stream, err := locallistener.Accept()
		if err != nil {
			return
		}
		proxyConn, err := net.Dial("tcp", *target)
		if err != nil {
			fmt.Errorf("Error creating Proxy TCP connection ! Error : %s\n", err)
			return
		}
		go handleRelay(stream, proxyConn)
	}
}
func handleRelay(src net.Conn, dst net.Conn) {
	stop := make(chan bool, 2)
	go relay(src, dst, stop)
	go relay(dst, src, stop)
	select {
	case <-stop:
		return
	}
}

func relay(src net.Conn, dst net.Conn, stop chan bool) {
	io.Copy(dst, src)
	dst.Close()
	src.Close()
	stop <- true
	return
}