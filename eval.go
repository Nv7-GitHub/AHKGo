package main

import "fmt"

func evaluateExpression(expr string) bool {
	fmt.Println(expr)
	return true
}

func evalExec(exec []func()) {
	for _, val := range exec {
		val()
	}
}
