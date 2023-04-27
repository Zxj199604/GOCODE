package main

import (
	"fmt"
)

// Animal 公共接口
type Animal interface {
	Speak() string
	Eat() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "I'm " + d.name
}

func (d Dog) Eat() string {
	return "I like eat poke!"
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return "I'm " + c.name
}

func (c Cat) Eat() string {
	return "I like eat fish!"
}

//interface{} 类型是没有方法的接口。由于没有 implements 关键字，所以所有类型都至少实现了 0 个方法，所以 所有类型都实现了空接口。
func DoSomothing(v interface{}) {
	fmt.Println(v)
}

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	fmt.Println("project01")
	dog := Dog{name: "dog!"}
	cat := Cat{name: "cat!"}
	animals := []Animal{}
	animals = append(animals, dog)
	animals = append(animals, cat)

	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i].Speak())
	}

	for _, animal := range animals {
		fmt.Println(animal.Eat())
	}
	// animals := []Animal{Dog{name: "dog"}, Cat{name: "cat"}}
	DoSomothing("ddd")
	DoSomothing(animals)

	//[]interface{} 要进行转换
	vals := make([]interface{}, len(animals))
	for i, v := range animals {
		vals[i] = v
	}
	PrintAll(vals)
}
