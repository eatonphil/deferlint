package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/eatonphil/deferlint"
)

func main() {
	singlechecker.Main(deferlint.Analyzer)
}
