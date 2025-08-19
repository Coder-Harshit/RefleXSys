package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Coder-Harshit/RefleXSys/common"
	"gopkg.in/yaml.v3"
)

func main() {
	conf, err := loadConfig()
	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}
	fmt.Println("Master Server Listing on Port :", conf.Port, conf.LogLevel)
	http.HandleFunc("/report", displayReport)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}

func displayReport(respw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var signals []common.Signal
		if err := json.NewDecoder(req.Body).Decode(&signals); err != nil {
			http.Error(respw, "Invalid JSON", http.StatusBadRequest)
			return
		}
		fmt.Printf("[RECEIVED BATCH] %d signals\n", len(signals))
		for _, sig := range signals {
			fmt.Printf("%+v\n", sig)
		}
		respw.WriteHeader(http.StatusOK)
	} else {
		http.Error(respw, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func loadConfig() (*common.ServerConfig, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	var config common.ServerConfig
	if len(data) == 0 {
		// empty config file ... create a default one and read it
		config.Port = 9101
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
