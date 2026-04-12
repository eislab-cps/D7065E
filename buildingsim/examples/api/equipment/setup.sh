#!/bin/bash
# Setup equipment via REST API
# Usage: ./setup_equipment.sh [host:port]

BASE=${1:-localhost:9090}

echo "=== Adding Equipment ==="

# Door lock
curl -s -X POST http://$BASE/api/equipment -H 'Content-Type: application/json' -d '{
  "id": "door-1",
  "name": "Main Entrance Door",
  "type": "door_lock",
  "category": "access_control",
  "level": "level0",
  "room": "1542",
  "status": "running"
}' | python3 -m json.tool

# Add sensor: door position
curl -s -X POST http://$BASE/api/equipment/door-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "door-1-pos",
  "name": "Door Position",
  "type": "door_position",
  "data_type": "binary",
  "unit": ""
}' | python3 -m json.tool

# Add actuator: lock control
curl -s -X POST http://$BASE/api/equipment/door-1/actuators -H 'Content-Type: application/json' -d '{
  "id": "door-1-lock",
  "name": "Lock Control",
  "type": "lock_control",
  "state": "locked"
}' | python3 -m json.tool

# Temperature sensor
curl -s -X POST http://$BASE/api/equipment -H 'Content-Type: application/json' -d '{
  "id": "temp-1",
  "name": "Room Temperature Sensor",
  "type": "temperature_sensor",
  "category": "monitoring",
  "level": "level0",
  "room": "1547",
  "status": "running"
}' | python3 -m json.tool

curl -s -X POST http://$BASE/api/equipment/temp-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "temp-1-reading",
  "name": "Temperature",
  "type": "temperature",
  "data_type": "text",
  "unit": "°C",
  "value": "21.5"
}' | python3 -m json.tool

# AC unit with multiple sensors and actuators
curl -s -X POST http://$BASE/api/equipment -H 'Content-Type: application/json' -d '{
  "id": "ac-1",
  "name": "AC Unit Conference Room",
  "type": "ac_unit",
  "category": "hvac",
  "level": "level0",
  "room": "1556",
  "status": "running"
}' | python3 -m json.tool

curl -s -X POST http://$BASE/api/equipment/ac-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "ac-1-temp",
  "name": "Supply Air Temperature",
  "type": "temperature",
  "data_type": "text",
  "unit": "°C",
  "value": "18.2"
}' | python3 -m json.tool

curl -s -X POST http://$BASE/api/equipment/ac-1/actuators -H 'Content-Type: application/json' -d '{
  "id": "ac-1-fan",
  "name": "Fan Speed",
  "type": "fan_speed",
  "state": "medium"
}' | python3 -m json.tool

curl -s -X POST http://$BASE/api/equipment/ac-1/actuators -H 'Content-Type: application/json' -d '{
  "id": "ac-1-power",
  "name": "Power",
  "type": "on_off",
  "state": "on"
}' | python3 -m json.tool

# Smoke detector
curl -s -X POST http://$BASE/api/equipment -H 'Content-Type: application/json' -d '{
  "id": "smoke-1",
  "name": "Smoke Detector Hallway",
  "type": "smoke_detector",
  "category": "safety_fire",
  "level": "level0",
  "room": "1544",
  "status": "running"
}' | python3 -m json.tool

curl -s -X POST http://$BASE/api/equipment/smoke-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "smoke-1-level",
  "name": "Smoke Level",
  "type": "smoke_level",
  "data_type": "text",
  "unit": "ppm",
  "value": "0.3"
}' | python3 -m json.tool

curl -s -X POST http://$BASE/api/equipment/smoke-1/sensors -H 'Content-Type: application/json' -d '{
  "id": "smoke-1-alarm",
  "name": "Alarm Status",
  "type": "alarm_status",
  "data_type": "binary",
  "unit": ""
}' | python3 -m json.tool

# Notify all sessions about new equipment
curl -s -X POST http://$BASE/api/equipment/notify | python3 -m json.tool

echo ""
echo "=== Equipment loaded ==="
curl -s http://$BASE/api/equipment | python3 -c "import sys,json; eq=json.load(sys.stdin); print(f'{len(eq)} equipment items loaded')"
