package main

import (
	"github.com/Chaice1/Linter/internal/analyze"

	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyze.Analyzer}, nil
}

func main() {}
