package sysinfo

type Options struct {
	Mode              string
	LoadAvg           float64
	Horizontal        bool
	ProcessorCount    int
	BatteryPercentage int
}
