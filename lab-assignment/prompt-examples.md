# Prompt Examples for AI-Assisted Development

This document shows how to use AI coding tools (Claude Code, Cursor, Copilot) effectively throughout the lab assignment. Each example demonstrates a real prompt and explains why it works.

The key principle: the better the specification, the better the output. Vague prompts produce vague code. Precise prompts with architecture documents, API contracts, and test cases produce working systems.

## Architecture Review

After writing the architecture document, use AI to review it for gaps and inconsistencies before the approval checkpoint.

**Prompt:**
```
Review my architecture document in docs/architecture.md. Check for:
- Missing components: are there any data flows that have no receiving component?
- Interface mismatches: do the JSON schemas match between producers and consumers?
- Single points of failure: what happens if each component crashes?
- Missing requirements coverage: which requirements from docs/requirements.md have no corresponding component?
List specific issues with file references and line numbers.
```

**Prompt:**
```
Look at my C4 container diagram in docs/architecture.md and my docker-compose.yml.
Are they consistent? Does every container in the diagram appear in docker-compose.yml?
Does every network connection in the diagram have a matching port/protocol in docker-compose.yml?
```

**Prompt:**
```
Read my architecture document and my requirements table. For each requirement,
tell me which component is responsible for satisfying it and which test would verify it.
Flag any requirement that has no clear owner or no clear test.
```

## Implementation Planning

Before coding, generate an implementation plan from the approved architecture.

**Prompt:**
```
Based on my architecture in docs/architecture.md, create an implementation plan.
For each component (sensor process, AI agent, data pipeline, actuator process):
1. List the files to create
2. List the dependencies (Python packages, Docker base image)
3. List the API endpoints it calls or exposes
4. Estimate the implementation order (what depends on what)
Start with the component that has the fewest dependencies.
```

**Prompt:**
```
I need to implement the temperature sensor process described in docs/architecture.md.
It should:
- Run as a standalone Docker container
- Register with BuildSim API on startup (POST /api/equipment, POST /api/equipment/{id}/sensors)
- Push temperature readings every 5 seconds (PUT /api/sensors/{id}/value)
- Store readings in InfluxDB
- Handle BuildSim being unavailable (retry with backoff)
- Re-register if BuildSim restarts

Generate the Python code, Dockerfile, and add it to docker-compose.yml.
The BuildSim API documentation is in ../buildingsim/docs/api/equipment.md.
```

## Test Generation

Generate tests from the specification, not from the implementation.

**Prompt:**
```
Read my requirements in docs/requirements.md. For each requirement, generate a pytest
test case. The test should:
- Be named test_{requirement_id}_{description}
- Include a docstring with the requirement text
- Test the actual behavior, not mock everything
- Use the BuildSim API at http://localhost:9090

For example, requirement FR-01 "Detect fire within 30 seconds" should become a test
that sets a smoke sensor value above threshold and verifies the agent responds within 30s.
```

**Prompt:**
```
Write integration tests for my sensor process. The tests should verify:
1. The process registers equipment with BuildSim on startup
2. Sensor values appear in BuildSim after the process starts
3. If BuildSim is restarted, the process re-registers within 60 seconds
4. If the process crashes and restarts, it does not create duplicate sensors

Use pytest with a real BuildSim instance (not mocks). The API docs are in
../buildingsim/docs/api/equipment.md.
```

**Prompt:**
```
Generate fault injection tests for my system. Based on my architecture in docs/architecture.md:
1. Kill each container one at a time and verify the system recovers
2. Send invalid sensor data (negative temperature, NaN) and verify it is rejected
3. Simulate network delay between the AI agent and BuildSim
4. Send contradicting sensor data (smoke high in one sensor, low in adjacent room)
For each test, specify what the expected behavior should be.
```

## AI Agent Development

Use prompts that reference the architecture and the BuildSim API together.

**Prompt:**
```
Implement the safety agent described in docs/architecture.md. It should:
- Use LangChain with the local Gemma model at http://gpu-server:11434
- Define tools that map to BuildSim API endpoints:
  - read_sensors: GET /api/equipment
  - set_actuator: PUT /api/actuators/{id}/state
  - highlight_rooms: PUT /api/sessions/{id}/highlights
  - find_route: GET /api/graph/route?from_name=X&to_name=Y&type=walkable
- Use the ReAct pattern (Thought, Action, Observation)
- Include safety guardrails: always activate sprinklers if smoke > 0.7, regardless of LLM output
- Log every decision with timestamp and reasoning

The BuildSim API docs are in ../buildingsim/docs/api/. The MCP server example
is in ../buildingsim/mcp/server.py.
```

**Prompt:**
```
My AI agent sometimes makes wrong decisions. Add an evaluation framework:
1. Define 10 scenarios (fire in room X, temperature spike, sensor failure)
2. For each scenario, define the expected agent actions
3. Run each scenario against the agent and compare actual vs expected actions
4. Report precision, recall, and response time for each scenario type
```

## Data Pipeline

**Prompt:**
```
Set up the data pipeline described in my architecture. I need:
1. An InfluxDB container with a database called "building"
2. A Python service that subscribes to MQTT topic "sensors/#" and writes to InfluxDB
3. SQL queries (via InfluxQL) for:
   - Average temperature per room over the last hour
   - Maximum CO2 reading per floor today
   - Rooms where temperature changed more than 5 degrees in 10 minutes
4. A Grafana dashboard showing real-time temperature per floor

Add all services to docker-compose.yml. The MQTT broker is already defined as "mosquitto".
```

## Deployment and Docker

**Prompt:**
```
Based on my architecture, generate a complete docker-compose.yml that includes:
- BuildSim server (use the pre-built binary at ../buildingsim/bin/buildsim)
- MQTT broker (Mosquitto)
- 3 sensor processes (temperature, smoke, CO2)
- 1 actuator process (HVAC + doors)
- InfluxDB for sensor storage
- The AI agent
- Grafana for dashboards

Each service should:
- Have a health check
- Restart on failure
- Log to stdout (collected by Docker)
- Use environment variables for configuration (no hardcoded URLs)
```

**Prompt:**
```
My docker-compose.yml works but the services start in the wrong order.
The sensor processes try to connect to BuildSim before it is ready.
Add proper startup ordering with health checks so that:
1. BuildSim starts first and is healthy (responds to GET /api/building)
2. MQTT broker starts and is healthy
3. Then sensor processes, actuator processes, and data pipeline start
4. Finally the AI agent starts (needs all others to be ready)
```

## Code Review

**Prompt:**
```
Review my sensor process code in sensor-process/src/main.py for:
- Race conditions (is it safe to run multiple instances?)
- Error handling (what happens when BuildSim is down?)
- Resource leaks (are HTTP connections closed properly?)
- Security (any hardcoded credentials, SQL injection, command injection?)
- Resilience (does it recover from transient failures?)
```

**Prompt:**
```
Review my AI agent code. Check specifically:
- Does the agent handle LLM timeouts or failures gracefully?
- Can the agent get stuck in an infinite tool-calling loop?
- Are there safety guardrails that cannot be overridden by LLM output?
- Is every agent decision logged for audit?
- What happens if the agent receives contradictory sensor data?
```

## Report Writing

**Prompt:**
```
Based on my test results in docs/test-results.md and my architecture in docs/architecture.md,
help me write the evaluation section of my report. For each architectural decision:
1. What was the decision and why
2. What alternative was considered
3. What the test results show about the decision
4. What would be done differently next time
Focus on trade-offs, not just "it works".
```

## Tips for Effective Prompting

1. **Always reference specific files.** "Read docs/architecture.md" is better than "look at my architecture".

2. **Include the API documentation.** Point the AI to the BuildSim API docs so it generates correct REST calls.

3. **Specify what not to do.** "Do not mock the database, use a real InfluxDB instance" prevents the AI from taking shortcuts.

4. **Ask for one thing at a time.** Generate the sensor process first, test it, then generate the agent. Do not ask for the entire system in one prompt.

5. **Use tests as specifications.** Write the test first, then ask the AI to write code that passes it.

6. **Review everything.** AI-generated code looks correct but may contain subtle bugs. Run the tests. Check the error handling. Verify it actually connects to the right endpoints.
