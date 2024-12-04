package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	prev, cur, i, j := 0, 0, 0, 0
	count := len(nums1) + len(nums2)
	for i+j <= count/2 {
		prev = cur
		if len(nums1) <= i {
			cur = nums2[j]
			j++
		} else if len(nums2) <= j {
			cur = nums1[i]
			i++
		} else {
			if nums1[i] < nums2[j] {
				cur = nums1[i]
				i++
			} else {
				cur = nums2[j]
				j++
			}
		}
	}

	if count%2 == 0 {
		return float64(prev+cur) / 2
	}
	return float64(cur)
}

var nums1 = []int{1, 2}
var nums2 = []int{3, 4}
var nums3 = []int{1, 3}
var nums4 = []int{2}

var nums5 = []int{}
var nums6 = []int{1}

func main() {
	var result float64
	result = findMedianSortedArrays(nums1, nums2)
	if result == 2.5 {
		println("TestCase1 pass")
	} else {
		println(result)
	}
	result = findMedianSortedArrays(nums3, nums4)
	if result == 2 {
		println("TestCase2 pass")
	} else {
		println(result)
	}
	result = findMedianSortedArrays(nums5, nums6)
	if result == 1 {
		println("TestCase3 pass")
	} else {
		println(result)
	}
}
