package main

import "fmt"

func main() {
	human := NewHuman("Petya", 10, humanDance)

	human.Speaks()

	//Демонстрация переопределения метода
	human.Dance("Break Dance")
	fmt.Println("")

	action := Action{
		Human: human,
		count: 4,
	}

	act := NewAction("Petya", 10, 4, customDance)

	action.Speaks()
	fmt.Println("")

	action.DoSomething()
	fmt.Println("")

	//Демонстрация переопределения метода
	act.Dance("Electric Boogie")
}

type Human struct {
	Name         string
	YearsOld     uint8
	callBackFunc func(danceStyle string)
}

func NewHuman(name string, yearsOld uint8, dance func(danceStyle string)) Human {
	return Human{Name: name, YearsOld: yearsOld, callBackFunc: dance}
}

func (h *Human) Dance(danceStyle string) {
	h.callBackFunc(danceStyle)
}

// Реализация для Human
func humanDance(danceStyle string) {
	fmt.Printf("He or She dance %s\n", danceStyle)
}

func (h *Human) Speaks() {
	fmt.Printf("My name is %s. I am %d years old.\n", h.Name, h.YearsOld)
}

type Action struct {
	Human
	count int
}

func NewAction(name string, yearsOld uint8, count int, dance func(danceStyle string)) Action {
	return Action{Human: Human{
		Name:         name,
		YearsOld:     yearsOld,
		callBackFunc: dance,
	}, count: count}
}

func (a *Action) DoSomething() {
	for i := 0; i < a.count; i++ {
		a.Speaks()
	}
}

// Реализация для Action
func customDance(danceStyle string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("I like to dance %s\n", danceStyle)
	}
}
