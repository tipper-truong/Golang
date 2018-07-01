package main

import "fmt"

type Person struct {
	Name    string
	Address Address // embedded types, has-a relationship
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("I’m at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

type Citizen struct {
	Country string
	Person // pseudo is-a relationship
}

func (c *Citizen) Nationality() {
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

func main() {
	c := Citizen{}
	c.Name = "Steve"
	c.Country = "America"
	c.Talk()
	c.Nationality()
}
