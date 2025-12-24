package sysinfo

// Options represents the configuration options for creating and displaying
// system information bars
type Options struct {
	// Mode specifies what system metric to display
	// Valid values are typically "cpu" or "battery"
	Mode string

	// Format specifies the output format string for the bar display
	// If empty, defaults to a simple rune representation
	Format string

	// Horizontal determines whether to use horizontal or vertical bar runes
	// When true, uses horizontal bar characters; when false, uses vertical bar characters
	Horizontal bool

	// ProcessorCount specifies the number of logical processors to use for CPU calculations
	// If negative (default), automatically detects the number of CPUs using runtime.NumCPU()
	ProcessorCount int

	// LoadAvg specifies the CPU load percentage to display directly
	// If negative (default), the system will measure actual CPU load via external commands
	LoadAvg float64

	// BatteryPercentage specifies the battery level percentage to display directly
	// If negative (default), the system will measure actual battery level via external commands
	BatteryPercentage int
}
