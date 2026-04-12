#!/bin/bash
# Set sensor values via REST API
# Usage: ./set_sensor_values.sh [host:port]

BASE=${1:-localhost:9090}

echo "=== Setting Sensor Values ==="

# Set temperature
echo "Setting temperature to 23.7°C..."
curl -s -X PUT http://$BASE/api/sensors/temp-1-reading/value -H 'Content-Type: application/json' -d '{
  "data_type": "text",
  "value": "23.7"
}' | python3 -m json.tool

# Set door to open (binary)
echo "Setting door to open..."
curl -s -X PUT http://$BASE/api/sensors/door-1-pos/value -H 'Content-Type: application/json' -d '{
  "data_type": "binary",
  "binary_value": true
}' | python3 -m json.tool

# Trigger smoke alarm
echo "Triggering smoke alarm..."
curl -s -X PUT http://$BASE/api/sensors/smoke-1-alarm/value -H 'Content-Type: application/json' -d '{
  "data_type": "binary",
  "binary_value": true
}' | python3 -m json.tool

curl -s -X PUT http://$BASE/api/sensors/smoke-1-level/value -H 'Content-Type: application/json' -d '{
  "data_type": "text",
  "value": "15.8"
}' | python3 -m json.tool

# Set actuator state
echo "Unlocking door..."
curl -s -X PUT http://$BASE/api/actuators/door-1-lock/state -H 'Content-Type: application/json' -d '{
  "state": "unlocked"
}' | python3 -m json.tool

echo "Setting AC fan to high..."
curl -s -X PUT http://$BASE/api/actuators/ac-1-fan/state -H 'Content-Type: application/json' -d '{
  "state": "high"
}' | python3 -m json.tool

# Bump version to notify browsers
curl -s -X POST http://$BASE/api/equipment/notify | python3 -m json.tool

echo "=== Done ==="
