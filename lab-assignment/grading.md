# Grading Criteria

Grading is based on how well you **connect knowledge** across architecture, data engineering, AI integration, and testing. Each grade level builds on the previous — you must meet all criteria for a grade to receive it.

The oral examination can adjust your grade up or down by one step (see [oral-examination.md](oral-examination.md)).

### Real hardware

You have access to the actual A-Building. If you integrate real IoT devices (Raspberry Pi, ESP32, real sensors) alongside your simulated ones, this can strengthen your grade — but only if your architecture cleanly handles both real and simulated data sources through the same interfaces. The point is not the hardware itself, but demonstrating that your architecture works with real devices without changes to the AI agent or data pipeline.

---

## Grade 3 — Design, implement, and explain a working system

### What you deliver

- **Sensor and actuator processes** running as separate containers, covering the needs of your use case
- **One AI/ML component** that makes decisions based on sensor data (not just thresholds — must use a trained model or LLM reasoning)
- **Architecture document** with C4 diagrams (context + container level) explaining what each component does and why
- **Test plan** with at least unit tests and one end-to-end scenario test, executed with results
- **Dashboard** showing system state (BuildSim session API or separate)
- **Written report** describing the system

### What you must explain

- Draw your architecture on the whiteboard from memory
- Explain what each component does and how they communicate
- **Justify every design choice**: why this communication pattern? Why this data storage? Why this AI approach? What alternatives did you consider and why did you reject them?
- Explain what your AI component does and how it makes decisions
- Show that your tests pass and explain what they verify

### What "working" means

The system must run end-to-end: sensor processes produce data → data reaches the AI → AI makes a decision → actuator state changes → the effect is visible in BuildSim. If any link in this chain doesn't work, the system is not working.

---

## Grade 4 — Evaluate trade-offs and demonstrate quality

*Everything in Grade 3, plus:*

### What you deliver additionally

- **Multiple sensor and actuator types** working together across the building
- **Data pipeline**: sensor data stored in a time-series DB or data lake, queryable with SQL, and used for ML model training or evaluation
- **ML/AI model evaluated with metrics** (accuracy, false alarm rate, latency, etc.) on data from the actual sensor pipeline
- **Fault injection tests**: kill a sensor process during operation, verify the system detects and handles it
- **Integration tests**: verify the full data flow from sensor → storage → AI → actuator

### What you must explain additionally

- Discuss trade-offs between architectural alternatives: why this pattern over others? What are the latency/reliability/scalability implications?
- Show ML/AI evaluation results: what metrics, what data, what do the results mean?
- Demonstrate resilience: crash a sensor process live, show the system recovering
- Connect architecture decisions to non-functional requirements

---

## Grade 5 — Synthesize, critically analyze, and find breaking points

*Everything in Grade 4, plus:*

### What you deliver additionally

- **Multiple AI techniques combined** with justification (e.g., anomaly detection model feeding into an LLM agent for reasoning and action planning)
- **Stress testing with documented breaking points**: where does the system fail and why?
- **Architecture discussion**: how would the design evolve for more sensors, more buildings? (design discussion in report, not necessarily implemented)

### Stress testing requirements

You must systematically find the limits of your system:

- **Failure simulation**: orchestrate your own sensor/actuator/agent processes to fail — kill a sensor container, have a sensor send corrupt data, crash the AI agent mid-decision, simulate a slow or unresponsive actuator
- **Boundary finding**: at what sensor update rate does the system lag? How many simultaneous events before the agent makes bad decisions? What happens when sensors contradict each other?
- **AI adversarial testing**: can the AI be fooled by gradual sensor drift? What's the false positive/negative boundary? How does the agent behave with 50% of sensors offline?
- **Cascading failure**: if one component fails, what else breaks? Does the system oscillate?

### What you must explain additionally

- Present a **breaking point analysis**: not just "it works" but "here's where it fails, why, and what we'd change"
- Critical self-assessment: what are the weaknesses of your architecture?
- Connect to domain knowledge: building regulations, physical constraints, real-world implications
- Demonstrate understanding of the full CPS feedback loop: sensor → data → model → decision → actuation → physical effect → sensor

---

## Summary

| Aspect | Grade 3 | Grade 4 | Grade 5 |
|--------|---------|---------|---------|
| **Scope** | Sensors + actuators for the use case | Multiple types working together | Multi-agent coordination |
| **Architecture** | C4 diagrams, explains choices | Trade-off analysis | Scalability discussion |
| **AI** | One model, explains why | Evaluated with metrics | Multiple techniques combined |
| **Testing** | Unit + e2e scenario | Fault injection + integration | Stress testing + breaking points |
| **Report** | Describes the system | Analyzes trade-offs | Critical analysis of failures |
| **Oral exam** | Draw, explain, and justify architecture | Discuss trade-offs in depth | Defend under critical questioning |
