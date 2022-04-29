package main

import (
	"strconv"
	"unicode"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct{}

func (n NestedInteger) IsInteger() bool           { return false }
func (n NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)     {}
func (n *NestedInteger) Add(elem NestedInteger)   {}
func (n NestedInteger) GetList() []*NestedInteger { return nil }

func deserialize(s string) *NestedInteger {
	if s[0] != '[' {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		ni := &NestedInteger{}
		ni.SetInteger(num)
		return ni
	}
	stack, num, negative := []*NestedInteger{}, 0, false
	for i, v := range s {
		if v == '-' {
			negative = true
		} else if unicode.IsDigit(v) {
			num = num*10 + int(v-'0')
		} else if v == '[' {
			stack = append(stack, &NestedInteger{})
		} else if v == ']' || v == ',' {
			if unicode.IsDigit(rune(s[i-1])) {
				if negative {
					num = -num
				}
				ni := NestedInteger{}
				ni.SetInteger(num)
				stack[len(stack)-1].Add(ni)
			}
			num, negative = 0, false
			if v == ']' && len(stack) > 1 {
				stack[len(stack)-2].Add(*stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}
	return stack[len(stack)-1]
}
func main() {

}
