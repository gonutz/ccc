package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/gonutz/cc"
)

func usage() {
	fmt.Fprint(os.Stderr, `usage: `+os.Args[0]+` -n SEED < input > output

    XORs standard input with a stream of random bytes generated
    using the given seed and writes the result to standard output.`)
}

var seed = flag.Int("n", -1, "seed for the generated random number stream")

func main() {
	if len(os.Args) == 2 {
		runGui(os.Args[1])
		return
	}

	flag.Parse()
	random := rand.New(rand.NewSource(int64(*seed)))
	r := cc.NewXORReader(os.Stdin, cc.NewFuncReader(func() byte {
		return byte(random.Intn(256))
	}))
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
