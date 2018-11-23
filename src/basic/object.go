package basic

import (
	"fmt"
)

type Person struct {
	name string
	age int32
}

type Student struct {
	Person
	school string
}

func hello(student *Student){
	fmt.Println("Hello,My name is", student.name, ", I am ", student.age, "years old now!", "and My School is ",student.school)
}

func Object(){
	william := Student{Person{"wiliam",25},"Tinghua"}
	hello(&william)
	var joyce Student
	joyce.name = "joyce"
	joyce.age = 25
	joyce.school = "Peking"
	hello(&joyce)
}