package sysinfo

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// processorCount returns the number of logical processors to use for CPU
// calculations
//
// The function respects the ProcessorCount field in Options:
// - If ProcessorCount is >= 0, it returns that value
// - If ProcessorCount is negative (default), it falls back to runtime.NumCPU()
//
// This approach allows for overriding the automatic detection of CPU cores,
// which can be useful for testing or when working with containerized
// environments
func processorCount(options Options) float64 {
	processorCount := options.ProcessorCount
	if processorCount < 0 {
		processorCount = runtime.NumCPU()
	}
	return float64(processorCount)
}

// cpuLoad calculates the current CPU load percentage
//
// The function handles two scenarios:
//  1. If LoadAvg is >= 0 in options, it returns that value directly
//  2. If LoadAvg is negative (default), it executes a system command to measure
//     actual CPU usage
//
// For the command execution approach:
// - Uses 'ps' command to get CPU percentages of all processes
// - Parses the output to sum up all CPU percentages
// - Divides by 100 to convert to a fraction (0.0 to 1.0)
//
// Error handling:
// - Returns error if command execution fails
// - Returns error if parsing individual CPU percentages fails
// - Returns 0.0 for invalid/empty values during parsing
//
// Edge cases:
// - Empty lines in command output are skipped
// - Invalid numeric values in command output cause parsing errors
// - Command execution failures result in propagated errors
//
// Note: This function may have performance implications as it spawns external
// processes and could be a bottleneck in high-frequency usage scenarios
func cpuLoad(options Options) (float64, error) {
	cpuLoad := options.LoadAvg
	if cpuLoad < 0 {
		// Execute ps command to get CPU load information
		cmd := exec.Command("ps", "-A", "-o", "%cpu=0.0")
		cmd.Env = append(os.Environ(), "LANG=C")
		out, err := cmd.Output()
		if err != nil {
			return 0.0, err
		}
		result := strings.Split(string(out), "\n")
		sum := 0.0
		for _, load := range result {
			load := strings.Trim(load, " ")
			if load == "" {
				continue
			}
			f, err := strconv.ParseFloat(load, 64)
			if err != nil {
				return 0.0, err
			}
			sum += f
		}
		cpuLoad = sum / 100
	}
	return cpuLoad, nil
}
