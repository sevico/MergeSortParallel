package main

import (
	"bufio"
	"fmt"
	"os"
)
import "../../pipeline"
func MergeDemo() {
	p:=pipeline.Merge(pipeline.InMemorySort(pipeline.ArraySource(3,2,6,7,4, 9)),
		pipeline.InMemorySort(pipeline.ArraySource(7,4,0,3,2,8,19,4)))
	for v:=range p{
		fmt.Println(v)
	}
}
func main() {
	//MergeDemo()
	const filename="small.in"
	const n =64
	file,err:=os.Create(filename)
	defer file.Close()
    if err!= nil {
    	panic(err)
	}
	p:=pipeline.RandomSource(n)
	writer:=bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()
	file,err=os.Open(filename)
	if err!= nil {
		panic(err)
	}
	defer file.Close()
	p=pipeline.ReaderSource(bufio.NewReader(file),-1)
	for v:= range p{
		fmt.Println(v)

	}

}
