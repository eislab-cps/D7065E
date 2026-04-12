#!/bin/bash
# Find the most recently active session (browser)
# Usage: source find_session.sh [host:port] [session_id]
# Sets: BASE, SESSION

BASE=${1:-localhost:9090}

if [ -n "$2" ]; then
  SESSION="$2"
else
  SESSION=$(curl -s http://$BASE/api/sessions | python3 -c "
import sys, json
sessions = json.load(sys.stdin)
if sessions:
    sessions.sort(key=lambda s: s.get('last_ws_active', ''), reverse=True)
    print(sessions[0]['id'])
" 2>/dev/null)
fi

if [ -z "$SESSION" ]; then
  echo "No active session found. Open the browser first, or pass a session ID."
  exit 1
fi

echo "Using session: $SESSION"
