package main

import (
	"com.swk/goIntro/pipeline"
	"fmt"
)

func main() {
	p:=pipeline.Merge(pipeline.ImMemorySort(pipeline.ArraySource(3,2,6,7,4, 9)),
			pipeline.ImMemorySort(pipeline.ArraySource(7,4,0,3,2,8,19,4)))
	for v:=range p{
		fmt.Println(v)
	}
	
}
