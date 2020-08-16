package main

import "fmt"

func main() {
	numa := [3]int{78, 79 ,80}
	//创建⼀个切⽚，包含整个数组的所有元素
	nums1 := numa[:]
	nums2 := numa[:]
	fmt.Println("array before change 1",numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
}
