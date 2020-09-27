package main

/*
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Example 4:
Input: s = ""
Output: 0
*/

// This one works a lot better, but still not the best you can get.
func lengthOfLongestSubstring(s string) int {
	seen := make(map[string]int, len(s))
	var counter int
	var max int
	var newStart int

	for i, c := range s {
		if j, ok := seen[string(c)]; !ok {
			counter++
		} else {
			// If we saw this in a previous subString, set j to beginning of
			// this substring
			if j < newStart {
				j = newStart
			}

			//  Distance since last appearance is the starting count value
			counter = i - j
			newStart = i - counter
		}

		if counter > max {
			max = counter
		}
		seen[string(c)] = i
	}

	return max
}

// This one works but is insanely slow
// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 1 {
// 		return 1
// 	}
//
// 	var max int
// 	sub := s
//
// 	seen := make(map[string]int, len(sub))
//
// 	for i, _ := range s {
// 		sub = s[i:]
// 		for j, subC := range sub {
// 			if _, ok := seen[string(subC)]; !ok {
// 				seen[string(subC)] = j
// 			} else {
// 				if len(seen) > max {
// 					max = len(seen)
// 				}
// 				seen = make(map[string]int, len(sub))
// 				break
// 			}
// 		}
// 	}
//
// 	return max
// }

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
// }
