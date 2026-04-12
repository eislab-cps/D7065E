#!/bin/bash
# Add a 5G base station with sensors and coverage zone
# Usage: ./add_5g.sh [host:port] [session_id]

BASE=${1:-localhost:9090}

echo "=== Adding 5G Base Station ==="

# Create equipment
curl -s -X POST http://$BASE/api/equipment -H 'Content-Type: application/json' -d '{
  "id": "5g-bs-1",
  "name": "5G Base Station A-Building",
  "type": "base_station_5g",
  "category": "network",
  "level": "level1",
  "room": "A2306",
  "status": "running"
}' | python3 -m json.tool

# Add sensors
curl -s -X POST http://$BASE/api/equipment/5g-bs-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "5g-bs-1-signal",
  "name": "Signal Strength",
  "type": "signal_strength",
  "data_type": "text",
  "unit": "dBm",
  "value": "-65"
}'

curl -s -X POST http://$BASE/api/equipment/5g-bs-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "5g-bs-1-users",
  "name": "Connected Users",
  "type": "connected_users",
  "data_type": "text",
  "unit": "",
  "value": "42"
}'

curl -s -X POST http://$BASE/api/equipment/5g-bs-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "5g-bs-1-band",
  "name": "Frequency Band",
  "type": "frequency_band",
  "data_type": "text",
  "unit": "GHz",
  "value": "3.5"
}'

curl -s -X POST http://$BASE/api/equipment/5g-bs-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "5g-bs-1-temp",
  "name": "Unit Temperature",
  "type": "temperature",
  "data_type": "text",
  "unit": "°C",
  "value": "38.2"
}'

# Add actuator
curl -s -X POST http://$BASE/api/equipment/5g-bs-1/actuators -H 'Content-Type: application/json' -d '{
  "id": "5g-bs-1-power",
  "name": "Power Control",
  "type": "on_off",
  "state": "on"
}'

# Notify browsers about new equipment
curl -s -X POST http://$BASE/api/equipment/notify > /dev/null

echo ""
echo "=== Setting 5G Coverage Zone ==="

# Find browser session
DIR="$(cd "$(dirname "$0")" && pwd)"
source "$DIR/../session/find_session.sh" "$@"

curl -s -X PUT http://$BASE/api/coverage -H 'Content-Type: application/json' -d '[
  {
    "id": "5g-coverage-1",
    "name": "5G Coverage - A-Building",
    "room": "A2306",
    "radius": 45,
    "color": "#ff00cc",
    "opacity": 0.1,
    "level": "",
    "height": 50
  }
]' | python3 -m json.tool

echo ""
echo "=== Done ==="
echo "Equipment: 5g-bs-1 in room A2306"
echo "Coverage: purple sphere centered on A2306"
