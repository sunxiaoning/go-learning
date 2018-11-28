package basic

import (
	"fmt"
)

type obj interface{}

type Teacher struct {
	Person
	student []string
}

type Door interface {
	open()
	close()
}

type Teach interface {
	teach()
}

type Parent struct {
	Person
	child string
}

func (teacher Teacher)teach(){
	fmt.Println("teacher ",teacher.name, "is teaching!")
}

func (parent Parent)teach(){
	fmt.Println("parent ",parent.name, "is teaching!")
}

func (person Person) open(){
	fmt.Println("person",person.name, "open the door!")
}

/**
	实现接口时，不能使用对象指针类型
*/
func (person Person) close(){
	fmt.Println("person",person.name, "close the door!")
}

func Interface(){

	// 空interface任意类型变量
	var a1 obj
	a1 = 3
	fmt.Println(a1)
	a1 = "hello"
	fmt.Println(a1)
	var teach Teach
	teacherZhang := Teacher{Person{"zhang",32},[]string{"xiaoming","lihua"}}
	teach = teacherZhang
	teach.teach()
	teach = Parent {Person{"wang",45}, "tom"}
	teach.teach()

	// 实现一个接口时必须实现其所有的方法
	var door Door
	door = teacherZhang
	door.open()
	door.close()
}