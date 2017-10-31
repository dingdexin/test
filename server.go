//服务端解包过程
package main

import (
	"fmt"
	"io"
	"net"
	"os"
	_ "sync/atomic"
	"time"

	//	_ "../protocol"
)

var clientNum int32 = 0
var connchan chan net.Conn

func main() {

	netListen, err := net.Listen("tcp", ":30101")
	CheckError(err)

	defer netListen.Close()
	count := 0
	Log("Waiting for clients")
	for {
		_, err := netListen.Accept()
		count++
		fmt.Println("count ", count)
		if err != nil {
			continue
		}

	}
}

func dispatchConnection() {
	for {
		select {
		case conn, _ := <-connchan:
			go handleConnection(conn)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			if io.EOF == err {
				Log(conn.RemoteAddr().String(), " connection close: ")
				return
			} else {
				break
			}

		}

	}
	select {}
}

func reader(readerChannel chan []byte) { //channel的消费者。
	for {
		select {
		case data := <-readerChannel:
			fmt.Println("收到一帧", data) //此处就是完整的一帧，后续可以解析
		}
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
