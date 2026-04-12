#!/bin/bash
# Set viewport: zoom to a room
# Usage: ./set_viewport.sh [host:port] [session_id]

DIR="$(cd "$(dirname "$0")" && pwd)"
source "$DIR/find_session.sh" "$@"

echo ""
echo "=== Setting Viewport (zoom into room A2306, 3D mode) ==="
curl -s -X PUT http://$BASE/api/sessions/$SESSION/viewport -H 'Content-Type: application/json' -d '{
  "room": "A2306",
  "zoom": 2.0,
  "mode": "3d"
}' | python3 -m json.tool
