package main

import (
	"flag"
	"fmt"

	sysinfo "github.com/flori/sysinfo"
)

var options sysinfo.Options

func init() {
	flag.Float64Var(
		&options.LoadAvg,
		"load",
		-1.0,
		"load",
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
	flag.StringVar(
		&options.Format,
		"format",
		"",
		"format the output, e. g. \"CPU: %s \"",
	)
	flag.Parse()
}

func main() {
	fmt.Print(sysinfo.NewBar(options))
}
