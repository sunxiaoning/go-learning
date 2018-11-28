package basic

import (
	"fmt"
)

type Person struct {
	name string
	age int32
}

/**
	面向对象
*/
func (p *Person) eat(){
	fmt.Println("person",p.name," eat!")
}

/**
	属性继承
*/
type Student struct {
	Person
	school string
}

/**
	面向过程
*/
func hello(student *Student){
	fmt.Println("Hello,My name is", student.name, ", I am ", student.age, "years old now!", "and My School is ",student.school)
}

/**
	面向对象
*/
func (student *Student) hi(){
	fmt.Println("Hello,My name is", student.name, ", I am ", student.age, "years old now!", "and My School is ",student.school)
}

/**
	面向对象覆盖
*/
func (student *Student) eat(){
	fmt.Println("student",student.name," eat!")
}

func Object(){
	william := Student{Person{"wiliam",25},"Tinghua"}
	hello(&william)
	var joyce Student
	joyce.name = "joyce"
	joyce.age = 25
	joyce.school = "Peking"
	hello(&joyce)
	john := Person{"john",28}
	john.eat()
	jack := Student{Person{"jack",30},"ShengDa"}
	jack.eat()
	jack.hi()

}