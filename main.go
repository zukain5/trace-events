package main

import (
	"os"

	"github.com/zukain5/trace-events/trace"
)

func main() {
	f, _ := os.Open("trace.out")
	result, err := trace.Parse(f, "")
	if err != nil {
		panic("Parse Error.")
	}
	events := result.Events
	trace.OutputCsv(events, "output.csv")
}
