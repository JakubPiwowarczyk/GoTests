package person

import (
	"fmt"
	"github.com/google/uuid"
)

var People = []Person{
	{uuid.New(), "Janek"},
	{uuid.New(), "John"},
	{uuid.New(), "Johan"},
	{uuid.New(), "Jonasz"},
}

type Person struct {
	Id   uuid.UUID
	Name string
}

func (p *Person) String() string {
	return fmt.Sprintf("Person{%v, %v}", p.Id, p.Name)
}

func New(name string) {
	person := Person{uuid.New(), name}
	People = append(People, person)
}
