package common

import (
	"time"
)

type Signal struct {
	HostID            string    `json:"host_id"`
	Hostname          string    `json:"host_name"`
	TotalMem          uint64    `json:"total_memory"`
	UsedMem           uint64    `json:"used_memory"`
	UsedMemPercentage float64   `json:"used_memory_percentage"`
	CPUPercentage     float64   `json:"cpu_used_percentage"`
	Timestamp         time.Time `json:"timestamp"`
}

type ServerConfig struct {
	Port     uint16 `yaml:"port"`
	LogLevel string `yaml:"logging_level"`
}

type Thresholds struct {
	CPUPercentage     float64 `yaml:"cpu_percentage"`
	UsedMemPercentage float64 `yaml:"memory_percentage"`
}

type AgentConfig struct {
	RelayURL       string     `yaml:"relay_url"`
	ReportInterval int        `yaml:"report_interval"`
	Thresholds     Thresholds `yaml:"thresholds"`
}

type RelayConfig struct {
	Port          uint16 `yaml:"port"`
	LogLevel      string `yaml:"logging_level"`
	FlushInterval int    `yaml:"flush_interval"`
	MasterUrl     string `yaml:"master_url"`
}
