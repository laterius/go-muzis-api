package stringutil

import (
	"math"
	"fmt"
)



func PrintStructLeaning(){
	//circle := Circle{1,2,3}
	//fmt.Println(circleArea1(&circle))
	//fmt.Println(circle.area())
	a := new(Android)
	a.Person.Talk()
}


type Android struct {
	Person
	Model string
}

type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

type Circle struct{
	x float64
	y float64
	z float64
}

func circleArea1(c *Circle) float64{
	return math.Pi * c.z * c.z
}

func (c *Circle) area() float64{
	return math.Pi * c.z * c.z
}