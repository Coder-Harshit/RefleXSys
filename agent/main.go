package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"

	"time"

	"net/http"
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

func main() {
	url := "http://localhost:8080/report"

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

		signal := signal{
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
		time.Sleep(1 * time.Second) // 1s sleep (at least)
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
