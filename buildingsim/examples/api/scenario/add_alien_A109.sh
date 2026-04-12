#!/bin/bash
# Add an alien (xenomorph) to room A109
# Usage: ./add_alien_A109.sh [host:port]

DIR="$(cd "$(dirname "$0")" && pwd)"
source "$DIR/../session/find_session.sh" "$@"

echo "=== Adding alien to A109 ==="
curl -s -X PUT http://$BASE/api/occupancy -H 'Content-Type: application/json' -d '{
  "97": {
    "persons": [{"id": "man-1", "name": "Johan"}, {"id": "woman-1", "name": "Alice"}],
    "aliens": [{"id": "xeno-1"}]
  }
}' | python3 -m json.tool
