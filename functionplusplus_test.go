package main

import "testing"

func TestSum(t *testing.T) {
	expected := []struct{
		params []int
		output int 
		}{
		{ []int{5, 10, 15}, 30 },
		{ []int{1, 1, 1}, 3 },
		{ []int{25, 26}, 51 },
		{ []int{20, 30, 40}, 90 },
		{ []int{25028492478347,9283938, 92738278372, 98293723742, 9389237823, 823928302932038, 9239283028321, 26}, (25028492478347+9283938+92738278372+98293723742+9389237823+823928302932038+9239283028321+26) },



	}
	for _, p := range expected{
		if Sum(p.params...) !=	p.output{
			t.Errorf("sum didn't fullfilled what it was asked for")
		}
	}
}
func TestGetMax(t *testing.T) {
	expected := []struct{
		params []int
		output int 
		}{
		{ []int{5, 10, 15}, 15 },
		{ []int{1, 1, 1}, 1 },
		{ []int{25, 26}, 26 },
		{ []int{20, 30, 40}, 40 },
		{ []int{}, 0 },
	}
	for _, p := range expected{
		if GetMax(p.params...) !=	p.output{
			t.Errorf("getmax didn't fullfilled what it was asked for")
		}
	}
}