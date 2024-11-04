package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//https://go.dev/blog/slices-intro

	// slices are same as array with some extra power and convenience
	// NOTE: A slice can also be formed by “slicing” an existing slice or array.
	fruits := []string{"mango", "chickoo", "banana", "apple"} // this creates an array, then builds a slice that references it

	fmt.Println(fruits)
	fmt.Printf("%T\n", fruits)

	fruits_trimmed := fruits[1:]
	fmt.Println(fruits_trimmed)

	// NOTE: when we trim an array or slice it returns a referece view to slice/array
	// now if we change the 0th index of fruits_trimmed, this will also get reflects into fruits

	fruits_trimmed[0] = "pineapple"
	fmt.Println("fruits_trimmed[0]", fruits_trimmed[0], "fruits[1]", fruits[1])

	/* NOTE: slice header (is a part of the slice) is typically 24 bytes in 64-bit architecture
	1. Pointer (8 bytes) - Points to the start of the underlying array where the data is stored.
	2. Length (8 bytes) - Holds the number of elements in the slice.
	3. Capacity (8 bytes) - Holds the maximum number of elements the slice can accommodate without reallocating.
	*/

	// Slice length and capacity

	// NOTE: The length of a array is capacity of the array
	// The length is the number of elements referred to by the slice.
	// The capacity is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).
	v := []int{1, 34, 4, 0, 1}
	PrintLenCap(v)

	v = v[0:2]
	// v[2] = 2 // NOTE: this will give index out of bound, because 2 is greater than the length of the slice
	PrintLenCap(v)

	// create slice using make

	s := make([]int, 5) // allocates a contiguous memory of size 5 and also creates a slice of length 5
	s[4] = 2
	s[2] = 4
	s[1] = 2
	PrintLenCap(s) // [0 2 4 0 2] len 5 cap: 5

	s = s[:3]
	PrintLenCap(s) // [0 2 4] len 3 cap: 5

	// grow slice upto its capacity
	s = s[1:4]
	PrintLenCap(s) // [2 4 0] len 3 cap: 4

	// memory allocation in array and slices
	fmt.Println("Memory allocation in array")
	arr := [3][3]int{
		{1, 3, 4},
		{5, 6, 8},
		{2, 0, 9},
	} // int uses 8 bytes in 64-bit architecture

	fmt.Printf("address of arr: %d\n", uintptr(unsafe.Pointer(&arr)))
	fmt.Printf("address of arr[0]: %d\n", uintptr(unsafe.Pointer(&arr[0])))
	fmt.Printf("address of arr[0][0]: %d\n", uintptr(unsafe.Pointer(&arr[0][0])))
	fmt.Printf("address of arr[0][1]: %d\n", uintptr(unsafe.Pointer(&arr[0][1])))
	fmt.Printf("address of arr[0][2]: %d\n", uintptr(unsafe.Pointer(&arr[0][2])))
	fmt.Printf("address of arr[1]: %d\n", uintptr(unsafe.Pointer(&arr[1])))
	fmt.Printf("address of arr[1][0]: %d\n", uintptr(unsafe.Pointer(&arr[1][0])))
	fmt.Printf("address of arr[1][1]: %d\n", uintptr(unsafe.Pointer(&arr[1][1])))
	fmt.Printf("address of arr[1][2]: %d\n", uintptr(unsafe.Pointer(&arr[1][2])))
	fmt.Printf("address of arr[2]: %d\n", uintptr(unsafe.Pointer(&arr[2])))
	fmt.Printf("address of arr[2][0]: %d\n", uintptr(unsafe.Pointer(&arr[2][0])))
	fmt.Printf("address of arr[2][1]: %d\n", uintptr(unsafe.Pointer(&arr[2][1])))
	fmt.Printf("address of arr[2][2]: %d\n", uintptr(unsafe.Pointer(&arr[2][2])))

	fmt.Println("Memory allocation in slices")

	slice := [][]int{
		{1, 3, 4},
		{5, 6, 8},
		{2, 0, 9},
	} // int uses 8 bytes in 64-bit architecture

	// here each slice element in 2d slice is an reference to some array

	fmt.Printf("address of slice: %d\n", uintptr(unsafe.Pointer(&slice)))
	fmt.Printf("address of slice[0]: %d\n", uintptr(unsafe.Pointer(&slice[0])))
	fmt.Printf("address of slice[0][0]: %d\n", uintptr(unsafe.Pointer(&slice[0][0])))
	fmt.Printf("address of slice[0][1]: %d\n", uintptr(unsafe.Pointer(&slice[0][1])))
	fmt.Printf("address of slice[0][2]: %d\n", uintptr(unsafe.Pointer(&slice[0][2])))
	fmt.Printf("address of slice[1]: %d\n", uintptr(unsafe.Pointer(&slice[1])))
	fmt.Printf("address of slice[1][0]: %d\n", uintptr(unsafe.Pointer(&slice[1][0])))
	fmt.Printf("address of slice[1][1]: %d\n", uintptr(unsafe.Pointer(&slice[1][1])))
	fmt.Printf("address of slice[1][2]: %d\n", uintptr(unsafe.Pointer(&slice[1][2])))
	fmt.Printf("address of slice[2]: %d\n", uintptr(unsafe.Pointer(&slice[2])))
	fmt.Printf("address of slice[2][0]: %d\n", uintptr(unsafe.Pointer(&slice[2][0])))
	fmt.Printf("address of slice[2][1]: %d\n", uintptr(unsafe.Pointer(&slice[2][1])))
	fmt.Printf("address of slice[2][2]: %d\n", uintptr(unsafe.Pointer(&slice[2][2])))

	/* NOTE:
	slice[0], slice[1], slice[2] are the slice headers that points to some actuall array (row)
	since slice header is 24 bytes, you can see &slice[2] =  &slice[1] + 24, &slice[1] =  &slice[0] + 24
	also the address location of each row are contiguous but independent to other rows
	*/

	// Growing Slices (beyond its capacity)

	smallSlice := []int{1, 2, 4}
	PrintLenCap(smallSlice)
	smallSlice = smallSlice[:2]
	smallSlice = append(smallSlice, 3, 5)
	PrintLenCap(smallSlice)

	// implement append from scratch
	A := []int{1, 2, 3}
	PrintLenCap(A)

	A = Append(A[1:], 1, 4, 5)
	PrintLenCap(A)

	// passing slice to a function in Go == passing array to a function in C/C++
	// this is possible because slice itself holds the reference to the actual array
	B := []int{1, 3, 4}
	PrintLenCap(B)
	updateValueAtIndex(1, 1000, B)
	PrintLenCap(B)
}

func updateValueAtIndex(index int, value int, slice []int) {
	slice[index] = value
}

func PrintLenCap(s []int) {
	fmt.Println(s, "len", len(s), "cap:", cap(s))
}

func Append(slice []int, data ...int) []int {
	m := len(slice) // slice ref m size

	if slice == nil {
		slice = make([]int, len(data))
		m = 0
	}
	n := m + len(data) // total ref size needed for slice

	// if cap of slice is less to put the data into the arr
	if n > cap(slice) {
		// create a new slice ref for n capacity
		c := cap(slice[:])
		new_cap := 2 * c
		if n > new_cap {
			new_cap = n
		}
		new_slice := make([]int, new_cap)
		copy(new_slice, slice) // copy slice ref data into new_slice
		slice = new_slice      // update the ref to slice again
	}

	// add the new data to the end of the slice
	copy(slice[m:n], data)
	return slice
}
