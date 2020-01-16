package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	for {
		var (
			buff   = make([]byte, 1024)
			source string
		)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Print(string(buff[0:n]))

		if string(buff[0:n]) == "Enter command:" {
			fmt.Scanln(&source)
			if toServer(source, conn) {
				return
			}
		}

		if string(buff[0:n]) == "Enter username:" {
			fmt.Scanln(&source)
			if toServer(source, conn) {
				return
			}
		}

		if string(buff[0:n]) == "Entering task... Select executor:\n" {
			n, _ := conn.Read(buff)
			fmt.Print(string(buff[0:n]))
			fmt.Scanln(&source)
			if toServer(source, conn) {
				return
			}
			n, _ = conn.Read(buff)
			fmt.Print(string(buff[0:n]))
			fmt.Scanln(&source)
			if toServer(source, conn) {
				return
			}
		}

		/*

			fmt.Print("From server:")
			n, err = conn.Read(buff)
			if err !=nil {
				break
			}
			fmt.Print(string(buff[0:n]))
			fmt.Println()*/
	}
	fmt.Println("\nServer closed...")
}

func toServer(source string, conn net.Conn) bool {
	if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
