package main

import "fmt"

func main() {

	// Next Greater Element -> NGE
	// her sayidan sonra gelen ilk buyuk sayi nedir?
	arr := []int{5, 7, 1, 3, 8, 2}

	var stack Stack
	result := make([]int, len(arr))

	for i := range result {
		result[i] = -1
	}

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] > arr[stack.Peek()] {
			top := stack.Peek()
			stack.Pop()
			result[top] = arr[i]
		}
		stack.Push(i)
	}

	fmt.Println("result:\t", result)
	fmt.Println("stack:\t", stack)
}

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}
