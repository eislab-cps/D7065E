#!/usr/bin/env python3
"""
BuildSim MCP Server — exposes the BuildSim REST API as MCP tools.

Allows Claude and other MCP-compatible AI models to directly control the building:
read sensors, command actuators, highlight rooms, set coverage zones, find routes, etc.

Usage:
    pip install mcp httpx
    python server.py [--buildsim-url http://localhost:9090]
"""

import json
import argparse
import logging
import sys
import httpx
from mcp.server.fastmcp import FastMCP

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(message)s",
    handlers=[
        logging.FileHandler("mcp-buildsim.log"),
        logging.StreamHandler(sys.stderr),
    ]
)
log = logging.getLogger("buildsim-mcp")

parser = argparse.ArgumentParser()
parser.add_argument("--buildsim-url", default="http://localhost:9090")
args, _ = parser.parse_known_args()

BASE = args.buildsim_url
log.info(f"BuildSim MCP server starting, API: {BASE}")
mcp = FastMCP("BuildSim", instructions="""
You control a simulated 3D building via the BuildSim API. The building has 3 floors
(level0, level1, level2) with ~300 rooms each, including corridors, offices, and labs.

You can read sensor data, control actuators, highlight rooms on the 3D map,
set coverage zones, place people and aliens in rooms, compute navigation routes,
and move the camera viewport.

A browser is viewing the building in 3D. When you set highlights, coverage, viewport,
or occupancy, the browser updates in real-time via WebSocket.
""")

class LoggingClient:
    """HTTP client wrapper that logs all requests."""
    def __init__(self, base_url, timeout=10):
        self._client = httpx.Client(base_url=base_url, timeout=timeout)

    def get(self, path, **kwargs):
        log.info(f"GET {path} {kwargs if kwargs else ''}")
        resp = self._client.get(path, **kwargs)
        log.info(f"  → {resp.status_code}")
        return resp

    def post(self, path, **kwargs):
        log.info(f"POST {path}")
        resp = self._client.post(path, **kwargs)
        log.info(f"  → {resp.status_code}")
        return resp

    def put(self, path, **kwargs):
        log.info(f"PUT {path}")
        resp = self._client.put(path, **kwargs)
        log.info(f"  → {resp.status_code}")
        return resp

    def delete(self, path, **kwargs):
        log.info(f"DELETE {path}")
        resp = self._client.delete(path, **kwargs)
        log.info(f"  → {resp.status_code}")
        return resp

client = LoggingClient(base_url=BASE, timeout=10)


def _find_session():
    """Find the most recently active browser session."""
    resp = client.get("/api/sessions")
    resp.raise_for_status()
    sessions = resp.json()
    if not sessions:
        resp = client.post("/api/sessions")
        resp.raise_for_status()
        sid = resp.json()["id"]
        log.info(f"Created new session: {sid}")
        return sid
    sessions.sort(key=lambda s: s.get("last_ws_active", ""), reverse=True)
    sid = sessions[0]["id"]
    log.info(f"Using session: {sid}")
    return sid


# === Building Data ===

@mcp.tool()
def get_building() -> str:
    """Get building metadata: name and available floor levels."""
    resp = client.get("/api/building")
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def get_rooms(level: str = "level0") -> str:
    """Get all rooms on a floor with their names, types, areas, and centers.

    Args:
        level: Floor level (level0, level1, level2)
    """
    resp = client.get(f"/api/building/floors/{level}")
    data = resp.json()
    rooms = [{"id": r["id"], "name": r["name"], "type": r.get("type", "room"),
              "area": r["area"], "center": r["center"]} for r in data["rooms"]]
    return json.dumps(rooms, indent=2)


@mcp.tool()
def search_room(name: str) -> str:
    """Search for a room by name across all floors. Returns room details and which floor it's on.

    Args:
        name: Room name or partial name to search for (e.g. "A2306", "1542")
    """
    results = []
    for level in ["level0", "level1", "level2"]:
        resp = client.get(f"/api/building/floors/{level}")
        for r in resp.json()["rooms"]:
            if name.lower() in r["name"].lower():
                results.append({"name": r["name"], "level": level, "id": r["id"],
                                "type": r.get("type", "room"), "area": r["area"],
                                "center": r["center"]})
    return json.dumps(results, indent=2)


# === Equipment ===

@mcp.tool()
def list_equipment(level: str = "", room: str = "", category: str = "") -> str:
    """List all equipment, optionally filtered by floor, room, or category.

    Args:
        level: Filter by floor (level0, level1, level2). Empty for all.
        room: Filter by room name. Empty for all.
        category: Filter by category (monitoring, hvac, safety_fire, access_control, etc). Empty for all.
    """
    params = {}
    if level: params["level"] = level
    if room: params["room"] = room
    if category: params["category"] = category
    resp = client.get("/api/equipment", params=params)
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def add_equipment(id: str, name: str, type: str, category: str, level: str, room: str, status: str = "running") -> str:
    """Add a new equipment item to the building.

    Args:
        id: Unique equipment ID
        name: Display name
        type: Equipment type (e.g. temperature_sensor, door_lock, smoke_detector, ac_unit, wifi_access_point)
        category: Category (monitoring, hvac, safety_fire, access_control, electrical, plumbing, network, fixed_equipment)
        level: Floor level (level0, level1, level2)
        room: Room name (e.g. "A2306", "1542")
        status: Status (running, stopped, warning, alarm)
    """
    resp = client.post("/api/equipment", json={
        "id": id, "name": name, "type": type, "category": category,
        "level": level, "room": room, "status": status
    })
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def add_sensor(equipment_id: str, sensor_id: str, name: str, type: str, data_type: str = "text", unit: str = "", value: str = "0") -> str:
    """Add a sensor to an equipment item.

    Args:
        equipment_id: ID of the equipment to add the sensor to
        sensor_id: Unique sensor ID
        name: Sensor display name
        type: Sensor type (e.g. temperature, humidity, smoke_level, door_position, signal_strength)
        data_type: "text" for numeric/string values, "binary" for on/off
        unit: Unit of measurement (e.g. "°C", "ppm", "%", "dBm")
        value: Initial value
    """
    resp = client.post(f"/api/equipment/{equipment_id}/sensors", json={
        "id": sensor_id, "name": name, "type": type,
        "data_type": data_type, "unit": unit, "value": value
    })
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def set_sensor_value(sensor_id: str, value: str, data_type: str = "text") -> str:
    """Set a sensor's current value.

    Args:
        sensor_id: Sensor ID
        value: New value (e.g. "23.5" for temperature, "true"/"false" for binary)
        data_type: "text" or "binary"
    """
    body = {"data_type": data_type, "value": value}
    if data_type == "binary":
        body = {"data_type": "binary", "binary_value": value.lower() in ("true", "1", "yes")}
    resp = client.put(f"/api/sensors/{sensor_id}/value", json=body)
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def add_actuator(equipment_id: str, actuator_id: str, name: str, type: str, state: str = "off") -> str:
    """Add an actuator to an equipment item.

    Args:
        equipment_id: ID of the equipment
        actuator_id: Unique actuator ID
        name: Actuator display name
        type: Actuator type (e.g. lock_control, fan_speed, on_off, dimmer)
        state: Initial state
    """
    resp = client.post(f"/api/equipment/{equipment_id}/actuators", json={
        "id": actuator_id, "name": name, "type": type, "state": state
    })
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def set_actuator_state(actuator_id: str, state: str) -> str:
    """Set an actuator's state.

    Args:
        actuator_id: Actuator ID
        state: New state (e.g. "locked", "unlocked", "on", "off", "high", "low")
    """
    resp = client.put(f"/api/actuators/{actuator_id}/state", json={"state": state})
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def notify_equipment_change() -> str:
    """Notify all browser sessions that equipment has changed. Call this after adding/modifying equipment so the browser refreshes."""
    resp = client.post("/api/equipment/notify")
    return json.dumps(resp.json(), indent=2)


# === Session Control (viewport, highlights, occupancy, coverage, route) ===

@mcp.tool()
def set_viewport(room: str, zoom: float = 2.0, mode: str = "3d") -> str:
    """Move the browser camera to focus on a room. The camera animates smoothly.

    Args:
        room: Room name to center on (e.g. "A2306")
        zoom: Zoom level (1.0 = default, higher = closer)
        mode: "3d" or "2d"
    """
    session = _find_session()
    resp = client.put(f"/api/sessions/{session}/viewport", json={
        "room": room, "zoom": zoom, "mode": mode
    })
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def highlight_rooms(rooms: str) -> str:
    """Highlight rooms on the 3D map with colors. Replaces any existing highlights.

    Args:
        rooms: JSON array of highlights, e.g. '[{"room_id": 5, "color": "#ff0000", "opacity": 0.8}]'. Use room IDs (integers). Set to '[]' to clear.
    """
    session = _find_session()
    highlights = json.loads(rooms)
    resp = client.put(f"/api/sessions/{session}/highlights", json=highlights)
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def set_occupancy(occupancy: str) -> str:
    """Set people and aliens in rooms. The browser shows person/group/xenomorph icons.

    Args:
        occupancy: JSON object mapping room IDs to occupants, e.g. '{"5": {"persons": [{"id": "p1", "name": "Alice"}], "aliens": []}}'.  Set to '{}' to clear.
    """
    session = _find_session()
    data = json.loads(occupancy)
    resp = client.put(f"/api/sessions/{session}/occupancy", json=data)
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def set_coverage_zones(zones: str) -> str:
    """Set coverage zones (translucent spheres) on the 3D map. Useful for WiFi coverage, risk zones, sensor range visualization. Clipped to building bounds.

    Args:
        zones: JSON array of zones, e.g. '[{"id": "wifi-1", "name": "WiFi AP", "room": "A2306", "radius": 30, "color": "#00aaff", "opacity": 0.12, "level": "", "height": 40}]'. Set room to center on a room. Set level to "" for all floors. Set to '[]' to clear.
    """
    session = _find_session()
    data = json.loads(zones)
    resp = client.put(f"/api/sessions/{session}/coverage", json=data)
    return json.dumps(resp.json(), indent=2)


@mcp.tool()
def find_route(from_room: str, to_room: str) -> str:
    """Compute the shortest walking route between two rooms using Dijkstra on the walkable navigation graph. Supports cross-floor routing via stairs/elevators.

    Args:
        from_room: Source room name (e.g. "A2306")
        to_room: Destination room name (e.g. "1542")
    """
    resp = client.get("/api/graph/route", params={
        "from_name": from_room, "to_name": to_room, "type": "walkable"
    })
    if resp.status_code == 404:
        return json.dumps({"error": "No route found between " + from_room + " and " + to_room})
    result = resp.json()
    return json.dumps({
        "from": from_room, "to": to_room,
        "distance": result["distance"],
        "waypoints": len(result["path"]),
        "floors": list(set(n.get("level", "?") for n in result["path"]))
    }, indent=2)


@mcp.tool()
def display_route(from_room: str, to_room: str) -> str:
    """Compute and display a route on the 3D map. The route is shown as a smooth green tube.

    Args:
        from_room: Source room name
        to_room: Destination room name
    """
    resp = client.get("/api/graph/route", params={
        "from_name": from_room, "to_name": to_room, "type": "walkable"
    })
    if resp.status_code == 404:
        return json.dumps({"error": "No route found"})
    route = resp.json()
    session = _find_session()
    client.put(f"/api/sessions/{session}/route", json=route)
    return json.dumps({
        "status": "route displayed",
        "from": from_room, "to": to_room,
        "distance": route["distance"],
        "waypoints": len(route["path"])
    }, indent=2)


@mcp.tool()
def clear_all() -> str:
    """Clear all session visualizations: highlights, coverage zones, occupancy, and routes."""
    session = _find_session()
    client.put(f"/api/sessions/{session}/highlights", json=[])
    client.put(f"/api/sessions/{session}/coverage", json=[])
    client.put(f"/api/sessions/{session}/occupancy", json={})
    client.put(f"/api/sessions/{session}/route", json={"path": [], "distance": 0})
    return json.dumps({"status": "cleared"})


if __name__ == "__main__":
    mcp.run()
