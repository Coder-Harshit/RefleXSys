package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Printf("Total Memory = %v\n", vmem.Total)
	fmt.Printf("Used Memory = %v (%.2f%%)\n", vmem.Used, vmem.UsedPercent)
}
