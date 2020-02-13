package sysinfo

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func batteryFull(options Options) float64 {
	batPer := options.BatteryPercentage
	if batPer < 0 {
		cmd := exec.Command("sh", "-c", "pmset -g ps | perl -ne '/(\\d+)%/ && print $1'")
		cmd.Env = append(os.Environ(), "LANG=C")
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		batteryValue := string(out)
		if batteryValue != "" {
			f, err := strconv.ParseInt(batteryValue, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			batPer = int(f)
		}
	}
	return float64(batPer) / 100
}
