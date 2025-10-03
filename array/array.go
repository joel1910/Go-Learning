package main

import "fmt"

func main() {

	//int array
	arr := [5]int{2, 4, 6, 8}
	fmt.Println("Array :", arr)
	fmt.Println("Length :", len(arr))

	//string array
	strarr := [10]string{"joel", "raj", "weslin", "shabana", "fathima"}
	fmt.Println("String Array :", strarr)
	fmt.Println("String Length :", len(strarr))

	//change and add values
	arr1 := [5]int{10, 20, 30, 40}
	arr1[4] = 50  //add
	arr1[2] = 100 //change
	fmt.Println("Array1:", arr1)
	fmt.Println("Length:", len(arr1))

	arr2 := arr1 // copied
	arr2[0] = 99
	fmt.Println("Original:", arr1) // unchanged
	fmt.Println("Copy:", arr2)

	//here array is fixed - only used for fixed data (i.e) Week days[7], so we go for slice - Its dynamic

	//slice - pointer,length,capacity

	//pointer
	//point to underlying array

	arr3 := [5]int{10, 20, 30, 40, 50}
	slice := arr3[1:4] // slice points to arr elements 20,30,40
	fmt.Println("Slice:", slice)
	fmt.Println("Length of Slice:", len(slice))
	slice[0] = 99                            // changes the array, because slice points to it
	fmt.Println("Array after change:", arr3) // [10 99 30 40 50]

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", nums, "Length:", len(nums)) // Length = 5

	sub := nums[0:3]                                    //[1,2,3]
	fmt.Println("Sub-slice:", sub, "Length:", len(sub)) // Length = 3

	arr5 := [5]int{10, 20, 30, 40, 50}
	slice1 := arr5[1:3] // elements [20,30]

	fmt.Println("Slice:", slice1, "Length:", len(slice1), "Capacity:", cap(slice1))
	// len=2 (20,30), cap=4 ([20,30,40,50] cap start from index 1 to end of actual array)

	//append and shrink

	num1 := []int{1, 2, 3, 4}
	fmt.Println("Before Append :", num1)
	num1 = append(num1, 5, 6)
	fmt.Println("After Append :", num1)
	num1 = num1[1:2]
	fmt.Println("Shrink :", num1)

	// Shrink Syntax
	//arr[start:]     from start index to end of array
	//arr[:end]       from beginning to end-1
	//arr[start:end]  from start to end-1

	//copy array

	actualarray := []int{1, 2, 3, 4}
	copyarray := make([]int, 3)
	n := copy(copyarray, actualarray)

	fmt.Println("Copy Count :", n)
	fmt.Println("Actual Array :", actualarray)
	fmt.Println("Copied Array :", copyarray)

}
