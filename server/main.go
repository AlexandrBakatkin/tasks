package main

import (
	"fmt"
	"github.com/Bakatkin/tasks/person"
	"net"
)

func main() {
	const (
		startmsg = "Server is starting..."
	)

	tasks := make(map[string]person.Task)
	pers := make([]person.Person, 0)
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
		go handleConnection(conn, &pers, &tasks)
	}
}

func handleConnection(conn net.Conn, pers *[]person.Person, tasks *map[string]person.Task) {
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
			addUser(username, pers)
			conn.Write([]byte("OK\n"))
		case "addtask":
			conn.Write([]byte("Entering task... Select executor:\n"))
			showUsers(pers, conn)
			n, _ := conn.Read(input)
			executor := string(input[0:n])
			fmt.Println(executor)
			conn.Write([]byte("Enter task:\n"))
			n, _ = conn.Read(input)
			task := string(input[0:n])
			fmt.Println(task)
			addTask(executor, task, pers, tasks)
			//addTask()

		case "shwtask":
			conn.Write([]byte("Command showtask..."))
		case "shwuser":
			showUsers(pers, conn)
		case "end":
			fmt.Println("End")
			close(conn)
		}
	}
}

func addTask(executor string, task string, pers *[]person.Person, tasks *map[string]person.Task) {
	for _, n := range *pers {
		if n.Name == executor {
			//tasks[executor] = person.Task{Text:task, Performer:n}
		}
	}
}

func close(conn net.Conn) error {
	return conn.Close()
}

func showUsers(pers *[]person.Person, conn net.Conn) {
	var s string
	for i, str := range *pers {
		fmt.Println(str.GetName())
		if i != 0 {
			s = str.GetName() + " " + s
		} else {
			s = str.GetName()
		}
	}
	s = s + "\n"
	conn.Write([]byte(s))
}

func addUser(username string, pers *[]person.Person) {
	var person = person.Person{Name: username}
	*pers = append(*pers, person)
}
