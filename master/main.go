package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type signal struct {
	HostID            string    `json:"host_id"`
	Hostname          string    `json:"host_name"`
	TotalMem          uint64    `json:"total_memory"`
	UsedMem           uint64    `json:"used_memory"`
	UsedMemPercentage float64   `json:"used_memory_percentage"`
	CPUPercentage     float64   `json:"cpu_used_percentage"`
	Timestamp         time.Time `json:"timestamp"`
}

type Config struct {
	Port     uint16 `yaml:"port"`
	LogLevel string `yaml:"logging_level"`
}

func main() {
	conf, err := loadConfig()
	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}
	fmt.Println("Server Listing on Port :", conf.Port, conf.LogLevel)
	http.HandleFunc("/report", displayReport)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}

func displayReport(respw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// fmt.Println("Got Signal!")
		var sig signal
		if err := json.NewDecoder(req.Body).Decode(&sig); err != nil {
			http.Error(respw, "Invalid JSON", http.StatusBadRequest)
			return
		}
		fmt.Printf("[RECEIVED] Host: %s | CPU: %.2f%% | RAM: %.2f%% | Time: %v\n",
			sig.Hostname, sig.CPUPercentage, sig.UsedMemPercentage, sig.Timestamp)
		respw.WriteHeader(http.StatusOK)
	}
}

func loadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	var config Config
	if len(data) == 0 {
		// empty config file ... create a default one and read it
		config.Port = 8080
		config.LogLevel = "info"

		_, err = yaml.Marshal(&config)
		if err != nil {
			return nil, fmt.Errorf("error creating default config file: %v", err)
		}
	} else {
		err = yaml.Unmarshal([]byte(data), &config)
		if err != nil {
			return nil, fmt.Errorf("error reading file: %v", err)
		}
	}
	return &config, nil
}
