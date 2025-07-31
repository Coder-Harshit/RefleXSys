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
	log.Fatal(http.ListenAndServe(":8080", nil))
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
