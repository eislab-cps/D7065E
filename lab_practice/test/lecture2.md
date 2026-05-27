# Lecture 2 — Edge Intelligence & CPS Architecture

Standalone notes for the second lecture of D7065E. Read this on its own or alongside `lectures/lecture-2-cps-architectures.md`.

---

## Part 1 — What a Cyber-Physical System Actually Is

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 380" role="img" aria-label="Anatomy of a CPS" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">The standard anatomy of a CPS</text>
<!-- Physical layer at top -->
<g>
  <rect x="40" y="60" width="640" height="60" rx="10" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
  <text x="60" y="84" font-size="11" font-weight="700" letter-spacing="2" fill="#8b3a1f">PHYSICAL WORLD</text>
  <text x="60" y="104" font-size="10.5" fill="#6b6660">rooms · ducts · equipment · people</text>
  <g transform="translate(360,90)">
    <circle cx="0" cy="0" r="12" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3.5999999999999996" fill="#8b3a1f"/>
    <path d="M -9.6 -12 Q -14.399999999999999 -16.8 -9.6 -21.599999999999998" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 9.6 -12 Q 14.399999999999999 -16.8 9.6 -21.599999999999998" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g>
  <g transform="translate(420,90)">
    <circle cx="0" cy="0" r="12" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3.5999999999999996" fill="#8b3a1f"/>
    <path d="M -9.6 -12 Q -14.399999999999999 -16.8 -9.6 -21.599999999999998" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 9.6 -12 Q 14.399999999999999 -16.8 9.6 -21.599999999999998" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g>
  <g transform="translate(480,90)">
    <circle cx="0" cy="0" r="12" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3.5999999999999996" fill="#8b3a1f"/>
    <path d="M -9.6 -12 Q -14.399999999999999 -16.8 -9.6 -21.599999999999998" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 9.6 -12 Q 14.399999999999999 -16.8 9.6 -21.599999999999998" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g>
  <!-- actuator icon -->
  <g transform="translate(580,82)">
    <rect x="0" y="0" width="24" height="14" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.3"/>
    <circle cx="6" cy="7" r="3" fill="#8b3a1f"/>
  </g>
</g>
<!-- Sensors down / Actuator up -->
<line x1="380" y1="130" x2="380" y2="158" stroke="#8b3a1f" stroke-width="1.8" stroke-linecap="round"/>
<polygon points="380,165 376,158 384,158" fill="#8b3a1f"/>
<text x="370" y="152" text-anchor="end" font-size="10" fill="#8b3a1f">readings</text>
<line x1="560" y1="165" x2="560" y2="137" stroke="#8b3a1f" stroke-width="1.8" stroke-linecap="round"/>
<polygon points="560,130 564,137 556,137" fill="#8b3a1f"/>
<text x="572" y="152" font-size="10" fill="#8b3a1f">commands</text>
<!-- Broker -->
<g transform="translate(280,170)">
  <rect width="200" height="50" rx="10" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
  <text x="100" y="22" text-anchor="middle" font-size="12" font-weight="700" fill="#2a5a7a">MESSAGE BROKER</text>
  <text x="100" y="38" text-anchor="middle" font-size="10" fill="#6b6660">MQTT · publish / subscribe</text>
</g>
<!-- Storage + Brain side by side -->
<g transform="translate(80,240)">
  <rect width="170" height="100" rx="10" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
  <text x="85" y="24" text-anchor="middle" font-size="11" font-weight="700" fill="#3a5a3a">TIME-SERIES DB</text>
  <g>
    <ellipse cx="85" cy="44" rx="23" ry="6" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <path d="M 62 44 L 62 84 Q 85 92 108 84 L 108 44" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="85" cy="44" rx="23" ry="6" fill="none" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="85" cy="54" rx="21" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
    <ellipse cx="85" cy="66" rx="20" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
  </g>
</g>
<g transform="translate(280,240)">
  <rect width="200" height="100" rx="10" fill="#f0d9d1" stroke="#2a2622" stroke-width="1.5"/>
  <text x="100" y="24" text-anchor="middle" font-size="11" font-weight="700" fill="#8b3a1f">AI AGENT</text>
  <g><circle cx="100" cy="64" r="22" fill="#8b3a1f" opacity="0.9"/>
    <text x="100" y="69" text-anchor="middle" font-size="15.399999999999999" font-weight="700" fill="#fdfbf7">AI</text></g>
  <text x="100" y="92" text-anchor="middle" font-size="10" fill="#6b6660">reason · choose · command</text>
</g>
<g transform="translate(510,240)">
  <rect width="170" height="100" rx="10" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
  <text x="85" y="24" text-anchor="middle" font-size="11" font-weight="700" fill="#7a5a1a">DASHBOARD</text>
  <rect x="20" y="36" width="130" height="50" rx="4" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
  <polyline points="28,76 50,60 70,68 90,46 110,58 130,48 142,54" fill="none" stroke="#8b3a1f" stroke-width="1.5"/>
</g>
<!-- Lines from broker to layer below -->
<line x1="330" y1="220" x2="171.94913647423596" y2="239.15768042736534" stroke="#6b6660" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="165,240 171.46781100415902,235.1867452992305 172.4304619443129,243.12861555550018" fill="#6b6660"/>
<line x1="380" y1="220" x2="380" y2="233" stroke="#6b6660" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="380,240 376,233 384,233" fill="#6b6660"/>
<line x1="430" y1="220" x2="588.050863525764" y2="239.15768042736534" stroke="#6b6660" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="595,240 587.5695380556871,243.12861555550018 588.532188995841,235.1867452992305" fill="#6b6660"/>
</svg>
</div><figcaption>Every CPS has the same anatomy: a physical layer at the top, sensors pushing readings down, a message broker fanning them out, storage and intelligence on the side, actuators pushing decisions back up.</figcaption></figure>

### The sense–compute–act loop

A cyber-physical system, or CPS for short, is software that reads the physical world, decides something, and changes the physical world. Then it reads the physical world again to see what changed. That cycle repeats forever.

```
   ┌──────────────┐         ┌──────────────┐         ┌──────────────┐
   │   SENSE      │  ────►  │   COMPUTE    │  ────►  │     ACT      │
   │              │         │              │         │              │
   │ A sensor     │         │ Software     │         │ An actuator  │
   │ measures     │         │ decides      │         │ changes the  │
   │ something    │         │ what to do   │         │ physical     │
   │ in the world │         │              │         │ world        │
   └──────────────┘         └──────────────┘         └──────┬───────┘
          ▲                                                 │
          │                                                 │
          └───────────  physical world changes  ────────────┘
```

A thermometer plus a heater plus a brain that says "turn on the heater when cold." That is the simplest CPS.

The cycle does not happen once. It happens millions of times a day. A safety-critical loop, like a fire suppression system, may run several times a second. A climate-control loop may run once every few minutes. The right speed depends on how fast the physical world can change.

### What makes a CPS hard

Three things separate CPS engineering from regular software:

**The physical world does not wait.** A web service can take 500 milliseconds to respond and nobody minds. A fire detector that takes 500 milliseconds to react to smoke is unacceptable. The clock is the physics of the room, not the developer's patience.

**The physical world has limits.** A fan cannot spin faster than its motor allows. A door cannot close faster than physically possible. A temperature cannot drop by 10 degrees in one second because there is not enough cooling power. Every command must respect those limits.

**The physical world has consequences.** A wrong decision in a web app means a wrong page renders. A wrong decision in a building control system can start a fire, lock people in during an emergency, or freeze a server room. Bugs cost more than money.

### A building as a CPS

A modern building has thousands of sensors and actuators. The physical layer includes the rooms themselves, HVAC ducts, temperatures, smoke concentrations, lighting levels, people moving around, and air flowing through vents. The cyber layer includes the software that observes all of that and decides what to do.

The two layers meet at one specific seam: the **API of the building control system**. Sensors report through this API. Actuator commands come back through this API. In production, that API is usually a field bus protocol — BACnet, KNX, or Modbus. In this course, the same role is played by **BuildSim**, a REST and WebSocket API that simulates a real building. The architecture above the API is identical whether the bottom layer is a simulated building or a real one. Only the wiring at the very bottom changes.

---

## Part 2 — Where the Brain Should Live: Edge vs Cloud

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 360" role="img" aria-label="Edge vs cloud latency comparison" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Same loop, two architectures</text>
<!-- LEFT: Edge -->
<g transform="translate(20,70)">
  <rect width="320" height="270" rx="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <text x="160" y="28" text-anchor="middle" font-size="12" font-weight="700" letter-spacing="2" fill="#8b3a1f">EDGE</text>
  <g transform="translate(50,80)">
    <circle cx="0" cy="0" r="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="4.199999999999999" fill="#8b3a1f"/>
    <path d="M -11.2 -14 Q -16.799999999999997 -19.599999999999998 -11.2 -25.2" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 11.2 -14 Q 16.799999999999997 -19.599999999999998 11.2 -25.2" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g>
  <text x="50" y="108" text-anchor="middle" font-size="10" fill="#2a2622">sensor</text>
  <line x1="72" y1="80" x2="131" y2="80" stroke="#8b3a1f" stroke-width="1.6" stroke-linecap="round"/>
<polygon points="138,80 131,84 131,76" fill="#8b3a1f"/>
  <rect x="138" y="60" width="56" height="40" rx="6" fill="#8b3a1f"/>
  <text x="166" y="84" text-anchor="middle" font-size="11" font-weight="700" fill="#fdfbf7">AI</text>
  <text x="166" y="108" text-anchor="middle" font-size="10" fill="#2a2622">edge server</text>
  <line x1="196" y1="80" x2="255" y2="80" stroke="#8b3a1f" stroke-width="1.6" stroke-linecap="round"/>
<polygon points="262,80 255,84 255,76" fill="#8b3a1f"/>
  <g transform="translate(260,68)"><rect width="30" height="24" rx="4" fill="#f3ece6" stroke="#2a2622" stroke-width="1.3"/><circle cx="8" cy="12" r="4" fill="#8b3a1f"/></g>
  <text x="276" y="108" text-anchor="middle" font-size="10" fill="#2a2622">actuator</text>
  <text x="160" y="170" text-anchor="middle" font-size="40" font-weight="700" fill="#8b3a1f">5 ms</text>
  <text x="160" y="200" text-anchor="middle" font-size="11" fill="#6b6660">total round-trip</text>
  <line x1="40" y1="220" x2="280" y2="220" stroke="#6b6660" stroke-width="0.8"/>
  <text x="50" y="240" font-size="10" fill="#2a2622">✓ fast</text>
  <text x="50" y="256" font-size="10" fill="#2a2622">✓ works offline</text>
  <text x="170" y="240" font-size="10" fill="#2a2622">✓ private</text>
  <text x="170" y="256" font-size="10" fill="#2a2622">✓ predictable</text>
</g>
<!-- RIGHT: Cloud -->
<g transform="translate(360,70)">
  <rect width="340" height="270" rx="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <text x="170" y="28" text-anchor="middle" font-size="12" font-weight="700" letter-spacing="2" fill="#6b6660">CLOUD</text>
  <g transform="translate(38,80)">
    <circle cx="0" cy="0" r="12" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3.5999999999999996" fill="#8b3a1f"/>
    <path d="M -9.6 -12 Q -14.399999999999999 -16.8 -9.6 -21.599999999999998" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 9.6 -12 Q 14.399999999999999 -16.8 9.6 -21.599999999999998" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g>
  <path d="M 56 80 Q 90 60 124 80" fill="none" stroke="#6b6660" stroke-width="1.4" stroke-dasharray="4 3"/>
  <rect x="124" y="68" width="40" height="24" rx="4" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
  <text x="144" y="84" text-anchor="middle" font-size="8" fill="#6b6660">router</text>
  <path d="M 164 80 Q 200 50 240 80" fill="none" stroke="#6b6660" stroke-width="1.4" stroke-dasharray="4 3"/>
  <g transform="translate(232.5,48.5) scale(1.1)">
    <path d="M 20 50 Q 0 50 0 35 Q 0 22 14 20 Q 16 6 32 6 Q 50 0 60 14 Q 78 14 80 30 Q 92 32 90 46 Q 88 60 70 60 L 28 60 Q 14 60 20 50 Z" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
  </g>
  <path d="M 264 90 Q 200 130 164 100" fill="none" stroke="#6b6660" stroke-width="1.4" stroke-dasharray="4 3"/>
  <path d="M 124 100 Q 90 130 56 100" fill="none" stroke="#6b6660" stroke-width="1.4" stroke-dasharray="4 3"/>
  <g transform="translate(28,116)"><rect width="30" height="24" rx="4" fill="#f3ece6" stroke="#2a2622" stroke-width="1.3"/><circle cx="8" cy="12" r="4" fill="#8b3a1f"/></g>
  <text x="170" y="170" text-anchor="middle" font-size="40" font-weight="700" fill="#6b6660">200+ ms</text>
  <text x="170" y="200" text-anchor="middle" font-size="11" fill="#6b6660">round-trip across the internet</text>
  <line x1="40" y1="220" x2="300" y2="220" stroke="#6b6660" stroke-width="0.8"/>
  <text x="50" y="240" font-size="10" fill="#2a2622">✗ slow</text>
  <text x="50" y="256" font-size="10" fill="#2a2622">✗ breaks offline</text>
  <text x="180" y="240" font-size="10" fill="#2a2622">✗ data leaves</text>
  <text x="180" y="256" font-size="10" fill="#2a2622">✓ scales easily</text>
</g>
</svg>
</div><figcaption>Putting the same control loop in two places gives wildly different latency. For anything safety-critical, the brain belongs on the edge.</figcaption></figure>

Many newcomers reach for "the cloud" as a default place to run software. For a CPS, that is usually the wrong default. The right place is the **edge** — a computer physically close to the building.

### Why local processing matters

Four reasons drive everything to the edge.

**Latency.** A round trip from a building in Sweden to a Google data centre in Belgium takes roughly 50 to 200 milliseconds when everything works. When the network is congested or the internet is degraded, it can be much worse. A fire suppression system that needs to respond within one second cannot tolerate a 200 ms baseline plus a model that takes another 200 ms plus a return trip. The budget runs out before the system has even started thinking.

**Bandwidth.** Imagine 500 sensors in a building, each producing a 200-byte reading every five seconds. That's about 20 kilobytes per second. Trivial to send to the cloud. Now imagine each sensor is a 1080p camera at 30 frames per second. That's 500 megabytes per second per camera. Sending that to the cloud continuously is physically impossible and expensive even if it were possible. The edge processes, compresses, and filters first.

**Privacy.** Video feeds and occupancy patterns are sensitive personal data. Cross-border data transfers are restricted by laws like the European GDPR. Processing on-premises eliminates the risk that someone's movement patterns leave the country.

**Reliability.** Cloud services go down. Internet connections go down. A building that cannot control its HVAC because Amazon Web Services is having an outage is an embarrassment. The edge keeps running when the wider network does not.

### The compute hierarchy

"Edge" is not a single place. It is a continuum with four tiers, each trading speed against compute power.

| Tier | Where it sits | Latency | Compute power | Typical examples |
|---|---|---|---|---|
| Device | On the sensor or actuator itself | <1 ms | Very low | Arduino, ESP32, Raspberry Pi |
| Edge | Local gateway, building server room | 1–10 ms | Medium | Intel NUC, Jetson, small industrial PC |
| Fog | Building or campus level | 10–50 ms | High | Rack server, mini data centre |
| Cloud | Regional data centre | 50–200 ms | Effectively unlimited | AWS, Azure, Google Cloud |

The terminology comes from two ideas that grew up separately. "Edge computing" was named by Shi et al. in their 2016 paper for IEEE IoT Journal. "Fog computing" was named by Cisco in 2012 and formalised by the OpenFog Consortium. For a single building, the distinction between edge and fog is mostly academic. The line gets sharper when one organisation manages many buildings.

A natural distribution in production looks like this. Simple threshold alarms and sensor firmware run at the device level. Real-time machine learning inference and agent decision-making run at the edge level. Long-term storage, model training, fleet analytics, and dashboards run in the cloud.

### The hybrid pattern: train in the cloud, run at the edge

The dominant real-world architecture is hybrid. Training a machine-learning model requires lots of data and significant compute time. That work is too slow to do at the edge and benefits from cheap cloud GPUs. Once trained, the model is downloaded to edge servers where it makes predictions in tens of milliseconds. Periodically, fresh data is shipped back to the cloud, models are retrained, and the new model is pushed down to the edge.

```
   EDGE                                       CLOUD
   ────                                       ─────
   ingest sensor readings              ────►  long-term storage
   real-time ML inference (<50ms)             model training (hours)
   AI agent decisions (<1s)                   fleet analytics
   local store (30-day window)                dashboards
                                       ◄────  updated model

```

Tools like TensorFlow Lite and ONNX Runtime exist specifically to make this pattern smooth: train in the cloud with full TensorFlow or PyTorch, convert the model to a smaller, faster format, deploy to the edge.

### Compute continuum orchestration

The hybrid pattern creates a practical headache. Workloads need to run in different places — a sensor pre-processing job on a Raspberry Pi, a model-training job on a GPU cluster, a dashboard on a cloud VM. Each platform has different capabilities and different APIs. Manually deploying to each is painful.

**Compute continuum orchestrators** abstract this away. The developer submits a workload description that says "this job needs a GPU" or "this job needs to run at the edge." The orchestrator decides which machine fits, dispatches the workload there, and tracks it.

**ColonyOS** is one example, developed in Sweden and used in industrial settings. A "colony" is the orchestration boundary. **Executors** register the kinds of work they can do — GPU, edge, IoT, CPU. **Function specifications** describe what to compute, written as a small JSON document. A **broker** matches each function specification against available executors. Every message is cryptographically signed, so executors can run on untrusted infrastructure (zero-trust).

For a building control system, this enables a pattern like the following:
- Sensor data preprocessing runs on an edge executor in the building's server room.
- Anomaly model training is submitted as a function specification that the broker dispatches to a GPU executor in the cloud.
- The trained model is automatically deployed back to edge executors.
- If the GPU executor is busy, training queues or overflows to another GPU.

This is the architecture the lecture refers to as "the edge–cloud continuum." It scales beyond a single machine without forcing every component into a specific cloud.

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

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 320" role="img" aria-label="REST vs MQTT communication" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Two ways for components to talk</text>
<!-- LEFT: REST -->
<g transform="translate(40,70)">
  <rect width="290" height="220" rx="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <text x="145" y="28" text-anchor="middle" font-size="12" font-weight="700" letter-spacing="2" fill="#8b3a1f">REST · request / response</text>
  <rect x="30" y="58" width="80" height="50" rx="6" fill="#f3ece6" stroke="#2a2622" stroke-width="1.4"/>
  <text x="70" y="88" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Client</text>
  <rect x="180" y="58" width="80" height="50" rx="6" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
  <text x="220" y="88" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Server</text>
  <line x1="112" y1="76" x2="171" y2="76" stroke="#8b3a1f" stroke-width="1.6" stroke-linecap="round"/>
<polygon points="178,76 171,80 171,72" fill="#8b3a1f"/>
  <text x="145" y="68" text-anchor="middle" font-size="10" fill="#8b3a1f">GET /temp</text>
  <line x1="178" y1="98" x2="119" y2="98" stroke="#8b3a1f" stroke-width="1.6" stroke-linecap="round"/>
<polygon points="112,98 119,94 119,102" fill="#8b3a1f"/>
  <text x="145" y="116" text-anchor="middle" font-size="10" fill="#8b3a1f">{ value: 22.4 }</text>
  <text x="145" y="160" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">"Like a phone call."</text>
  <text x="145" y="186" text-anchor="middle" font-size="10" fill="#2a2622">synchronous · 1-to-1 · request must wait</text>
</g>
<!-- RIGHT: MQTT -->
<g transform="translate(360,70)">
  <rect width="320" height="220" rx="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <text x="160" y="28" text-anchor="middle" font-size="12" font-weight="700" letter-spacing="2" fill="#2a5a7a">MQTT · publish / subscribe</text>
  <!-- publisher -->
  <rect x="20" y="58" width="68" height="40" rx="6" fill="#f3ece6" stroke="#2a2622" stroke-width="1.4"/>
  <text x="54" y="82" text-anchor="middle" font-size="10.5" font-weight="600" fill="#2a2622">Sensor</text>
  <line x1="90" y1="78" x2="125" y2="78" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="132,78 125,82 125,74" fill="#8b3a1f"/>
  <!-- broker -->
  <rect x="132" y="58" width="60" height="100" rx="8" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
  <text x="162" y="80" text-anchor="middle" font-size="10" font-weight="700" fill="#2a5a7a">BROKER</text>
  <text x="162" y="100" text-anchor="middle" font-size="9" fill="#6b6660">topic:</text>
  <text x="162" y="114" text-anchor="middle" font-size="9" font-family="monospace" fill="#2a2622">sensors/#</text>
  <!-- subscribers -->
  <line x1="192" y1="68" x2="235.13593527016357" y2="59.37281294596729" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="242,58 235.9203998107163,63.29513564873097 234.35147072961084,55.45049024320361" fill="#8b3a1f"/>
  <line x1="192" y1="108" x2="235" y2="108" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="242,108 235,112 235,104" fill="#8b3a1f"/>
  <line x1="192" y1="148" x2="235.13593527016357" y2="156.6271870540327" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="242,158 234.35147072961084,160.5495097567964 235.9203998107163,152.704864351269" fill="#8b3a1f"/>
  <rect x="242" y="44" width="64" height="28" rx="4" fill="#e2ebde" stroke="#2a2622" stroke-width="1.2"/>
  <text x="274" y="62" text-anchor="middle" font-size="10" fill="#2a2622">DB writer</text>
  <rect x="242" y="94" width="64" height="28" rx="4" fill="#f4ead9" stroke="#2a2622" stroke-width="1.2"/>
  <text x="274" y="112" text-anchor="middle" font-size="10" fill="#2a2622">Detector</text>
  <rect x="242" y="144" width="64" height="28" rx="4" fill="#f0d9d1" stroke="#2a2622" stroke-width="1.2"/>
  <text x="274" y="162" text-anchor="middle" font-size="10" fill="#2a2622">Dashboard</text>
  <text x="160" y="194" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">"Like a notice board."</text>
</g>
</svg>
</div><figcaption>REST is a phone call: synchronous, one-to-one. MQTT is a notice board: one publisher, many subscribers, none of them blocking each other.</figcaption></figure>

Once components are placed at the right tier, the next question is how they communicate. Each communication pattern has different latency, scalability, coupling, and failure characteristics. Choosing the wrong one is one of the most common architectural mistakes.

### Request–response (REST over HTTP)

In this pattern, one component sends a request and waits for the answer. The simplest model and the default for most web APIs.

REST stands for Representational State Transfer. It uses the verbs of HTTP — GET, POST, PUT, DELETE — to express intent. Resources are identified by URLs. Responses carry standard status codes (200 for success, 404 for not found, 500 for server error). The BuildSim API is a REST API: `GET /api/equipment` lists equipment, `POST /api/actuators/{id}/state` commands an actuator.

Request–response works well for control commands (set the actuator's state, then verify the state changed), one-off queries (get the current temperature), and configuration (register a new sensor). It works badly for high-frequency data streams. Polling a REST endpoint at 10 Hz consumes resources for almost nothing, since most polls return the same value. It also works badly for fan-out, where one event needs to reach many subscribers — each subscriber would have to poll independently.

The crucial property of request–response is **tight coupling**: the client must know the server's address, the server must be reachable, and both sides block while the call is in progress. If the server is down, the client fails. There is no buffering, no retry semantics by default, no asynchrony.

### Publish–subscribe (MQTT, Kafka, RabbitMQ)

A different model entirely. Publishers send messages to a **topic** without knowing who is listening. Subscribers register interest in topics and receive all matching messages. A **broker** sits in the middle, routing messages from publishers to subscribers.

```
   publishers                broker                 subscribers
   ──────────                ──────                 ───────────
   temp sensor  ──►   topic "sensors/floor1/temp"  ──►  AI agent
                                                  ──►  database
                                                  ──►  dashboard
   smoke sensor ──►   topic "sensors/floor1/smoke" ──►  AI agent
                                                  ──►  database
```

The temperature sensor does not know whether one, ten, or zero consumers are listening. An AI agent can subscribe to a wildcard pattern, `sensors/+/smoke`, and receive all smoke readings from all floors without those sensors knowing the agent exists.

This pattern offers **loose coupling**. Adding a new consumer requires no change to publishers. Subscribers can come and go. The broker buffers messages if a subscriber is briefly unavailable.

**MQTT** is the standard pub/sub protocol for IoT. It uses tiny message headers, just two bytes, which matters for battery-powered devices and constrained networks. It supports three quality-of-service levels: at-most-once delivery, at-least-once delivery, and exactly-once delivery. Mosquitto is the standard open-source MQTT broker.

**Kafka** is a different beast. It looks like pub/sub from the outside but is internally a distributed, durable, append-only log. Every message published to a topic is stored, and subscribers can replay history from any point. This makes Kafka the natural choice for analytics pipelines and event sourcing, at the cost of being heavier to operate than MQTT.

**RabbitMQ** sits between MQTT and Kafka. It is a general-purpose message broker that supports pub/sub, work queues, request-reply patterns, and complex routing. It speaks AMQP natively but can also speak MQTT through a plugin. Good when MQTT's simplicity is not enough but Kafka's scale is overkill.

Pub/sub is the right fit for sensor data distribution (one publisher, many consumers), system events (the building's fire alarm being triggered), and loose coupling between independently developed components.

### Event-driven architecture

An **event** is an immutable record of something that already happened. Not a command ("turn on the fan") but a fact ("the temperature exceeded 28°C at 14:32:05"). Components that care about that fact react independently.

Event-driven architecture, abbreviated EDA, is the pattern where components communicate only through events. Two key properties fall out of this.

**Temporal decoupling.** A component that reacts to an event does not have to be running at the moment the event is produced. If the event store keeps a history, the consumer can catch up later. Systems become resilient to partial failures.

**Event sourcing.** Storing every event in an append-only log means the entire state of the system can be reconstructed at any point in time. If a bug is discovered, replay the log up to the moment before the bug and analyse it. This is invaluable for debugging and auditability.

For a fire alarm, event-driven architecture is natural. The single event "fire alarm triggered in A2306" causes sprinklers to activate, doors to unlock, HVAC to shut down, and an alert to be sent to the fire department. Each is a separate reaction by a separate component, all triggered by one event.

### WebSocket

WebSocket provides a persistent, two-way TCP connection between a client and server. Once opened, either side can send data at any time. This is fundamentally different from HTTP, where the client must initiate every exchange.

The BuildSim API uses WebSocket to push live sensor readings and viewer state changes to the browser. The browser opens a WebSocket on page load and receives a continuous stream of updates without polling. Polling a REST endpoint at 10 Hz to achieve the same effect would be 600 requests per minute per client, most of them returning nothing new.

WebSocket is the right fit for real-time server-to-client push (sensor streaming, live dashboards), bidirectional real-time communication, and long-lived connections. It is the wrong fit for occasional queries, where the setup cost of opening the connection outweighs the benefit, and for stateless interactions, where every request is independent.

### gRPC

gRPC is a high-performance request–response protocol developed at Google. It uses HTTP/2 transport and Protocol Buffers for serialisation. The result is typed APIs with low latency and good multiplexing. It supports streaming in both directions, which lets it cover some of the WebSocket use cases too.

gRPC fits internal service-to-service communication where performance matters and both ends are controlled by the same team. It is less common in IoT, where MQTT dominates, but appears in modern microservice stacks.

### The pattern comparison table

```
Pattern          Coupling   Latency      Scalability   Durability   Building use
────────────     ────────   ──────────   ───────────   ──────────   ──────────────────────
REST             Tight      Low          Medium        No           Control commands, queries
MQTT pub/sub     Loose      Low          High          Optional     Sensor distribution
RabbitMQ         Loose      Low          High          Yes          Complex routing
Kafka            Loose      Medium       Very high     Yes          Multi-building log
WebSocket        Medium     Very low     Medium        No           Real-time streams
gRPC             Tight      Very low     High          No           Internal services
```

The most useful rule of thumb: REST when one component asks another a specific question and waits; pub/sub when one source feeds many consumers; WebSocket when the server needs to push to a client. Kafka enters the picture only when the system needs durable event history at high throughput.

---

## Part 4 — How Components are Organised: Service-Oriented Architecture

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 340" role="img" aria-label="Service-oriented architecture" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Each service owns one job</text>
  <g transform="translate(40,80)">
    <rect width="120" height="80" rx="12" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <text x="60" y="36" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Sensor service</text>
    <text x="60" y="54" text-anchor="middle" font-size="10" fill="#6b6660">reads BuildSim</text>
    <!-- container hint -->
    <text x="60" y="70" text-anchor="middle" font-size="9" font-family="monospace" fill="#6b6660">docker</text>
  </g>
  <g transform="translate(220,80)">
    <rect width="120" height="80" rx="12" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
    <text x="60" y="36" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Ingestion service</text>
    <text x="60" y="54" text-anchor="middle" font-size="10" fill="#6b6660">validates · stores</text>
    <!-- container hint -->
    <text x="60" y="70" text-anchor="middle" font-size="9" font-family="monospace" fill="#6b6660">docker</text>
  </g>
  <g transform="translate(400,80)">
    <rect width="120" height="80" rx="12" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <text x="60" y="36" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Detection service</text>
    <text x="60" y="54" text-anchor="middle" font-size="10" fill="#6b6660">anomaly model</text>
    <!-- container hint -->
    <text x="60" y="70" text-anchor="middle" font-size="9" font-family="monospace" fill="#6b6660">docker</text>
  </g>
  <g transform="translate(580,80)">
    <rect width="120" height="80" rx="12" fill="#f0d9d1" stroke="#2a2622" stroke-width="1.5"/>
    <text x="60" y="36" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Safety service</text>
    <text x="60" y="54" text-anchor="middle" font-size="10" fill="#6b6660">agent · decides</text>
    <!-- container hint -->
    <text x="60" y="70" text-anchor="middle" font-size="9" font-family="monospace" fill="#6b6660">docker</text>
  </g>
  <g transform="translate(40,230)">
    <rect width="120" height="80" rx="12" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
    <text x="60" y="36" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Storage</text>
    <text x="60" y="54" text-anchor="middle" font-size="10" fill="#6b6660">time-series DB</text>
    <!-- container hint -->
    <text x="60" y="70" text-anchor="middle" font-size="9" font-family="monospace" fill="#6b6660">docker</text>
  </g>
  <g transform="translate(220,230)">
    <rect width="120" height="80" rx="12" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <text x="60" y="36" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Dashboard</text>
    <text x="60" y="54" text-anchor="middle" font-size="10" fill="#6b6660">UI · operators</text>
    <!-- container hint -->
    <text x="60" y="70" text-anchor="middle" font-size="9" font-family="monospace" fill="#6b6660">docker</text>
  </g>
  <g transform="translate(400,230)">
    <rect width="120" height="80" rx="12" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
    <text x="60" y="36" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Actuator service</text>
    <text x="60" y="54" text-anchor="middle" font-size="10" fill="#6b6660">commands BuildSim</text>
    <!-- container hint -->
    <text x="60" y="70" text-anchor="middle" font-size="9" font-family="monospace" fill="#6b6660">docker</text>
  </g>
  <g transform="translate(580,230)">
    <rect width="120" height="80" rx="12" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <text x="60" y="36" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Audit log</text>
    <text x="60" y="54" text-anchor="middle" font-size="10" fill="#6b6660">every decision</text>
    <!-- container hint -->
    <text x="60" y="70" text-anchor="middle" font-size="9" font-family="monospace" fill="#6b6660">docker</text>
  </g>
<!-- connecting arrows (just a few key flows) -->
<line x1="160" y1="120" x2="213" y2="120" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="220,120 213,124 213,116" fill="#8b3a1f"/>
<line x1="340" y1="120" x2="393" y2="120" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="400,120 393,124 393,116" fill="#8b3a1f"/>
<line x1="520" y1="120" x2="573" y2="120" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="580,120 573,124 573,116" fill="#8b3a1f"/>
<line x1="340" y1="160" x2="453.8449562342426" y2="221.66601796021473" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="460,225 451.93982364007957,225.18318582636184 455.75008882840564,218.14885009406763" fill="#8b3a1f"/>
<line x1="280" y1="160" x2="166.15504376575743" y2="221.66601796021473" stroke="#6b6660" stroke-width="1.2" stroke-linecap="round"/>
<polygon points="160,225 164.24991117159442,218.14885009406763 168.06017635992043,225.18318582636184" fill="#6b6660"/>
<line x1="640" y1="160" x2="640" y2="218" stroke="#6b6660" stroke-width="1.2" stroke-linecap="round"/>
<polygon points="640,225 636,218 644,218" fill="#6b6660"/>
<line x1="640" y1="310" x2="526.6407830863536" y2="272.2135943621179" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="520,270 527.905694150421,268.41886116991583 525.3758720222862,276.0083275543199" fill="#8b3a1f"/>
</svg>
</div><figcaption>Each service is independently deployable, owns one job, and talks through stable interfaces. Replace one without touching the others.</figcaption></figure>

The patterns above describe **how** components talk. The next question is **how big** each component should be.

### Monolith versus microservices

A **monolithic** architecture puts all functionality into a single deployable unit. One process, one binary, one deploy. Simple to operate when small, painful when large.

A **microservices** architecture decomposes that single binary into small, independently deployable services that communicate over the network. Each service does one thing. Each can be updated, scaled, and operated independently.

For a CPS like a building control system, microservices fit well for four reasons.

Different components have different scaling requirements. The data pipeline needs storage. The AI agent needs CPU or GPU. The dashboard needs nothing. Mixing them in one process means provisioning for the worst-case need everywhere.

Different components evolve at different speeds. The sensor process is stable — once it works, it rarely changes. The agent logic, by contrast, evolves rapidly as new use cases are added. Microservices let the agent be redeployed many times a week without touching the sensor process.

Different components are written by different teams or in different languages. A research team writing the ML model in Python should not be forced to write the high-frequency sensor ingestion in Python too. The sensor ingestion can be in Go for performance; the agent in Python for ML ecosystem; they talk over the network.

Failure isolation. A bug in the dashboard does not take down the safety agent. In a monolith, a memory leak in the dashboard module crashes the whole binary, including safety.

The cost of microservices is operational complexity. Multiple containers must be orchestrated. Network failures between services must be handled gracefully. Deployment configuration multiplies. For a course project, this complexity is instructive rather than a burden. It teaches real-world architecture.

### Containers and Docker

A **container** packages a service together with all its dependencies — libraries, runtime, configuration files — into a single image that runs identically on any machine that has Docker installed. The "works on my laptop" problem disappears, because the laptop and the production server run the same container image.

A **Dockerfile** is the recipe for building a container image. It specifies the base operating system, the libraries to install, the application code to copy in, and the command to run at startup. Running the Dockerfile produces an **image**, a static snapshot. Running the image produces a **container**, a live instance with its own filesystem, network, and process tree.

**Docker Compose** is the tool for running multiple containers together. A `docker-compose.yml` file lists every container, its image, its port mappings, its dependencies on other containers, and its environment variables. Running `docker compose up` starts all of them in the right order. The compose file is essentially the executable version of a C4 container diagram.

### Each sensor and actuator as a service

In a service-oriented building control system, every sensor and every actuator is treated as a service with a well-defined interface. The architecture might include `smoke-sensor-service`, `temperature-sensor-service`, `hvac-actuator-service`, and `door-actuator-service`, each as a separate container. They share the same patterns: each publishes or subscribes on agreed topics; each registers itself with the building's API on startup; each can be killed and restarted without affecting the others.

This has two important consequences. First, the AI agent does not need to know whether the smoke sensor reads from a real device, a BuildSim simulator, a CSV file replay, or a generative ML model. As long as `sensors/smoke` carries the same JSON shape, the agent reads the same way. Second, scaling is per-service. If a building has 500 smoke sensors, you can spin up 500 instances of the same `smoke-sensor-service` image, each with different configuration.

### Eclipse Arrowhead

For industrial CPS, a reference architecture called the **Eclipse Arrowhead Framework** treats every sensor, actuator, and service as a registered, discoverable, and authorised service. Three core systems make this work:

- **Service Registry.** Every service registers itself; consumers look up producers by service type rather than by hardcoded address.
- **Authorisation System.** Enforces which services are allowed to talk to which. A guest sensor cannot pretend to be the building's primary smoke detector.
- **Orchestration System.** Routes requests from consumers to appropriate providers, picking the best available.

Arrowhead is used in Swedish industrial automation and is worth knowing as the production analogue of what we sketch with Docker Compose and MQTT in the course.

---

## Part 5 — Higher-Level Architectural Patterns

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 280" role="img" aria-label="Three higher-level architectural patterns" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Three shapes a system can take</text>
<!-- Layered -->
<g transform="translate(30,70)">
  <rect width="200" height="180" rx="12" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <text x="100" y="26" text-anchor="middle" font-size="11" font-weight="700" letter-spacing="1.5" fill="#8b3a1f">LAYERED</text>
  <rect x="22" y="44" width="156" height="22" rx="4" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/>
  <text x="100" y="59" text-anchor="middle" font-size="10" fill="#2a2622">UI · dashboard</text>
  <rect x="22" y="72" width="156" height="22" rx="4" fill="#dde7ec" stroke="#2a2622" stroke-width="1.2"/>
  <text x="100" y="87" text-anchor="middle" font-size="10" fill="#2a2622">Agents · logic</text>
  <rect x="22" y="100" width="156" height="22" rx="4" fill="#e2ebde" stroke="#2a2622" stroke-width="1.2"/>
  <text x="100" y="115" text-anchor="middle" font-size="10" fill="#2a2622">Storage</text>
  <rect x="22" y="128" width="156" height="22" rx="4" fill="#f4ead9" stroke="#2a2622" stroke-width="1.2"/>
  <text x="100" y="143" text-anchor="middle" font-size="10" fill="#2a2622">Sensors · actuators</text>
  <text x="100" y="168" text-anchor="middle" font-size="9.5" font-style="italic" fill="#6b6660">each layer talks down</text>
</g>
<!-- Microservices -->
<g transform="translate(250,70)">
  <rect width="220" height="180" rx="12" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <text x="110" y="26" text-anchor="middle" font-size="11" font-weight="700" letter-spacing="1.5" fill="#8b3a1f">MICROSERVICES</text>
  <rect x="22" y="43" width="36" height="26" rx="6" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/><rect x="92" y="43" width="36" height="26" rx="6" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/><rect x="162" y="43" width="36" height="26" rx="6" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/><rect x="22" y="93" width="36" height="26" rx="6" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/><rect x="92" y="93" width="36" height="26" rx="6" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/><rect x="162" y="93" width="36" height="26" rx="6" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/>
  <line x1="40" y1="68" x2="110" y2="42" stroke="#6b6660" stroke-width="1"/><line x1="110" y1="68" x2="180" y2="42" stroke="#6b6660" stroke-width="1"/><line x1="40" y1="68" x2="110" y2="92" stroke="#6b6660" stroke-width="1"/><line x1="110" y1="68" x2="110" y2="92" stroke="#6b6660" stroke-width="1"/><line x1="180" y1="68" x2="180" y2="92" stroke="#6b6660" stroke-width="1"/><line x1="40" y1="118" x2="110" y2="92" stroke="#6b6660" stroke-width="1"/><line x1="110" y1="118" x2="180" y2="92" stroke="#6b6660" stroke-width="1"/>
  <text x="110" y="148" text-anchor="middle" font-size="10" fill="#2a2622">small · independent · networked</text>
  <text x="110" y="168" text-anchor="middle" font-size="9.5" font-style="italic" fill="#6b6660">talk through APIs</text>
</g>
<!-- Event-driven -->
<g transform="translate(490,70)">
  <rect width="200" height="180" rx="12" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <text x="100" y="26" text-anchor="middle" font-size="11" font-weight="700" letter-spacing="1.5" fill="#8b3a1f">EVENT-DRIVEN</text>
  <rect x="40" y="80" width="120" height="22" rx="6" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
  <text x="100" y="95" text-anchor="middle" font-size="10" font-weight="600" fill="#2a2622">Event bus</text>
  <rect x="46" y="38" width="28" height="22" rx="4" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/><rect x="86" y="38" width="28" height="22" rx="4" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/><rect x="126" y="38" width="28" height="22" rx="4" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/>
  <line x1="60" y1="62" x2="60" y2="80" stroke="#8b3a1f" stroke-width="1.2"/><line x1="100" y1="62" x2="100" y2="80" stroke="#8b3a1f" stroke-width="1.2"/><line x1="140" y1="62" x2="140" y2="80" stroke="#8b3a1f" stroke-width="1.2"/>
  <rect x="46" y="128" width="28" height="22" rx="4" fill="#e2ebde" stroke="#2a2622" stroke-width="1.2"/><rect x="86" y="128" width="28" height="22" rx="4" fill="#e2ebde" stroke="#2a2622" stroke-width="1.2"/><rect x="126" y="128" width="28" height="22" rx="4" fill="#e2ebde" stroke="#2a2622" stroke-width="1.2"/>
  <line x1="60" y1="102" x2="60" y2="128" stroke="#8b3a1f" stroke-width="1.2"/><line x1="100" y1="102" x2="100" y2="128" stroke="#8b3a1f" stroke-width="1.2"/><line x1="140" y1="102" x2="140" y2="128" stroke="#8b3a1f" stroke-width="1.2"/>
  <text x="100" y="170" text-anchor="middle" font-size="9.5" font-style="italic" fill="#6b6660">publishers · subscribers · loose</text>
</g>
</svg>
</div><figcaption>Different shapes optimise for different things: layered for simplicity, microservices for scale, event-driven for loose coupling between components that come and go.</figcaption></figure>

The previous parts cover individual building blocks. This part describes whole-system patterns that combine them.

### Event-driven pub/sub architecture

The most common pattern for building control. Sensors publish readings to a message broker. Consumers — AI agents, databases, dashboards — subscribe to the topics they care about. The broker handles routing.

Advantages: loose coupling, easy extension (a new consumer is added without changing producers), and a natural fit for the one-to-many relationship between sensors and downstream systems. The default choice for sensor data collection, event distribution, and real-time alerting.

### Lambda architecture

The Lambda architecture splits data processing into two parallel paths. The **speed layer**, also called the hot path, processes events in real time and produces low-latency approximate results — current temperature, active alerts, anomaly scores in the last minute. The **batch layer**, also called the cold path, processes historical data in bulk and produces accurate long-term results — weekly energy reports, monthly trends, training data for the next model. A **serving layer** merges results from both paths when a query comes in.

For building control, the speed layer handles real-time ML inference and safety responses, while the batch layer trains models, generates reports, and populates the feature store for the next training run.

Lambda has been criticised because the same business logic ends up implemented twice — once in the streaming path and once in the batch path — which is hard to keep in sync. The **Kappa architecture**, proposed by Nathan Marz, simplifies this by using a single stream-processing layer for both real-time and historical processing, replaying historical data through the same code. Apache Flink supports both patterns.

### Multi-agent architecture

Instead of one agent that controls everything, several specialised agents each focus on a single objective. A typical decomposition:

- A **Safety Agent** monitors smoke and fire sensors and responds to emergencies. It has the highest priority — it always overrides.
- An **Energy Agent** optimises HVAC schedules to minimise consumption.
- A **Comfort Agent** maintains temperature and air quality within occupant preferences.
- A **Coordination Layer** resolves conflicts. Safety always wins. Energy and comfort negotiate.

Each agent is a separate process. Each can be updated independently. The coordination layer is itself a separate component, which can be implemented as priority-based (the simplest), auction-based (each agent submits a bid), or consensus-based (agents converge through messaging). Multi-agent is the dominant pattern when the system has competing objectives.

### Digital twin feedback loop

A **digital twin** is a real-time simulation of the physical system that runs in parallel with the real one. The idea is to use the twin as a sandbox: before sending a command to a real actuator, send it to the twin and observe the predicted outcome. If the twin says the outcome would be bad, the agent tries a different command.

BuildSim is itself a digital twin of a building. The agent can ask it questions like "if I turn HVAC zone 3 on, what will the temperature be in ten minutes?" The agent uses the answer to choose better actions before committing them to the real building. For higher-fidelity simulation, tools like the Building Controls Virtual Test Bed and EnergyPlus exist in research labs.

```
        ┌──────────┐
        │ AI Agent │
        └────┬─────┘
             │ proposed command
             ▼
        ┌──────────────────┐    simulated outcome     ┌─────────────────┐
        │ Digital Twin      │ ──────────────────────► │ outcome OK?     │
        │ (BuildSim)        │                          └────────┬────────┘
        └──────────────────┘                                    │
                                                  ┌─────────────┴────────┐
                                              yes │                      │ no
                                                  ▼                      ▼
                                          ┌─────────────┐         (try another
                                          │ Real        │          command)
                                          │ Building    │
                                          └─────────────┘
                                                 │
                                                 │ sensor feedback
                                                 ▼
                                            back to agent
```

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

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 380" role="img" aria-label="A worked HVAC control system" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">A full HVAC control system, placed</text>
<!-- Top: building/sensors -->
<g transform="translate(40,60)">
  <rect width="640" height="64" rx="10" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
  <text x="20" y="22" font-size="11" font-weight="700" letter-spacing="1.5" fill="#8b3a1f">SENSORS & ACTUATORS</text>
  <g transform="translate(160,42)">
    <circle cx="0" cy="0" r="11" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3.3000000000000003" fill="#8b3a1f"/>
    <path d="M -8.8 -11 Q -13.200000000000001 -15.400000000000002 -8.8 -19.8" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8.8 -11 Q 13.200000000000001 -15.400000000000002 8.8 -19.8" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><text x="160" y="64" text-anchor="middle" font-size="9" fill="#2a2622">temp</text>
  <g transform="translate(220,42)">
    <circle cx="0" cy="0" r="11" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3.3000000000000003" fill="#8b3a1f"/>
    <path d="M -8.8 -11 Q -13.200000000000001 -15.400000000000002 -8.8 -19.8" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8.8 -11 Q 13.200000000000001 -15.400000000000002 8.8 -19.8" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><text x="220" y="64" text-anchor="middle" font-size="9" fill="#2a2622">CO₂</text>
  <g transform="translate(280,42)">
    <circle cx="0" cy="0" r="11" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3.3000000000000003" fill="#8b3a1f"/>
    <path d="M -8.8 -11 Q -13.200000000000001 -15.400000000000002 -8.8 -19.8" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8.8 -11 Q 13.200000000000001 -15.400000000000002 8.8 -19.8" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><text x="280" y="64" text-anchor="middle" font-size="9" fill="#2a2622">occ.</text>
  <g transform="translate(420,30)"><rect width="28" height="20" rx="4" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.3"/><circle cx="8" cy="10" r="4" fill="#8b3a1f"/></g>
  <text x="434" y="64" text-anchor="middle" font-size="9" fill="#2a2622">fan</text>
  <g transform="translate(480,30)"><rect width="28" height="20" rx="4" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.3"/><circle cx="8" cy="10" r="4" fill="#8b3a1f"/></g>
  <text x="494" y="64" text-anchor="middle" font-size="9" fill="#2a2622">damper</text>
</g>
<!-- Middle: broker -->
<g transform="translate(220,150)">
  <rect width="280" height="38" rx="8" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
  <text x="140" y="22" text-anchor="middle" font-size="11" font-weight="700" fill="#2a5a7a">MQTT BROKER · sensors/* · actuators/*</text>
</g>
<!-- Bottom row: services -->
  <g transform="translate(40,220)">
    <rect width="120" height="110" rx="10" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <polyline points="20,38 32,28 44,32 56,22 68,30 80,18 92,24" fill="none" stroke="#8b3a1f" stroke-width="1.5"/>
    <text x="60" y="80" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Forecast service</text>
    <text x="60" y="96" text-anchor="middle" font-size="9.5" fill="#6b6660">next-hour temp</text>
  </g>
  <g transform="translate(200,220)">
    <rect width="120" height="110" rx="10" fill="#f0d9d1" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(56,32)"><g><circle cx="0" cy="0" r="18" fill="#8b3a1f" opacity="0.9"/>
    <text x="0" y="5" text-anchor="middle" font-size="12.6" font-weight="700" fill="#fdfbf7">AI</text></g></g>
    <text x="60" y="80" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">HVAC controller</text>
    <text x="60" y="96" text-anchor="middle" font-size="9.5" fill="#6b6660">optimise comfort</text>
  </g>
  <g transform="translate(360,220)">
    <rect width="120" height="110" rx="10" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="32" y="20" width="48" height="32" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/><line x1="32" y1="28" x2="80" y2="28" stroke="#2a2622" stroke-width="0.8"/><line x1="48" y1="20" x2="48" y2="52" stroke="#2a2622" stroke-width="0.6"/><line x1="64" y1="20" x2="64" y2="52" stroke="#2a2622" stroke-width="0.6"/><circle cx="56" cy="40" r="3" fill="#8b3a1f"/>
    <text x="60" y="80" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Schedule</text>
    <text x="60" y="96" text-anchor="middle" font-size="9.5" fill="#6b6660">occupancy rules</text>
  </g>
  <g transform="translate(520,220)">
    <rect width="120" height="110" rx="10" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(56,36)"><g>
    <ellipse cx="0" cy="-14" rx="20" ry="6" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <path d="M -20 -14 L -20 14 Q 0 22 20 14 L 20 -14" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="0" cy="-14" rx="20" ry="6" fill="none" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="0" cy="-4" rx="18" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
    <ellipse cx="0" cy="8" rx="17" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
  </g></g>
    <text x="60" y="80" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Time-series DB</text>
    <text x="60" y="96" text-anchor="middle" font-size="9.5" fill="#6b6660">history · 6 months</text>
  </g>
<!-- arrows: top -> broker -->
<line x1="220" y1="125" x2="293.27251487249083" y2="146.06584802584112" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="300,148 292.16728517297145,149.91012524156065 294.3777445720102,142.2215708101216" fill="#8b3a1f"/>
<line x1="500" y1="125" x2="426.72748512750917" y2="146.06584802584112" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="420,148 425.6222554279898,142.2215708101216 427.83271482702855,149.91012524156065" fill="#8b3a1f"/>
<!-- arrows: broker -> services -->
<line x1="280" y1="190" x2="106.904757466825" y2="218.8492070888625" stroke="#6b6660" stroke-width="1.3" stroke-linecap="round"/>
<polygon points="100,220 106.24716151760357,214.9036313935339 107.56235341604643,222.79478278419109" fill="#6b6660"/>
<line x1="320" y1="190" x2="266.2609903369994" y2="216.8695048315003" stroke="#6b6660" stroke-width="1.3" stroke-linecap="round"/>
<polygon points="260,220 264.47213595499954,213.29179606750066 268.0498447189992,220.44721359549996" fill="#6b6660"/>
<line x1="400" y1="190" x2="416.1170986264234" y2="214.1756479396351" stroke="#6b6660" stroke-width="1.3" stroke-linecap="round"/>
<polygon points="420,220 412.788897449072,216.394448724536 419.4452998037748,211.9568471547342" fill="#6b6660"/>
<line x1="440" y1="190" x2="573.1553831014581" y2="218.53329637888388" stroke="#6b6660" stroke-width="1.3" stroke-linecap="round"/>
<polygon points="580,220 572.3172667465346,222.4445060351935 573.9934994563816,214.62208672257424" fill="#6b6660"/>
</svg>
</div><figcaption>A full HVAC control system, with each component placed and the data flowing top to bottom: physical world → broker → services → back to the world.</figcaption></figure>

To make all of the above concrete, consider a fire detection system and walk through the architectural choices.

The use case: smoke sensors and temperature sensors are deployed in every room. When smoke is detected, the system must activate sprinklers in the affected room within one second, unlock fire doors on the evacuation path, and notify the building manager. False positives must be kept under 5 percent.

**Where does the brain run?** At the edge. The one-second latency budget rules out cloud reasoning on the critical path.

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

That sequence of choices — edge for inference, MQTT for sensor fan-out, REST for commands, microservices for failure isolation, multi-agent for competing objectives, digital twin for safety, ColonyOS for training orchestration — is the architecture this lecture trains its readers to design.

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
