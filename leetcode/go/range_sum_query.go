package main

type NumArray struct {
	prefixS []int
}

func Constructor(nums []int) NumArray {
	prefixS := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefixS = append(prefixS, prefixS[len(prefixS)-1]+nums[i])
	}

	return NumArray{prefixS: prefixS}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixS[right]
	} else {
		return this.prefixS[right] - this.prefixS[left-1]
	}
}
