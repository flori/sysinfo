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

func (bar *Bar) measure() int {
	options := bar.Options
	r := float64(len(bar.Runes))

	switch options.Mode {
	case "cpu":
		l := cpuLoad(options)
		c := processorCount(options)
		fraction := l / c
		if fraction > 1 {
			fraction = 1.0
		}
		if fraction >= 0 {
			return int((r - 1) * fraction)
		}
		break
	case "battery":
		fraction := batteryFull(options)
		if fraction >= 0 {
			return int((r - 1) * fraction)
		}
		break
	default:
		panic("unknown mode")
	}
	return -1
}

func (bar *Bar) String() string {
	measurement := bar.measure()
	if measurement < 0 {
		return ""
	}
	format := "%s"
	if bar.Options.Format != "" {
		format = bar.Options.Format
	}
	return fmt.Sprintf(format, string(bar.Runes[measurement]))
}
