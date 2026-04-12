# BuildSim

Building simulation server with REST API and WebSocket session control. A single Go binary embeds building floor plan data, a 3D web viewer, and SVG equipment icons.

## Quick Start

```bash
# Build
go build -o buildsim ./cmd

# Run
./buildsim start --port 9090

# Run with floor plan editing enabled
./buildsim start --port 9090 --edit
```

Open http://localhost:9090 in a browser to view the 3D building.

## Features

- **3D Building Viewer** -- Multi-floor building visualization with pan, zoom, and orbit controls
- **REST API** -- Full CRUD for equipment, sensors, and actuators
- **Session Control** -- Programmatic control of the viewer (viewport, room highlights, occupancy, routes) via REST + WebSocket
- **Navigation Graph** -- Room adjacency graph and walkable corridor graph with Dijkstra shortest path routing (single and cross-floor)
- **Equipment Icons** -- 48 detailed SVG icons for all equipment types, loaded dynamically
- **Occupancy** -- Person, group, and alien (xenomorph) icons rendered in rooms via session API

## Documentation

- [Architecture](docs/architecture.md) -- System diagrams, data model, session flow, project structure
- [Equipment Icons](docs/icons.md) -- Gallery of all 48 SVG icons

### API Reference

- [Building Data](docs/api/building.md) -- Floor plans, rooms, cross-floor edges
- [Equipment](docs/api/equipment.md) -- Equipment CRUD, sensors, actuators
- [Sessions](docs/api/sessions.md) -- Session lifecycle, viewport, highlights, occupancy, routes
- [Navigation Graph](docs/api/graph.md) -- Graph queries and Dijkstra routing
- [Icons](docs/api/icons.md) -- SVG icon endpoint
- [Config](docs/api/config.md) -- Server configuration

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

See [examples/api/](examples/api/) for ready-to-run bash scripts:

```bash
# Load equipment with sensors and actuators
./examples/api/equipment/setup.sh

# Set sensor values
./examples/api/equipment/set_sensors.sh

# Control browser viewport, highlights, occupancy
./examples/api/session/set_viewport.sh
./examples/api/session/set_highlights.sh
./examples/api/session/set_occupancy.sh

# Compute shortest path
./examples/api/graph/compute_route.sh
```

## CLI

```
buildsim start [flags]

Flags:
  --port, -p int   Port to listen on (default 9090)
  --edit           Enable floor plan editing tools
```

The server binds to all interfaces (0.0.0.0). Logs are written to both stdout and `buildsim.log`.
