package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"gopkg.in/yaml.v3"

	"net/http"
	"time"

	"github.com/Coder-Harshit/RefleXSys/common"
)

func main() {
	conf, err := loadConfig()
	errorCheck(err, "Config Failed to load")

	url := conf.RelayURL

	hid, err := host.HostID()
	errorCheck(err, "[HostID] Object creation Issue!")
	info, err := host.Info()
	errorCheck(err, "[HostInfo] Object creation Issue!")
	for {
		// infinite loop

		vmem, err := mem.VirtualMemory()
		errorCheck(err, "[MemoryObject] Object creation Issue!")

		timestamp := time.Now()

		cpupr, err := cpu.Percent(time.Second, false)
		errorCheck(err, "[CPUPercentage] Object creation Issue!")

		signal := common.Signal{
			HostID:            hid,
			Hostname:          info.Hostname,
			TotalMem:          vmem.Total,
			UsedMem:           vmem.Used,
			UsedMemPercentage: vmem.UsedPercent,
			CPUPercentage:     cpupr[0],
			Timestamp:         timestamp,
		}

		jsonSignal, err := json.Marshal(signal)
		errorCheck(err, "[JSON-Encoding] Issue")

		// displaySignal(jsonSignal)

		resp, err := http.Post(url, "application/json", bytes.NewReader(jsonSignal))
		errorCheck(err, "[POST] signal issue")
		resp.Body.Close()
		// fmt.Println(resp.StatusCode)
		time.Sleep(time.Duration(conf.ReportInterval) * time.Second) // 1s sleep (at least)
	}
}

// func displaySignal(sig signal) {
// 	fmt.Println("Signal {")
// 	fmt.Printf("\tHostID:\t\t%v\n", sig.HostID)
// 	fmt.Printf("\tHostName:\t%v\n", sig.Hostname)
// 	fmt.Printf("\tMemory Total:\t%v\n", sig.TotalMem)
// 	fmt.Printf("\tMemory Used:\t%v\n", sig.UsedMem)
// 	fmt.Printf("\tMemory Used:\t%.2f%%\n", sig.UsedMemPercentage)
// 	fmt.Printf("\tCPU Used:\t%.2f%%\n", sig.CPUPercentage)
// 	fmt.Printf("\tTimestamp:\t%v\n", sig.Timestamp)
// 	fmt.Println("}")
// 	fmt.Println()
// }

// func displaySignal(jsonSig []byte) {
// 	fmt.Println("Signal =>")
// 	fmt.Println(string(jsonSig))
// 	fmt.Println()
// }

func errorCheck(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err)
		// time.Sleep(5 * time.Second)
	}
}

func loadConfig() (*common.AgentConfig, error) {
	data, err := os.ReadFile("config.yaml")
	errorCheck(err, "Error Reading File")
	var config common.AgentConfig
	if len(data) == 0 {
		// empty config file ... create a default one and read it
		config.RelayURL = "http://localhost:8101/report"
		config.ReportInterval = 1

		_, err = yaml.Marshal(&config)
		errorCheck(err, "error creating default config file")

	} else {
		err = yaml.Unmarshal([]byte(data), &config)
		errorCheck(err, "unable to parse config file")
	}
	return &config, nil
}
