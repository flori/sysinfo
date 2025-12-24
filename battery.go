package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func batteryFull(options Options) (float64, error) {
	batPer := options.BatteryPercentage
	if batPer < 0 {
		cmd := exec.Command("sh", "-c", "pmset -g ps | perl -ne '/(\\d+)%/ && print $1'")
		cmd.Env = append(os.Environ(), "LANG=C")
		out, err := cmd.Output()
		if err != nil {
			return -1.0, fmt.Errorf("failed to get battery percentage: %w", err)
		}
		batteryValue := strings.TrimSpace(string(out))
		if batteryValue == "" {
			return -1.0, fmt.Errorf("empty battery value")
		} else {
			f, err := strconv.ParseInt(batteryValue, 10, 32)
			if err != nil {
				return -1.0, fmt.Errorf("failed to parse battery percentage: %w", err)
			}
			batPer = int(f)
		}
	}
	return float64(batPer) / 100, nil
}
