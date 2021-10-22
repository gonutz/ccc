package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/gonutz/ccc"
)

func usage() {
	prog := filepath.Base(os.Args[0])
	fmt.Fprintf(
		os.Stderr,
		`Usage: %[1]s [-n SEED] [-f] [FILE] [-o OFFSET] < input > output

  %[1]s XORs standard input with a stream of bytes and writes the result to
  standard output. It can use a GUI on Windows.

  Examples:

    Show GUI on Windows, must have no options, just a file name:
      %[1]s C:\Folder\File

    XOR files on command line, must have -f option:
      %[1]s -f file1 < file2 > output

    XOR random numbers on command line:
      %[1]s -n 123 < input > output

    Skip first 100 bytes of XOR stream, works with -n and -f:
      %[1]s -n 123 -o 100 < input > output

  Use -n SEED to generate a byte stream from random numbers, using the given
  seed.
  Use -f FILE to use a file as the byte stream. If the file is shorter than
  standard input, it is repeated indefinitely.
  Pass -o OFFSET to skip the first OFFSET bytes of the XOR stream.
  As a special case on Windows, if you pass only a file path as the single
  parameter, a GUI opens to XOR that file.

`,
		prog)
	flag.PrintDefaults()
}

var (
	seed       = flag.Int("n", -1, "seed for the generated random number stream")
	file       = flag.String("f", "", "file to xor with standard input")
	fileOffset = flag.Int("o", 0, "byte offset to start in xor stream")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	var hasSeed, hasFile, hasFileOffset bool
	flag.Visit(func(f *flag.Flag) {
		hasSeed = hasSeed || f.Name == "n"
		hasFile = hasFile || f.Name == "f"
		hasFileOffset = hasFileOffset || f.Name == "o"
	})

	var xor io.Reader

	if hasSeed && !hasFile {
		random := rand.New(rand.NewSource(int64(*seed)))
		xor = ccc.FuncReader(func() byte {
			return byte(random.Intn(256))
		})
	} else if hasFile && !hasSeed {
		f, err := os.Open(*file)
		check("opening file", err)
		defer f.Close()
		xor = ccc.NewLoopReader(f)
	} else if !hasSeed && !hasFile && !hasFileOffset && len(args) == 1 {
		runGui(args[0])
		return
	} else {
		flag.Usage()
		return
	}

	if *fileOffset > 0 {
		_, err := io.CopyN(ioutil.Discard, xor, int64(*fileOffset))
		check("skipping to offset", err)
	}

	_, err := io.Copy(os.Stdout, ccc.NewXORReader(os.Stdin, xor))
	check("XORing streams", err)
}

func check(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erorr %s: %v\n", msg, err)
		os.Exit(1)
	}
}
