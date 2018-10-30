package main

import (
	"fmt"

	"github.com/go-ego/gse"
)

var seg gse.Segmenter

func main() {
	hmm := seg.HMMCut("纽约时代广场")

	fmt.Println("hmm cut: ", hmm)
}
