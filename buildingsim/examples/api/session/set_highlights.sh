#!/bin/bash
# Set room highlights with colors
# Usage: ./set_highlights.sh [host:port] [session_id]

DIR="$(cd "$(dirname "$0")" && pwd)"
source "$DIR/find_session.sh" "$@"

echo ""
echo "=== Highlighting Rooms ==="
curl -s -X PUT http://$BASE/api/sessions/$SESSION/highlights -H 'Content-Type: application/json' -d '[
  {"room_id": 5, "color": "#ff0000", "opacity": 0.8},
  {"room_id": 10, "color": "#00ff00", "opacity": 0.5},
  {"room_id": 15, "color": "#ffaa00", "opacity": 0.6}
]' | python3 -m json.tool
