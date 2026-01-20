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

	s := slices.AppendSeq([]int {1,2}, seq)
	fmt.Println(s)
}