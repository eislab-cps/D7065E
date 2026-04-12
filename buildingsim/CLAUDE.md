# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Building simulation server with REST API and WebSocket-based session control. A Go binary embeds building floor plan data and a web viewer, exposing APIs for equipment management, navigation graph queries, and programmatic control of the map viewer.

## Build & Run

```bash
# Build
go build -o buildsim ./cmd

# Run (binds to 0.0.0.0)
./buildsim start --port 9090
```

## Architecture

### Go Server (`cmd/`, `pkg/`, `embed.go`)

- **`cmd/main.go`** — Cobra CLI entry point. `buildsim start --port 9090`.
- **`embed.go`** — `go:embed` directives for `data/` and `web/` directories.
- **`pkg/server/server.go`** — Gin HTTP server setup, route registration, building data loading.
- **`pkg/server/handlers/`** — REST API handlers:
  - `building.go` — read-only building/floor data
  - `graph.go` — navigation graph + Dijkstra route queries
  - `equipment.go` — equipment CRUD + version bump notifications
  - `sensor.go` — sensor CRUD + value setting (binary or text)
  - `actuator.go` — actuator CRUD + state setting
  - `session.go` — session lifecycle + viewport/highlights/occupancy/route state + WebSocket upgrade
- **`pkg/server/websocket/hub.go`** — WebSocket hub managing per-session client subscriptions, ping/pong keepalive, and 1-hour auto-purge of inactive sessions.
- **`pkg/model/`** — Data types: Building, FloorData, Room, Equipment, Sensor, Actuator, Session, Viewport, Person, Alien, NavGraph, RouteResult.
- **`pkg/store/memory.go`** — Thread-safe in-memory store for all runtime state (equipment, sessions). Building data is read-only after startup.
- **`pkg/graph/dijkstra.go`** — Dijkstra shortest path on the navigation graph.

### Web UI (`web/index.html`)

Single-file Three.js-based viewer/editor. Renders multi-floor building with pan/zoom, room selection, graph overlay, and editing tools.

### Python Scripts (`scripts/`)

- `transform.py` — PDF-to-JSON floor plan extraction (PyMuPDF, OpenCV, NumPy)
- `build_graph.py` — Navigation graph builder (Shapely). Default building: `abuilding`.

## REST API

| Endpoint | Method | Description |
|---|---|---|
| `/api/building` | GET | Building metadata |
| `/api/building/floors/{level}` | GET | Floor plan data |
| `/api/building/cross-floor-edges` | GET | Cross-floor connections |
| `/api/graph` | GET | Navigation graph (?level=) |
| `/api/graph/route` | GET | Dijkstra shortest path (?from=&to=&level=) |
| `/api/equipment` | GET/POST | List/create equipment |
| `/api/equipment/{id}` | GET/PUT/DELETE | Equipment CRUD |
| `/api/equipment/notify` | POST | Bump equipment version, notify all sessions |
| `/api/equipment/{id}/sensors` | GET/POST | List/add sensors |
| `/api/sensors/{id}` | DELETE | Remove sensor |
| `/api/sensors/{id}/value` | PUT | Set sensor value (binary or text) |
| `/api/equipment/{id}/actuators` | GET/POST | List/add actuators |
| `/api/actuators/{id}` | DELETE | Remove actuator |
| `/api/actuators/{id}/state` | PUT | Set actuator state |
| `/api/sessions` | POST | Create session |
| `/api/sessions/{id}` | GET/DELETE | Get/delete session |
| `/api/sessions/{id}/viewport` | PUT | Set viewport (pushed via WS) |
| `/api/sessions/{id}/highlights` | PUT | Set room highlights (pushed via WS) |
| `/api/sessions/{id}/occupancy` | PUT | Set room occupancy (version bump via WS) |
| `/api/sessions/{id}/route` | PUT | Set displayed route (pushed via WS) |
| `/api/sessions/{id}/coverage` | PUT | Set coverage zones (pushed via WS) |
| `/api/icons/{name}.svg` | GET | Equipment SVG icon |
| `/ws/{session_id}` | WS | WebSocket subscription |

## Data Layout

All embedded in the binary:
- `data/abuilding/level{0,1,2}/floorplan_data.json` — per-floor room/wall/graph data
- `data/abuilding/cross_floor_edges.json` — stair/elevator connections
- `data/equipment/icons/` — SVG icons for equipment types

## Dependencies

Go: `gin-gonic/gin`, `gorilla/websocket`, `spf13/cobra`, `google/uuid`
Python (scripts only): `PyMuPDF`, `opencv-contrib-python`, `numpy`, `shapely`

## Git

Do NOT add Co-Authored-By lines for Claude in git commits.
