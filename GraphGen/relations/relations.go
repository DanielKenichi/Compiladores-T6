package relations

import (
	"log"
	"slices"
)

type GroupName = string
type RelationName = string
type PersonName = string

type Relations struct {
	Groups  map[GroupName][]*Person
	Persons map[PersonName]*Person
}

type Person struct {
	Name          PersonName
	Relationships map[RelationName][]*Person
}

func NewRelations() *Relations {
	return &Relations{
		Groups:  make(map[string][]*Person),
		Persons: make(map[string]*Person),
	}
}

func NewPerson(personName string) *Person {
	return &Person{
		Name:          personName,
		Relationships: make(map[string][]*Person),
	}
}

func (r *Relations) AddNewGroup(groupName string) {
	r.Groups[groupName] = make([]*Person, 0)
}

func (r *Relations) AddNewPerson(personName string) {
	r.Persons[personName] = NewPerson(personName)
}

func (r *Relations) AddPersonToGroup(person *Person, groupName string) {
	groupMembers, ok := r.Groups[groupName]

	if !ok {
		log.Print("Something is wrong, group should be defined")
		return
	}

	if slices.Contains(groupMembers, person) {
		return
	}

	groupMembers = append(groupMembers, person)

	r.Groups[groupName] = groupMembers
}

func (r *Relations) GetGroupMembers(groupName string) []*Person {
	return r.Groups[groupName]
}

func (r *Relations) GetPerson(personName string) *Person {
	return r.Persons[personName]
}

func (p *Person) AddPersonToRelationship(person *Person, relation string) {

	persons, ok := p.Relationships[relation]

	if !ok {
		persons = make([]*Person, 0)
	}

	if slices.Contains(persons, person) {
		return
	}

	persons = append(persons, person)

	p.Relationships[relation] = persons
}
