# Oral Examination

The oral examination is individual, approximately 30 minutes. You will be expected to explain your system, draw diagrams, and answer questions demonstrating understanding of the concepts — not just the implementation.

## Format

1. **Presentation (10 min)** — present your system with a live demo
2. **Whiteboard session (15 min)** — draw and explain diagrams on the whiteboard
3. **Questions (5 min)** — examiner asks follow-up questions

You may not use slides or notes during the whiteboard session.

## What You Must Be Able To Do

### Draw and Explain Your Architecture
- Draw the system architecture on the whiteboard from memory
- Explain each component: what it does, why it exists, how it communicates
- Show the data flow: sensor → storage → AI → decision → actuator → physical effect
- Explain your communication pattern choices (pub/sub, REST, event-driven) and why

### Explain Your AI/Agent Design
- Draw the agent loop: perceive → reason → act → observe outcome
- Explain how your AI component makes decisions
- Describe the tool/API integration (how the agent reads sensors and controls actuators)
- Explain what happens when the AI is wrong or unavailable

### Explain Your Data Pipeline
- Draw how sensor data flows from generation to storage to ML model
- Explain your storage choices (time-series DB, data lake, etc.) and why
- Describe your feature engineering: what features, why, how computed

### Explain Your Testing Strategy
- Describe how you tested the system
- Explain a specific failure scenario and how the system handled it
- For grade 5: explain your breaking point analysis and what you found

### Demonstrate Conceptual Understanding
Be prepared to answer questions such as:
- What is the difference between edge and cloud processing? When would you choose each?
- What is pub/sub and when is it better than request/response?
- What is a data lake and why would you use one for sensor data?
- What is an AI agent and how does it differ from a rule-based system?
- What is MCP and how does it connect AI models to external tools?
- What does resilience mean in a distributed system? Give an example from your system.
- What is the difference between a unit test and an integration test?
- How do you test an AI agent's behavior?
- What happens in your system if a sensor sends wrong data?
- How would your architecture change if you had 10x more sensors?

## Grading During Oral Examination

The oral examination can **adjust your grade up or down by one step** from what your report and implementation suggest.

- If your report shows grade 4 work but you cannot explain the architecture on the whiteboard → grade 3
- If your report shows grade 3 work but you demonstrate deep understanding in the oral → grade 4
- You cannot pass (grade 3) if you cannot draw and explain your own system architecture

## Tips

- Practice drawing your architecture from memory before the exam
- Be honest about what you don't know — it's better than guessing
- If you used AI tools to generate code, you must still understand what the code does
- Focus on **why**, not just **what** — the examiner wants to see reasoning, not memorization
