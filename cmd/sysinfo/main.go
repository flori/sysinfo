package main

import (
	"flag"
	"fmt"

	load "github.com/flori/sysinfo"
)

var options load.Options

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
	flag.StringVar(
		&options.Mode,
		"mode",
		"cpu",
		"mode of operation cpu/battery, cpu is the default",
	)
	flag.IntVar(
		&options.BatteryPercentage,
		"battery-percentage",
		-1,
		"battery is filled to this percentage",
	)
	flag.Parse()
}

func main() {
	fmt.Print(load.NewBar(options))
}
