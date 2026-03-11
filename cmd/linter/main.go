package main

import (
	"github.com/Chaice1/Linter/internal/analyze"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyze.Analyzer)
}
