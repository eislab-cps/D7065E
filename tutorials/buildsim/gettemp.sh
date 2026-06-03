#!/usr/bin/env bash
# Print the current temperature reading for hvac-A109 from BuildSim.
# A sensor value is read from its parent equipment (there is no GET /api/sensors/{id}).
BASE=${1:-http://localhost:9090}

resp=$(curl -s "$BASE/api/equipment/hvac-A109")
if [ -z "$resp" ]; then
  echo "no response from $BASE -- is BuildSim running? (./buildsim start --port 9090)" >&2
  exit 1
fi

echo "$resp" | python3 -c '
import sys, json
eq = json.load(sys.stdin)
if "error" in eq or not eq.get("sensors"):
    sys.exit("hvac-A109 not found -- run ./setup.sh first")
for s in eq["sensors"]:
    if s["id"] == "A109-temp":
        print(s["value"], s["unit"])
        break
'
