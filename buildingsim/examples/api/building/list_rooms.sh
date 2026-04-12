#!/bin/bash
# List all rooms across all floors
# Usage: ./list_rooms.sh [host:port] [level]
#   ./list_rooms.sh                - all floors
#   ./list_rooms.sh localhost:9090 level0  - specific floor

BASE=${1:-localhost:9090}
LEVEL=$2

if [ -n "$LEVEL" ]; then
    LEVELS="$LEVEL"
else
    LEVELS="level0 level1 level2"
fi

for L in $LEVELS; do
    curl -s http://$BASE/api/building/floors/$L | python3 -c "
import sys, json
data = json.load(sys.stdin)
rooms = data['rooms']
corridors = [r for r in rooms if r.get('type') == 'corridor']
regular = [r for r in rooms if r.get('type') != 'corridor']
print(f'$L: {len(regular)} rooms, {len(corridors)} corridors')
for r in sorted(regular, key=lambda x: x['name']):
    print(f'  {r[\"name\"]:12s}  area={r[\"area\"]:6.0f}  center=({r[\"center\"][0]:6.1f}, {r[\"center\"][1]:6.1f})')
"
done
