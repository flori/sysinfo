package sysinfo

type Options struct {
	Mode              string
	Format            string
	Horizontal        bool
	ProcessorCount    int
	LoadAvg           float64
	BatteryPercentage int
}
