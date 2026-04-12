# D7065E — Embedded Intelligence at the Edge

Master's course at Luleå University of Technology, 7.5 ECTS.

[Course syllabus](https://www.ltu.se/en/education/syllabuses/course-syllabus?id=D7065E)

## About the Course

This course teaches how to design and build **autonomous intelligent systems** that sense, reason, and act in physical environments. The focus is on the intersection of edge computing, AI agents, and cyber-physical systems (CPS) — where software meets the real world.

You will learn to:
- Design **CPS architectures** that connect sensors, AI, and actuators into autonomous control loops
- Build **agentic AI systems** where AI agents observe sensor data, reason about building state, and take autonomous action through APIs
- Apply **data engineering** to collect, store, and transform sensor data into ML-ready features using data lakes and SQL pipelines
- Use **modern AI tools** (Claude Code, Cursor, LangChain, MCP) to implement systems from specifications using test-driven generation
- Evaluate **resilience and scalability** — what happens when components fail, sensors lie, or the AI is wrong

The lab assignment centers on **autonomous building control**: you instrument a simulated building with sensors and actuators, build AI agents that make real-time decisions, and demonstrate the system working end-to-end — including when things go wrong.

## Contents

- [Lectures](lectures/) — CPS architectures, AI-driven development, data engineering, agentic AI
- [Lab Assignment](lab-assignment/) — Autonomous building control
  - [Use Cases](lab-assignment/README.md#use-case-catalogue) — fire response, HVAC optimization, intrusion detection, predictive maintenance, and more
  - [Grading Criteria](lab-assignment/grading.md) — grades 3/4/5 with stress testing requirements
  - [Oral Examination](lab-assignment/oral-examination.md) — individual whiteboard exam
  - [Resources](lab-assignment/resources.md) — tools, frameworks, and links
- [BuildSim](buildingsim/) — Building simulation server
  - [API Documentation](buildingsim/docs/api/) — REST API with curl examples
  - [Architecture](buildingsim/docs/architecture.md) — system diagrams
  - [Equipment Icons](buildingsim/docs/icons.md) — 48 SVG icons
