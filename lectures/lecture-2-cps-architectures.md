# Edge Intelligence & CPS Architecture

## Cyber-Physical Systems Fundamentals

### The Sensing-Computing-Actuating Loop

A **cyber-physical system (CPS)** is defined by the tight, continuous coupling between computation and physical processes. Unlike a traditional software system, where the physical world is an afterthought, in a CPS the physical world is the primary concern. The computation exists to control the physical process, and the physical process evolves faster than a human can respond.

The core loop is: **sense, compute, act**. Sensors measure physical quantities (temperature, pressure, position, light). Computation processes these measurements, reasons about the current state, and decides on a response. Actuators execute the response, changing the physical state. The changed physical state produces new sensor readings, and the loop repeats. In a building, this loop might run every 5 seconds for routine climate control, or every 100 milliseconds for a fire suppression system.

What makes CPS engineering hard is that the physical world imposes real-time constraints, has physical limits (a fan cannot spin faster than its motor allows), is subject to noise and faults (sensors fail, network packets are dropped), and has safety consequences (a wrong command can start a fire or lock occupants in). These constraints drive every architectural decision.

### Building Control as a CPS

In BuildSim, the physical layer includes rooms at different temperatures, doors with states (open/locked/unlocked), HVAC units that heat and cool air, ventilation fans that move air between zones, smoke sensors with particle concentration readings, CO2 sensors tracking occupancy, and people moving through the building. The cyber layer includes sensor processes reading from the BuildSim API, a data pipeline storing and processing readings, an AI agent reasoning about the building state, and an actuator process sending commands back to BuildSim.

The physical and cyber layers are connected through the BuildSim API, a REST/WebSocket interface that abstracts the simulation. In a real building, this interface would be a field bus (BACnet, KNX, Modbus) connecting physical devices. The architecture is identical; only the bottom layer changes.

```mermaid
%%{init: {"theme": "neutral"}}%%
graph TD
    subgraph Physical["Physical Process (BuildSim)"]
        HVAC[HVAC Units]
        SENSORS[Temperature / Smoke / CO2 Sensors]
        DOORS[Doors & Locks]
        ROOMS[Rooms & Occupants]
    end

    subgraph Cyber["Cyber System"]
        SP[Sensor Process]
        DB[(Time-Series DB)]
        AI[AI Agent]
        ACT[Actuator Process]
    end

    SENSORS -->|readings via WebSocket| SP
    SP -->|store| DB
    DB -->|query| AI
    AI -->|commands| ACT
    ACT -->|REST POST| HVAC
    ACT -->|REST POST| DOORS
    HVAC -->|physical effect| ROOMS
    ROOMS -->|affects| SENSORS
```

## Edge-Cloud Continuum

### Why Edge Computing Matters for CPS

The conventional model of cloud computing (sending all data to a data centre, processing it there, and sending responses back) fails for many CPS use cases. The reasons are fundamental, not incidental.

**Latency.** A fire suppression system must detect fire and activate sprinklers in under one second. A round-trip to a cloud data centre takes 50–200 ms under ideal conditions, and much more under load or network congestion. If the internet connection drops during a fire, the system must still work. Safety-critical decisions cannot depend on a network connection outside the building.

**Bandwidth.** A building with 500 sensors each generating a reading every 5 seconds produces 100 readings per second. If each reading is a 200-byte JSON object, that is 20 KB/s, trivial. But if each sensor is a camera generating 1080p video at 30 fps, that is 500 MB/s per camera, impossible to stream to the cloud continuously. Edge processing compresses, filters, and aggregates data before it ever leaves the building.

**Privacy.** Video feeds and occupancy patterns are sensitive personal data. Many jurisdictions restrict cross-border data transfer under regulations like [GDPR](https://gdpr-info.eu/). Processing data on-premises eliminates the risk of sensitive data leaving the building.

**Reliability.** Cloud services go down. Internet connections go down. A building that cannot control its HVAC because AWS is having an outage is an embarrassing failure. Edge systems keep running when connectivity is lost.

> **Key term — Edge computing:** Processing data close to where it is generated, on the device itself or on local hardware, rather than sending it to a remote data centre. Reduces latency, bandwidth consumption, and dependency on connectivity.

### The Compute Hierarchy

The edge-cloud continuum has several layers, each with different properties:

| Layer | Location | Latency | Compute | Examples |
|-------|----------|---------|---------|---------|
| **Device** | On the sensor/actuator | <1 ms | Very low | Arduino, Raspberry Pi, ESP32 |
| **Edge** | Local gateway, server room | 1–10 ms | Medium | NUC, industrial PC, Jetson |
| **Fog** | Building/campus level | 10–50 ms | High | Rack server, mini data centre |
| **Cloud** | Regional data centre | 50–200 ms | Unlimited | AWS, Azure, GCP |

**Edge computing** refers to processing data near the source of generation rather than in a centralized data centre. The term was popularized by Shi et al. in their 2016 survey ["Edge Computing: Vision and Challenges"](https://doi.org/10.1109/JIOT.2016.2579198) (IEEE IoT Journal). In building control, edge computing typically means running sensor data processing, ML inference, and control logic on a local server in the building's server room, avoiding dependency on external network connectivity for real-time decisions. See also the [Edge Computing article on Wikipedia](https://en.wikipedia.org/wiki/Edge_computing).

**Fog computing** extends edge computing by introducing a hierarchical layer between edge devices and the cloud. The term was coined by [Cisco in 2012](https://en.wikipedia.org/wiki/Fog_computing) and formalized by the [OpenFog Consortium](https://www.iiconsortium.org/), now part of the Industrial Internet Consortium. While edge computing is typically a single device or gateway, fog computing refers to a distributed infrastructure spanning multiple nodes at the building or campus level. In practice, for a single building, the distinction between edge and fog is often academic, but it becomes relevant when managing a campus of buildings with shared compute resources. See Bonomi et al., ["Fog Computing and Its Role in the Internet of Things"](https://doi.org/10.1145/2342509.2342513) (MCC Workshop, 2012).

For building control, the typical distribution places sensor firmware and simple threshold alarms at the device level; data collection, real-time ML inference, and agent decision-making at the edge level; and model training, long-term analytics, fleet management, and dashboards in the cloud.

**Hybrid edge-cloud** is the dominant real-world pattern: run inference at the edge, train models in the cloud. The edge model is updated periodically from the cloud with a freshly trained version. This combines the latency benefits of edge with the computational power of cloud for training. [TensorFlow Lite](https://www.tensorflow.org/lite) and [ONNX Runtime](https://onnxruntime.ai/) are the standard tools for deploying ML models at the edge.

```mermaid
%%{init: {"theme": "neutral"}}%%
graph LR
    subgraph Edge["Edge (Building Server)"]
        INGEST[Sensor Ingestion]
        INFER[ML Inference\n< 50ms]
        AGENT[AI Agent\n< 1s]
        STORE[Local DB\n30-day window]
    end
    subgraph Cloud["Cloud (Data Centre)"]
        TRAIN[Model Training\nhours/days]
        ANALYTICS[Long-term Analytics\nyears of data]
        FLEET[Fleet Management\nmany buildings]
    end
    INFER -->|predictions| AGENT
    STORE -->|historical data, nightly| TRAIN
    TRAIN -->|updated model| INFER
    STORE -->|aggregated metrics| ANALYTICS
```

### Compute Continuum Orchestration

The edge-cloud continuum creates a practical problem: how to manage workloads that span multiple platforms. A sensor process runs on a Raspberry Pi, a data pipeline on an edge server, ML training on a GPU cluster, and a dashboard in the cloud. Each has different capabilities, different APIs, and different failure modes.

**Compute continuum orchestrators** solve this by providing a single abstraction layer across all platforms. Instead of deploying to specific machines, a workload description is submitted specifying what to run and what resources it needs, and the orchestrator places it on the best available platform.

[**ColonyOS**](https://github.com/colonyos/colonies) is an open-source meta-orchestrator designed for this purpose. It creates a "colony", a unified orchestration layer across heterogeneous infrastructure (edge devices, local servers, cloud VMs, HPC clusters). **Executors** register capabilities (GPU, edge, IoT) with the colony. **Function specifications** describe what to compute, not where. The broker matches workloads to executors based on requirements, and zero-trust security ensures executors can run on untrusted infrastructure.

For building control, this enables sensor data preprocessing to run on edge executors (low latency), ML model training to be submitted as a job that runs on the GPU server, and the trained model to be automatically deployed back to edge executors. If the GPU server is busy, training jobs queue or overflow to cloud.

```mermaid
%%{init: {"theme": "neutral"}}%%
graph TB
    subgraph Colony["ColonyOS Broker"]
        BROKER[Workload Scheduler]
    end
    subgraph Edge["Edge Executors"]
        E1[Sensor Processing]
        E2[Real-time Inference]
        E3[AI Agent]
    end
    subgraph GPU["GPU Server Executor"]
        G1[ML Model Training]
    end
    subgraph Cloud["Cloud Executors"]
        C1[Long-term Storage]
        C2[Dashboard]
    end
    E1 & E2 & E3 --> BROKER
    G1 --> BROKER
    C1 & C2 --> BROKER
    BROKER -->|"submit training job"| G1
    BROKER -->|"deploy model"| E2
```

This approach prepares the architecture for real deployment across distributed infrastructure, beyond running everything on a single machine with Docker Compose.

### When Edge Is Essential

Edge is essential when response time is under one second (fire suppression, emergency unlock), when the system must work without internet (safety-critical functions), when continuous data volume is too large to stream (video, high-frequency sensors), or when data is too sensitive to leave the premises.

Cloud is appropriate when latency greater than five seconds is acceptable (dashboard updates, reports), when compute requirements exceed local hardware (training a large model), when data needs to be aggregated across many buildings (fleet benchmarking), or for disaster recovery and backup storage.

A common mistake is to over-edge: running everything on a single device because "it's the edge," then discovering it cannot run an LLM fast enough. The edge layer in a real building is a rack server or a small cluster of industrial computers with significant compute. For this course, a developer laptop serves as the edge.

## Communication Patterns for CPS

Choosing the right communication pattern is one of the most consequential architectural decisions. It affects latency, scalability, coupling between components, and how the system behaves under failure.

### Request/Response (REST / HTTP)

In request/response, a client sends a request and blocks waiting for a response. This is the simplest pattern and the default for most web APIs, including the BuildSim API.

**REST** (Representational State Transfer) uses HTTP verbs (GET, POST, PUT, DELETE) to express intent. Resources are identified by URLs. Responses use standard HTTP status codes (200 OK, 400 Bad Request, 404 Not Found, 500 Internal Server Error). The [BuildSim API](https://github.com/eislab-cps/buildsim) is a REST API, using `GET /sensors` to list sensors and `POST /actuators/{id}` to command an actuator.

Request/response is appropriate for control commands (set actuator state), one-time queries (get current sensor reading), and configuration (register a new sensor), where the caller needs to know the result before proceeding. It should be avoided for high-frequency data streams (polling at 1 Hz is inefficient) or fan-out (one event needs to reach 20 subscribers).

A good REST API tutorial: [RESTful API Design Best Practices](https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/) (Stack Overflow Blog).

### Publish/Subscribe (MQTT / Kafka)

In pub/sub, **publishers** emit messages to **topics** without knowing who receives them. **Subscribers** declare interest in topics and receive all matching messages. A **broker** routes messages between publishers and subscribers.

This decouples producers from consumers entirely. A sensor process publishes to `sensors/floor2/temperature` without knowing whether one, ten, or zero consumers are listening. An AI agent subscribes to `sensors/+/smoke` (MQTT wildcard) and receives all smoke readings from all floors, without the sensor processes knowing the agent exists.

**MQTT** ([mqtt.org](https://mqtt.org/)) is the standard protocol for IoT sensor data. It is lightweight (2-byte header), supports QoS levels (at-most-once, at-least-once, exactly-once), and handles intermittent connectivity gracefully with retained messages and persistent sessions. [Mosquitto](https://mosquitto.org/) is the standard open-source broker. The [HiveMQ MQTT essentials series](https://www.hivemq.com/mqtt-essentials/) is an excellent free resource.

**Kafka** ([kafka.apache.org](https://kafka.apache.org/)) is a distributed log designed for high-throughput, durable event streaming. Unlike MQTT, Kafka stores messages durably, so subscribers can replay historical events. This is valuable for building a data lake (see Lecture 3). Kafka is appropriate at fleet scale, managing many buildings rather than a single one.

**RabbitMQ** ([rabbitmq.com](https://www.rabbitmq.com/)) is a general-purpose message broker that supports multiple messaging patterns: pub/sub, work queues, request/reply, and routing. It implements the AMQP protocol and offers more sophisticated routing than MQTT (exchanges, bindings, dead-letter queues) while being simpler to operate than Kafka. For building control, RabbitMQ is a good middle ground when MQTT is too simple (no message acknowledgement guarantees, no complex routing) and Kafka is too heavy (no need for durable log replay). It also supports [MQTT as a plugin](https://www.rabbitmq.com/docs/mqtt), allowing MQTT-speaking sensors to connect to a RabbitMQ broker alongside AMQP consumers.

```mermaid
%%{init: {"theme": "neutral"}}%%
graph LR
    subgraph Producers
        T1[Temp Sensor\nFloor 1]
        T2[Temp Sensor\nFloor 2]
        S1[Smoke Sensor\nFloor 2]
    end
    subgraph Broker["MQTT Broker (Mosquitto)"]
        TOP1[sensors/floor1/temperature]
        TOP2[sensors/floor2/temperature]
        TOP3[sensors/floor2/smoke]
    end
    subgraph Consumers
        AG[AI Agent]
        DB[Data Pipeline]
        DASH[Dashboard]
    end
    T1 -->|publish| TOP1
    T2 -->|publish| TOP2
    S1 -->|publish| TOP3
    TOP1 -->|subscribe| AG
    TOP2 -->|subscribe| AG
    TOP3 -->|subscribe| AG
    TOP1 -->|subscribe| DB
    TOP2 -->|subscribe| DB
    TOP3 -->|subscribe| DB
    TOP1 -->|subscribe| DASH
```

Pub/sub is appropriate for sensor data distribution (one sensor, many consumers), system events (door opened, fire alarm triggered), and loose coupling between independently developed components.

### Event-Driven Architecture

**Event-driven architecture (EDA)** is a broader pattern where system components communicate exclusively through events, immutable records of something that happened. An event is not a command ("turn on the fan") but a fact ("the temperature exceeded 28°C at 14:32:05"). Components react to events independently and asynchronously.

The key benefit of EDA is **temporal decoupling**: a component that reacts to an event does not need to be running when the event is produced. This makes systems resilient to partial failures. It also enables **event sourcing**: storing every event in an append-only log means the entire state of the system can be reconstructed at any point in time, invaluable for debugging.

In building control, EDA is the natural pattern for safety systems: "fire alarm triggered" is an event that independently causes sprinklers to activate, doors to unlock, HVAC to shut down, and an alert to be sent to the fire department, all four being reactions to one event, each handled by a different component.

A good introduction: [Martin Fowler on Event-Driven Architecture](https://martinfowler.com/articles/201701-event-driven.html).

### WebSocket

WebSocket provides a persistent, full-duplex TCP connection between a client and server. Unlike HTTP, where the client must initiate every exchange, WebSocket allows the server to push data to the client at any time. This is essential for real-time sensor streaming.

The BuildSim API uses WebSocket to stream live sensor readings to subscribed clients. A sensor process connects to the WebSocket endpoint and receives a continuous stream of readings without polling, dramatically more efficient than polling a REST endpoint at 10 Hz.

WebSocket is appropriate for real-time push from server to client (sensor streaming, live dashboards), bidirectional real-time communication, and connections that must stay open for extended periods. It should be avoided for occasional queries where connection setup cost is acceptable and for stateless interactions.

### Pattern Comparison

| Pattern | Coupling | Latency | Scalability | Durability | Use in building control |
|---------|----------|---------|-------------|-----------|------------------------|
| REST | Tight | Low | Medium | No | Control commands, API queries |
| MQTT pub/sub | Loose | Low | High | Optional | Sensor data distribution |
| RabbitMQ | Loose | Low | High | Yes | Complex routing, work queues |
| Kafka | Loose | Medium | Very high | Yes | Multi-building event log |
| WebSocket | Medium | Very low | Medium | No | Real-time sensor streams |
| gRPC | Tight | Very low | High | No | High-performance internal services |

## Service-Oriented Architecture

### Microservices vs. Monolith

A monolithic architecture puts all functionality in a single deployable unit. A microservices architecture decomposes functionality into small, independently deployable services that communicate over a network.

For CPS and building control, the microservices pattern fits well because different components have different scaling requirements (the data pipeline needs more storage than the AI agent needs compute), different components have different update frequencies (the sensor process is stable; the agent logic evolves rapidly), components may be written by different teams or use different languages, and failure isolation means a bug in the dashboard does not take down the safety agent.

The cost of microservices is operational complexity: multiple containers must be orchestrated, service discovery managed, network failures between services handled, and more deployment configuration maintained. For a course project, this complexity is instructive, teaching real-world CPS architecture.

> **Key term — Microservice:** A small, independently deployable service that does one thing well, communicates with other services over a network, and can be scaled, updated, and failed independently.

### Containerisation with Docker

Docker packages a service and all its dependencies (libraries, runtime, configuration) into a container image that runs identically on any machine. This solves the "works on my laptop" problem and is the standard deployment mechanism for microservices.

A **Dockerfile** is a recipe for building a container image, specifying the base OS, installed packages, application code, and startup command. A **container** is a running instance of an image with isolated filesystem, network, and process space. **Docker Compose** is a tool for defining and running multi-container applications; the `docker-compose.yml` is the executable version of a container diagram.

The [official Docker getting-started tutorial](https://docs.docker.com/get-started/) is the best introduction. The [Docker Compose documentation](https://docs.docker.com/compose/) covers multi-container setups.

### Eclipse Arrowhead Framework

[Eclipse Arrowhead](https://www.arrowhead.eu/) is a service-oriented framework specifically designed for industrial IoT and CPS. It treats every sensor, actuator, and service as a registered, discoverable, and authorised service. The core system services are the **Service Registry** (all services register themselves; consumers look up producers by service type), the **Authorisation System** (enforces which services can talk to which), and the **Orchestration System** (routes requests from consumers to appropriate providers).

Arrowhead is used in Swedish industrial automation and is relevant as a reference architecture for how production CPS systems organise service communication. The [Arrowhead Framework documentation](https://github.com/eclipse-arrowhead/core-java-spring) is the primary reference.

### Each Sensor and Actuator as a Service

In a service-oriented building control architecture, every sensor and actuator is a service: it has a well-defined interface, is independently deployable, and can be discovered and used by other services without direct coupling.

Rather than one monolithic process that reads all sensors and commands all actuators, the architecture provides `smoke-sensor-service` (reads smoke sensors, publishes to `sensors/smoke`), `temperature-sensor-service` (reads temperature sensors, publishes to `sensors/temperature`), `hvac-actuator-service` (subscribes to control commands on `actuators/hvac`, sends REST commands to BuildSim), and `door-actuator-service` (subscribes to `actuators/doors`, sends door commands to BuildSim). Each service can be restarted independently, scaled if needed, and replaced with a different implementation. The AI agent does not care whether the smoke sensor is a real device or a BuildSim simulator; it reads only from `sensors/smoke`.

## Architecture Patterns for Building Control

### Event-Driven Pub/Sub Architecture

The most common pattern for building control: sensors publish readings to a message broker, consumers (AI agents, databases, dashboards) subscribe to topics of interest. This is the building block of most modern IoT architectures.

Its advantages are loose coupling, simplicity of extension (a new consumer is added without changing producers), and a natural fit for the one-to-many relationship between sensors and consumers. It is most appropriate for sensor data collection, event distribution, and real-time alerting.

### Lambda Architecture

The Lambda architecture separates data processing into two paths. The **speed layer (hot path)** processes events in real time, producing low-latency approximate results (current temperature, active alerts). The **batch layer (cold path)** processes historical data in bulk, producing accurate long-term results (weekly energy reports, model training data). The **serving layer** merges results from both paths for queries.

For building control, the speed layer handles real-time ML inference and safety responses; the batch layer trains models, generates reports, and populates the feature store for the next training run.

The Lambda architecture has been [critiqued](https://www.oreilly.com/radar/questioning-the-lambda-architecture/) for duplicating logic across two code paths. The **Kappa architecture** (Nathan Marz) simplifies it by using a single stream processing layer for both real-time and historical processing, replaying historical data through the same stream processor. [Apache Flink](https://flink.apache.org/) supports both patterns.

### Multi-Agent Architecture

Multiple specialised agents, each responsible for a single objective, coordinate through a shared state or message passing. An example decomposition places a **Safety Agent** monitoring smoke and fire sensors and responding to emergencies with highest priority, an **Energy Agent** optimising HVAC schedules to minimise energy consumption, a **Comfort Agent** maintaining temperature and CO2 within occupant preference ranges, and a **Coordination Layer** resolving conflicts (Safety Agent always wins; Energy Agent and Comfort Agent negotiate). Each agent is a separate process, independently updatable. See Lecture 4 for a detailed treatment of multi-agent coordination.

### Digital Twin Feedback Loop

A **digital twin** is a real-time simulation of the physical system that runs in parallel with the real system. Before the AI agent sends a command to a real actuator, it sends the command to the digital twin and observes the simulated outcome. If the outcome is undesirable, a different command is tried.

BuildSim is itself a digital twin of a building, simulating the physical dynamics of room temperature, airflow, and occupancy. The AI agent can use BuildSim as an oracle: "if HVAC zone 3 is turned on, what will the temperature be in 10 minutes?" The [Building Controls Virtual Test Bed (BCVTB)](https://simulationresearch.lbl.gov/bcvtb) and [EnergyPlus](https://energyplus.net/) are examples of higher-fidelity building simulation tools used in research.

```mermaid
%%{init: {"theme": "neutral"}}%%
graph LR
    AG[AI Agent]
    DT[Digital Twin\nBuildSim]
    REAL[Real Building\nor BuildSim]
    EVAL{Outcome\nacceptable?}

    AG -->|proposed command| DT
    DT -->|simulated outcome| EVAL
    EVAL -->|yes| REAL
    EVAL -->|no| AG
    REAL -->|sensor feedback| AG
```

### Choosing a Pattern

No single pattern fits all use cases. A latency requirement under 100 ms points to local processing and direct calls; under one second to edge processing and pub/sub; over one second means cloud is acceptable. Where many consumers need the same data, pub/sub is preferred over request/response. When the system must work without connectivity, all critical logic belongs at the edge with local storage and an offline-first design. Simple threshold decisions fit rule engines; contextual trade-offs fit LLM agents. A single objective is served by a single agent; multiple competing objectives require multi-agent coordination.

## Recommended Reading

- Shi, W. et al. "Edge Computing: Vision and Challenges" (IEEE IoT Journal, 2016) — [doi.org/10.1109/JIOT.2016.2579198](https://doi.org/10.1109/JIOT.2016.2579198) — the foundational paper on edge computing
- Dastjerdi, A.V. & Buyya, R. "Fog Computing: Helping the Internet of Things Realize Its Potential" (IEEE Computer, 2016) — [doi.org/10.1109/MC.2016.301](https://doi.org/10.1109/MC.2016.301)
- Eclipse Arrowhead Framework — [arrowhead.eu](https://www.arrowhead.eu/) and [GitHub](https://github.com/eclipse-arrowhead/core-java-spring)
- HiveMQ MQTT Essentials (12-part series, free) — [hivemq.com/mqtt-essentials](https://www.hivemq.com/mqtt-essentials/) — read parts 1–6
- Fowler, M. "Patterns of Enterprise Application Architecture" — [martinfowler.com/eaaCatalog](https://martinfowler.com/eaaCatalog/) — free online reference for architectural patterns
- Fowler, M. "Microservices" article — [martinfowler.com/articles/microservices.html](https://martinfowler.com/articles/microservices.html) — the canonical introduction
- Richardson, C. "Microservices Patterns" — [microservices.io](https://microservices.io/) — free pattern catalog; especially relevant: service registry, API gateway, event-driven patterns
- Docker getting-started tutorial — [docs.docker.com/get-started](https://docs.docker.com/get-started/)
