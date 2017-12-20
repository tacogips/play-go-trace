package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	tf, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer tf.Close()

	err = trace.Start(tf)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < 100000; i++ {
		fmt.Fprintf(f, "line number : %d\n", i)
	}

}
