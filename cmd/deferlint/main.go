package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/eatonphil/deferlinter"
)

func main() {
	singlechecker.Main(deferlinter.Analyzer)
}
