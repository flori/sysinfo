package sysinfo

import "fmt"

// Constants defining the Unicode runes used for rendering horizontal and
// vertical bars
const (
	horizontalRunes = " ▏▎▍▌▋▊▉"
	verticalRunes   = " ▁▂▃▄▅▆▇█"
)

// Bar represents a visual bar display element
type Bar struct {
	// Options contains the configuration for this bar
	Options Options

	// Runes contains the Unicode characters used to render the bar
	Runes []rune
}

// NewBar creates and returns a new Bar instance with the specified options
// It initializes the Runes field based on whether the bar should be horizontal
// or vertical
func NewBar(options Options) *Bar {
	bar := Bar{Options: options}
	if options.Horizontal {
		bar.Runes = []rune(horizontalRunes)
	} else {
		bar.Runes = []rune(verticalRunes)
	}
	return &bar
}

// measure calculates the appropriate bar segment to display based on system
// metrics It returns the index of the rune to display and any error that
// occurred during measurement
//
// The function handles two main modes:
// - "cpu": Calculates CPU load as a fraction of total processors
// - "battery": Calculates battery level as a fraction (0.0 to 1.0)
//
// Edge cases handled:
// - CPU load exceeding total processor count is capped at 1.0
// - Negative values for fractions return errors
// - Unknown modes return errors
func (bar *Bar) measure() (int, error) {
	options := bar.Options
	r := float64(len(bar.Runes))

	switch options.Mode {
	case "cpu":
		// Get CPU load and processor count
		l, err := cpuLoad(options)
		if err != nil {
			return 0, fmt.Errorf("invalid cpu load")
		}
		c := processorCount(options)
		fraction := l / c
		// Cap fraction at 1.0 to prevent overflow
		if fraction > 1 {
			fraction = 1.0
		}
		// Validate non-negative fraction
		if fraction >= 0 {
			return int((r - 1) * fraction), nil
		}
		return 0, fmt.Errorf("invalid CPU load fraction")
	case "battery":
		// Get battery fraction
		fraction, err := batteryFull(options)
		if err != nil {
			return 0, err
		}
		// Validate non-negative fraction
		if fraction >= 0 {
			return int((r - 1) * fraction), nil
		}
		return 0, fmt.Errorf("invalid battery fraction")
	default:
		return 0, fmt.Errorf("unknown mode: %s", options.Mode)
	}
}

// String returns the string representation of the bar
// It measures the current system state and returns the appropriate rune
//
// Error handling:
// - If measurement fails or returns negative values, returns empty string
// - Uses custom format if specified in Options
//
// Corner cases:
// - Empty string returned when measurement fails
// - Format string can be used to wrap the returned rune in additional text
func (bar *Bar) String() string {
	measurement, err := bar.measure()
	if err != nil || measurement < 0 {
		return ""
	}
	format := "%s"
	if bar.Options.Format != "" {
		format = bar.Options.Format
	}
	return fmt.Sprintf(format, string(bar.Runes[measurement]))
}
