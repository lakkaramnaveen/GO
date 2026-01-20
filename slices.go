package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := func(yield func(int) bool) {
		for i :=0; i< 10; i+=2 {
			if !yield(i){
				return
			}
		}
	}

	s1 := slices.AppendSeq([]int{1, 2}, seq)
	s2 := slices.AppendSeq([]int{12, 21}, seq)

	fmt.Println(s1)
	fmt.Println(s2)

}