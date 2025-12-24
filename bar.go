package sysinfo

import "fmt"

const (
	horizontalRunes = " ▏▎▍▌▋▊▉"
	verticalRunes   = " ▁▂▃▄▅▆▇█"
)

type Bar struct {
	Options Options
	Runes   []rune
}

func NewBar(options Options) *Bar {
	bar := Bar{Options: options}
	if options.Horizontal {
		bar.Runes = []rune(horizontalRunes)
	} else {
		bar.Runes = []rune(verticalRunes)
	}
	return &bar
}

func (bar *Bar) measure() (int, error) {
	options := bar.Options
	r := float64(len(bar.Runes))

	switch options.Mode {
	case "cpu":
		l, err := cpuLoad(options)
		if err != nil {
			return -1, fmt.Errorf("invalid cpu load")
		}
		c := processorCount(options)
		fraction := l / c
		if fraction > 1 {
			fraction = 1.0
		}
		if fraction >= 0 {
			return int((r - 1) * fraction), nil
		}
		return -1, fmt.Errorf("invalid CPU load fraction")
	case "battery":
		fraction, err := batteryFull(options)
		if err != nil {
			return -1, err
		}
		if fraction >= 0 {
			return int((r - 1) * fraction), nil
		}
		return -1, fmt.Errorf("invalid battery fraction")
	default:
		return -1, fmt.Errorf("unknown mode: %s", options.Mode)
	}
}

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
