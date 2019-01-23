package main

import (
	"flag"
	"os"
)

var nFlag = flag.Bool("n", false, "print newline")

const NewLine = "\n"

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *nFlag { // -n is parsed, flag becomes true
				s += NewLine
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
