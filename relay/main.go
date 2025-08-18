package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Coder-Harshit/RefleXSys/common"
	"gopkg.in/yaml.v3"
)

type SignalBuffer struct {
	sync.Mutex
	signals []common.Signal
}

var buffer SignalBuffer

func main() {
	conf, err := loadConfig()
	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}
	fmt.Println("Relay Server Listing on Port :", conf.Port, conf.LogLevel)
	http.HandleFunc("/report", displayReport)
	go flushing(time.Duration(conf.FlushInterval)*time.Second, func(s []common.Signal) {
		if len(s) == 0 {
			return
		}
		jsonBatch, _ := json.Marshal(s)
		resp, err := http.Post(conf.MasterUrl, "application/json", bytes.NewReader(jsonBatch))
		if err != nil {
			fmt.Println("Error encountered while informing MASTER!")
			return
		}
		resp.Body.Close()

	})
	// go flushing(time.Duration(conf.FlushInterval)*time.Second, conf.MasterUrl)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil))
}

func displayReport(respw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var sig common.Signal
		if err := json.NewDecoder(req.Body).Decode(&sig); err != nil {
			http.Error(respw, "Invalid JSON", http.StatusBadRequest)
			return
		}
		fmt.Printf("[RECEIVED] %+v\n", sig)
		// fmt.Printf("buffer.signals: %v\n", buffer.signals)

		buffer.Lock()
		buffer.signals = append(buffer.signals, sig)
		buffer.Unlock()

		respw.WriteHeader(http.StatusOK)
	} else {
		http.Error(respw, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func flushing(interval time.Duration, sendFun func([]common.Signal)) {
	// func flushing(interval time.Duration, masterUrl string) {
	// To flush the buffer signals via the sendFun function
	for {
		// always running loop

		time.Sleep(interval)
		buffer.Lock()
		batch := buffer.signals
		buffer.signals = nil
		buffer.Unlock()

		if len(batch) > 0 {
			sendFun(batch)
			// informMaster(masterUrl, batch)
		}

	}
}

func informMaster(url string, batch []common.Signal) {
	if len(batch) == 0 {
		return
	}
	jsonBatch, _ := json.Marshal(batch)
	http.Post(url, "application/json", bytes.NewReader(jsonBatch))
}

func loadConfig() (*common.RelayConfig, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	var config common.RelayConfig
	if len(data) == 0 {
		// empty config file ... create a default one and read it
		config.Port = 8080
		config.LogLevel = "info"
		config.FlushInterval = 5

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
