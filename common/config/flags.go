package config

import (
	"flag"
)

type Flags struct {
	Environment string
	Service     string
	Port        int
}

var (
	flags = &Flags{}
	isSet = false
)

func NewFlags() *Flags {
	if !isSet {
		flag.StringVar(&flags.Environment, "env", "dev", "environment")
		flag.IntVar(&flags.Port, "port", 9010, "port")

		flag.Parse()

		isSet = true
	}
	return flags
}

// SetFlags allows manually setting the flags rather than getting them from the CLI directly, allowing overriding the loading behavior
func SetFlags(f *Flags) {
	flags.Environment = f.Environment
	flags.Port = f.Port
	flags.Service = f.Service

	isSet = true
}
