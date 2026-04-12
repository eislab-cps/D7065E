# BuildSim MCP Server

Model Context Protocol server that exposes BuildSim as tools for AI models. Allows Claude and other MCP-compatible AI to directly control the 3D building map.

## Setup

```bash
pip install mcp httpx
```

## Usage with Claude Code

Add to your Claude Code settings (`~/.claude/settings.json` or project `.claude/settings.json`):

```json
{
  "mcpServers": {
    "buildsim": {
      "command": "python3",
      "args": ["mcp/server.py", "--buildsim-url", "http://localhost:9090"]
    }
  }
}
```

Then start BuildSim and ask Claude to interact with the building.

## Available Tools

| Tool | Description |
|------|-------------|
| `get_building` | Get building metadata and floor levels |
| `get_rooms` | List all rooms on a floor |
| `search_room` | Search for a room by name across all floors |
| `list_equipment` | List equipment (filter by floor, room, category) |
| `add_equipment` | Add equipment to a room |
| `add_sensor` | Add a sensor to equipment |
| `set_sensor_value` | Update a sensor reading |
| `add_actuator` | Add an actuator to equipment |
| `set_actuator_state` | Control an actuator |
| `notify_equipment_change` | Notify browsers of equipment changes |
| `set_viewport` | Move camera to a room (smooth animation) |
| `highlight_rooms` | Color rooms on the map |
| `set_occupancy` | Place people and aliens in rooms |
| `set_coverage_zones` | Show coverage spheres (WiFi, risk zones) |
| `find_route` | Compute shortest path between rooms |
| `display_route` | Compute and show route on map |
| `clear_all` | Clear all visualizations |

## Example Prompts

Once connected, you can ask Claude things like:

- "Show me all rooms on level 1"
- "Add a temperature sensor in room A2306"
- "Highlight rooms 5 and 10 in red"
- "Find the shortest route from A2306 to 1542"
- "Place Johan in room A2306 and an alien in room 1542"
- "Show WiFi coverage centered on room A1306"
- "Zoom into room A2306"
- "Set up a fire detection system with smoke sensors on level 0"
