package main

import (
	"fmt"
	"regexp"
	"strings"
)

var funcParser = regexp.MustCompile(`(?ms)^([a-zA-Z0-9_.-]+?):.?{(.+?)}\nreturn`)
var ifParser = regexp.MustCompile(`(?ms)if (.+?) {(.+?)}`)
var onKey = regexp.MustCompile(`(?ms)^([a-zA-Z0-9_.-^]+?)::(.+?)return`)
var blankLines = regexp.MustCompile("\n\n")

func parseFuncs(script string) string {
	matches := funcParser.FindAllStringSubmatch(script, -1)
	for _, match := range matches {
		exc := createExecutable(match[2])
		functions[match[1]] = func() { evalExec(exc) }
		script = strings.Replace(script, match[0], "", 1)
	}
	return script
}

func parseLine(line string) func() {
	return func() { fmt.Println(line) }
}

func createExecutable(script string) []func() {
	script = parseFuncs(script)
	exec := make([]func(), 0)

	matches := ifParser.FindAllStringSubmatch(script, -1)
	for _, match := range matches {
		beforeCode := before(script, match[0])
		exec = append(exec, createExecutable(beforeCode)...)
		script = strings.Replace(script, beforeCode, "", 1)
		exec = append(exec, func() {
			exc := createExecutable(match[2])
			if evaluateExpression(match[1]) {
				evalExec(exc)
			}
		})
		script = strings.Replace(script, match[0], "", 1)
	}

	script = blankLines.ReplaceAllString(script, "\n")

	for _, val := range strings.Split(script, "\n") {
		if len(val) > 0 && string(val[0]) != ";" {
			val = strings.TrimSpace(val)
			exec = append(exec, parseLine(val))
		}
	}

	return exec
}

func before(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}
