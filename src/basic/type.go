package basic

import "fmt"

func Type(){
	var i int = 3
	fmt.Println("i =", i)
	var j, k = 10,100
	fmt.Println("j =", j, " ,k =", k)
	m,n := 8,9
	fmt.Println("m =", m , ", n = ", n)
	const PI = 3.14;
	fmt.Println("PI =", PI)
	enabled := true
	fmt.Println("enabled =",enabled)
	var canRead bool = true
	fmt.Println(canRead)
	str1 := "hello"
	str2 := "world"
	fmt.Println(str1 + ","+ str2)
	fmt.Println(`this is ok
					? is ok?`)
	fmt.Println(str1[1:])
	fmt.Println(str2[:3])
	fmt.Println(str1[1:3])
	
	const(
		a1,a2,a3,a4 = iota,iota,iota,iota
	)
	fmt.Println(a1,a2,a3,a4)
	const(
		a5 = iota
		a6 = iota
		a7 = iota
	)
	fmt.Println(a5,a6,a7);
	const a8 = iota;
	fmt.Println(a8)
	var array1 = [3]byte{'1','2','3'} 
	fmt.Println(array1[0],array1[1])
	 var arrray2 []byte
	 arrray2 = array1[:]
	 fmt.Println(arrray2[0], arrray2[1])
	 var map1 map[int]string
	 map1 = make(map[int]string)
	 map1[0] = "a1"
	 map1[1] = "a2"
	 fmt.Println(map1[0], map1[1])
	 map2 := map[string]int{"C":1,"JAVA":2}
	 fmt.Println(map2["JAVA"])
	 b1,ok := map2["C#"]
	 fmt.Println(b1, ok)


}