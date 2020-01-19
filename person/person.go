package person

type Person struct {
	ID   int
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetID() int {
	return person.ID
}

func (person Person) NewPerson(id int, name string) {
	person.ID = id
	person.Name = name
}
