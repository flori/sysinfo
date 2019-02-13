package main

import (
	"flag"
	"fmt"

	cpuload "github.com/flori/cpuload"
)

var options cpuload.Options

func init() {
	flag.Float64Var(
		&options.LoadAvg,
		"load-avg",
		-1.0,
		"load average",
	)
	flag.BoolVar(
		&options.Horizontal,
		"horizontal",
		false,
		"display horizontal bar instead of vertical",
	)
	flag.IntVar(
		&options.ProcessorCount,
		"processor-count",
		-1,
		"number of processor cores/threads",
	)
	flag.Parse()
}

func main() {
	fmt.Print(cpuload.NewBar(options))
}
