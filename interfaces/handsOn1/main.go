package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type SecretAgent struct {
	Person
	AgentNumber        int
	SuccessfulMisssion int
}

func main() {
	sa1 := SecretAgent{
		Person{"Arnaldo", 26},
		26263,
		24,
	}

	p1 := Person{"Alarico", 21}
	fmt.Println(p1.Name, "\n")
	p1.Speak()
	fmt.Println("agent Number :", sa1.AgentNumber, "\n")
	sa1.Speak()
	sa1.Person.Speak()

}

func (p Person) Speak() {
	fmt.Printf(" Hi I am Person %s \n", p.Name)
}

func (sa SecretAgent) Speak() {
	fmt.Printf(" Hi I am Secret Agent %s \n", sa.Person.Name)
}
