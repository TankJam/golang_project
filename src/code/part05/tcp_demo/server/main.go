package main

import (
	"bufio"
	"fmt"
	"net"
)

// Nagle 可以解决黏包问题

// tcp server
func process(conn net.Conn){
	defer conn.Close() // 处理完之后要关闭这个连接
	// 针对当前的连接座数据的发送和接收操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil{
			fmt.Printf("read from conn failed, err: %v\n", err)
			break
		}
		// 接收数据
		recv := string(buf[:n])
		fmt.Println("接收到客户端的数据: ", recv)
		conn.Write([]byte("ok"))  // 把收到的数据返回给客户端ok
	}
}

func main() {
	//1.开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err!= nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	for {
		// 2.等待客户端建立连接
		conn, err := listen.Accept()
		if err!= nil{
			fmt.Printf("accept faild, err:%v\n", err)
			continue
		}
		// 3.开启一个单独的goroutine去处理连接
		go process(conn)
	}
}

