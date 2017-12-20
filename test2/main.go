package main

import (
	"bufio"
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

	w := bufio.NewWriter(f)
	defer w.Flush()

	for i := 0; i < 100; i++ {
		fmt.Fprintf(w, "line number : %d\n", i)
	}

}
