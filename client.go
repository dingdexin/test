//客户端发送封包
package main

import (
	_ "crypto/tls"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var ip *string
var port *string

func send() {
	server := *ip + ":" + *port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("connected!")

	select {}
}

func main() {

	num := flag.Int("num", 1500, "tcp client num")
	ip = flag.String("ip", "139.219.97.225", "tcp server ip") //139.219.97.225
	port = flag.String("port", "30101", "  port")
	flag.Parse()
	fmt.Println("ok:", *num)
	for i := 0; i < *num; i++ {
		go send()
		time.Sleep(time.Millisecond * 5)
	}
	select {}

}
