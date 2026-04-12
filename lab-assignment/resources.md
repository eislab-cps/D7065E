# Technical Resources

## BuildSim Simulator

Full documentation, API reference with curl examples, and example scripts:

- [BuildSim README](../buildingsim/README.md) -- quick start, architecture overview
- [API: Building Data](../buildingsim/docs/api/building.md) -- floor plans, rooms, cross-floor edges
- [API: Equipment](../buildingsim/docs/api/equipment.md) -- equipment, sensor, actuator CRUD
- [API: Sessions](../buildingsim/docs/api/sessions.md) -- viewport, highlights, occupancy, coverage zones, WebSocket
- [API: Navigation Graph](../buildingsim/docs/api/graph.md) -- Dijkstra routing, walkable graph
- [API: Icons](../buildingsim/docs/api/icons.md) -- equipment icon gallery
- [Architecture](../buildingsim/docs/architecture.md) -- system diagrams, data model
- [Example Scripts](../buildingsim/examples/api/) -- ready-to-run bash scripts for equipment, sessions, scenarios

## Agentic AI & Agent Frameworks

Your system must include AI agent(s) that autonomously observe, reason, and act. An agent is a software component that:
1. **Perceives** the environment (reads sensor data from the BuildSim API)
2. **Reasons** about the state (using an LLM, ML model, or planning algorithm)
3. **Acts** on the environment (sets actuator states, triggers alerts via the API)
4. **Learns** from outcomes (updates models, adjusts policies)

You should use an agent framework. Recommended options:

| Framework | Description | Link |
|-----------|-------------|------|
| **LangChain** | Python framework for LLM-powered agents with tool use, chains, and memory | [langchain.com](https://www.langchain.com/) |
| **LangGraph** | State machine-based agent orchestration built on LangChain, supports multi-agent workflows | [langchain-ai.github.io/langgraph](https://langchain-ai.github.io/langgraph/) |
| **Model Context Protocol (MCP)** | Open standard for connecting AI models to external tools and data sources. Build an MCP server that exposes BuildSim as tools. | [modelcontextprotocol.io](https://modelcontextprotocol.io/) |
| **CrewAI** | Multi-agent framework where agents have roles, goals, and can delegate tasks to each other | [crewai.com](https://www.crewai.com/) |
| **AutoGen** | Microsoft's framework for multi-agent conversations and collaboration | [microsoft.github.io/autogen](https://microsoft.github.io/autogen/) |
| **Anthropic Agent SDK** | SDK for building agents with Claude, including tool use and computer control | [docs.anthropic.com](https://docs.anthropic.com/en/docs/agents-and-tools/agent-sdk) |

### MCP Integration

You could build an MCP server that wraps the BuildSim API as tools, allowing any MCP-compatible AI model to directly control the building:

- Tool: `read_sensors` -- reads all sensor values from the building
- Tool: `set_actuator` -- sets an actuator state (lock door, activate sprinkler)
- Tool: `highlight_rooms` -- highlights rooms on the 3D map
- Tool: `find_route` -- computes shortest path between rooms
- Tool: `set_coverage` -- visualizes a coverage/risk zone

This enables any LLM with MCP support to autonomously control the building by calling these tools.

### Agent Architecture Patterns

Consider these patterns for your agent design:

```
ReAct (Reasoning + Acting)
  Agent observes sensors → reasons about state → picks action → observes result → repeats

Plan-and-Execute
  Agent creates a plan (sequence of steps) → executes each step → re-plans if needed

Multi-Agent
  Multiple specialized agents (safety, energy, comfort) → coordination layer resolves conflicts

Human-in-the-Loop
  Agent proposes actions → human approves/rejects → agent learns from feedback
```

## ML/AI Model Resources

| Resource | Use Case | Link |
|----------|----------|------|
| **scikit-learn** | Anomaly detection, clustering, regression | [scikit-learn.org](https://scikit-learn.org/) |
| **PyTorch** | Neural networks, time-series models, GNNs | [pytorch.org](https://pytorch.org/) |
| **TensorFlow/Keras** | Neural networks, time-series forecasting | [tensorflow.org](https://www.tensorflow.org/) |
| **Prophet** | Time-series forecasting (occupancy, temperature) | [facebook.github.io/prophet](https://facebook.github.io/prophet/) |
| **Ollama** | Run local LLMs (Gemma 4 available on lab GPUs) | [ollama.com](https://ollama.com/) |
| **Hugging Face** | Pre-trained models, transformers, datasets | [huggingface.co](https://huggingface.co/) |

## Local AI Model Access

A local Gemma 4 model is available on the lab GPU servers (NVIDIA RTX 5090) via Ollama. You can use it from any language via the Ollama HTTP API, or through LangChain/LangGraph with the Ollama integration.

## Infrastructure & Communication Patterns

| Technology | Description | Link |
|------------|-------------|------|
| **Docker / Docker Compose** | Containerize each sensor/actuator process | [docs.docker.com](https://docs.docker.com/) |
| **MQTT** | Lightweight pub/sub messaging for IoT sensors | [mqtt.org](https://mqtt.org/) |
| **Apache Kafka** | Distributed event streaming for data pipelines | [kafka.apache.org](https://kafka.apache.org/) |
| **Redis** | In-memory data store, pub/sub, time-series | [redis.io](https://redis.io/) |
| **InfluxDB** | Time-series database for sensor data | [influxdata.com](https://www.influxdata.com/) |
| **DuckDB** | Embedded analytical SQL engine, reads Parquet directly | [duckdb.org](https://duckdb.org/) |
| **Grafana** | Dashboard and visualization | [grafana.com](https://grafana.com/) |
| **Eclipse Arrowhead** | SOA framework for industrial IoT | [arrowhead.eu](https://www.arrowhead.eu/) |
| **ColonyOS** | Meta-orchestrator for edge-cloud compute continuums | [github.com/colonyos/colonies](https://github.com/colonyos/colonies) |

## Data Engineering

| Technology | Description | Link |
|------------|-------------|------|
| **Apache Parquet** | Columnar storage format for sensor data | [parquet.apache.org](https://parquet.apache.org/) |
| **MinIO** | S3-compatible object storage for data lake | [min.io](https://min.io/) |
| **dbt** | SQL-based data transformation pipelines | [getdbt.com](https://www.getdbt.com/) |
| **Apache Airflow** | Workflow orchestration for ETL pipelines | [airflow.apache.org](https://airflow.apache.org/) |
| **TimescaleDB** | PostgreSQL extension for time-series | [timescale.com](https://www.timescale.com/) |
| **ClickHouse** | Column-oriented DB for fast analytical queries | [clickhouse.com](https://clickhouse.com/) |
