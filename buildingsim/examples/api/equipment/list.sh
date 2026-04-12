#!/bin/bash
# List all equipment with sensors and actuators
# Usage: ./list.sh [host:port] [filter]
#   ./list.sh                          - all equipment
#   ./list.sh localhost:9090 level0    - filter by level
#   ./list.sh localhost:9090 A109      - filter by room

BASE=${1:-localhost:9090}
FILTER=$2

curl -s http://$BASE/api/equipment | python3 -c "
import sys, json
equipment = json.load(sys.stdin)
filt = '$FILTER'
if filt:
    equipment = [e for e in equipment if filt in (e.get('level',''), e.get('room',''), e.get('type',''), e.get('category',''))]

print(f'{len(equipment)} equipment items')
for eq in sorted(equipment, key=lambda e: (e.get('level',''), e.get('room',''))):
    sensors = eq.get('sensors', [])
    actuators = eq.get('actuators', [])
    print(f'  {eq[\"id\"]:30s}  {eq[\"type\"]:25s}  {eq[\"level\"]:6s}  room={eq[\"room\"]:10s}  status={eq[\"status\"]}')
    for s in sensors:
        val = s.get('value', str(s.get('binary_value','')))
        print(f'    sensor:   {s[\"id\"]:30s}  {s[\"type\"]:15s}  {val} {s.get(\"unit\",\"\")}')
    for a in actuators:
        print(f'    actuator: {a[\"id\"]:30s}  {a[\"type\"]:15s}  state={a[\"state\"]}')
"
