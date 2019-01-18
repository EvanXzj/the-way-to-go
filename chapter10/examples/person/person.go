package person

type Person struct {
	firstname, lastname string
}

func (p *Person) FirstName() string {
	return p.firstname
}

func (p *Person) SetFirstName(name string) {
	p.firstname = name
}
