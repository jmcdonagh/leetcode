/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single
digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the
number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listLength(l *ListNode) int {
	var length int

	for n := l; n != nil; n = n.Next {
		length += 1
	}

	return length
}

func padList(l *ListNode, padLength int) {
	for n := l; n != nil; n = n.Next {
		l = n
	}

	for i := 0; i < padLength; i++ {
		l.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		l = l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var remainder int

	var n3 = &ListNode{
		Val:  0,
		Next: nil,
	}

	l1Length := listLength(l1)
	l2Length := listLength(l2)

	if l1Length > l2Length {
		padList(l2, l1Length - l2Length)
	}

	if l2Length > l1Length {
		padList(l1, l2Length - l1Length)
	}

	result := n3

	for n1, n2 := l1, l2; n1 != nil && n2 != nil; n1, n2 = n1.Next, n2.Next {
		val := n1.Val + n2.Val + n3.Val
		remainder = 0
		if val > 9 {
			remainder = 1
			val -= 10
		}

		n3.Val = val

		if remainder > 0 || n1.Next != nil {
			n3.Next = &ListNode{
				Val:  remainder,
				Next: nil,
			}
		}

		n3 = n3.Next
	}

	return result
}

func main() {
	n1 := ListNode{
		Val:  2,
		Next: nil,
	}

	n2 := ListNode{
		Val:  4,
		Next: nil,
	}

	n3 := ListNode{
		Val:  3,
		Next: nil,
	}

	n4 := ListNode{
		Val:  5,
		Next: nil,
	}

	n5 := ListNode{
		Val:  6,
		Next: nil,
	}

	n6 := ListNode{
		Val:  4,
		Next: nil,
	}
	n1.Next = &n2
	n2.Next = &n3

	n4.Next = &n5
	n5.Next = &n6

	fmt.Println("Initial")
	for r := addTwoNumbers(&n1, &n4); r != nil; r = r.Next {
		fmt.Println(r.Val)
	}

	n7 := ListNode{
		Val:  5,
		Next: nil,
	}

	n8 := ListNode{
		Val:  5,
		Next: nil,
	}

	fmt.Println("Two")
	for r := addTwoNumbers(&n7, &n8); r != nil; r = r.Next {
		fmt.Println(r.Val)
	}

	n9 := ListNode{
		Val:  1,
		Next: nil,
	}

	n10 := ListNode{
		Val:  8,
		Next: nil,
	}

	n11 := ListNode{
		Val:  0,
		Next: nil,
	}

	n9.Next = &n10
	fmt.Println("Three")
	for r := addTwoNumbers(&n9, &n11); r != nil; r = r.Next {
		fmt.Println(r.Val)
	}

	n12 := ListNode{
		Val:  1,
		Next: nil,
	}

	n13 := ListNode{
		Val:  8,
		Next: nil,
	}

	n14 := ListNode{
		Val:  3,
		Next: nil,
	}

	n15 := ListNode{
		Val:  7,
		Next: nil,
	}

	n16 := ListNode{
		Val:  1,
		Next: nil,
	}

	n12.Next = &n13
	n13.Next = &n14
	n15.Next = &n16

	fmt.Println("Four")
	for r := addTwoNumbers(&n12, &n15); r != nil; r = r.Next {
		fmt.Println(r.Val)
	}
}
