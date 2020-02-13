package sysinfo

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
		l := loadAvg(options)
		c := processorCount(options)
		fraction := l / c
		if fraction > 1 {
			fraction = 1.0
		}
		return int((r - 1) * fraction)
	case "battery":
		fraction := batteryFull(options)
		return int((r - 1) * fraction)
	default:
		panic("unknown mode")
	}
}

func (bar *Bar) String() string {
	return string(bar.Runes[bar.measure()])
}
