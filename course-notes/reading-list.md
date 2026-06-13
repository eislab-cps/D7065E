# Reading List

How to use this page: **Required** items are part of the course. **Recommended** items deepen a topic you are actively working on — read them when the project gets you there, not all at once. **Optional** items are for the curious. Each chapter of the [course notes](README.md) also ends with a full reference list of the papers and standards behind the material.

---

## Start here (required)

**Introduction to Distributed Systems — Martin Kleppmann (Cambridge, Lecture 1.1)** · *Required*
A 20-minute introduction to what makes distributed systems hard — the mindset behind every architecture decision in the project.
Video: https://www.youtube.com/playlist?list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB · Notes: https://www.cl.cam.ac.uk/teaching/2122/ConcDisSys/dist-sys-notes.pdf

**Intro to Large Language Models — Andrej Karpathy** · *Required*
The standard one-hour introduction to how modern LLMs work and what they can and cannot do. Watch before designing your autonomous component.
https://www.youtube.com/watch?v=zjkBMFhNj_g

**Building Effective Agents — Anthropic** · *Required*
The best concise introduction to agent design patterns: workflows vs. agents, when to add autonomy, and when not to. Directly applicable to your control agent.
https://www.anthropic.com/engineering/building-effective-agents

---

## Architecture and systems engineering

*Supports chapter 1 of the course notes and the architecture document.*

**The C4 Model — Simon Brown** · *Recommended (before the proposal)*
The lightweight architecture notation used throughout the course. Your architecture document requires C4 context and container diagrams, so read this before drawing them.
https://c4model.com/

**Microservices Resource Guide — Martin Fowler** · *Recommended*
Practical discussion of service decomposition, APIs, and system boundaries — the reasoning behind the "every component is its own process" requirement.
https://martinfowler.com/microservices/

**Site Reliability Engineering — Google (selected chapters)** · *Optional*
How Google thinks about reliability, monitoring, and failure handling. The chapters on monitoring distributed systems and handling overload pair well with grading dimension D.
https://sre.google/books/

---

## Distributed systems

*Supports chapter 2 of the course notes.*

**Distributed Systems lecture series — Martin Kleppmann (Cambridge)** · *Recommended*
The full 8-lecture series behind the required first lecture: RPC, clocks and ordering, replication, consensus. Dip into individual lectures as the project raises the questions.
https://www.youtube.com/playlist?list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB

**Designing Data-Intensive Applications — Martin Kleppmann** · *Recommended*
Widely considered the modern introduction to data systems. The chapters on data models, encoding, messaging, and stream processing map directly onto your pipeline design.
https://dataintensive.net/

---

## Cyber-physical systems and digital twins

*Supports chapters 1–2 of the course notes and the simulator you will build.*

**System-Level Simulation — Wikipedia** · *Optional*
Why CPS engineering validates whole systems in simulation before touching hardware — the rationale for the BuildSim-based project.
https://en.wikipedia.org/wiki/System-level_simulation

**Digital Twin — Wikipedia** · *Optional*
Background on simulation-driven engineering; BuildSim plays exactly this role in your architecture.
https://en.wikipedia.org/wiki/Digital_twin

---

## Data engineering

*Supports chapter 3 of the course notes and the data-pipeline requirement.*

**MQTT Essentials — HiveMQ** · *Recommended (before building the pipeline)*
A short, well-illustrated series on MQTT: topics, QoS levels, retained messages, last will. Read it before choosing your broker setup.
https://www.hivemq.com/mqtt-essentials/

**TimescaleDB Documentation** · *Recommended*
Time-series data modelling and storage — hypertables, continuous aggregates, retention. The pragmatic choice for the project's hot path.
https://docs.timescale.com/

**Event Streaming 101 — Confluent Developer** · *Optional*
A good overview of event-driven architectures and durable logs, for students interested in Kafka-style pipelines beyond the course's MQTT baseline.
https://developer.confluent.io/

---

## Autonomous systems and AI agents

*Supports chapter 4 of the course notes and the autonomous component.*

**ReAct: Synergizing Reasoning and Acting in Language Models — Yao et al. (2023)** · *Recommended*
The paper behind the Thought–Action–Observation pattern used by most agent frameworks — and the reason your agent's reasoning trace doubles as an audit log.
https://arxiv.org/abs/2210.03629

**State of AI Agents — LangChain** · *Optional*
A survey-based snapshot of how agents are actually built and deployed in industry: architectures, tooling, and what goes wrong.
https://www.langchain.com/stateofaiagents

---

## Testing and reliability

*Supports grading dimension D — critical evaluation.*

**Release It! (2nd ed.) — Michael Nygard** · *Recommended (before the evaluation phase)*
The classic on how systems fail in production: cascading failures, stability patterns, fault isolation. A goldmine of ideas for your fault-injection experiments.
https://pragprog.com/titles/mnee2/release-it-second-edition/

**Google Testing Blog** · *Optional*
Practical testing strategies for complex systems, from unit-test hygiene to integration-test design.
https://testing.googleblog.com/
