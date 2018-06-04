package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lapitsky/spreadsheet/cells"
	"github.com/lapitsky/spreadsheet/csvparser"
	"github.com/lapitsky/spreadsheet/expreval"
)

func main() {
	input := os.Stdin

	warningsPtr := flag.Bool("w", false, "show warnings")
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 1 {
		var err error
		input, err = os.Open(flag.Args()[0])
		if err != nil {
			panic(err)
		}
	} else if flag.NArg() > 1 {
		panic("Too many files, only one file parameter is allowed")
	}

	if csvValues, err := csvparser.Read(input); err == nil {
		cells := cells.NewCells(csvValues)
		evaluateCells(cells)
		if *warningsPtr {
			outputWarnings(cells)
		}
	} else {
		panic(err)
	}
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] [file]\n", os.Args[0])
	flag.PrintDefaults()
}

func evaluateCells(cs *cells.Cells) {
	for i := range cs.Cells {
		for j := range cs.Cells[i] {
			val := expreval.Evaluate(cs, &cs.Cells[i][j])
			if j == 0 {
				fmt.Printf("%s", val)
			} else {
				fmt.Printf(",%s", val)
			}
		}
		fmt.Println("")
	}
}

func outputWarnings(cs *cells.Cells) {
	for i := range cs.Cells {
		for j, c := range cs.Cells[i] {
			if c.Value.Error() != nil {
				fmt.Fprintf(os.Stderr, "%s%d: %s\n", string('A'+j), i+1, c.Value.Error())
			}
		}
	}
}
