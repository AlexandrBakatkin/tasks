package main

import (
	"fmt"
	"github.com/Bakatkin/tasks/database"
	"net"
	"strconv"
)

func main() {
	const (
		startmsg = "Server is starting..."
	)
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println(startmsg)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer close(conn)

	for {
		conn.Write([]byte("Enter command:"))
		input := make([]byte, 1024)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		command := string(input[0:n])
		fmt.Println(string(input))
		switch command {
		case "adduser":
			conn.Write([]byte("Enter username:"))
			n, _ := conn.Read(input)
			username := string(input[0:n])
			database.AddUser(username)
			conn.Write([]byte("OK\n"))
		case "addtask":
			conn.Write([]byte("Entering task... Select executor:\n"))
			database.ShowUsers(conn)
			n, _ := conn.Read(input)
			executor := string(input[0:n])
			fmt.Println(executor)
			conn.Write([]byte("Enter task:\n"))
			n, _ = conn.Read(input)
			task := string(input[0:n])
			addTask(executor, task)

		case "shwtask":
			showTask(conn)
		case "shwuser":
			database.ShowUsers(conn)
		case "getuser":
			conn.Write([]byte("Enter username:"))
			n, _ := conn.Read(input)
			username := string(input[0:n])
			p, _ := database.GetUser(username)
			fmt.Println("ID: ", p.ID, "; Name: ", p.Name)
			str := "ID: " + strconv.Itoa(p.ID) + "; Name: " + p.Name + "\n"
			conn.Write([]byte(str))
		case "end":
			fmt.Println("End")
			close(conn)
		}
	}
}

func showTask(conn net.Conn) {
	conn.Write([]byte(database.AllTasks()))
}

func addTask(executor string, task string) {
	if database.FindUser(executor) {
		p, _ := database.GetUser(executor)
		database.AddTask(p, task)
	}
}

func close(conn net.Conn) error {
	return conn.Close()
}
