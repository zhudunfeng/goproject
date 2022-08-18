package main

func DuplicateNum(nums []int) int {
	//快慢指针找出数组中重复的数
	//确定快慢指针初次相遇位置
	slow, fast := 0, 0
	// slow = nums[slow]
	// fast = nums[nums[fast]]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	//快指针回到起点，重新追逐慢指针
	fast = 0
	for {
		if slow == fast {
			break
		}
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

func ChangeNum(nums []int) {
	nums[0] = 100
}
