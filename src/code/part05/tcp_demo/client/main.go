package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	// 1.与服务端建立连接
	conn ,err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil{
		fmt.Printf("filed, err: %v\n", err)
		return

	}
	// 2.利用该连接进行数据的发送和接收

	input := bufio.NewReader(os.Stdin)
	for {
		// 对需要发送给服务端的消息根据换行符读取
		s, _ := input.ReadString('\n')
		// 去除量变空格
		s = strings.TrimSpace(s)
		// 判断如果输入的是Q，则退出
		if strings.ToUpper(s) == "Q"{
			return
		}

		// 给服务端发送消息
		_, err := conn.Write([]byte(s))
		if err != nil{
			fmt.Printf("send fialed, err: %v\n", err)

			return
		}

		// 从服务端接收回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil{
			fmt.Printf("read fialed, err: %v\n", err)
			return
		}

		fmt.Println("接收到服务端回复: ", string(buf[:n]))


	}
}
