# Workshop Challenges

Each lecture is an interactive workshop. Two challenges are given in advance — students read the lecture material and prepare, then the group discusses solutions together during the session.

---

## Lecture 1: Introduction & MBSE

### Challenge A: Decompose a Fire Safety System

A building has 300 rooms across 3 floors. You must design a fire detection and response system. Before the workshop:

1. Write 5 testable requirements (functional and non-functional) for the system
2. Draw a C4 Context diagram showing the system, its actors, and external systems
3. Draw a C4 Container diagram showing the major deployable components

**Discussion questions:**
- What is the boundary between the fire system and the building management system?
- Which requirements are safety-critical and what does that imply for the architecture?
- How do you trace each requirement to a test case?

### Challenge B: Who Is Responsible When the AI Is Wrong?

An AI agent decides not to activate sprinklers because its anomaly detection model classifies a real fire as a false alarm. The building suffers significant damage.

1. Who is responsible — the developer, the building owner, the AI model provider?
2. What design decisions could have prevented this failure?
3. Write 3 requirements that address this scenario
4. How would you test that your system handles this case correctly?

**Discussion questions:**
- Should safety-critical decisions ever be fully autonomous?
- What is the role of human-in-the-loop for life-safety systems?
- How does this change your architecture document?

---

## Lecture 2: CPS Architectures

### Challenge A: Design a Communication Architecture

You have 50 temperature sensors, 20 smoke detectors, 10 door locks, and 5 HVAC units spread across 3 floors. An AI agent needs to read all sensor values and control all actuators. Design the communication architecture.

Before the workshop, propose two alternative designs:
1. **Design 1:** All sensors POST directly to the AI agent via REST
2. **Design 2:** All sensors publish to an MQTT broker, AI agent subscribes

For each design, answer:
- What happens when the AI agent is restarting for 10 seconds?
- What happens when a sensor sends 100 readings per second instead of 1?
- What happens when you add a second AI agent that also needs sensor data?
- What is the maximum latency from sensor event to actuator response?

**Discussion questions:**
- Is there a design that combines the strengths of both?
- Where does the data pipeline fit in?
- What would change if one of the sensors was safety-critical?

### Challenge B: Edge vs. Cloud Decision

Your building control system needs to make two types of decisions:
1. **Fire response:** smoke detected → activate sprinklers (deadline: < 1 second)
2. **Energy optimization:** predict tomorrow's occupancy → adjust HVAC schedule (deadline: before midnight)

For each decision:
1. Where should the computation run (sensor device, edge server, cloud)?
2. What happens if the network to the cloud is down?
3. What ML model would you use and how large is it?
4. What data does it need and where is that data stored?

Draw a deployment diagram showing where each component runs.

**Discussion questions:**
- Can the same architecture handle both decisions? Should it?
- What is the minimum hardware you need at the edge?
- How do you update the ML model at the edge without downtime?

---

## Lecture 3: Data Engineering for CPS

### Challenge A: Design a Sensor Data Pipeline

You have 80 sensors reporting every 5 seconds. That is approximately 1.4 million readings per day. Design a data pipeline that supports:
1. Real-time alerting (smoke above threshold → alert within 2 seconds)
2. Historical queries ("what was the temperature in room A2306 last Tuesday at 14:00?")
3. ML model training ("give me 30 days of hourly temperature averages per room")

Before the workshop:
1. Draw a data flow diagram from sensor to each of the 3 use cases
2. Choose specific technologies for each stage (ingestion, storage, query, processing)
3. Estimate storage requirements for 1 year of data

**Discussion questions:**
- Can you use the same storage for real-time and historical queries? Should you?
- What happens when a sensor is offline for 2 hours — how do you handle the gap?
- How do you detect that a sensor is sending wrong data?

### Challenge B: Feature Engineering Challenge

Given this raw sensor data from one room over 24 hours:
- Temperature: reading every 5 seconds (17,280 values)
- CO2: reading every 10 seconds (8,640 values)
- Door open/close events: 47 events with timestamps
- Occupancy counter: count every minute (1,440 values)

Design features for an ML model that predicts "will this room need ventilation boost in the next 30 minutes?"

1. List at least 8 features you would compute
2. Write the SQL query (or pseudocode) for 3 of them
3. What time window do you need for each feature?
4. How would you store these features for fast retrieval during real-time inference?

**Discussion questions:**
- Which features are most predictive and why?
- How do you handle different sampling rates (5s vs. 10s vs. event-based)?
- Should feature computation happen in the database (SQL) or in Python? Trade-offs?

---

## Lecture 4: AI-Driven Development

### Challenge A: Specification-to-Tests

You receive this requirement:

> *R-HVAC-03: The system shall reduce HVAC energy consumption by at least 15% compared to a fixed-schedule baseline, while maintaining room temperature within 20-24°C during occupied hours (07:00-18:00 weekdays).*

Before the workshop:
1. Write 5 test cases that verify this requirement (unit, integration, and end-to-end)
2. For each test, specify: what is the input, what is the expected output, how do you measure it
3. Which tests can run in under 1 second? Which need a simulation?

Then, write a prompt for an AI coding tool that would generate the HVAC control agent, using your tests as context.

**Discussion questions:**
- How do you test the "15% reduction" — what is the baseline and how do you simulate it?
- Can an AI coding tool generate meaningful ML model code from this specification?
- What parts of this requirement are ambiguous? How do you resolve the ambiguity?

### Challenge B: Debug AI-Generated Code

The following AI-generated sensor process has at least 5 bugs. Find them all.

```python
import requests, time, random

BASE = "http://localhost:9090"
SENSOR_ID = "temp-01"

# Register sensor
requests.post(f"{BASE}/api/equipment", json={
    "id": SENSOR_ID, "name": "Temp", "type": "temperature_sensor",
    "level": "level0", "room": "A2306", "status": "running"
})

requests.post(f"{BASE}/api/equipment/{SENSOR_ID}/sensors", json={
    "id": SENSOR_ID, "name": "Temperature",
    "data_type": "text", "unit": "C", "value": "21.0"
})

temp = 21.0
while True:
    temp += random.random() * 2 - 1
    requests.put(f"{BASE}/api/sensors/{SENSOR_ID}/value",
                 json={"data_type": "text", "value": temp})
    time.sleep(5)
```

Before the workshop:
1. List every bug you can find (functional, resilience, correctness)
2. For each bug, write a test that would catch it
3. Write a corrected version

**Discussion questions:**
- Would an AI code review tool (e.g., "review this code") catch these bugs?
- Which bugs are dangerous in a CPS context vs. just annoying?
- How would you prompt an AI to generate a more robust version?

---

## Lecture 5: Agentic AI for Autonomous Systems

### Challenge A: Design an Agent for Conflicting Objectives

A building has two competing concerns:
- **Energy agent:** wants to turn off HVAC in empty rooms to save energy
- **Air quality agent:** wants to keep ventilation running because CO2 drops slowly and people might return

The current state: room A2306 has been empty for 20 minutes. CO2 is 650 ppm (acceptable but rising slowly). The HVAC is consuming 2.5 kW. Weather forecast says 35°C outside. A meeting is scheduled in this room in 40 minutes.

Before the workshop:
1. Design both agents: what sensors do they read, what actuators do they control, what model do they use?
2. Design a coordination protocol: how do they resolve the conflict?
3. Write the ReAct trace for each agent (Thought → Action → Observation → Thought → ...)
4. What should happen? Justify your answer.

**Discussion questions:**
- Should one agent have priority? When?
- What if the calendar data is wrong (meeting cancelled but not updated)?
- How do you prevent oscillation (energy turns off, comfort turns on, energy turns off, ...)?
- Would a single agent with multiple objectives be simpler? What are the trade-offs?

### Challenge B: Build an MCP Tool Specification

Design an MCP (Model Context Protocol) server that exposes BuildSim as tools for an AI agent. Before the workshop:

1. Define 8 tools with their names, descriptions, and input/output schemas. For example:

```json
{
  "name": "get_room_temperature",
  "description": "Get the current temperature reading for a specific room",
  "input_schema": {
    "type": "object",
    "properties": {
      "room": {"type": "string", "description": "Room name, e.g. A2306"}
    },
    "required": ["room"]
  }
}
```

2. Write a scenario where the AI agent uses at least 4 of these tools in sequence to handle a fire alarm
3. What happens if a tool call fails (BuildSim is down)? Design the error handling.
4. How does the agent know which tools are available?

**Discussion questions:**
- Which tools are read-only (safe to retry) vs. write (dangerous to retry)?
- Should the agent be allowed to call any tool at any time, or should there be constraints?
- How do you test that the MCP server correctly wraps the BuildSim API?
- Could you use the same MCP server with different AI models (Gemma, Claude, GPT)? What changes?
