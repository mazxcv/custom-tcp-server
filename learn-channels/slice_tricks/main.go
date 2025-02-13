package slicetricks

import (
	"fmt"
	"sort"
)

type Numbers []int

type byInc struct {
	Numbers
}

func (n byInc) Len() int {
	return len(n.Numbers)
}

func (n byInc) Swap(i, j int) {
	n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i]
}

func (n byInc) Less(i, j int) bool {
	return n.Numbers[i] < n.Numbers[j]
}

type byDec struct {
	Numbers
}

func (n byDec) Len() int {
	return len(n.Numbers)
}

func (n byDec) Swap(i, j int) {
	n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i]
}

func (n byDec) Less(i, j int) bool {
	return n.Numbers[i] > n.Numbers[j]
}

func (n Numbers) Len() int {
	return len(n)
}

func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Numbers) Less(i, j int) bool {
	return n[i] < n[j]
}

type Person struct {
	Name string
	Age  int
}

func slice() {
	number := Numbers{2, 3, 1, 4, 5}
	sort.Sort(byInc{number})
	fmt.Println(number)
	sort.Sort(byDec{number})
	fmt.Println(number)
}
