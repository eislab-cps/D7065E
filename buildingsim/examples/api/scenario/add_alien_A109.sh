#!/bin/bash
# Add an alien to room A109 (alongside man and woman)
# Usage: ./add_alien_A109.sh [host:port]

BASE=${1:-localhost:9090}

echo "=== Adding alien to A109 ==="
curl -s -X PUT http://$BASE/api/occupancy -H 'Content-Type: application/json' -d '{
  "97": {
    "persons": [{"id": "man-1", "name": "Johan", "icon": "man"}, {"id": "woman-1", "name": "Alice", "icon": "woman"}],
    "aliens": [{"id": "xeno-1"}]
  }
}' | python3 -m json.tool
