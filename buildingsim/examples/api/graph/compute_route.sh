#!/bin/bash
# Compute shortest path using Dijkstra via REST API
# Usage: ./compute_route.sh [host:port]

BASE=${1:-localhost:9090}

echo "=== Computing Route (room 0 -> room 5, level0) ==="
curl -s "http://$BASE/api/graph/route?from=0&to=5&level=level0" | python3 -m json.tool

echo ""
echo "=== Computing Route (room 0 -> room 50, level0) ==="
curl -s "http://$BASE/api/graph/route?from=0&to=50&level=level0" | python3 -m json.tool

echo ""
echo "=== Computing Route (room 10 -> room 100, level1) ==="
curl -s "http://$BASE/api/graph/route?from=10&to=100&level=level1" | python3 -m json.tool

echo ""
echo "=== Full Navigation Graph (level0 summary) ==="
curl -s "http://$BASE/api/graph?level=level0" | python3 -c "
import sys, json
g = json.load(sys.stdin)
print(f'Nodes: {len(g[\"nodes\"])}')
print(f'Edges: {len(g[\"edges\"])}')
print('First 5 nodes:')
for n in g['nodes'][:5]:
    print(f'  {n[\"id\"]}: {n[\"name\"]} ({n[\"x\"]}, {n[\"y\"]})')
"
