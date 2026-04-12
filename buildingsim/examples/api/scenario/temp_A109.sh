#!/bin/bash
# Set and read temperature for room A109
# Usage: ./temp_A109.sh [set VALUE] [host:port]
#   ./temp_A109.sh           - read current temperature
#   ./temp_A109.sh set 23.5  - set temperature to 23.5°C

BASE=${2:-${1:-localhost:9090}}
ACTION=${1:-read}

if [ "$ACTION" = "set" ]; then
    VALUE=$2
    BASE=${3:-localhost:9090}
    echo "Setting A109 temperature to ${VALUE}°C..."
    curl -s -X PUT http://$BASE/api/sensors/temp-A109-val/value \
        -H 'Content-Type: application/json' \
        -d "{\"data_type\": \"text\", \"value\": \"$VALUE\"}" | python3 -m json.tool
    curl -s -X POST http://$BASE/api/equipment/notify > /dev/null
    echo "Done."
else
    echo "=== Temperature in A109 ==="
    curl -s http://$BASE/api/equipment/temp-A109 | python3 -c "
import sys, json
eq = json.load(sys.stdin)
for s in eq.get('sensors', []):
    if s['type'] == 'temperature':
        print(f'{s[\"value\"]} {s[\"unit\"]}  (last update: {s[\"timestamp\"]})')
"
fi
