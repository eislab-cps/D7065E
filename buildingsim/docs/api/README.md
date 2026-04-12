# BuildSim API Reference

All endpoints are relative to `http://localhost:9090`.

| Document | Endpoints |
|----------|-----------|
| [Building Data](building.md) | `GET /api/building`, `GET /api/building/floors/{level}`, `GET /api/building/cross-floor-edges` |
| [Equipment](equipment.md) | `POST/GET/PUT/DELETE /api/equipment`, `POST /api/equipment/bulk`, sensors, actuators |
| [Sessions](sessions.md) | `POST/GET/DELETE /api/sessions`, viewport, highlights, occupancy, coverage, route, WebSocket |
| [Navigation Graph](graph.md) | `GET /api/graph`, `GET /api/graph/route` (Dijkstra) |
| [Icons](icons.md) | `GET /api/icons/{name}.svg` |
| [Config](config.md) | `GET /api/config` |
