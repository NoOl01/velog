package velog_config

type LogLevel int
type Format int

const (
	Disabled Format = iota
	Text
	Json
)

type Config struct {
	Format    string
	Separator string
	Debug     bool
	Console   ConsoleConfig
	File      FileConfig
}

type ConsoleConfig struct {
	Enabled bool
	Output  OutputFormat
}

type FileConfig struct {
	Enabled       bool
	Directory     string
	FileName      string
	FileExtension string
	AutoDelete    int
	AutoArchive   int
	Output        OutputFormat
}

type OutputFormat struct {
	Format Format
	Json   string
}

const (
	INFO LogLevel = iota
	DEBUG
	WARNING
	ERROR
	PANIC
)
