package person

type Task struct {
	ID        int
	Text      string
	Performer Person
}

func (task Task) NewTask(id int, text string, person Person) {
	task.ID = id
	task.Text = text
	task.Performer = person
}
