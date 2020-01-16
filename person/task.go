package person

type Task struct {
	Text      string
	Performer Person
}

func (task *Task) AddTask(text string, person *Person) {
	task.Text = text
	task.Performer = *person
}
