package sysinfo

import (
	"log"
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

func loadAvg(options Options) float64 {
	loadAvg := options.LoadAvg
	if loadAvg < 0 {
		cmd := exec.Command("ps", "-A", "-o", "%cpu=0.0")
		cmd.Env = append(os.Environ(), "LANG=C")
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
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
				log.Fatal(err)
			}
			sum += f
		}
		loadAvg = sum / 100
	}
	return loadAvg
}
