package config

import (
	"github.com/go-ini/ini"
)

type LocationConfig struct {
	Java         string
	NusaTenggara string
	Sulawesi     string
	Papua        string
	Kalimantan   string
	Sumatra      string
}

func NewLocationConfig(cfg *ini.File) LocationConfig {
	// Load locations for tweets filtering
	java := cfg.Section("location").Key("java").String()
	nusaTenggara := cfg.Section("location").Key("nusa_tenggara").String()
	sulawesi := cfg.Section("location").Key("sulawesi").String()
	papua := cfg.Section("location").Key("papua").String()
	kalimantan := cfg.Section("location").Key("kalimantan").String()
	sumatra := cfg.Section("location").Key("sumatra").String()

	locationConfig := LocationConfig{java, nusaTenggara, sulawesi, papua, kalimantan, sumatra}
	return locationConfig
}

func (lc LocationConfig) Locations() []string {
	locations := []string{lc.Java, lc.Kalimantan, lc.NusaTenggara, lc.Sulawesi, lc.Papua, lc.Sumatra}
	return locations
}
