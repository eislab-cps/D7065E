#!/usr/bin/env bash
# Create the HVAC equipment, then add a temperature sensor and a setpoint actuator.
# Sensors/actuators must be added via their own endpoints so they are addressable
# by PUT /api/sensors/{id}/value and PUT /api/actuators/{id}/state.
BASE=${1:-http://localhost:9090}

curl -s -X POST "$BASE/api/equipment" -H 'Content-Type: application/json' \
  -d '{"id":"hvac-A109","name":"HVAC A109","type":"ac_unit","category":"hvac","level":"level0","room":"A109","status":"running"}' >/dev/null

curl -s -X POST "$BASE/api/equipment/hvac-A109/sensors" -H 'Content-Type: application/json' \
  -d '{"id":"A109-temp","name":"Temperature","type":"temperature","data_type":"text","unit":"°C","value":"18.0"}' >/dev/null

curl -s -X POST "$BASE/api/equipment/hvac-A109/actuators" -H 'Content-Type: application/json' \
  -d '{"id":"A109-set","name":"Setpoint","type":"setpoint","state":"21"}' >/dev/null

echo "created hvac-A109 (sensor A109-temp, actuator A109-set)"
