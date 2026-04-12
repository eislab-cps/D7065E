# D7065E — Embedded Intelligence at the Edge

Master's course at Luleå University of Technology, 7.5 ECTS.

[Course syllabus](https://www.ltu.se/en/education/syllabuses/course-syllabus?id=D7065E)

## About the Course

This course teaches how to design and build **autonomous intelligent systems** that sense, reason, and act in physical environments. The focus is on the intersection of edge computing, AI agents, and cyber-physical systems (CPS).

You will learn to:
- Design **CPS architectures** that connect sensors, AI, and actuators into autonomous control loops
- Apply **data engineering** to collect, store, and transform sensor data into ML-ready features using data lakes and SQL pipelines
- Build **agentic AI systems** where AI agents observe sensor data, reason about building state, and take autonomous action through APIs
- Use **modern AI tools** (Claude Code, Cursor, LangChain, MCP) to implement systems from specifications using test-driven generation
- Evaluate **resilience and scalability** — what happens when components fail, sensors lie, or the AI is wrong

The lab assignment centers on **autonomous building control**: you instrument a simulated building with sensors and actuators, develop AI agents that make real-time decisions.

## Specification-First Development

Most developers write code first, then tests, then discover the design was wrong. With AI coding tools this can sometimes gets worse, where AI generates clean-looking code that satisfies the prompt, not the problem.

This course uses **Model-Based Systems Engineering (MBSE)**: you design the system with precise models (component diagrams, sequence diagrams, data flows, requirements) before writing any code. These models become the input to AI coding tools. Tests are derived from the specification, not invented after the fact.

In this case, the architecture artifacts become he prompt, your tests are the contract, and AI is the contractor. A clear specification produces working code. A vague specification results in days of debugging.

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
