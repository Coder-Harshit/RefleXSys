package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"

	"time"
)

type signal struct {
	hostid            string
	hostname          string
	totalMem          uint64
	usedMem           uint64
	usedMemPercentage float64
	cpuPercentage     float64
	timestamp         time.Time
}

func main() {
	hid, err := host.HostID()
	errorCheck(err)
	info, err := host.Info()
	errorCheck(err)
	for {
		// infinite loop

		vmem, err := mem.VirtualMemory()
		errorCheck(err)

		timestamp := time.Now()

		cpupr, err := cpu.Percent(time.Second, false)
		errorCheck(err)

		signal := signal{
			hostid:            hid,
			hostname:          info.Hostname,
			totalMem:          vmem.Total,
			usedMem:           vmem.Used,
			usedMemPercentage: vmem.UsedPercent,
			cpuPercentage:     cpupr[0],
			timestamp:         timestamp,
		}
		displaySignal(signal)

		time.Sleep(1 * time.Second) // 1s sleep (at least)
	}
}

// func displaySignal(sig signal) {
// 	stype := reflect.TypeOf(sig)
// 	sval := reflect.ValueOf(sig)
// 	fmt.Println("Signal {")
// 	for i := 0; i < sval.NumField(); i++ {
// 		if stype.Field(i).Name == "timestamp" {
// 			fmt.Println("\t", "timestamp", ": ", sig.timestamp)
// 		} else {
// 			fmt.Println("\t", stype.Field(i).Name, ": ", sval.Field(i))
// 		}
// 	}
// 	fmt.Println("}")
// 	fmt.Println()
// }

func displaySignal(sig signal) {
	fmt.Println("Signal {")
	fmt.Printf("\tHostID:\t\t%v\n", sig.hostid)
	fmt.Printf("\tHostName:\t%v\n", sig.hostname)
	fmt.Printf("\tMemory Total:\t%v\n", sig.totalMem)
	fmt.Printf("\tMemory Used:\t%v\n", sig.usedMem)
	fmt.Printf("\tMemory Used:\t%.2f%%\n", sig.usedMemPercentage)
	fmt.Printf("\tCPU Used:\t%.2f%%\n", sig.cpuPercentage)
	fmt.Printf("\tTimestamp:\t%v\n", sig.timestamp)
	fmt.Println("}")
	fmt.Println()
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println("Error", err)
	}
}
