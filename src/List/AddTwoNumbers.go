package List

import "fmt"

/**
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Example:
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		if l2 != nil {
			l2 = l2.reverse()
			return l2
		} else {
			return nil
		}
	}
	l1 = l1.reverse()
	l1.printAll()
	result := l1.addList(l2)
	result.printAll()
	return result
}

//reverse the list
func (l *ListNode) reverse() *ListNode {
	cur := l
	if !cur.hasNext() {
		return cur
	}
	var prev *ListNode = cur
	cur = cur.next()
	prev.Next = nil
	var next *ListNode
	for {
		if cur == nil {
			return prev
		}
		next = cur.next()
		cur.Next = prev
		prev = cur
		cur = next
	}
}

//add two list
func (l *ListNode) addList(src *ListNode) *ListNode {
	origin := l
	if src == nil {
		return origin.copy()
	}

	carry := 0
	cur := &ListNode{Val: carry, Next: nil}
	target := cur
	for {
		if origin != nil {
			cur.Val = cur.Val + origin.Val
			origin = origin.next()
			fmt.Println(origin)
		}
		if src != nil {
			cur.Val = cur.Val + src.Val
			src = src.next()
			fmt.Println(src)
		}
		if cur.Val >= 10 {
			carry = 1
			cur.Val = cur.Val - 10
		}
		if origin == nil && src == nil {
			if carry == 1 {
				cur.Next = &ListNode{Val: carry, Next: nil}
			}
			break
		}
		cur.Next = &ListNode{Val: carry, Next: nil}
		fmt.Println("add cur: ", cur)
		carry = 0
		cur = cur.next()
	}
	target = target.reverse()
	return target
}

//usage of judging the end
func (l *ListNode) hasNext() bool {
	if l.Next == nil {
		return false
	}
	return true
}

//get next node
func (l *ListNode) next() *ListNode {
	return l.Next
}

//copy whole list
func (l *ListNode) copy() *ListNode {
	if l == nil {
		return nil
	}
	target := &ListNode{Val: l.Val, Next: nil}
	for {
		if !l.hasNext() {
			break
		}
		l = l.next()
		cur := &ListNode{Val: l.Val, Next: nil}
		target.Next = cur
		target = target.next()
	}
	return target
}

func construct(args ...int) (result *ListNode) {
	if len(args) <= 0 {
		return
	}
	cur := &ListNode{Val:args[0], Next:nil}
	result = cur
	for i :=1; i<len(args); i++ {
		cur.Next = &ListNode{Val:args[i], Next:nil}
		cur = cur.next()
	}
	return
}

func (l *ListNode) toSlice() (s []int) {
	for {
		if l==nil {
			return
		}
		s = append(s, l.Val)
		l = l.next()
	}
}

func (l *ListNode) printAll() {
	for {
		if l==nil {
			return
		}
		fmt.Println("print: ", l)
		l = l.next()
	}
}