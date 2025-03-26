package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	serverURL      = "http://localhost:8080/api/switch/status"
	authToken      = "your-secret-token"
	switchID       = "switch-001"
	pollInterval   = 3 * time.Hour
	automationPoll = 30 * time.Minute
)

type SwitchStatus struct {
	ID        string `json:"id"`
	State     bool   `json:"state"`
	Timestamp int64  `json:"timestamp"`
}

func sendStatus(state bool) {
	status := SwitchStatus{
		ID:        switchID,
		State:     state,
		Timestamp: time.Now().Unix(),
	}
	jsonData, err := json.Marshal(status)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending status:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status sent successfully, Response:", resp.Status)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	interval := pollInterval
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Отправляем первый статус сразу после запуска
	state := rand.Intn(2) == 1
	fmt.Println("Initial status:", state)
	sendStatus(state)

	for {
		select {
		case <-ticker.C:
			state := rand.Intn(2) == 1 // Random ON/OFF state
			fmt.Println("Sending status:", state)
			sendStatus(state)
		}
	}
}
