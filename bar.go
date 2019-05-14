package cpuload

const (
	horizontalRunes = " ▏▎▍▌▋▊▉"
	verticalRunes   = " ▁▂▃▄▅▆▇█"
)

type Bar struct {
	LoadAvg        float64
	Runes          []rune
	ProcessorCount float64
}

func NewBar(options Options) *Bar {
	bar := Bar {
		LoadAvg:        loadAvg(options),
		ProcessorCount: processorCount(options),
	}
	if options.Horizontal {
		bar.Runes = []rune(horizontalRunes)
	} else {
		bar.Runes = []rune(verticalRunes)
	}
	return &bar
}

func (bar *Bar) String() string {
	l := bar.LoadAvg
	c := bar.ProcessorCount
	fraction := l / c
	if fraction > 1 {
		fraction = 1.0
	}
	r := float64(len(bar.Runes))
	i := int((r - 1) * fraction)
	return string(bar.Runes[i])
}
