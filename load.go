package cpuload

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func processorCount(options Options) float64 {
	processorCount := options.ProcessorCount
	if processorCount < 0 {
		out, err := exec.Command("getconf", "_NPROCESSORS_ONLN").Output()
		if err != nil {
			log.Fatal(err)
		}
		i, err := strconv.ParseInt(strings.TrimSpace(string(out)), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		processorCount = int(i)
	}
	return float64(processorCount)
}

func loadAvg(options Options) float64 {
	loadAvg := options.LoadAvg
	if loadAvg < 0 {
		cmd := exec.Command("uptime")
		cmd.Env = append(os.Environ(), "LANG=C")
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		result := strings.Split(string(out), " ")
		l := len(result)
		if l > 3 {
			loadValue := strings.Trim(result[l-3], ",")
			f, err := strconv.ParseFloat(loadValue, 64)
			if err != nil {
				log.Fatal(err)
			}
			loadAvg = f
		} else {
			log.Fatal(errors.New("Invalid output of uptime"))
		}
	}
	return loadAvg
}
