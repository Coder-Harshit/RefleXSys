package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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
	fmt.Println("Server Listing on Port :8080")
	http.HandleFunc("/report", displayReport)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func displayReport(respw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// fmt.Println("Got Signal!")
		var sig signal
		json.NewDecoder(req.Body).Decode(&sig)
		fmt.Println(sig)
		respw.WriteHeader(http.StatusOK)
	}
}
