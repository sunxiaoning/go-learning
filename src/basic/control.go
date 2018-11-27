package basic

import (
	"fmt"
)

/**

	选择排序，又称打擂台算法，关键点指定一个最大或最小值，
	其他元素同他比，如果有更小或更大的则替换擂台
*/

func chooseSort(array []int){
	var temp int
	for i := 0; i < len(array); i++{
		minNum := i;
		for j := i + 1; j < len(array); j++{
			if array[minNum] > array[j] {
				minNum = j
			}
		}
		temp = array[i]
		array[i] = array[minNum]
		array[minNum] = temp
	}
}


/**
	冒泡排序，排序一次，待排序数组长度减1
	左边数组为待排序数组，右边排序为已排序数组
*/
func bubbleSort(array []int){
	var temp int;
	for i := 0; i < len(array); i++{
		for j := 0; j < len(array) - i - 1; j++{
			if array[j] < array[j+1] {
				temp = array[j];
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

func Controll() {
	array := []int {1,3,100,199,3,1,0,0,2,-1,666,1,100,44,2,1,1,0,4,3,1232,1,1,-100,-3,5,2}
	// chooseSort(array)
	bubbleSort(array)
	for _,num := range array {
		fmt.Println(num)
	}
	
}