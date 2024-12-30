package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func checkError(error error) {
	if error != nil {
		panic("Error: " + error.Error()) // terminate program
	}
}

func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}

func doServerStuff(conn net.Conn) {
	var buf []byte
	var error error

	for {
		buf = make([]byte, 512)
		_, error = conn.Read(buf)
		checkError(error)

		input := string(buf)
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down")
			os.Exit(0)
		}

		// op commando WHO:  write out mapUsers
		if strings.Contains(input, ": WHO") {
			DisplayList()
		}

		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]
		mapUsers[string(clName)] = 1
		fmt.Printf("Received data: --%v--", string(buf))
	}
}

func main() {
	var listener net.Listener
	var err error
	var conn net.Conn
	mapUsers = make(map[string]int)

	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err = net.Listen("tcp", "localhost:50000")
	checkError(err)

	// 监听并接受来自客户端的连接
	for {
		conn, err = listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}
