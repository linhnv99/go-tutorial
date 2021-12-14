package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Bird interface {
	Fly()
}
//Biểu thị Cat implement Animal
func NewCat() Animal {
	return &Cat{}
}

func NewBird() Bird {
	return &Falcon{}
}

type Cat struct {
	name string
}

type Falcon struct {
	name string
}

func (c Cat) Eat() {
	fmt.Println("Cat: ", c.name)
}

func (c Cat) Sleep() {
	panic("implement me")
}

func (f Falcon) Eat() {
	fmt.Println("Falcon: ", f.name)

}

func (f Falcon) Sleep() {
	fmt.Println("Flying", f.name)
}

func (f *Falcon) Fly() {
	f.name = "Flajsdhkasd"
	fmt.Println("Flying")
}

func main() {
	c := Cat{"Linh"}
	c.Eat()
	f := &Falcon{"abc"}
	f.Fly()
	fmt.Println(f.name)
}
