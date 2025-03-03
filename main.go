package main

import "fmt"

// MenuComponent - component interface
type MenuComponent interface {
	Serve()
}

// Rice - leaf node
type Rice struct {
	Name string
}

func (r *Rice) Serve() {
	fmt.Println("serving item name : ", r.Name)
}

// FreshJuice - leaf node
type FreshJuice struct {
	Name string
}

func (f *FreshJuice) Serve() {
	fmt.Println("serving item name : ", f.Name)
}

// Buffet  - Composite node
type Buffet struct {
	Name  string
	Items []MenuComponent
}

func (b *Buffet) Add(item MenuComponent) {
	b.Items = append(b.Items, item)
}

func (b *Buffet) Serve() {
	fmt.Println("serving buffet name:", b.Name)
	for _, item := range b.Items {
		item.Serve()
	}
}

func main() {
	yellowRice := &Rice{Name: "yellow rice"}
	mangoJuice := &FreshJuice{Name: "Mango Juice"}

	dinnerBuffet := &Buffet{
		Name: "Dinner Buffet",
	}

	dinnerBuffet.Add(yellowRice)
	dinnerBuffet.Add(mangoJuice)

	dinnerBuffet.Serve()
}
