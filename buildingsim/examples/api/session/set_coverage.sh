#!/bin/bash
# Set coverage zones (e.g. WiFi coverage, sensor range)
# Usage: ./set_coverage.sh [host:port] [session_id]

DIR="$(cd "$(dirname "$0")" && pwd)"
source "$DIR/find_session.sh" "$@"

echo ""
echo "=== Setting Coverage Zones ==="
curl -s -X PUT http://$BASE/api/coverage -H 'Content-Type: application/json' -d '[
  {
    "id": "wifi-1",
    "name": "WiFi AP Central",
    "room": "A1306",
    "radius": 35,
    "color": "#00aaff",
    "opacity": 0.12,
    "level": "",
    "height": 40
  },
  {
    "id": "wifi-2",
    "name": "WiFi AP East",
    "room": "A1123",
    "radius": 30,
    "color": "#00aaff",
    "opacity": 0.12,
    "level": "",
    "height": 40
  },
  {
    "id": "wifi-3",
    "name": "WiFi AP South",
    "room": "A182",
    "radius": 30,
    "color": "#00aaff",
    "opacity": 0.12,
    "level": "",
    "height": 40
  },
  {
    "id": "smoke-zone",
    "name": "Smoke Detection Zone",
    "room": "1540",
    "radius": 15,
    "color": "#ff4400",
    "opacity": 0.1,
    "level": "level0",
    "height": 0
  },
  {
    "id": "motion-1",
    "name": "Motion Sensor Range",
    "room": "1556",
    "radius": 12,
    "color": "#44ff44",
    "opacity": 0.15,
    "level": "level0",
    "height": 0
  }
]' | python3 -m json.tool

echo ""
echo "=== To clear coverage ==="
echo "curl -X PUT http://$BASE/api/coverage -H 'Content-Type: application/json' -d '[]'"
