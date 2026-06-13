# D7065E Course Notes: Embedded Intelligence at the Edge

## Course Plan

| # | Chapter | Topic |
|---|--------|-------|
| 1 | [Introduction & Model-Based Systems Engineering](course-notes1.md) | Course overview, embedded intelligence at the edge, Model-Based Systems Engineering, architecture viewpoints, C4 model |
| 2 | [Edge Intelligence & CPS Architecture](course-notes2.md) | CPS fundamentals, edge-cloud continuum, communication patterns (REST/MQTT), SOA, architectural patterns |
| 3 | [Data Engineering for Cyber-Physical Systems](course-notes3.md) | Sensor data pipelines, ingestion, stream/batch processing, medallion architecture, time-series |
| 4 | [Agentic AI for Autonomous Systems](course-notes4.md) | AI agents, the agent loop, tool use, ReAct, multi-agent coordination, safety guardrails |

Practical, hands-on guides live separately in [`../tutorials`](../tutorials/), and curated external resources are collected in the [reading list](reading-list.md).

Chapters 1-4 provide the foundation before the architecture approval checkpoint (week 3).

## Figures

All figures are draw.io sources in [`diagrams/`](diagrams/), exported to `figures/*.png` with the
[Makefile](Makefile) (`make` — requires the draw.io desktop app). They share the visual style of the
Lecture 1 deck: navy `#032040` background, node fill `#0B3B66`, stroke `#89A5BE`, accent `#FF8247`.
Edit the `.drawio` file, run `make`, and the PNG referenced by the chapter updates.
