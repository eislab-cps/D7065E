#!/bin/bash
# Add a temperature sensor to room A109
# Usage: ./add_temp_sensor_A109.sh [host:port]

BASE=${1:-localhost:9090}

echo "=== Adding temperature sensor to A109 ==="

curl -s -X POST http://$BASE/api/equipment -H 'Content-Type: application/json' -d '{
  "id": "temp-A109",
  "name": "Temperature Sensor A109",
  "type": "temperature_sensor",
  "category": "monitoring",
  "level": "level0",
  "room": "A109",
  "status": "running"
}' | python3 -m json.tool

curl -s -X POST http://$BASE/api/equipment/temp-A109/sensors -H 'Content-Type: application/json' -d '{
  "id": "temp-A109-val",
  "name": "Temperature",
  "type": "temperature",
  "data_type": "text",
  "unit": "°C",
  "value": "21.0"
}' | python3 -m json.tool

curl -s -X POST http://$BASE/api/equipment/notify > /dev/null
echo "Done."
