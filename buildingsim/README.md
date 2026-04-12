# BuildSim

Building simulation server with REST API and WebSocket session control. A single Go binary embeds building floor plan data, a 3D web viewer, and SVG equipment icons.

## Quick Start

```bash
make          # build to bin/buildsim
make run      # build and start on port 9090
make run-edit # build and start with editing enabled
make test     # run all tests with race detector
```

Open http://localhost:9090 in a browser to view the 3D building.

## Features

| Feature | Description |
|---------|-------------|
| 3D Building Viewer | Multi-floor building visualization with pan, zoom, and orbit controls |
| REST API | Full CRUD for equipment, sensors, and actuators |
| Session Control | Programmatic control of the viewer (viewport, highlights, occupancy, routes, coverage zones) via REST + WebSocket |
| Navigation Graph | Walkable corridor graph with Dijkstra shortest path routing, single and cross-floor |
| Equipment Icons | 49 detailed SVG icons for all equipment types, loaded dynamically |
| Occupancy | Person, group, and alien (xenomorph) icons rendered in rooms via session API |
| MCP Server | Model Context Protocol server for AI-controlled building manipulation |

## REST API

### Building Data

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/building` | GET | Building metadata (name, levels) |
| `/api/building/floors/{level}` | GET | Floor plan data (rooms, walls, graphs) |
| `/api/building/cross-floor-edges` | GET | Stair/elevator connections between floors |

### Equipment, Sensors, Actuators

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/equipment` | GET | List all equipment (filter by level, room, type, category) |
| `/api/equipment` | POST | Create equipment |
| `/api/equipment/bulk` | POST | Bulk create equipment with sensors/actuators |
| `/api/equipment/{id}` | GET | Get equipment by ID |
| `/api/equipment/{id}` | PUT | Update equipment |
| `/api/equipment/{id}` | DELETE | Delete equipment |
| `/api/equipment/notify` | POST | Bump version, notify all browser sessions |
| `/api/equipment/{id}/sensors` | GET/POST | List or add sensors |
| `/api/sensors/{id}` | DELETE | Remove sensor |
| `/api/sensors/{id}/value` | PUT | Set sensor value (text or binary) |
| `/api/equipment/{id}/actuators` | GET/POST | List or add actuators |
| `/api/actuators/{id}` | DELETE | Remove actuator |
| `/api/actuators/{id}/state` | PUT | Set actuator state |

### Sessions

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/sessions` | GET | List active sessions |
| `/api/sessions` | POST | Create session |
| `/api/sessions/{id}` | GET | Get session state |
| `/api/sessions/{id}` | DELETE | Delete session |
| `/api/sessions/{id}/viewport` | PUT | Set viewport, zoom to room (pushed via WebSocket) |
| `/api/sessions/{id}/highlights` | PUT | Set room highlights with colors (pushed via WebSocket) |
| `/api/sessions/{id}/occupancy` | PUT | Set room occupancy, persons/aliens (version bump via WebSocket) |
| `/api/sessions/{id}/route` | PUT | Set displayed route (pushed via WebSocket) |
| `/api/sessions/{id}/coverage` | PUT | Set coverage zones, translucent spheres (pushed via WebSocket) |

### Navigation Graph

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/graph` | GET | Navigation graph (?level=&type=adjacency\|walkable) |
| `/api/graph/route` | GET | Dijkstra shortest path (?from_name=&to_name=&type=walkable) |

### Other

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/icons/{name}.svg` | GET | Equipment SVG icon |
| `/api/config` | GET | Server configuration (edit mode) |
| `/ws/{session_id}` | WebSocket | Session state subscription |

Full API documentation with curl examples: [docs/api/](docs/api/)

## Documentation

| Document | Description |
|----------|-------------|
| [API Reference](docs/api/) | REST API with curl examples |
| [Architecture](docs/architecture.md) | System diagrams, data model, session flow |
| [Equipment Icons](docs/icons.md) | Gallery of all 49 SVG icons |
| [MCP Server](mcp/) | AI-controlled building manipulation via MCP |

## Equipment Types

| Category | Types |
|----------|-------|
| Access Control | door_lock, card_reader |
| HVAC | temperature_sensor, humidity_sensor, ac_unit, ventilation_fan, radiator, thermostat, ahu |
| Safety & Fire | smoke_detector, fire_extinguisher, emergency_light, sprinkler, fire_alarm_panel, aed |
| Electrical | light_fixture, distribution_panel, emergency_generator, ups |
| Plumbing | water_valve, water_leak_sensor, pump |
| Monitoring | security_camera, motion_sensor, co2_sensor, co_sensor, air_quality_sensor, noise_sensor, light_sensor, vibration_sensor, water_flow_meter, gas_leak_detector, radon_detector, occupancy_counter |
| Network | network_switch, wifi_access_point, base_station_5g, bms_controller, iot_gateway |
| Vertical Transport | elevator, escalator |
| Fixed Equipment | compressor, coffee_machine, printer, projector |

## Example Scripts

| Script | Description |
|--------|-------------|
| `examples/api/building/list_rooms.sh` | List all rooms by floor |
| `examples/api/equipment/list.sh` | List equipment with sensors/actuators |
| `examples/api/equipment/setup.sh` | Add equipment with sensors |
| `examples/api/equipment/set_sensors.sh` | Set sensor values |
| `examples/api/session/set_viewport.sh` | Zoom to a room |
| `examples/api/session/set_highlights.sh` | Highlight rooms with colors |
| `examples/api/session/set_occupancy.sh` | Place people and aliens |
| `examples/api/session/set_coverage.sh` | Set coverage zones |
| `examples/api/graph/compute_route.sh` | Compute shortest path |
| `examples/api/scenario/populate_building.py` | Bulk load 2500 equipment |
| `examples/api/scenario/add_5g.sh` | Add 5G base station with coverage |

## CLI

| Flag | Description | Default |
|------|-------------|---------|
| `--port, -p` | Port to listen on | 9090 |
| `--edit` | Enable floor plan editing tools | false |

The server binds to all interfaces (0.0.0.0). Logs are written to both stdout and `buildsim.log`.
