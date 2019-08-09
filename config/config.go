package config

// Cfg defines a configuration structure.
type Cfg struct {
	Verbose  bool
	Archiver string
	Fetcher  string
}

// Config provides concrete configuration usage.
var Config = Cfg{}
