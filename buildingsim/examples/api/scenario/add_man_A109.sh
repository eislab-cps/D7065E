#!/bin/bash
# Add a man to room A109
# Usage: ./add_man_A109.sh [host:port]

DIR="$(cd "$(dirname "$0")" && pwd)"
source "$DIR/../session/find_session.sh" "$@"

echo "=== Adding man to A109 ==="
curl -s -X PUT http://$BASE/api/sessions/$SESSION/occupancy -H 'Content-Type: application/json' -d '{
  "97": {
    "persons": [{"id": "man-1", "name": "Johan"}],
    "aliens": []
  }
}' | python3 -m json.tool
