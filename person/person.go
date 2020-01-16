package person

type Person struct {
	Name string
}

func (person *Person) GetName() string {
	return person.Name
}
