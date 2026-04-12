#!/bin/bash
# Add a woman to room A109
# Usage: ./add_woman_A109.sh [host:port]

DIR="$(cd "$(dirname "$0")" && pwd)"
source "$DIR/../session/find_session.sh" "$@"

echo "=== Adding woman to A109 ==="
curl -s -X PUT http://$BASE/api/occupancy -H 'Content-Type: application/json' -d '{
  "97": {
    "persons": [{"id": "man-1", "name": "Johan", "icon": "man"}, {"id": "woman-1", "name": "Alice", "icon": "woman"}],
    "aliens": []
  }
}' | python3 -m json.tool
