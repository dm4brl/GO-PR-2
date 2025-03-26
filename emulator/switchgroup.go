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
	serverURL      = "http://localhost:8080/api/switch/status" // URL сервера
	authToken      = "your-secret-token"
	numSwitches    = 5 // Количество свитчей в кластере
	pollInterval   = 3 * time.Hour
	automationPoll = 30 * time.Minute
)

type SwitchStatus struct {
	ID        string `json:"id"`
	State     bool   `json:"state"`
	Timestamp int64  `json:"timestamp"`
}

// Функция отправки статуса на сервер
func sendStatus(state bool, switchID string) {
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

	fmt.Println("Status sent successfully for", switchID, ", Response:", resp.Status)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Массив с уникальными идентификаторами свитчей
	switchIDs := []string{}
	for i := 1; i <= numSwitches; i++ {
		switchIDs = append(switchIDs, fmt.Sprintf("switch-%03d", i)) // Генерация ID типа switch-001, switch-002 и т.д.
	}

	// Настроим периодичность отправки статусов
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	// Отправляем начальные статусы для всех свитчей
	for _, switchID := range switchIDs {
		state := rand.Intn(2) == 1
		fmt.Println("Initial status for", switchID, ":", state)
		sendStatus(state, switchID)
	}

	// Отправляем статусы по расписанию
	for {
		select {
		case <-ticker.C:
			for _, switchID := range switchIDs {
				state := rand.Intn(2) == 1 // случайный ON/OFF
				fmt.Println("Sending status for", switchID, ":", state)
				sendStatus(state, switchID)
			}
		}
	}
}
