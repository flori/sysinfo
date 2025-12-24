package main

import (
	"flag"
	"fmt"

	sysinfo "github.com/flori/sysinfo"
)

// options holds the configuration options for the sysinfo bar display
// These are populated from command-line flags during initialization
var options sysinfo.Options

// init is called before main() and sets up command-line flag parsing
// It defines all available command-line options that can be used to configure
// the behavior of the sysinfo bar display
func init() {
	// Flag for specifying CPU load percentage
	// Default value is -1.0 which means the system will measure actual CPU load
	flag.Float64Var(
		&options.LoadAvg,
		"load",
		-1.0,
		"load",
	)

	// Flag to determine display orientation
	// When true, displays horizontal bars; when false, displays vertical bars
	flag.BoolVar(
		&options.Horizontal,
		"horizontal",
		false,
		"display horizontal bar instead of vertical",
	)

	// Flag to specify number of processor cores/threads for CPU calculations
	// Default value is -1 which means automatic detection using runtime.NumCPU()
	flag.IntVar(
		&options.ProcessorCount,
		"processor-count",
		-1,
		"number of processor cores/threads",
	)

	// Flag to set the mode of operation
	// Valid values are typically "cpu" or "battery", with "cpu" as default
	flag.StringVar(
		&options.Mode,
		"mode",
		"cpu",
		"mode of operation cpu/battery, cpu is the default",
	)

	// Flag to specify battery percentage directly
	// Default value is -1 which means the system will measure actual battery level
	flag.IntVar(
		&options.BatteryPercentage,
		"battery-percentage",
		-1,
		"battery is filled to this percentage",
	)

	// Flag to format the output string
	// Allows custom formatting like "CPU: %s " where %s will be replaced with the bar character
	flag.StringVar(
		&options.Format,
		"format",
		"",
		"format the output, e. g. \"CPU: %s \"",
	)

	// Parse all defined command-line flags
	// This must be called after all flag variables are defined
	flag.Parse()
}

// main is the entry point of the application
// It creates a new Bar instance with the parsed command-line options
// and prints the resulting bar display to standard output
func main() {
	// Create a new bar with the configured options
	// The String() method of Bar returns the formatted display string
	// which is then printed to stdout using fmt.Print
	fmt.Print(sysinfo.NewBar(options))
}
