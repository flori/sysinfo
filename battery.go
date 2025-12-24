package sysinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// batteryFull calculates and returns the battery percentage as a fraction (0.0
// to 1.0)
//
// The function follows this logic:
// - If BatteryPercentage in options is >= 0, it returns that value converted to fraction
// - If BatteryPercentage is negative (default), it executes a system command to fetch actual battery info
//
// For the command execution approach on macOS:
// - Uses 'pmset -g ps' to get power status information
// - Pipes output through perl to extract percentage value using regex
// - Parses the extracted value to convert to a fraction
//
// Error handling:
// - Returns 0 with wrapped error if command execution fails
// - Returns 0 with wrapped error if extracted value is empty
// - Returns 0 with wrapped error if parsing fails
//
// Edge cases:
// - Empty command output results in "empty battery value" error
// - Invalid numeric parsing results in error propagation
// - Command execution failures are wrapped with descriptive context
// - The function assumes battery percentage values are in range 0-100
//
// Note: This function is macOS-specific and will not work on other operating
// systems The command relies on macOS power management tools and perl regex
// parsing
//
// Corner cases:
// - Battery percentage values outside 0-100 range may cause unexpected behavior
// - Command output format changes could break the regex pattern
// - Environment variable LANG=C ensures consistent command output formatting
func batteryFull(options Options) (float64, error) {
	batPer := options.BatteryPercentage
	if batPer < 0 {
		// Execute command to get battery percentage on macOS
		cmd := exec.Command("sh", "-c", "pmset -g ps | perl -ne '/(\\d+)%/ && print $1'")
		cmd.Env = append(os.Environ(), "LANG=C")
		out, err := cmd.Output()
		if err != nil {
			return 0, fmt.Errorf("failed to get battery percentage: %w", err)
		}
		batteryValue := strings.TrimSpace(string(out))
		if batteryValue == "" {
			return 0, fmt.Errorf("empty battery value")
		} else {
			f, err := strconv.ParseInt(batteryValue, 10, 32)
			if err != nil {
				return 0, fmt.Errorf("failed to parse battery percentage: %w", err)
			}
			batPer = int(f)
		}
	}
	return float64(batPer) / 100, nil
}
