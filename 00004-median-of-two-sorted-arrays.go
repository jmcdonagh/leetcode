/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

Follow up: The overall run time complexity should be O(log (m+n)).
*/
package main

// This merges the array then gets the median. Although it is inefficient to do
// it this way, this it's easy to follow. Alternative form below (which can also
// be simplified greatly, but shows the general technique that is faster than
// this.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))

	var x, y, z int
	outer:
	for x <= len(nums1) || y <= len(nums2) {
		switch {
		case x == len(nums1) && y == len(nums2):
			break outer
		case x == len(nums1) && y < len(nums2):
			for ; y < len(nums2); y++ {
				merged[z] = nums2[y]
				z++
			}
		case y == len(nums2) && x < len(nums1):
			for ; x < len(nums1); x++ {
				merged[z] = nums1[x]
				z++
			}
		case nums1[x] < nums2[y]:
			merged[z] = nums1[x]
			x++
			z++
		case nums2[y] < nums1[x]:
			merged[z] = nums2[y]
			y++
			z++
		case nums1[x] == nums2[y]:
			merged[z], merged[z+1] = nums1[x], nums2[y]
			x++
			y++
			z += 2
		}
	}

	mergedLength := len(merged)
	var median float64
	if mergedLength % 2 == 1 {
		median = float64(merged[mergedLength/2.0])
	} else {
		one := float64(merged[mergedLength / 2.0 - 1])
		two := float64(merged[mergedLength / 2.0])
		median = (one + two) / 2.0
	}
	return median
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	mergedLength := len(nums1) + len(nums2)
// 	middle := float64(mergedLength) / 2.0 + 0.5
// 	var median float64
//
// 	var x, y, z int
// 	var prev, cur int
//
// 	for float64(z) <= middle {
// 		switch {
// 		case float64(z) == middle:
// 			return float64(cur)
// 		case x == len(nums1) && y < len(nums2):
// 			for ; y < len(nums2); y++ {
// 				prev = cur
// 				cur = nums2[y]
// 				z++
// 				if float64(z) >= middle {
// 					break
// 				}
// 			}
// 		case y == len(nums2) && x < len(nums1):
// 			for ; x < len(nums1); x++ {
// 				prev = cur
// 				cur = nums1[x]
// 				z++
// 				if float64(z) >= middle {
// 					break
// 				}
// 			}
// 		case nums1[x] < nums2[y]:
// 			prev = cur
// 			cur = nums1[x]
// 			x++
// 			z++
// 		case nums2[y] < nums1[x]:
// 			prev = cur
// 			cur = nums2[y]
// 			y++
// 			z++
// 		case nums1[x] == nums2[y]:
// 			prev = cur
// 			cur = nums1[x]
// 			x++
// 			z++
// 			if float64(z) >= middle {
// 				break
// 			}
// 		}
// 	}
//
// 	sum := prev + cur
// 	median = float64(sum) / 2.0
// 	return median
// }


// func main() {
// 	fmt.Println(findMedianSortedArrays([]int{0,0,0,0,0}, []int{-1,0,0,0,0,0,1}))
// }
