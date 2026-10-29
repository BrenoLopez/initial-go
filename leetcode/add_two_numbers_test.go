package main

import (
	"reflect"
	"testing"
)

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Explanation: 342 + 465 = 807.
// Output: [7,0,8]
func TestAddTwoNumbersListSameLength(t *testing.T) {
	//Arrange
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	//Act
	result := addTwoNumbers(l1, l2)
	//Assert
	expected := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %v expected %v", result, expected)
	}
}

// Input: l1 = [0], l2 = [9,9,9,9]
// Explanation: 9999999 + 9999 = 10009998.
// Output: [8,9,9,9,0,0,0,1]
func TestAddTwoNumbersListWithZeroValue(t *testing.T) {
	//Arrange
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}
	//Act
	result := addTwoNumbers(l1, l2)
	//Assert
	expected := &ListNode{Val: 0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %v expected %v", result, expected)
	}
}

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Explanation: 9999999 + 9999 = 10009998.
// Output: [8,9,9,9,0,0,0,1]
func TestAddTwoNumbersListDifferentLength(t *testing.T) {
	//Arrange
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	//Act
	result := addTwoNumbers(l1, l2)
	//Assert
	expected := &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %v expected %v", result, expected)
	}
}
