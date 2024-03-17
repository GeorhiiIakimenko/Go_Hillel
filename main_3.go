package main

import (
	"fmt"
)

type Animal struct {
	runSpeed int
	name     string
	voice    string
}

func (a *Animal) WinnerSay() {
	fmt.Printf("**Крутезний крик %s**\n", a.name)
}

func (a *Animal) LoserSay() {
	fmt.Printf("**Крик лузера від %s**\n", a.name)
}

type Turtle struct {
	Animal
	uniqueProperty1 int    // Наприклад, кількість плавників
	uniqueProperty2 string // Наприклад, колір панциря
}

type Tiger struct {
	Animal
	uniqueProperty1 int    // Наприклад, кількість смуг
	uniqueProperty2 string // Наприклад, розмір лап
}

type Fish struct {
	Animal
	uniqueProperty1 int    // Наприклад, довжина тіла
	uniqueProperty2 string // Наприклад, тип чіплячка
}

type Race struct {
	Distance int
	Turtle   *Turtle
	Tiger    *Tiger
	Fish     *Fish
}

func (r *Race) CreateTeam() {
	r.Turtle = &Turtle{Animal{runSpeed: 20, name: "Черепаха", voice: "Брррр"}, uniqueProperty1: 4, uniqueProperty2: "Зелений"}
	r.Tiger = &Tiger{Animal{runSpeed: 70, name: "Тигр", voice: "Рррр"}, uniqueProperty1: 6, uniqueProperty2: "Великий"}
	r.Fish = &Fish{Animal{runSpeed: 50, name: "Риба", voice: "Буль-буль"}, uniqueProperty1: 1, uniqueProperty2: "Малий"}
}

func (r *Race) Start() {
	var time int
	for {
		if time%r.Turtle.runSpeed == 0 {
			r.Turtle.WinnerSay()
		}
		if time%r.Tiger.runSpeed == 0 {
			r.Tiger.WinnerSay()
		}
		if time%r.Fish.runSpeed == 0 {
			r.Fish.WinnerSay()
		}

		time++
		if time > r.Distance {
			break
		}
	}

	fmt.Printf("Черепаха пробігла - %d ітерацій\n", time/r.Turtle.runSpeed)
	fmt.Printf("Тигр пробіг - %d ітерацій\n", time/r.Tiger.runSpeed)
	fmt.Printf("Риба пробігла - %d ітерацій\n", time/r.Fish.runSpeed)

	r.Turtle.LoserSay()
}
func main() {
	race := Race{Distance: 1000}
	race.CreateTeam()
	race.Start()
}
