package main

import "fmt"

func Sum(nums ...int) int{
	if len(nums) == 2{
		return nums[0] + nums[1]
	}else if len(nums) == 1{
		return Sum(nums[0], 0)
	}
	return nums[0] + Sum(nums[1:]...)
	
}
func GetMax(nums ...int)int{
	if len(nums) == 0{
		return 0
	}
	max := nums[0]
	for _, n := range nums{
		if n > max{
			max = n
		}
	}
	return max
}
func GetValues(a int) (double, triple, quad int){
	double = a * 2
	triple = a * 3
	quad = a * 4
	return
}
func FunctionsPlusPlus(){
	fmt.Println(Sum(3, 4, 3, 10))
	fmt.Println(Sum(GetValues(1)))
}