#!/usr/bin/env python3
"""Populate the building with sensors and actuators in every room."""

import requests
import sys
import time

BASE = sys.argv[1] if len(sys.argv) > 1 else "http://localhost:9090"
s = requests.Session()

start = time.time()
equipment = []

def eq(id, name, type, cat, level, room, sensors=None, actuators=None):
    equipment.append({
        "id": id, "name": name, "type": type, "category": cat,
        "level": level, "room": room, "status": "running",
        "sensors": sensors or [], "actuators": actuators or []
    })

for level in ["level0", "level1", "level2"]:
    resp = s.get(f"{BASE}/api/building/floors/{level}")
    rooms = [r["name"] for r in resp.json()["rooms"] if r.get("type") != "corridor"]
    print(f"{level}: {len(rooms)} rooms")

    for i, room in enumerate(rooms):
        # Temperature sensor in every room
        eq(f"temp-{level}-{room}", f"Temp {room}", "temperature_sensor", "monitoring", level, room,
           sensors=[{"id": f"temp-{level}-{room}-val", "name": "Temperature", "type": "temperature",
                     "data_type": "text", "unit": "°C", "value": "21.5"}])

        # Door lock in every room
        eq(f"door-{level}-{room}", f"Door {room}", "door_lock", "access_control", level, room,
           sensors=[{"id": f"door-{level}-{room}-pos", "name": "Position", "type": "door_position",
                     "data_type": "binary", "unit": "", "value": ""}],
           actuators=[{"id": f"door-{level}-{room}-lock", "name": "Lock", "type": "lock_control",
                       "state": "locked"}])

        # Smoke detector every 3rd room
        if i % 3 == 0:
            eq(f"smoke-{level}-{room}", f"Smoke {room}", "smoke_detector", "safety_fire", level, room,
               sensors=[{"id": f"smoke-{level}-{room}-lvl", "name": "Smoke Level", "type": "smoke_level",
                         "data_type": "text", "unit": "ppm", "value": "0.2"}])

        # Motion sensor every 4th room
        if i % 4 == 0:
            eq(f"motion-{level}-{room}", f"Motion {room}", "motion_sensor", "monitoring", level, room,
               sensors=[{"id": f"motion-{level}-{room}-det", "name": "Motion", "type": "motion",
                         "data_type": "binary", "unit": "", "value": ""}])

        # HVAC every 5th room
        if i % 5 == 0:
            eq(f"hvac-{level}-{room}", f"HVAC {room}", "ac_unit", "hvac", level, room,
               actuators=[
                   {"id": f"hvac-{level}-{room}-fan", "name": "Fan", "type": "fan_speed", "state": "medium"},
                   {"id": f"hvac-{level}-{room}-pwr", "name": "Power", "type": "on_off", "state": "on"}
               ])

        # Sprinkler every 6th room
        if i % 6 == 0:
            eq(f"sprink-{level}-{room}", f"Sprinkler {room}", "sprinkler", "safety_fire", level, room,
               actuators=[{"id": f"sprink-{level}-{room}-act", "name": "Activate", "type": "on_off", "state": "off"}])

# WiFi APs
for level, room in [("level0","A1306"),("level0","A1123"),("level1","A2306"),
                     ("level1","A2318"),("level2","A3003"),("level2","A3318")]:
    eq(f"wifi-{level}-{room}", f"WiFi AP {room}", "wifi_access_point", "network", level, room,
       sensors=[
           {"id": f"wifi-{level}-{room}-sig", "name": "Signal", "type": "signal_strength", "data_type": "text", "unit": "dBm", "value": "-58"},
           {"id": f"wifi-{level}-{room}-usr", "name": "Users", "type": "connected_users", "data_type": "text", "unit": "", "value": "10"}
       ])

# Single bulk request
print(f"\nSending {len(equipment)} equipment in one request...")
resp = s.post(f"{BASE}/api/equipment/bulk", json=equipment)
result = resp.json()
elapsed = time.time() - start

print(f"Done in {elapsed:.2f}s: {result}")
