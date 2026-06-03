package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const base = "http://localhost:9090"

type equipment struct {
	Actuators []struct {
		ID    string `json:"id"`
		State string `json:"state"`
	} `json:"actuators"`
}

func main() {
	temp := 18.0
	for {
		// 1. read the heating setpoint from BuildSim (fail if it is not running)
		resp, err := http.Get(base + "/api/equipment/hvac-A109")
		if err != nil {
			log.Fatalf("cannot reach BuildSim at %s: %v", base, err)
		}
		var eq equipment
		json.NewDecoder(resp.Body).Decode(&eq)
		resp.Body.Close()

		setpoint := 21.0
		for _, a := range eq.Actuators {
			if a.ID == "A109-set" {
				if v, err := strconv.ParseFloat(a.State, 64); err == nil {
					setpoint = v
				}
			}
		}

		// 2. very simple physics: move 20% of the way toward the setpoint
		temp += 0.2 * (setpoint - temp)

		// 3. write the new temperature back, and fail if BuildSim does not accept it
		body, _ := json.Marshal(map[string]string{
			"data_type": "text", "value": fmt.Sprintf("%.1f", temp),
		})
		req, _ := http.NewRequest(http.MethodPut,
			base+"/api/sensors/A109-temp/value", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		put, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("cannot write to BuildSim: %v", err)
		}
		if put.StatusCode != http.StatusOK {
			log.Fatalf("BuildSim rejected the write (%d), did you run ./setup.sh?", put.StatusCode)
		}
		put.Body.Close()
		http.Post(base+"/api/equipment/notify", "", nil) // refresh the viewer

		fmt.Printf("setpoint %.1f -> temp %.1f\n", setpoint, temp)
		time.Sleep(2 * time.Second)
	}
}
