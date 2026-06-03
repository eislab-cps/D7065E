# Using BuildSim

A hands-on tutorial for D7065E.

BuildSim is the building provided for the project: a 3D viewer and a REST API in a single
binary. It **holds the building state** (rooms, equipment, sensor values, actuator states);
the project reads and writes that state. This tutorial runs BuildSim, puts a sensor and an
actuator in a room, and connects them with a **very simple physics simulator** so that
actuator commands change the sensor readings, the closed loop at the heart of the project.

<figure class="diagram">
<img src="figures/buildsim-fig01.svg" alt="BuildSim sits in the middle: the browser viewer connects over WebSocket, and the project's processes read and write state over the REST API">
<figcaption><em>BuildSim holds the shared state. The browser viewer connects over WebSocket; the project's own processes read and write that state over the REST API.</em></figcaption>
</figure>

For the full endpoint reference, see `buildingsim/docs/api/`; for how the API maps onto the
lab architecture, see `buildingsim/docs/lab-quickstart.md`.

---

## Part 1 — Run it

BuildSim is a single Go binary. Build and start it:

```bash
go build -o buildsim ./cmd
./buildsim start --port 9090
```

Open <http://localhost:9090> in **Chrome or Firefox** (Safari blocks the WebSocket). The 3D
building appears. The runtime state is **in memory**, so it resets every time BuildSim
restarts; keep a small script to recreate the equipment.

---

## Part 2 — Put a sensor and an actuator in a room

Equipment is placed in a room by its **name** (for example `A109`). Create the equipment,
then add a sensor and an actuator to it. Add them through their **own endpoints**, not as part
of a bulk create, so they are addressable by `PUT /api/sensors/{id}/value` and
`PUT /api/actuators/{id}/state`:

```bash
BASE=http://localhost:9090

curl -X POST $BASE/api/equipment -H 'Content-Type: application/json' -d \
  '{"id":"hvac-A109","name":"HVAC A109","type":"ac_unit","category":"hvac","level":"level0","room":"A109","status":"running"}'

curl -X POST $BASE/api/equipment/hvac-A109/sensors -H 'Content-Type: application/json' -d \
  '{"id":"A109-temp","name":"Temperature","type":"temperature","data_type":"text","unit":"°C","value":"18.0"}'

curl -X POST $BASE/api/equipment/hvac-A109/actuators -H 'Content-Type: application/json' -d \
  '{"id":"A109-set","name":"Setpoint","type":"setpoint","state":"21"}'
```

These three commands are bundled in [`buildsim/setup.sh`](buildsim/setup.sh).

Read it back (a sensor value is read from its parent equipment, there is no
`GET /api/sensors/{id}`):

```bash
curl -s http://localhost:9090/api/equipment/hvac-A109
```

To print just the temperature value, use [`buildsim/gettemp.sh`](buildsim/gettemp.sh) (it
reports clearly if BuildSim is not running or the equipment was never created):

```bash
cd buildsim
./gettemp.sh
# 18.0 °C
```

Sensor values and actuator states are always **text strings** (`"18.0"`, `"21"`), so the
code parses and formats them as numbers.

---

## Part 3 — A very simple physics simulator (close the loop)

The simulator is the program that makes actuator commands affect future sensor readings.
This one models a single idea: **the room temperature drifts toward the heating setpoint.**
Each tick it reads the setpoint actuator from BuildSim, nudges the temperature toward it,
and writes the new temperature back to the sensor.

`main.go`:

```go
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
```

The complete runnable example, `main.go`, `go.mod`, and `setup.sh`, is in the
[`buildsim/`](buildsim/) folder. With BuildSim running, create the equipment and
start the simulator:

```bash
cd buildsim
./setup.sh          # create hvac-A109 with its sensor and actuator
go run .            # start the simulator
# setpoint 21.0 -> temp 18.6
# setpoint 21.0 -> temp 19.1
# setpoint 21.0 -> temp 19.5  ...converging on 21
```

Now **close the loop**: change the setpoint actuator from another terminal and watch the
temperature follow it.

```bash
curl -X PUT http://localhost:9090/api/actuators/A109-set/state \
  -H 'Content-Type: application/json' -d '{"state": "25"}'
# the simulator now drives temp up toward 25
```

That is the whole idea of a cyber-physical loop: an **actuator** command (the setpoint)
changes the **physical** model (the temperature), which changes the next **sensor** reading,
which an autonomous agent would read to decide the next command. A real simulator adds more
effects (occupancy, heat loss, CO₂), but the shape stays exactly this.

---

## Part 4 — See it in the 3D viewer (optional)

The browser viewer is driven through a **session**. Find the active session, then colour the
room so the temperature is visible on the 3D model:

```bash
SID=$(curl -s http://localhost:9090/api/sessions | python3 -c 'import sys,json; print(json.load(sys.stdin)[0]["id"])')
curl -X PUT http://localhost:9090/api/sessions/$SID/highlights \
  -H 'Content-Type: application/json' -d '[{"room_id": 12, "color": "#e53935", "opacity": 0.5}]'
```

Highlights use the room's integer **id**, not its name. Map name to id once from the floor
data:

```bash
curl -s http://localhost:9090/api/building/floors/level0 \
  | python3 -c 'import sys,json; print({r["name"]: r["id"] for r in json.load(sys.stdin)["rooms"]}["A109"])'
```

Highlights, coverage zones, and the viewport all update live over WebSocket, so a dashboard
is largely a matter of mapping state to colours. See `buildingsim/docs/api/sessions.md` for
the full session API.

---

## Things to remember

- **Values are text.** Sensor `value` and actuator `state` are strings; parse and format them.
- **Two room identifiers.** Equipment uses the room **name** (`A109`); occupancy and
  highlights use the integer **id** (map via `/api/building/floors/{level}`).
- **`notify` refreshes the viewer.** `POST /api/equipment/notify` after writing values; the
  value is stored either way.
- **No persistence, no physics.** BuildSim stores the current state only and does not model
  physics, the simulator above is the project's job.
