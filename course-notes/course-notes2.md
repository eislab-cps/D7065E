# Edge Intelligence & CPS Architecture

Standalone notes for the second chapter of D7065E.

---

## Part 1 — What a Cyber-Physical System Actually Is

<figure class="diagram">
<img src="figures/course-notes2-fig01.png" alt="Anatomy of a CPS">
<figcaption><em>Every CPS has the same anatomy: a physical layer at the top, sensors pushing readings down, a message broker fanning them out, storage and intelligence on the side, actuators pushing decisions back up.</em></figcaption>
</figure>

### The sense–compute–act loop

A cyber-physical system, or CPS for short, is software that reads the physical world, decides something, and changes the physical world. Then it reads the physical world again to see what changed. That cycle repeats forever.

The term itself comes from the embedded-systems research community of the mid-2000s, and the canonical statement of its challenges identifies the defining property of a CPS as the *integration* of computation with physical processes — not computation that merely happens to be near hardware ([Lee, 2008](#lee2008)). The central observation is worth internalising early: mainstream programming abstractions deliberately hide time. A correct Python program is one that eventually produces the right value; how long it takes is considered a performance detail. In a CPS, that abstraction breaks. Timing is part of the *semantics* of the system — a correct answer delivered too late is a wrong answer. This single shift, from "what is computed" to "what is computed *and when*", explains most of what makes the rest of this chapter necessary.

<figure class="diagram">
<img src="figures/course-notes2-fig07.png" alt="The sense–compute–act loop">
<figcaption><em>Sense, compute, act — and the physical world's response becomes the next reading.</em></figcaption>
</figure>

A thermometer plus a heater plus a brain that says "turn on the heater when cold." That is the simplest CPS.

The cycle does not happen once. It happens millions of times a day. A safety-critical loop, like a fire suppression system, may run several times a second. A climate-control loop may run once every few minutes. The right speed depends on how fast the physical world can change.

### What makes a CPS hard

Three things separate CPS engineering from regular software. All three are instances of the same point — the physical world refuses to behave like a software abstraction ([Lee, 2008](#lee2008)):

**The physical world does not wait.** A web service can take 500 milliseconds to respond and nobody minds. A fire detector that takes 500 milliseconds to react to smoke is unacceptable. The clock is the physics of the room, not the developer's patience.

**The physical world has limits.** A fan cannot spin faster than its motor allows. A door cannot close faster than physically possible. A temperature cannot drop by 10 degrees in one second because there is not enough cooling power. Every command must respect those limits.

**The physical world has consequences.** A wrong decision in a web app means a wrong page renders. A wrong decision in a building control system can start a fire, lock people in during an emergency, or freeze a server room. Bugs cost more than money.

A fourth difficulty hides behind the first three: **the physical world is concurrent and continuous**, while software is sequential and discrete. Temperatures drift while the code is still deciding; two doors can move at the same instant even though the program processes their events one at a time. This mismatch — discrete, sequential computation coupled to continuous, parallel physics — is the core intellectual challenge of CPS, and it cannot be solved by simply making computers faster ([Lee, 2008](#lee2008)). It has to be addressed architecturally: by deciding where computation runs, how components exchange information, and how the system degrades when parts of it fail. Those three decisions are precisely the subject of the remaining parts of this chapter.

### A building as a CPS

A modern building has thousands of sensors and actuators. The physical layer includes the rooms themselves, HVAC ducts, temperatures, smoke concentrations, lighting levels, people moving around, and air flowing through vents. The cyber layer includes the software that observes all of that and decides what to do.

The two layers meet at one specific seam: the **API of the building control system**. Sensors report through this API. Actuator commands come back through this API. In production, that API is usually a field bus protocol — BACnet, KNX, or Modbus. In this course, the same role is played by **BuildSim**, a REST and WebSocket API that simulates a real building. The architecture above the API is identical whether the bottom layer is a simulated building or a real one. Only the wiring at the very bottom changes.

---

## Part 2 — Where the Brain Should Live: Edge vs Cloud

<figure class="diagram">
<img src="figures/course-notes2-fig02.png" alt="Edge vs cloud latency comparison">
<figcaption><em>Putting the same control loop in two places gives wildly different latency. For anything safety-critical, the brain belongs on the edge.</em></figcaption>
</figure>

Many newcomers reach for "the cloud" as a default place to run software. For a CPS, that is usually the wrong default. The right place is the **edge** — a computer physically close to the building.

This is not a fashion statement; it is the conclusion of a decade of systems research. Edge computing means performing computation at the edge of the network, downstream from cloud services and upstream from IoT devices, and the cloud-centric model collapses under exactly the workloads a CPS produces ([Shi et al., 2016](#shi2016)). The idea traces back to *cloudlets* — small, trusted, datacentre-grade machines placed one network hop from the devices they serve — and edge computing has been framed as the third major wave of computing after centralised mainframes and the distributed cloud ([Satyanarayanan, 2017](#satyanarayanan2017)). The intellectual point in both lines of work is the same: the speed of light and the capacity of access networks are physical constants, so no amount of cloud engineering can make a remote data centre behave as if it were in the building's server room.

### Why local processing matters

Four reasons drive everything to the edge. The arguments below are not course folklore but the consensus of the field ([Shi et al., 2016](#shi2016); [Satyanarayanan, 2017](#satyanarayanan2017)).

**Latency.** A round trip from a building in Sweden to a Google data centre in Belgium takes roughly 50 to 200 milliseconds when everything works. When the network is congested or the internet is degraded, it can be much worse. A fire suppression system that needs to respond within one second cannot tolerate a 200 ms baseline plus a model that takes another 200 ms plus a return trip. The budget runs out before the system has even started thinking.

**Bandwidth.** Imagine 500 sensors in a building, each producing a 200-byte reading every five seconds. That's about 20 kilobytes per second. Trivial to send to the cloud. Now imagine each sensor is a 1080p camera at 30 frames per second. That's 500 megabytes per second per camera. Sending that to the cloud continuously is physically impossible and expensive even if it were possible. The edge processes, compresses, and filters first.

**Privacy.** Video feeds and occupancy patterns are sensitive personal data. In the European Union, their processing falls under the General Data Protection Regulation, which restricts transfers of personal data to third countries and embeds the principles of data minimisation and purpose limitation directly in law ([GDPR, 2016](#gdpr2016)). Processing on-premises is the architectural expression of those principles: raw video never leaves the building, only derived, non-identifying signals do — "room occupied: yes" instead of a face. Edge processing turns a legal compliance problem into a system design choice, which is almost always the cheaper place to solve it.

**Reliability.** Cloud services go down. Internet connections go down. A building that cannot control its HVAC because Amazon Web Services is having an outage is an embarrassment. The edge keeps running when the wider network does not.

### The compute hierarchy

"Edge" is not a single place. It is a continuum with four tiers, each trading speed against compute power.

| Tier | Where it sits | Latency | Compute power | Typical examples |
|---|---|---|---|---|
| Device | On the sensor or actuator itself | <1 ms | Very low | Arduino, ESP32, Raspberry Pi |
| Edge | Local gateway, building server room | 1–10 ms | Medium | Intel NUC, Jetson, small industrial PC |
| Fog | Building or campus level | 10–50 ms | High | Rack server, mini data centre |
| Cloud | Regional data centre | 50–200 ms | Effectively unlimited | AWS, Azure, Google Cloud |

The terminology comes from two ideas that grew up separately. "Edge computing" was articulated as a research vision building on the earlier cloudlet work ([Shi et al., 2016](#shi2016); [Satyanarayanan, 2017](#satyanarayanan2017)). "Fog computing" was coined at Cisco, from the observation that IoT applications need a *hierarchy* of compute between device and cloud — not a single edge box but a layered continuum that supports mobility, geographic distribution, and location-aware processing ([Bonomi et al., 2012](#bonomi2012)). The fog idea was later formalised by the OpenFog Consortium and adopted as an IEEE standard, IEEE 1934-2018, which defines a reference architecture with explicit pillars such as security, scalability, openness, and hierarchy ([IEEE, 2018](#ieee1934)). For a single building, the distinction between edge and fog is mostly academic. The line gets sharper when one organisation manages many buildings: the per-building gateway is edge, the campus-level aggregation tier is fog, and the fleet-wide analytics platform is cloud.

A natural distribution in production looks like this. Simple threshold alarms and sensor firmware run at the device level. Real-time machine learning inference and agent decision-making run at the edge level. Long-term storage, model training, fleet analytics, and dashboards run in the cloud.

### The hybrid pattern: train in the cloud, run at the edge

The dominant real-world architecture is hybrid. Training a machine-learning model requires lots of data and significant compute time. That work is too slow to do at the edge and benefits from cheap cloud GPUs. Once trained, the model is downloaded to edge servers where it makes predictions in tens of milliseconds. Periodically, fresh data is shipped back to the cloud, models are retrained, and the new model is pushed down to the edge.

<figure class="diagram">
<img src="figures/course-notes2-fig08.png" alt="The hybrid pattern — train in the cloud, run at the edge">
<figcaption><em>The edge owns the real-time loop; the cloud owns storage, training, and analytics. Data flows up, models flow back down.</em></figcaption>
</figure>

Tools like LiteRT (formerly TensorFlow Lite) and ONNX Runtime exist specifically to make this pattern smooth: train in the cloud with full TensorFlow or PyTorch, convert the model to a smaller, faster format, deploy to the edge. The conversion step typically involves quantisation — reducing the numerical precision of the model's weights, for example from 32-bit floats to 8-bit integers — which trades a small amount of accuracy for a large reduction in memory footprint and inference time. That trade-off is exactly the kind a CPS architect must make explicitly: an anomaly detector that is 1 percent less accurate but runs in 20 milliseconds on the edge beats a perfect model that needs a 300-millisecond round trip to the cloud, because in a CPS the deadline is part of correctness ([Lee, 2008](#lee2008)).

### Compute continuum orchestration

The hybrid pattern creates a practical headache. Workloads need to run in different places — a sensor pre-processing job on a Raspberry Pi, a model-training job on a GPU cluster, a dashboard on a cloud VM. Each platform has different capabilities and different APIs. Manually deploying to each is painful.

**Compute continuum orchestrators** abstract this away. The developer submits a workload description that says "this job needs a GPU" or "this job needs to run at the edge." The orchestrator decides which machine fits, dispatches the workload there, and tracks it.

**ColonyOS** is one example, developed in Sweden and used in industrial settings ([Kristiansson, 2024](#kristiansson2024)). It describes itself as a *meta-operating system*: rather than replacing Kubernetes, Slurm, or a cloud provider, it sits above them and treats each as just another place where work can run. A "colony" is the orchestration boundary. **Executors** register the kinds of work they can do — GPU, edge, IoT, CPU. **Function specifications** describe what to compute, written as a small JSON document. A **broker** matches each function specification against available executors. Every message is cryptographically signed, so executors can run on untrusted infrastructure (zero-trust).

For a building control system, this enables a pattern like the following:
- Sensor data preprocessing runs on an edge executor in the building's server room.
- Anomaly model training is submitted as a function specification that the broker dispatches to a GPU executor in the cloud.
- The trained model is automatically deployed back to edge executors.
- If the GPU executor is busy, training queues or overflows to another GPU.

This is the architecture these notes refer to as "the edge–cloud continuum." It scales beyond a single machine without forcing every component into a specific cloud.

### When edge is essential vs when cloud is fine

Use edge processing when:
- Response time must be under one second (fire suppression, emergency unlock).
- The system must keep working without internet (safety-critical functions).
- Continuous data volume is too large to stream (video, high-frequency sensors).
- Data is sensitive and cannot leave the premises.

Use cloud processing when:
- Latency over five seconds is acceptable (dashboard updates, weekly reports).
- Compute requirements exceed local hardware (training a large model).
- Data must be aggregated across many buildings (fleet benchmarking).
- The job is rare or scheduled (disaster recovery, backup).

A common mistake is to over-edge — running everything on a single device because "it's the edge," then discovering that the device cannot run a large language model fast enough. The edge in a real building is not a Raspberry Pi. It is a rack server or a small cluster of industrial computers with serious compute. In a course-scale setup, a developer laptop plays this role.

---

## Part 3 — How Components Talk to Each Other

<figure class="diagram">
<img src="figures/course-notes2-fig03.png" alt="REST vs MQTT communication">
<figcaption><em>REST is a phone call: synchronous, one-to-one. MQTT is a notice board: one publisher, many subscribers, none of them blocking each other.</em></figcaption>
</figure>

Once components are placed at the right tier, the next question is how they communicate. Each communication pattern has different latency, scalability, coupling, and failure characteristics. Choosing the wrong one is one of the most common architectural mistakes.

### Request–response (REST over HTTP)

In this pattern, one component sends a request and waits for the answer. The simplest model and the default for most web APIs.

REST stands for Representational State Transfer. The name comes from Roy Fielding's doctoral dissertation, which is worth knowing about because REST is often misunderstood as "JSON over HTTP" when it is actually an *architectural style*: a named set of constraints — client–server separation, statelessness, cacheability, a uniform interface — chosen because of the system properties they induce, such as scalability, simplicity, and independent evolution of client and server ([Fielding, 2000](#fielding2000)). Fielding derived REST by analysing why the Web itself scaled, and the dissertation remains one of the clearest demonstrations in the literature that architectural constraints are not bureaucracy but the mechanism by which a system acquires its qualities. In practice, REST uses the verbs of HTTP — GET, POST, PUT, DELETE — to express intent. Resources are identified by URLs. Responses carry standard status codes (200 for success, 404 for not found, 500 for server error). The statelessness constraint matters most for a CPS: because every request carries everything the server needs, any replica of the server can answer it, which is what makes REST services easy to restart, scale, and load-balance. The BuildSim API is a REST API: `GET /api/equipment` lists equipment, `POST /api/actuators/{id}/state` commands an actuator.

Request–response works well for control commands (set the actuator's state, then verify the state changed), one-off queries (get the current temperature), and configuration (register a new sensor). It works badly for high-frequency data streams. Polling a REST endpoint at 10 Hz consumes resources for almost nothing, since most polls return the same value. It also works badly for fan-out, where one event needs to reach many subscribers — each subscriber would have to poll independently.

The crucial property of request–response is **tight coupling**: the client must know the server's address, the server must be reachable, and both sides block while the call is in progress. If the server is down, the client fails. There is no buffering, no retry semantics by default, no asynchrony.

### Publish–subscribe (MQTT, Kafka, RabbitMQ)

A different model entirely. Publishers send messages to a **topic** without knowing who is listening. Subscribers register interest in topics and receive all matching messages. A **broker** sits in the middle, routing messages from publishers to subscribers.

<figure class="diagram">
<img src="figures/course-notes2-fig09.png" alt="Publish–subscribe: publishers, topics, subscribers">
<figcaption><em>Publishers write to topics; subscribers read from them. Neither side knows the other exists — the broker does the routing.</em></figcaption>
</figure>

The temperature sensor does not know whether one, ten, or zero consumers are listening. An AI agent can subscribe to a wildcard pattern, `sensors/+/smoke`, and receive all smoke readings from all floors without those sensors knowing the agent exists.

This pattern offers **loose coupling**. Adding a new consumer requires no change to publishers. Subscribers can come and go. The broker buffers messages if a subscriber is briefly unavailable.

**MQTT** is the standard pub/sub protocol for IoT — literally a standard: MQTT 5.0 is ratified by OASIS, the same standards body behind AMQP ([OASIS, 2019](#oasis2019)). It was originally designed in 1999 at IBM for telemetry over unreliable satellite links, which explains its design priorities: tiny message headers, just two bytes in the minimal case, which matters for battery-powered devices and constrained networks. Its quality-of-service semantics deserve precision, because they encode a genuine engineering trade-off. **QoS 0** (at most once) is fire-and-forget: lowest overhead, but a message lost in transit is gone — acceptable for a temperature reading that will be superseded in five seconds anyway. **QoS 1** (at least once) guarantees delivery via acknowledgement and retransmission, but the same message may arrive twice, so consumers must be idempotent — appropriate for an alarm event, where a duplicate is annoying but a loss is dangerous. **QoS 2** (exactly once) uses a four-step handshake to guarantee single delivery, at the highest protocol cost — reserved for messages where a duplicate is as harmful as a loss, such as a billing event. Choosing a QoS level per topic is a small but real architectural decision. Mosquitto is the standard open-source MQTT broker.

**Kafka** is a different beast. It looks like pub/sub from the outside but is internally a distributed, durable, append-only log. Kafka was built at LinkedIn precisely because neither traditional message brokers nor log aggregators could handle both real-time consumption and high-throughput persistence at once; the key insight is to make the *log itself* the first-class abstraction ([Kreps et al., 2011](#kreps2011)). Every message published to a topic is stored at an offset in the log, and subscribers track their own position rather than the broker tracking deliveries. Two consequences follow. Consumers can replay history from any point — a new analytics service deployed today can process last month's sensor data as if it were arriving live. And the broker stays cheap, because it never needs to remember who has read what. This makes Kafka the natural choice for analytics pipelines and event sourcing, at the cost of being heavier to operate than MQTT: Kafka assumes datacentre-class machines and stable networks, the opposite of MQTT's design point.

**RabbitMQ** sits between MQTT and Kafka. It is a general-purpose message broker that supports pub/sub, work queues, request-reply patterns, and complex routing. It speaks AMQP natively but can also speak MQTT through a plugin. Good when MQTT's simplicity is not enough but Kafka's scale is overkill.

Pub/sub is the right fit for sensor data distribution (one publisher, many consumers), system events (the building's fire alarm being triggered), and loose coupling between independently developed components.

### Event-driven architecture

An **event** is an immutable record of something that already happened. Not a command ("turn on the fan") but a fact ("the temperature exceeded 28°C at 14:32:05"). Components that care about that fact react independently.

Event-driven architecture, abbreviated EDA, is the pattern where components communicate only through events. Two key properties fall out of this.

**Temporal decoupling.** A component that reacts to an event does not have to be running at the moment the event is produced. If the event store keeps a history, the consumer can catch up later. Systems become resilient to partial failures.

**Event sourcing.** Storing every event in an append-only log means the entire state of the system can be reconstructed at any point in time. If a bug is discovered, replay the log up to the moment before the bug and analyse it. This is invaluable for debugging and auditability.

For a fire alarm, event-driven architecture is natural. The single event "fire alarm triggered in A2306" causes sprinklers to activate, doors to unlock, HVAC to shut down, and an alert to be sent to the fire department. Each is a separate reaction by a separate component, all triggered by one event.

### WebSocket

WebSocket provides a persistent, two-way TCP connection between a client and server, standardised by the IETF as RFC 6455 ([IETF, 2011](#rfc6455)). Once opened, either side can send data at any time. This is fundamentally different from HTTP, where the client must initiate every exchange. The protocol was designed to coexist with the existing web: the connection begins as an ordinary HTTP request that is then *upgraded* to a WebSocket, so it passes through the same ports and infrastructure as normal web traffic. After the upgrade, the overhead drops to a few bytes per message frame, which is why a single browser tab can comfortably consume hundreds of sensor updates per second.

The BuildSim API uses WebSocket to push live sensor readings and viewer state changes to the browser. The browser opens a WebSocket on page load and receives a continuous stream of updates without polling. Polling a REST endpoint at 10 Hz to achieve the same effect would be 600 requests per minute per client, most of them returning nothing new.

WebSocket is the right fit for real-time server-to-client push (sensor streaming, live dashboards), bidirectional real-time communication, and long-lived connections. It is the wrong fit for occasional queries, where the setup cost of opening the connection outweighs the benefit, and for stateless interactions, where every request is independent.

### gRPC

gRPC is a high-performance request–response protocol developed at Google. It uses HTTP/2 transport and Protocol Buffers for serialisation — a compact binary format in which message schemas are declared in `.proto` files and compiled into typed client and server code. The result is typed APIs with low latency and good multiplexing. The contrast with REST is instructive: where REST trades performance for uniformity and loose evolution (any client that speaks HTTP and JSON can participate), gRPC trades that openness for speed and compile-time type safety, which is a good bargain only when both ends of the connection are built and deployed together. It supports streaming in both directions, which lets it cover some of the WebSocket use cases too.

gRPC fits internal service-to-service communication where performance matters and both ends are controlled by the same team. It is less common in IoT, where MQTT dominates, but appears in modern microservice stacks.

### The pattern comparison table

| Pattern | Coupling | Latency | Scalability | Durability | Building use |
|---|---|---|---|---|---|
| REST | Tight | Low | Medium | No | Control commands, queries |
| MQTT pub/sub | Loose | Low | High | Optional | Sensor distribution |
| RabbitMQ | Loose | Low | High | Yes | Complex routing |
| Kafka | Loose | Medium | Very high | Yes | Multi-building log |
| WebSocket | Medium | Very low | Medium | No | Real-time streams |
| gRPC | Tight | Very low | High | No | Internal services |

The most useful rule of thumb: REST when one component asks another a specific question and waits; pub/sub when one source feeds many consumers; WebSocket when the server needs to push to a client. Kafka enters the picture only when the system needs durable event history at high throughput.

---

## Part 4 — How Components are Organised: Service-Oriented Architecture

<figure class="diagram">
<img src="figures/course-notes2-fig04.png" alt="Service-oriented architecture">
<figcaption><em>Each service is independently deployable, owns one job, and talks through stable interfaces. Replace one without touching the others.</em></figcaption>
</figure>

The patterns above describe **how** components talk. The next question is **how big** each component should be.

### Monolith versus microservices

A **monolithic** architecture puts all functionality into a single deployable unit. One process, one binary, one deploy. Simple to operate when small, painful when large.

A **microservices** architecture decomposes that single binary into small, independently deployable services that communicate over the network. Each service does one thing. Each can be updated, scaled, and operated independently. The term and its modern definition come from a 2014 essay by James Lewis and Martin Fowler: suites of independently deployable services organised around business capabilities, with smart endpoints and dumb pipes, and decentralised ownership of data ([Lewis & Fowler, 2014](#lewis2014)).

The intellectual roots are much older than the name. In 1972, David Parnas asked what criteria should govern how a system is split into modules and gave an answer that still holds: decompose around *design decisions that are likely to change*, hiding each such decision inside one module behind a stable interface ([Parnas, 1972](#parnas1972)). A module's interface should reveal as little as possible about its internals, so that when the internals change, nothing else has to. Microservices are Parnas's information hiding pushed to its logical extreme — the module boundary becomes a process boundary, the interface becomes a network API, and "likely to change independently" becomes "deployable independently". Reading the two papers together is illuminating: Lewis and Fowler's organisational arguments are Parnas's technical arguments, forty-two years later, with the network in between.

For a CPS like a building control system, microservices fit well for four reasons.

Different components have different scaling requirements. The data pipeline needs storage. The AI agent needs CPU or GPU. The dashboard needs nothing. Mixing them in one process means provisioning for the worst-case need everywhere.

Different components evolve at different speeds. The sensor process is stable — once it works, it rarely changes. The agent logic, by contrast, evolves rapidly as new use cases are added. Microservices let the agent be redeployed many times a week without touching the sensor process.

Different components are written by different teams or in different languages. A research team writing the ML model in Python should not be forced to write the high-frequency sensor ingestion in Python too. The sensor ingestion can be in Go for performance; the agent in Python for ML ecosystem; they talk over the network.

Failure isolation. A bug in the dashboard does not take down the safety agent. In a monolith, a memory leak in the dashboard module crashes the whole binary, including safety.

The cost of microservices is operational complexity. Multiple containers must be orchestrated. Network failures between services must be handled gracefully. Deployment configuration multiplies. A call that was an in-process function invocation in the monolith becomes a remote call that can time out, arrive twice, or arrive out of order, and the team inherits the burden of distributed-systems thinking whether it wants it or not ([Lewis & Fowler, 2014](#lewis2014)). The honest trade-off is therefore not "microservices good, monolith bad" but "modularity enforced by the network versus modularity enforced by discipline". A well-structured monolith with clean module boundaries is a perfectly respectable architecture for a small system ([Parnas, 1972](#parnas1972)) — and a far better one than a poorly cut set of microservices, which combines the coupling of a monolith with the failure modes of a distributed system. For a course project, the microservices complexity is instructive rather than a burden. It teaches real-world architecture.

### Containers and Docker

A **container** packages a service together with all its dependencies — libraries, runtime, configuration files — into a single image that runs identically on any machine that has Docker installed. The "works on my laptop" problem disappears, because the laptop and the production server run the same container image.

A **Dockerfile** is the recipe for building a container image. It specifies the base operating system, the libraries to install, the application code to copy in, and the command to run at startup. Running the Dockerfile produces an **image**, a static snapshot. Running the image produces a **container**, a live instance with its own filesystem, network, and process tree.

**Docker Compose** is the tool for running multiple containers together. A `docker-compose.yml` file lists every container, its image, its port mappings, its dependencies on other containers, and its environment variables. Running `docker compose up` starts all of them in the right order. The compose file is essentially the executable version of a C4 container diagram.

### Each sensor and actuator as a service

In a service-oriented building control system, every sensor and every actuator is treated as a service with a well-defined interface. The architecture might include `smoke-sensor-service`, `temperature-sensor-service`, `hvac-actuator-service`, and `door-actuator-service`, each as a separate container. They share the same patterns: each publishes or subscribes on agreed topics; each registers itself with the building's API on startup; each can be killed and restarted without affecting the others.

This has two important consequences. First, the AI agent does not need to know whether the smoke sensor reads from a real device, a BuildSim simulator, a CSV file replay, or a generative ML model. As long as `sensors/smoke` carries the same JSON shape, the agent reads the same way. Second, scaling is per-service. If a building has 500 smoke sensors, you can spin up 500 instances of the same `smoke-sensor-service` image, each with different configuration.

### Eclipse Arrowhead

For industrial CPS, a reference architecture called the **Eclipse Arrowhead Framework** treats every sensor, actuator, and service as a registered, discoverable, and authorised service. Arrowhead was developed largely at Luleå University of Technology — this university — under the leadership of Jerker Delsing, and grew out of one of the largest European industrial automation research projects of its decade ([Delsing, 2017](#delsing2017)). Its central concept is the **local cloud**: a self-contained, protected island of services (one building, one factory cell, one vehicle) that has everything it needs to operate autonomously — its own registry, its own authorisation, its own orchestration — and that interacts with other local clouds only through controlled gateways. The local cloud is the service-oriented mirror of the edge argument from Part 2: just as computation should stay close to the physics for latency and reliability, service governance should stay local so that the building keeps discovering, authorising, and orchestrating its own services even when the wide-area network is gone. Three core systems make this work:

- **Service Registry.** Every service registers itself; consumers look up producers by service type rather than by hardcoded address.
- **Authorisation System.** Enforces which services are allowed to talk to which. A guest sensor cannot pretend to be the building's primary smoke detector.
- **Orchestration System.** Routes requests from consumers to appropriate providers, picking the best available.

Arrowhead is used in Swedish industrial automation and is worth knowing as the production analogue of what we sketch with Docker Compose and MQTT in the course. The mapping is direct: our compose file plays the role of a primitive service registry (services find each other by container name), our shared MQTT topics play the role of orchestrated service exchange, and what we leave out — authorisation — is exactly the part an industrial deployment cannot leave out ([Delsing, 2017](#delsing2017)).

---

## Part 5 — Higher-Level Architectural Patterns

<figure class="diagram">
<img src="figures/course-notes2-fig05.png" alt="Three higher-level architectural patterns">
<figcaption><em>Different shapes optimise for different things: layered for simplicity, microservices for scale, event-driven for loose coupling between components that come and go.</em></figcaption>
</figure>

The previous parts cover individual building blocks. This part describes whole-system patterns that combine them.

### Event-driven pub/sub architecture

The most common pattern for building control. Sensors publish readings to a message broker. Consumers — AI agents, databases, dashboards — subscribe to the topics they care about. The broker handles routing.

Advantages: loose coupling, easy extension (a new consumer is added without changing producers), and a natural fit for the one-to-many relationship between sensors and downstream systems. The default choice for sensor data collection, event distribution, and real-time alerting.

### Lambda architecture

The Lambda architecture, proposed by Nathan Marz, splits data processing into two parallel paths ([Marz & Warren, 2015](#marz2015)). The **speed layer**, also called the hot path, processes events in real time and produces low-latency approximate results — current temperature, active alerts, anomaly scores in the last minute. The **batch layer**, also called the cold path, processes historical data in bulk and produces accurate long-term results — weekly energy reports, monthly trends, training data for the next model. A **serving layer** merges results from both paths when a query comes in.

For building control, the speed layer handles real-time ML inference and safety responses, while the batch layer trains models, generates reports, and populates the feature store for the next training run.

The deeper motivation is often missed: Lambda is built on the principle of an immutable master dataset — raw data is never updated in place, only appended — so that human error can always be corrected by recomputing from the original facts ([Marz & Warren, 2015](#marz2015)). This is the same philosophy as event sourcing from Part 3, applied at analytics scale.

Lambda has been criticised because the same business logic ends up implemented twice — once in the streaming path and once in the batch path — which is hard to keep in sync. The most pointed critique came from Jay Kreps, one of Kafka's creators: maintaining two codebases that must produce identical answers is an operational tax with no compensating benefit once stream processors became capable of reprocessing history ([Kreps, 2014](#kreps2014)). His proposal, the **Kappa architecture**, simplifies Lambda by using a single stream-processing layer for both real-time and historical processing: when logic changes, replay the durable log — which Kafka conveniently provides — through the same code, and switch over when the new output catches up. (Lambda itself was proposed by Nathan Marz, as noted above; the debate between the two architectures is one of the rare cases where the primary sources are short, readable, and genuinely in dialogue with each other.) Apache Flink supports both patterns; its design treats batch processing as a special case of stream processing over a finite stream, which is the Kappa idea made into an engine ([Carbone et al., 2015](#carbone2015)).

### Multi-agent architecture

Instead of one agent that controls everything, several specialised agents each focus on a single objective. The pattern long predates large language models: the multi-agent paradigm was established in the AI literature of the 1990s, which defines an agent by four properties — *autonomy* (it operates without direct intervention), *reactivity* (it perceives and responds to its environment), *pro-activeness* (it pursues goals, not just stimuli), and *social ability* (it interacts with other agents) ([Wooldridge & Jennings, 1995](#wooldridge1995)). A building control agent satisfies all four, which is why the thirty-year-old theory transfers so cleanly: the LLM changes how an agent reasons, not what an agent *is*. A typical decomposition:

- A **Safety Agent** monitors smoke and fire sensors and responds to emergencies. It has the highest priority — it always overrides.
- An **Energy Agent** optimises HVAC schedules to minimise consumption.
- A **Comfort Agent** maintains temperature and air quality within occupant preferences.
- A **Coordination Layer** resolves conflicts. Safety always wins. Energy and comfort negotiate.

Each agent is a separate process. Each can be updated independently. The coordination layer is itself a separate component, which can be implemented as priority-based (the simplest), auction-based (each agent submits a bid), or consensus-based (agents converge through messaging). These coordination mechanisms come straight from the classical literature — negotiation and coordination among self-interested agents is the core of the agents' social ability ([Wooldridge & Jennings, 1995](#wooldridge1995)). Multi-agent is the dominant pattern when the system has competing objectives, and the trade-off it embodies is decomposition again: many simple agents with a coordination problem versus one complicated agent with an internal-consistency problem. Experience in both AI and software architecture suggests the first problem is the easier one to debug, because each agent can be tested against its single objective in isolation.

### Digital twin feedback loop

A **digital twin** is a real-time simulation of the physical system that runs in parallel with the real one. The concept is usually credited to Michael Grieves, who introduced it in the context of product lifecycle management. The twin is defined as three things kept in correspondence — the physical product, its virtual counterpart, and the data connection that ties them together — and its purpose is to surface *emergent* behaviour: the unpredicted and undesirable things a complex system does that no component-level analysis would reveal ([Grieves & Vickers, 2017](#grieves2017)). In manufacturing, the idea has since matured from vision to standard practice, with three phases where twins pay off — design (test before building), operation (monitor and predict), and maintenance (detect degradation before failure) ([Tao et al., 2019](#tao2019)).

For our purposes, the operational use dominates. The idea is to use the twin as a sandbox: before sending a command to a real actuator, send it to the twin and observe the predicted outcome. If the twin says the outcome would be bad, the agent tries a different command. The trade-off a twin introduces is fidelity versus cost: a more accurate physical model gives more trustworthy predictions but is more expensive to build, calibrate, and keep synchronised with the real building — and a twin that drifts out of sync is worse than no twin, because it gives the agent confident wrong answers.

BuildSim is itself a digital twin of a building. The agent can ask it questions like "if I turn HVAC zone 3 on, what will the temperature be in ten minutes?" The agent uses the answer to choose better actions before committing them to the real building. For higher-fidelity simulation, tools like the Building Controls Virtual Test Bed and EnergyPlus exist in research labs.

<figure class="diagram">
<img src="figures/course-notes2-fig10.png" alt="Digital twin feedback loop">
<figcaption><em>Commands are rehearsed against the twin first. Only an outcome that looks good is committed to the real building, whose sensor feedback flows back to the agent.</em></figcaption>
</figure>

### Choosing a pattern

No single pattern fits every use case. A few rules of thumb help:

- A latency requirement below 100 milliseconds points to local processing with direct in-process calls.
- A latency budget below one second is comfortable on the edge with pub/sub.
- A latency budget above one second admits the cloud.
- When many consumers need the same data, prefer pub/sub over request–response.
- When the system must keep working without connectivity, all critical logic must be at the edge with local storage and an offline-first design.
- Simple threshold decisions fit rule engines. Contextual trade-offs fit LLM agents.
- A single objective is served by a single agent. Multiple competing objectives need multi-agent coordination.

---

## Part 6 — A Worked Example

<figure class="diagram">
<img src="figures/course-notes2-fig06.png" alt="A worked HVAC control system">
<figcaption><em>A full HVAC control system, with each component placed and the data flowing top to bottom: physical world → broker → services → back to the world.</em></figcaption>
</figure>

To make all of the above concrete, walk through the architectural choices for a concrete use case. The figure above shows the result of this exercise for an HVAC system; the walkthrough below makes the same decisions for a fire detection system.

The use case: smoke sensors and temperature sensors are deployed in every room. When smoke is detected, the system must activate sprinklers in the affected room within one second, unlock fire doors on the evacuation path, and notify the building manager. False positives must be kept under 5 percent.

**Where does the brain run?** At the edge. The one-second latency budget rules out cloud reasoning on the critical path — this is the edge-computing argument applied with actual numbers ([Shi et al., 2016](#shi2016)). Note also the reliability corollary: the sprinkler decision must not depend on the internet being up, so the entire critical chain from smoke topic to sprinkler command lives on the building's edge server.

**How do components communicate?**
- Sensor readings: pub/sub via MQTT. Many consumers want the same readings (the database, the anomaly detector, the dashboard), and the natural fan-out makes MQTT a clean fit. Topics use a hierarchy: `sensors/level0/A2306/smoke`.
- Actuator commands: REST. One agent sends one command to one actuator and wants to know the result.
- Browser updates: WebSocket. The dashboard needs live updates without polling.

**How are components organised?** As microservices in Docker containers.
- One container per sensor type (`smoke-sensor-service`, `temp-sensor-service`).
- One container per actuator type (`sprinkler-service`, `door-service`).
- One MQTT broker container.
- One time-series database container.
- One anomaly-detection container (the ML model, edge-deployed).
- One safety-agent container (the LLM-driven reasoning, edge-deployed but calling out to the cloud LLM when needed).

**Which higher-level pattern?** Event-driven pub/sub for the data flow, with multi-agent on the decision side. A safety agent reacts to anomaly events. An energy agent monitors HVAC schedules. A coordination layer resolves conflicts. A digital twin (BuildSim) lets the agent test a sprinkler command in simulation before committing.

**Where does training happen?** In the cloud, on a GPU server, scheduled overnight. The trained anomaly model is downloaded to the edge and replaces the previous model on the next agent restart.

**What is the deployment plan?** A `docker-compose.yml` on the building's edge server brings up all containers. A ColonyOS executor running on the same server picks up nightly training jobs from a colony broker. The GPU server runs another executor that takes the heavy training jobs.

That sequence of choices — edge for inference ([Shi et al., 2016](#shi2016)), MQTT for sensor fan-out ([OASIS, 2019](#oasis2019)), REST for commands ([Fielding, 2000](#fielding2000)), microservices for failure isolation ([Lewis & Fowler, 2014](#lewis2014)), multi-agent for competing objectives ([Wooldridge & Jennings, 1995](#wooldridge1995)), digital twin for safety ([Grieves & Vickers, 2017](#grieves2017)), ColonyOS for training orchestration ([Kristiansson, 2024](#kristiansson2024)) — is the architecture these notes train readers to design. Each choice is defensible from the literature, but the more important skill is noticing that each one *was* a choice: a different set of requirements (no safety deadline, one consumer per sensor, a single objective) would have justified a different, simpler architecture.

---

## Part 7 — Vocabulary Reference

Every term used in this chapter, defined.

| Term | Definition |
|---|---|
| **CPS (Cyber-Physical System)** | A system where software directly senses and changes the physical world in a closed loop |
| **Sense-compute-act loop** | The continuous cycle: read sensor, decide action, command actuator, observe effect |
| **Edge computing** | Running software on hardware physically close to where the data is generated, instead of in a remote data centre |
| **Fog computing** | A distributed layer of compute spanning edge and cloud, typically at building or campus scale |
| **Compute continuum** | The unified abstraction across device, edge, fog, and cloud resources |
| **REST** | A request–response API style using HTTP verbs (GET, POST, PUT, DELETE) and JSON over HTTP |
| **MQTT** | A lightweight publish/subscribe protocol designed for IoT, using topics and a broker |
| **Kafka** | A distributed durable event log that looks like pub/sub from the outside but stores messages for replay |
| **RabbitMQ** | A general-purpose message broker that supports many messaging patterns through the AMQP protocol |
| **WebSocket** | A persistent bidirectional connection that allows the server to push data to the client |
| **gRPC** | A high-performance typed RPC framework using HTTP/2 and Protocol Buffers |
| **Event-driven architecture** | A pattern where components communicate only through immutable event records, never direct commands |
| **Event sourcing** | Storing every event in an append-only log, enabling full state reconstruction at any past moment |
| **Monolith** | A single deployable unit containing all of an application's functionality |
| **Microservice** | A small independently deployable service that does one thing and communicates with others over the network |
| **Container** | An isolated runtime that packages a service with all its dependencies; the standard packaging for microservices |
| **Docker** | The most common implementation of containers, with a build format (Dockerfile) and runtime |
| **Docker Compose** | A tool for defining and running multi-container applications, configured by a YAML file |
| **Eclipse Arrowhead** | A service-oriented framework for industrial IoT, with built-in service registry, authorisation, and orchestration |
| **ColonyOS** | A meta-orchestrator that dispatches workloads across edge, GPU, and cloud executors based on resource requirements |
| **Lambda architecture** | A pattern that splits data processing into a real-time hot path and a historical batch cold path |
| **Kappa architecture** | A simplification of Lambda using a single stream-processing layer for both real-time and historical data |
| **Multi-agent system** | An architecture with several specialised AI agents that coordinate through priority, auction, or consensus |
| **Digital twin** | A real-time simulation running in parallel with the physical system, used to predict outcomes of proposed actions |

---

## Part 8 — Summary in Five Sentences

1. A cyber-physical system is software that senses the physical world, decides, acts on it, and observes the result, in a continuous loop fast enough that the physics cannot run away from it.
2. The brain belongs at the edge, close to the sensors, because latency, bandwidth, privacy, and reliability all push against putting it in the cloud.
3. Communication patterns differ in coupling and durability — pick REST for control commands, pub/sub for sensor data, WebSocket for server-push to browsers, Kafka when history matters.
4. Components are small, single-purpose, independently deployable services in containers, communicating over well-defined interfaces, with no shared state.
5. Higher-level patterns combine these building blocks — event-driven pub/sub for data flow, Lambda or Kappa for the speed/batch split, multi-agent for competing objectives, digital twin for safety-critical command verification.

These five ideas, properly internalised, are the foundation for every architectural decision in the rest of the course.

---

## Part 9 — References

### Literature

- <a id="bonomi2012"></a>Bonomi, F., Milito, R., Zhu, J., & Addepalli, S. (2012). Fog Computing and Its Role in the Internet of Things. *Proceedings of the First ACM SIGCOMM Workshop on Mobile Cloud Computing (MCC '12)*.
- <a id="carbone2015"></a>Carbone, P., Katsifodimos, A., Ewen, S., Markl, V., Haridi, S., & Tzoumas, K. (2015). Apache Flink: Stream and Batch Processing in a Single Engine. *IEEE Data Engineering Bulletin*, 38(4).
- <a id="delsing2017"></a>Delsing, J. (Ed.) (2017). *IoT Automation: Arrowhead Framework*. CRC Press.
- <a id="fielding2000"></a>Fielding, R. T. (2000). *Architectural Styles and the Design of Network-based Software Architectures*. PhD dissertation, University of California, Irvine.
- <a id="grieves2017"></a>Grieves, M., & Vickers, J. (2017). Digital Twin: Mitigating Unpredictable, Undesirable Emergent Behavior in Complex Systems. In *Transdisciplinary Perspectives on Complex Systems*. Springer.
- <a id="ieee1934"></a>IEEE (2018). *IEEE 1934-2018 — IEEE Standard for Adoption of OpenFog Reference Architecture for Fog Computing*.
- <a id="rfc6455"></a>IETF (2011). *RFC 6455 — The WebSocket Protocol*.
- <a id="kreps2014"></a>Kreps, J. (2014). Questioning the Lambda Architecture. *O'Reilly Radar*.
- <a id="kreps2011"></a>Kreps, J., Narkhede, N., & Rao, J. (2011). Kafka: a Distributed Messaging System for Log Processing. *Proceedings of the NetDB Workshop*.
- <a id="kristiansson2024"></a>Kristiansson, J. (2024). ColonyOS — A Meta-Operating System for Distributed Computing Across Heterogeneous Platforms. *arXiv:2403.16486*.
- <a id="lee2008"></a>Lee, E. A. (2008). Cyber Physical Systems: Design Challenges. *Proceedings of the 11th IEEE International Symposium on Object-Oriented Real-Time Distributed Computing (ISORC)*.
- <a id="lewis2014"></a>Lewis, J., & Fowler, M. (2014). Microservices. https://martinfowler.com/articles/microservices.html
- <a id="marz2015"></a>Marz, N., & Warren, J. (2015). *Big Data: Principles and Best Practices of Scalable Realtime Data Systems*. Manning.
- <a id="oasis2019"></a>OASIS (2019). *MQTT Version 5.0*. OASIS Standard.
- <a id="parnas1972"></a>Parnas, D. L. (1972). On the Criteria To Be Used in Decomposing Systems into Modules. *Communications of the ACM*, 15(12).
- <a id="gdpr2016"></a>Regulation (EU) 2016/679 of the European Parliament and of the Council (General Data Protection Regulation).
- <a id="satyanarayanan2017"></a>Satyanarayanan, M. (2017). The Emergence of Edge Computing. *IEEE Computer*, 50(1).
- <a id="shi2016"></a>Shi, W., Cao, J., Zhang, Q., Li, Y., & Xu, L. (2016). Edge Computing: Vision and Challenges. *IEEE Internet of Things Journal*, 3(5).
- <a id="tao2019"></a>Tao, F., Zhang, H., Liu, A., & Nee, A. Y. C. (2019). Digital Twin in Industry: State-of-the-Art. *IEEE Transactions on Industrial Informatics*, 15(4).
- <a id="wooldridge1995"></a>Wooldridge, M., & Jennings, N. R. (1995). Intelligent Agents: Theory and Practice. *The Knowledge Engineering Review*, 10(2).

### Software, standards, and online resources

- ColonyOS — https://colonyos.io
- Docker — https://docs.docker.com
- Docker Compose — https://docs.docker.com/compose/
- Eclipse Arrowhead — https://arrowhead.eu
- Eclipse Mosquitto (MQTT broker) — https://mosquitto.org
- gRPC — https://grpc.io
- LiteRT (formerly TensorFlow Lite) — https://ai.google.dev/edge/litert
- MQTT — https://mqtt.org
- ONNX Runtime — https://onnxruntime.ai
- Protocol Buffers — https://protobuf.dev
- RabbitMQ — https://www.rabbitmq.com
