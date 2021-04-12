package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	start int
}

// 类似于Java中的toString
func (p Human) String() string {
	return fmt.Sprintf("%v (%v years)-------------", p.Name, p.Age)
}

type Interval struct {
	start int
	end   int
	Human
}

func (inter Interval) String() string {
	return fmt.Sprintf("name: %v, age: %v, start: %v, end: %v",
		inter.Name, inter.Age, inter.start, inter.end)
}

func main() {

	man := new(Human)
	man.Name = "Tom"
	man.Age = 12

	a := Human{"爱丽丝", 18, 100}
	z := Human{"FBI张三", 48, 100}
	fmt.Println(a, z)

	str := fmt.Sprintf("%v-----%v", man.Name, man.Age)
	fmt.Println(str)
	inter := Interval{0, 3, Human{"Nilson", 15, 100}}
	fmt.Println(inter)

	inter1 := Interval{start: 10, end: 20, Human: Human{Name: "爱丽丝", Age: 100,start: 100}}
	fmt.Println(inter1)
}
