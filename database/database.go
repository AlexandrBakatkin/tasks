package database

import (
	"database/sql"
	"github.com/Bakatkin/tasks/person"
	_ "github.com/go-sql-driver/mysql"
)

func FindUser(user string) bool {
	var str string
	db, err := sql.Open("mysql", "root:svpunk@/tasks_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row, err := db.Query("SELECT person_id FROM tasks_db.persons_tbl WHERE person_name = ?;", user)
	if err != nil {
		panic(err)
	}
	for row.Next() {
		row.Scan(&str)
	}
	return str != ""
}

func GetUser(user string) (person.Person, error) {
	p := person.Person{}
	db, err := sql.Open("mysql", "root:svpunk@/tasks_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT person_id, person_name FROM tasks_db.persons_tbl WHERE person_name = ?;", user)
	if err != nil {
		panic(err)
	}
	for row.Next() {

		row.Scan(&p.ID, &p.Name)
	}
	return p, err
}

func AddUser(username string) {
	db, err := sql.Open("mysql", "root:svpunk@/tasks_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `tasks_db`.`persons_tbl` (`person_name`) VALUES (?);", username)
	if err != nil {
		panic(err)
	}
}

func ShowUsers() string {
	var s string
	db, err := sql.Open("mysql", "root:svpunk@/tasks_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT person_name FROM tasks_db.persons_tbl;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var t string
		rows.Scan(&t)
		s = s + t + "|"
	}
	defer rows.Close()
	s = s + "\n"
	return s
}

func AddTask(executor person.Person, task string) {
	db, err := sql.Open("mysql", "root:svpunk@/tasks_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO `tasks_db`.`tasks_tbl` (`task_name`, `task_perfomer`) VALUES (?, ?);", task, executor.Name)
}

func AllTasks() string {
	var s string
	db, err := sql.Open("mysql", "root:svpunk@/tasks_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT task_name FROM tasks_db.tasks_tbl;")
	for rows.Next() {
		var t string
		rows.Scan(&t)
		s = s + t + "|"
	}
	return s + "\n"
}
