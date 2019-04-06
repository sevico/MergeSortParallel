package main

import (
	"bufio"
	"com.swk/goIntro/pipeline"
	"fmt"
	"os"
)

func main() {
	
	p:= createPipeline("small.in",512,4)
	writeToFile(p,"small.out")
	printFile("small.out")
}
func printFile(filename string)  {
	file,err:=os.Open(filename)
	if err!= nil {
		panic(nil)
	}
	defer file.Close()
	p:=pipeline.ReaderSource(file,-1)
	for v:=range p{
		fmt.Println(v)

	}
}
func writeToFile(p <-chan int, filename string) {
	file,err:=os.Create(filename)
	if err!= nil {
		panic(err)
	}
	defer file.Close()
	writer:=bufio.NewWriter(file)
	defer writer.Flush()
	pipeline.WriterSink(writer,p)
}
func createPipeline(filename string,fileSize,chunkCount int) <-chan int {
	chunkSize := fileSize/chunkCount
	sortResult :=[]<-chan int{}
	for i:=0;i<chunkCount;i++{
		file,err:=os.Open(filename)
		//defer file.Close()
		if err!= nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source:=pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
        sortResult=append(sortResult,pipeline.InMemorySort(source))

	}
	return pipeline.MergeN(sortResult...)

}
