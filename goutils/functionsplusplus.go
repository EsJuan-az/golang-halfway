package goutils

import "fmt"

func sum(nums ...int) int{
	if len(nums) == 2{
		return nums[0] + nums[1]
	}else if len(nums) == 1{
		return sum(nums[0], 0)
	}
	return nums[0] + sum(nums[1:]...)
	
}
func getValues(a int) (double, triple, quad int){
	double = a * 2
	triple = a * 3
	quad = a * 4
	return
}
func FunctionsPlusPlus(){
	fmt.Println(sum(3, 4, 3, 10))
	fmt.Println(sum(getValues(1)))
}