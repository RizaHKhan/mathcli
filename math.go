package main

import (
	"fmt"
	"os"

	"github.com/maja42/goval"
)

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	expression := os.Args[1]
	eval := goval.NewEvaluator()
	result, err := eval.Evaluate(expression, nil, nil)
	Error(err)

	fmt.Println(result)
}
