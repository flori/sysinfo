package sysinfo

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func processorCount(options Options) float64 {
	processorCount := options.ProcessorCount
	if processorCount < 0 {
		processorCount = runtime.NumCPU()
	}
	return float64(processorCount)
}

func cpuLoad(options Options) (float64, error) {
	cpuLoad := options.LoadAvg
	if cpuLoad < 0 {
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
