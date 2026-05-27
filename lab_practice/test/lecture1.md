# Lecture 1 — Introduction and Model-Based Systems Engineering

Standalone notes for the first lecture of D7065E. Read on its own or alongside `lectures/lecture-1-introduction.md`.

---

## Part 1 — Embedded Intelligence at the Edge

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 320" role="img" aria-label="Edge intelligence vs cloud" class="dgm">
<text x="360" y="32" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">The brain lives where the action is</text>
<g>
    <polygon points="50,120 105,90 160,120" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="50" y="120" width="110" height="100" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="62" y="140" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="94" y="140" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="126" y="140" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="87" y="180" width="36" height="30" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  </g>
<g transform="translate(180,170)">
  <rect x="0" y="0" width="60" height="40" rx="6" fill="#8b3a1f" stroke="#2a2622" stroke-width="1.5"/>
  <text x="30" y="26" text-anchor="middle" font-size="14" font-weight="700" fill="#fdfbf7">AI</text>
  <text x="30" y="58" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Edge server</text>
  <text x="30" y="72" text-anchor="middle" font-size="10" fill="#6b6660">~5 ms</text>
</g>
<path d="M 165 145 Q 175 165 195 175" fill="none" stroke="#8b3a1f" stroke-width="2"/>
<polygon points="195,175 187,172 192,180" fill="#8b3a1f"/>
<path d="M 195 200 Q 175 210 165 220" fill="none" stroke="#8b3a1f" stroke-width="2"/>
<polygon points="165,220 173,217 168,225" fill="#8b3a1f"/>
<path d="M 260 190 Q 320 170 380 190 Q 440 210 500 190 Q 540 175 560 175" fill="none" stroke="#6b6660" stroke-width="1.5" stroke-dasharray="5 4"/>
<text x="420" y="160" text-anchor="middle" font-size="11" fill="#6b6660">~200 ms across the internet</text>
<text x="420" y="222" text-anchor="middle" font-size="10" font-style="italic" fill="#6b6660">optional · never on the critical path</text>
<g transform="translate(547,135) scale(1.4)">
    <path d="M 20 50 Q 0 50 0 35 Q 0 22 14 20 Q 16 6 32 6 Q 50 0 60 14 Q 78 14 80 30 Q 92 32 90 46 Q 88 60 70 60 L 28 60 Q 14 60 20 50 Z" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
  </g>
<text x="610" y="240" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Cloud</text>
<text x="610" y="256" text-anchor="middle" font-size="10.5" fill="#6b6660">training, analytics</text>
</svg>
</div><figcaption>Local intelligence sits on the same network as the sensors and actuators it controls. The cloud helps with slow tasks, but the critical loop never depends on it.</figcaption></figure>

A modern building, a modern car, a modern factory — none of these are simple machines anymore. Each contains hundreds of sensors and actuators, and a continuous stream of small decisions is needed to keep it running safely, efficiently, and comfortably. The branch of computer science that deals with making those decisions on the same network as the physical system, rather than in a remote data centre far away, is called **embedded intelligence at the edge**.

The phrase has two halves worth unpacking.

**Embedded** means the software lives inside or right next to the physical system it controls. The control logic of a car's anti-lock brakes lives in the car. The thermostat's brain lives in the thermostat itself. There is no question of phoning home to ask permission. The software is part of the thing.

**At the edge** means on the local network — close to the sensors and the actuators — not in the cloud. A useful image: the difference between a self-driving car that can stop itself when something runs into the road, and one that has to ask a server in another country whether to brake. The first one is safe. The second one is a horror film waiting to happen.

Two ideas follow naturally. First, decisions happen fast and locally, in milliseconds, because the physical world cannot wait for a round trip across the internet. Second, the system keeps working when the wider network does not. If the building's broadband line is cut during a storm, the fire detector still detects fires.

The course's running project, autonomous building control, makes both points concrete. A simulated building called BuildSim provides rooms, sensors, and actuators. The job is to instrument that building with software that observes its state, decides what to do, and changes the state — all running on the same network as the building itself, without depending on any cloud service to make the decisions. Cloud services can help with slow tasks, like overnight model training, but they are never on the critical path.

Before any code is written, a discipline called Model-Based Systems Engineering is applied: the system is fully specified using structured diagrams and tables. The specification is reviewed and approved at week three. The approved design becomes the basis for both implementation and the final oral examination.

---

## Part 2 — A Building as a Cyber-Physical System

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 400" role="img" aria-label="CPS feedback loop" class="dgm">
<text x="360" y="32" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">The two halves of a cyber-physical system</text>
<g transform="translate(40,70)">
  <rect width="260" height="290" rx="14" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
  <text x="130" y="30" text-anchor="middle" font-size="11" font-weight="700" letter-spacing="2" fill="#8b3a1f">PHYSICAL</text>
  <g>
    <polygon points="75,90 130,60 185,90" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="75" y="90" width="110" height="90" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="87" y="110" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="119" y="110" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="151" y="110" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="112" y="150" width="36" height="20" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  </g>
  <!-- thermometer -->
  <line x1="55" y1="120" x2="55" y2="170" stroke="#8b3a1f" stroke-width="2"/>
  <circle cx="55" cy="175" r="6" fill="#8b3a1f"/>
  <!-- smoke -->
  <path d="M 200 100 Q 205 90 200 80 Q 195 70 200 60" fill="none" stroke="#6b6660" stroke-width="1.5"/>
  <text x="130" y="240" text-anchor="middle" font-size="11.5" fill="#2a2622">rooms · ducts · pipes</text>
  <text x="130" y="258" text-anchor="middle" font-size="11.5" fill="#2a2622">temperature · smoke · people</text>
  <text x="130" y="276" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">obeys physics</text>
</g>
<g transform="translate(420,70)">
  <rect width="260" height="290" rx="14" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
  <text x="130" y="30" text-anchor="middle" font-size="11" font-weight="700" letter-spacing="2" fill="#2a5a7a">CYBER</text>
  <g transform="translate(80,55)">
    <rect width="100" height="130" rx="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
    <g><circle cx="50" cy="40" r="22" fill="#8b3a1f" opacity="0.9"/>
    <text x="50" y="45" text-anchor="middle" font-size="15.399999999999999" font-weight="700" fill="#fdfbf7">AI</text></g>
    <line x1="14" y1="78" x2="86" y2="78" stroke="#6b6660" stroke-width="1"/>
    <line x1="14" y1="90" x2="86" y2="90" stroke="#6b6660" stroke-width="1"/>
    <line x1="14" y1="102" x2="62" y2="102" stroke="#6b6660" stroke-width="1"/>
    <line x1="14" y1="114" x2="78" y2="114" stroke="#6b6660" stroke-width="1"/>
  </g>
  <text x="130" y="240" text-anchor="middle" font-size="11.5" fill="#2a2622">sensor processes · pipelines</text>
  <text x="130" y="258" text-anchor="middle" font-size="11.5" fill="#2a2622">AI agents · actuator commands</text>
  <text x="130" y="276" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">obeys code</text>
</g>
<g>
  <line x1="308" y1="150" x2="401" y2="150" stroke="#8b3a1f" stroke-width="2.2" stroke-linecap="round"/>
<polygon points="408,150 401,154 401,146" fill="#8b3a1f"/>
  <text x="358" y="142" text-anchor="middle" font-size="11" font-weight="600" fill="#8b3a1f">sense</text>
  <text x="358" y="170" text-anchor="middle" font-size="10" fill="#6b6660">sensors report state</text>
  <line x1="408" y1="280" x2="315" y2="280" stroke="#8b3a1f" stroke-width="2.2" stroke-linecap="round"/>
<polygon points="308,280 315,276 315,284" fill="#8b3a1f"/>
  <text x="358" y="272" text-anchor="middle" font-size="11" font-weight="600" fill="#8b3a1f">act</text>
  <text x="358" y="298" text-anchor="middle" font-size="10" fill="#6b6660">actuators change state</text>
  <text x="358" y="220" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">repeats forever</text>
</g>
</svg>
</div><figcaption>A continuous feedback loop joins the world the system lives in to the software that observes and changes it. Every component in a CPS sits somewhere on this loop.</figcaption></figure>

### The two halves of a CPS

A **cyber-physical system**, abbreviated CPS, is a system where computation and physical processes are tightly intertwined. Two halves live side by side and communicate forever.

The **physical half** is everything you can touch. In a building, that means rooms, walls, doors, HVAC ducts, ventilation fans, smoke and air, temperature gradients, water pipes, electrical loads, and the people walking around inside. Physical state evolves according to the laws of physics and the patterns of human behaviour. The temperature in a room does not stop changing just because no computer is watching.

The **cyber half** is the software. Sensor processes that read the physical state and report it. Data pipelines that store the readings. AI agents that reason about what should happen next. Actuator processes that issue commands. The cyber half is fast and flexible but has no direct access to physics — it only knows what sensors report, and it only acts through what actuators can do.

A useful image: the physical half is the room, the cyber half is the brain. The brain cannot feel the temperature directly. It can only ask the thermometer (a sensor) what the temperature is, and it can only warm the room by telling the heater (an actuator) to turn on. Everything in between is just words and electricity.

### The continuous feedback loop

A CPS is defined by the loop that joins the two halves. The cycle runs forever, and every component sits somewhere on it.

```
   1. Physical world produces a state         (smoke level rises)
                  │
                  ▼
   2. Sensor measures the state              (smoke = 0.82)
                  │
                  ▼
   3. Software reasons about the measurement (this looks like a fire)
                  │
                  ▼
   4. Actuator changes the world             (sprinkler on)
                  │
                  ▼
   5. World responds to the change           (smoke level drops)
                  │
                  ▼
   6. Sensor measures the new state          (smoke = 0.4)
                  │
                  ▼
   7. Software notices and adapts            (working; consider easing off)
                  │
                  ▼
   8. (Back to step 1.)
```

Riding a bicycle is the simplest everyday version of this loop. You sense your balance (eyes and inner ear are the sensors). You decide whether you're leaning too far left or right (the brain is the reasoner). You correct (your muscles are the actuators). And your balance changes — the world responds. The loop runs many times per second, automatically, and if any part of it breaks, you fall over.

A CPS is a bicycle for an entire building. The loop has to be fast enough that the physics can't run away (a fire-suppression loop in well under a second, a climate-control loop comfortable with one update every few minutes). It has to be reliable enough to keep working when a sensor crashes, a network drops, or a model gives the wrong answer. And it has to be correct enough that a wrong decision does not start a fire, lock people in during an emergency, or freeze a server room.

### Automation versus autonomy

A grandfather clock is automated. Every hour, it strikes the hour. It follows a fixed rule that was wound up once and changes nothing. Automation is predictable and easy to verify but brittle — the rule has to anticipate every situation, and any situation it didn't anticipate becomes a bug.

A modern smart thermostat is autonomous. It knows that the building is empty on Sunday morning, that the weather forecast is mild, that the family is on holiday. It predicts that nobody will want a warm room for the next eight hours. It learns from history. It adapts as patterns change.

The difference is bigger than it sounds. Two ways to think about it:

| Automation | Autonomy |
|---|---|
| `if temp < 20°C: turn on heater` | "It's 9 a.m. on a Tuesday, the office will fill in 30 minutes, outdoor temp is dropping, pre-heat now so the room reaches 22°C by then." |
| One rule, fixed forever | Many models, adapting over time |
| Predictable but brittle | Smarter but less predictable |

Real-world autonomous building systems include Johnson Controls' OpenBlue, Siemens Desigo CC, and the cooling controller that DeepMind built for Google's data centres — which reduced cooling-energy consumption by about 40 percent by reasoning over far more variables than a rule-based controller can hold.

The course's autonomous building control project asks for the autonomous version, not the automated one. Threshold rules ("if smoke > 0.7, alarm") are not sufficient. A real ML model or LLM reasoning step is required.

---

## Part 3 — Specification Before Code: Model-Based Systems Engineering

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 320" role="img" aria-label="MBSE: prose vs structured specification" class="dgm">
<text x="360" y="32" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">From "vibe" to working recipe</text>
<!-- Vague prose paper -->
<g transform="translate(60,80)">
  <rect width="180" height="200" rx="4" fill="#fdfbf7" stroke="#6b6660" stroke-width="1.5"/>
  <line x1="18" y1="30" x2="162" y2="30" stroke="#6b6660" stroke-width="1" opacity="0.5"/>
  <line x1="18" y1="46" x2="150" y2="46" stroke="#6b6660" stroke-width="1" opacity="0.5"/>
  <line x1="18" y1="62" x2="162" y2="62" stroke="#6b6660" stroke-width="1" opacity="0.5"/>
  <line x1="18" y1="78" x2="120" y2="78" stroke="#6b6660" stroke-width="1" opacity="0.5"/>
  <line x1="18" y1="94" x2="155" y2="94" stroke="#6b6660" stroke-width="1" opacity="0.5"/>
  <text x="90" y="130" text-anchor="middle" font-size="40" fill="#8b3a1f" opacity="0.6">?</text>
  <line x1="18" y1="158" x2="162" y2="158" stroke="#6b6660" stroke-width="1" opacity="0.5"/>
  <line x1="18" y1="174" x2="140" y2="174" stroke="#6b6660" stroke-width="1" opacity="0.5"/>
  <text x="90" y="180" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">"Add a bit of salt</text>
  <text x="90" y="195" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">and bake until done."</text>
</g>
<text x="150" y="298" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">prose specification</text>
<line x1="265" y1="180" x2="353" y2="180" stroke="#8b3a1f" stroke-width="2.5" stroke-linecap="round"/>
<polygon points="360,180 353,184 353,176" fill="#8b3a1f"/>
<text x="312" y="170" text-anchor="middle" font-size="10" font-weight="600" fill="#8b3a1f">MBSE</text>
<!-- Structured spec -->
<g transform="translate(380,80)">
  <rect width="280" height="200" rx="4" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <!-- requirements table -->
  <rect x="14" y="18" width="120" height="80" rx="3" fill="#f3ece6" stroke="#2a2622" stroke-width="1"/>
  <text x="74" y="36" text-anchor="middle" font-size="9" font-weight="700" fill="#8b3a1f">REQUIREMENTS</text>
  <line x1="22" y1="46" x2="126" y2="46" stroke="#2a2622" stroke-width="0.8"/>
  <line x1="22" y1="58" x2="100" y2="58" stroke="#2a2622" stroke-width="0.8"/>
  <line x1="22" y1="70" x2="116" y2="70" stroke="#2a2622" stroke-width="0.8"/>
  <line x1="22" y1="82" x2="92" y2="82" stroke="#2a2622" stroke-width="0.8"/>
  <!-- arch diagram -->
  <g transform="translate(148,18)">
    <rect width="118" height="80" rx="3" fill="#dde7ec" stroke="#2a2622" stroke-width="1"/>
    <text x="59" y="34" text-anchor="middle" font-size="9" font-weight="700" fill="#2a5a7a">ARCHITECTURE</text>
    <rect x="14" y="44" width="22" height="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="0.8"/>
    <rect x="48" y="44" width="22" height="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="0.8"/>
    <rect x="82" y="44" width="22" height="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="0.8"/>
    <line x1="36" y1="51" x2="48" y2="51" stroke="#2a2622" stroke-width="0.8"/>
    <line x1="70" y1="51" x2="82" y2="51" stroke="#2a2622" stroke-width="0.8"/>
    <rect x="32" y="64" width="54" height="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="0.8"/>
  </g>
  <!-- interfaces -->
  <rect x="14" y="108" width="120" height="74" rx="3" fill="#e2ebde" stroke="#2a2622" stroke-width="1"/>
  <text x="74" y="126" text-anchor="middle" font-size="9" font-weight="700" fill="#3a5a3a">INTERFACES</text>
  <text x="22" y="142" font-size="8" font-family="monospace" fill="#2a2622">{ ts, room, value }</text>
  <text x="22" y="156" font-size="8" font-family="monospace" fill="#2a2622">PUT /act/{id}</text>
  <text x="22" y="170" font-size="8" font-family="monospace" fill="#2a2622">QoS: 1</text>
  <!-- behaviour -->
  <rect x="148" y="108" width="118" height="74" rx="3" fill="#f4ead9" stroke="#2a2622" stroke-width="1"/>
  <text x="207" y="126" text-anchor="middle" font-size="9" font-weight="700" fill="#7a5a1a">BEHAVIOUR</text>
  <circle cx="166" cy="146" r="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="0.8"/>
  <circle cx="190" cy="146" r="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="0.8"/>
  <circle cx="214" cy="146" r="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="0.8"/>
  <circle cx="238" cy="146" r="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="0.8"/>
  <line x1="172" y1="146" x2="184" y2="146" stroke="#2a2622" stroke-width="0.8"/>
  <line x1="196" y1="146" x2="208" y2="146" stroke="#2a2622" stroke-width="0.8"/>
  <line x1="220" y1="146" x2="232" y2="146" stroke="#2a2622" stroke-width="0.8"/>
  <text x="207" y="172" text-anchor="middle" font-size="8" font-family="monospace" fill="#2a2622">Start → Run → Stop</text>
</g>
<text x="520" y="298" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">structured artefacts</text>
</svg>
</div><figcaption>MBSE replaces ambiguous documents with structured, precise models. Two engineers reading the same model build the same mental picture.</figcaption></figure>

### Why prose breaks down

The default way to design software is to write a Word document describing what should be built, hand it to the developers, and hope they get the same picture from it.

This almost always fails. Prose is ambiguous, and the ambiguity is invisible until somebody tries to act on the document. Two engineers reading the same paragraph form different mental pictures. The product manager imagined a database; the developer built a spreadsheet. The architect assumed the broker held messages for a minute; the operations engineer configured it to drop them after a second. Three weeks into implementation, somebody discovers a contradiction. Six weeks in, the design and the code have parted ways completely.

A useful image: writing prose specifications is like cooking from a recipe that says "add a bit of salt and bake until done." Five cooks following this recipe produce five different dishes. The recipe didn't capture the actual instructions; it captured a vibe.

A working recipe says: "1.5 teaspoons of salt, bake at 180°C for 35 minutes." Five cooks following this recipe produce the same dish. The difference between a vibe and a working recipe is precision.

### What MBSE is

**Model-Based Systems Engineering**, abbreviated MBSE, replaces prose documents with structured, precise models of the system. A model is unambiguous — either it specifies a thing or it doesn't. Two engineers reading the same model build the same mental picture.

The models are not single documents but a small collection of structured artefacts:

- A **requirements table** listing every requirement with a unique ID, a type, a priority, and an acceptance criterion.
- An **architecture diagram** with labelled boxes for every component and labelled arrows for every connection.
- **Interface specifications** — JSON schemas for every message, REST endpoints with request and response shapes, MQTT topic names and payload formats.
- **Behaviour models** — sequence diagrams showing message exchanges over time, state machine diagrams showing what each component does in each state.
- A **validation matrix** linking each requirement to the test that proves it works.

A useful image: an architect designing a house doesn't write a 200-page essay about the house. They draw blueprints. The blueprints have measurements, labels, and conventions that any contractor in the world can read the same way. MBSE applies the same idea to software systems.

### Why this matters more for CPS

Two engineers building a typical web application can sometimes get away with ambiguity, because the consequences of a misunderstanding are usually a wrong-looking page that gets fixed in the next sprint.

A cyber-physical system is different. The interactions between physical processes and software components are complex, timing-sensitive, and have safety consequences. A sequence diagram showing a smoke sensor publishing a reading, a data pipeline storing it, an anomaly model evaluating it, and a safety agent commanding a sprinkler makes the design concrete in a way that no paragraph can. A state diagram of the sensor process specifies exactly what happens when the network drops and what happens when it recovers. These artefacts catch bugs that prose hides.

### The MBSE process

The work flows through six activities. They are not strictly sequential — work loops back as understanding deepens — but they are ordered in importance and dependence.

**Requirements analysis.** Write down what the system must do as testable statements, each with a unique ID. Three kinds appear: **functional** ("detect fire conditions within 30 seconds"), **non-functional** ("survive a sensor-process crash without losing data"), and **regulatory** ("comply with Swedish BBR fire protection requirements"). A useful image: requirements are the customer's wishlist, formalised so it cannot be misread. "Make the cake delicious" is not a requirement. "The cake must be 30 cm in diameter, contain no peanuts, and be ready by 6 p.m." is.

**Functional decomposition.** Take each high-level requirement and break it into the smaller operations needed to satisfy it. "Detect fire" decomposes into collecting smoke and temperature readings, validating them, applying a detection model, raising an alert, commanding sprinklers, and notifying occupants. A useful image: decomposing a recipe. "Make a cake" decomposes into "measure flour, beat eggs, mix, pour into tin, bake at 180°C for 35 minutes, cool, frost." Each step is small enough to assign to a worker.

**Architecture design.** Decide what software components exist, how they're organised, and where they run. A component diagram shows the parts. A deployment view shows which machine each part lives on. A useful image: the floor plan of a house. Walls, rooms, doors — but no furniture yet.

**Interface design.** Specify exactly how components talk to each other. The REST endpoint URL and the request and response shapes. The MQTT topic name and the payload format. The protocol, the message size, the rate. Ambiguity here is the most expensive kind. Two components that look connected on the diagram but disagree about the message format will work in isolation and fail when they meet. A useful image: a power plug and a wall socket. If the prongs don't match the holes, the lamp doesn't light. Better to discover that mismatch on paper than on the day the lamp is supposed to ship.

**Behaviour modeling.** Capture how the system acts over time. A **sequence diagram** is a comic strip with software components as the characters — it shows one specific scenario, with messages going back and forth in order. A **state machine** is a map of the moods a component can be in — `Starting`, `Running`, `Reconnecting`, `Stopped` — and the events that flip it between them. Both are precise; both expose timing bugs that static diagrams hide.

**Validation.** Close the loop by linking every requirement to a design element that satisfies it and a test case that verifies it. If a requirement has no owning component, the design has a gap. If it has no test, the requirement cannot be enforced. A useful image: tasting the soup before serving. Recipes that produce dishes nobody tests are recipes that produce surprises.

### Why design happens before code

Errors discovered during design cost minutes. Errors discovered during implementation cost days. Errors discovered after deployment cost weeks, or — in safety-critical systems — lives.

The cost curve is steep, and the MBSE process catches as many errors as possible while they are still cheap.

A second reason has emerged more recently. When code is generated with AI tools, the quality of the output is bounded by the quality of the input. A vague prompt produces clean-looking code that satisfies the prompt but not the underlying problem. A precise specification produces working code that solves the actual problem. The specification is also the test contract — the same artefact that describes the system also describes how to know whether it was built correctly.

A useful image: imagine asking a contractor to "build me a house." You will get a house, technically. It will not be the house you wanted. Now imagine handing the contractor a complete set of blueprints with measurements, materials, and electrical layout. You will get the house you wanted. The contractor is talented in both cases. The difference is the contract.

---

## Part 4 — Architecture Viewpoints

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 360" role="img" aria-label="Architecture viewpoints — one system, many lenses" class="dgm">
<text x="360" y="32" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">One system, many lenses</text>
<g>
    <polygon points="305,180 360,150 415,180" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="305" y="180" width="110" height="100" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="317" y="200" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="349" y="200" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="381" y="200" width="22" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <rect x="342" y="240" width="36" height="30" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  </g>
<text x="360" y="304" text-anchor="middle" font-size="11" fill="#2a2622" font-weight="600">The system</text>
<!-- 5 lenses around -->
  <g>
    <ellipse cx="90" cy="90" rx="56" ry="36" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <text x="90" y="88" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Business</text>
    <text x="90" y="103" text-anchor="middle" font-size="10" fill="#6b6660">goals · regs</text>
  </g>
  <g>
    <ellipse cx="580" cy="90" rx="56" ry="36" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
    <text x="580" y="88" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Functional</text>
    <text x="580" y="103" text-anchor="middle" font-size="10" fill="#6b6660">components</text>
  </g>
  <g>
    <ellipse cx="560" cy="240" rx="56" ry="36" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <text x="560" y="238" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Behavioural</text>
    <text x="560" y="253" text-anchor="middle" font-size="10" fill="#6b6660">sequences</text>
  </g>
  <g>
    <ellipse cx="100" cy="240" rx="56" ry="36" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
    <text x="100" y="238" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Information</text>
    <text x="100" y="253" text-anchor="middle" font-size="10" fill="#6b6660">data flow</text>
  </g>
  <g>
    <ellipse cx="340" cy="80" rx="56" ry="36" fill="#f0d9d1" stroke="#2a2622" stroke-width="1.5"/>
    <text x="340" y="78" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Deployment</text>
    <text x="340" y="93" text-anchor="middle" font-size="10" fill="#6b6660">where it runs</text>
  </g>
<!-- lines connecting lenses to building (subtle) -->
<line x1="146" y1="90" x2="360" y2="200" stroke="#6b6660" stroke-width="1" stroke-dasharray="3 3" opacity="0.55"/><line x1="524" y1="90" x2="360" y2="200" stroke="#6b6660" stroke-width="1" stroke-dasharray="3 3" opacity="0.55"/><line x1="504" y1="240" x2="360" y2="200" stroke="#6b6660" stroke-width="1" stroke-dasharray="3 3" opacity="0.55"/><line x1="156" y1="240" x2="360" y2="200" stroke="#6b6660" stroke-width="1" stroke-dasharray="3 3" opacity="0.55"/><line x1="360" y1="80" x2="360" y2="200" stroke="#6b6660" stroke-width="1" stroke-dasharray="3 3" opacity="0.55"/>
</svg>
</div><figcaption>Different stakeholders need different views — no single diagram serves them all. Each viewpoint filters out everything not relevant to one specific concern.</figcaption></figure>

### Why one diagram is never enough

A single diagram cannot serve every audience. A developer writing code needs to see software components. An operator deploying the system needs to see machines and networks. A safety engineer needs to see failure modes. A building manager needs to see which components are responsible for regulatory compliance.

If one diagram tries to show everything at once, it becomes unreadable — a bowl of spaghetti with too many arrows. If it focuses on one perspective, it hides everything else. The solution is multiple diagrams, each tailored to one specific concern.

A useful image: a house seen by different people. The future home-owner needs a floor plan to imagine living there. The electrician needs a wiring diagram. The plumber needs a pipe diagram. The structural engineer needs a load-bearing-walls diagram. All four describe the same house, but each one filters out everything that doesn't matter for that audience. Trying to draw all four on one sheet of paper produces a mess.

The same idea applied to software systems is called **architecture viewpoints**. The concept is formalised in the international standard IEEE 42010, in Kruchten's influential 4+1 View Model (1995), and in enterprise frameworks like ArchiMate.

### What a viewpoint actually is

A **viewpoint** is a way of looking at the system that filters out everything not relevant to one stakeholder concern. Each viewpoint is a lens. Each catches design errors the others miss.

A useful image: a subway map and a street map of the same city. The subway map shows lines and stops — everything else is suppressed. The street map shows streets and addresses — the subway is barely visible. Both are correct. Neither one is enough on its own. A tourist needs both.

### ArchiMate's three layers

The ArchiMate framework organises architecture across three layers. Each layer answers a different stakeholder question.

| Layer | Concern | Building control examples |
|---|---|---|
| Business | Processes, actors, goals, regulations | The building manager monitors safety; the building must comply with BBR fire code |
| Application | Software components, data flows, interfaces | An AI agent, an anomaly detector, a data pipeline, the BuildSim API client |
| Technology | Infrastructure, devices, containers, networks | Docker containers on an edge server, a GPU server, an MQTT broker, a TimescaleDB instance |

The same system can be described at all three layers, and each description is useful to a different audience. A senior manager cares about the business layer. A developer cares about the application layer. A site reliability engineer cares about the technology layer.

### The five required viewpoints for a CPS architecture document

A complete architecture description for a building control system contains five viewpoints. Each answers a different question.

| Viewpoint | Question it answers | What it shows |
|---|---|---|
| Context | What interacts with the system? | The system as a single box, with users and external systems around it |
| Functional | What are the parts, and how do they connect? | All components inside, with their interfaces |
| Information | What data exists, and how does it flow? | Data models, storage, transformations |
| Behavioral | What happens when a specific event occurs? | The sequence of messages for one scenario |
| Deployment | What runs where? | Containers, hardware, network topology |

Each viewpoint is the lens that catches a different kind of mistake. A correct functional view can hide a broken deployment view — two components that look connected on paper might actually require a network link that does not exist. A clean data flow can mask a behavioural problem — a sequence diagram of a specific scenario can reveal a timing issue that no static diagram exposes. A complete business view can expose a compliance requirement that no component has been assigned to satisfy.

A useful image: five different home inspections of the same house. The structural inspector checks the foundation. The electrical inspector checks the wiring. The plumbing inspector checks the pipes. The roof inspector checks for leaks. The pest inspector checks for termites. Each one is looking for something different. Skipping any of them means certain defects will only be discovered after you move in.

---

## Part 5 — Modeling Notations

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 300" role="img" aria-label="Modeling notations spectrum" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Three notations along a spectrum of formality</text>
<!-- Whiteboard panel -->
<g transform="translate(70,60)">
  <rect width="130" height="135" rx="8" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <rect x="14" y="16" width="102" height="62" rx="3" fill="#fdfbf7" stroke="#6b6660" stroke-width="1" stroke-dasharray="3 2"/>
  <path d="M 28 38 Q 38 28 50 38 Q 62 48 75 32" fill="none" stroke="#2a2622" stroke-width="1.4" stroke-linecap="round"/>
  <circle cx="32" cy="62" r="7" fill="none" stroke="#2a2622" stroke-width="1.2"/>
  <line x1="39" y1="62" x2="58" y2="62" stroke="#2a2622" stroke-width="1.2"/>
  <rect x="58" y="54" width="22" height="16" fill="none" stroke="#2a2622" stroke-width="1.2"/>
  <line x1="92" y1="22" x2="102" y2="32" stroke="#8b3a1f" stroke-width="1.2"/>
  <line x1="102" y1="22" x2="92" y2="32" stroke="#8b3a1f" stroke-width="1.2"/>
  <text x="65" y="105" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Whiteboard</text>
  <text x="65" y="120" text-anchor="middle" font-size="10" fill="#6b6660">brainstorming</text>
</g>
<!-- C4 + Mermaid (highlighted) -->
<g transform="translate(295,55)">
  <rect width="130" height="145" rx="8" fill="#f3ece6" stroke="#8b3a1f" stroke-width="2"/>
  <rect x="14" y="14" width="102" height="20" rx="3" fill="#fdfbf7" stroke="#8b3a1f" stroke-width="1"/>
  <text x="65" y="28" text-anchor="middle" font-size="9" fill="#8b3a1f">Context</text>
  <rect x="14" y="38" width="48" height="20" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <rect x="68" y="38" width="48" height="20" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <line x1="62" y1="48" x2="68" y2="48" stroke="#6b6660" stroke-width="1"/>
  <text x="65" y="74" text-anchor="middle" font-size="8" fill="#6b6660">Containers</text>
  <rect x="14" y="80" width="30" height="14" rx="2" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <rect x="50" y="80" width="30" height="14" rx="2" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <rect x="86" y="80" width="30" height="14" rx="2" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <text x="65" y="115" text-anchor="middle" font-size="11" font-weight="600" fill="#8b3a1f">C4 + Mermaid</text>
  <text x="65" y="130" text-anchor="middle" font-size="10" fill="#8b3a1f">project documentation</text>
</g>
<!-- UML / SysML -->
<g transform="translate(520,60)">
  <rect width="130" height="135" rx="8" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
  <rect x="14" y="14" width="46" height="36" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <line x1="14" y1="24" x2="60" y2="24" stroke="#2a2622" stroke-width="1"/>
  <line x1="14" y1="34" x2="60" y2="34" stroke="#2a2622" stroke-width="1"/>
  <rect x="70" y="14" width="46" height="36" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <line x1="70" y1="24" x2="116" y2="24" stroke="#2a2622" stroke-width="1"/>
  <line x1="70" y1="34" x2="116" y2="34" stroke="#2a2622" stroke-width="1"/>
  <line x1="60" y1="32" x2="70" y2="32" stroke="#2a2622" stroke-width="1"/>
  <polygon points="70,28 76,32 70,36" fill="none" stroke="#2a2622" stroke-width="1"/>
  <rect x="14" y="58" width="46" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <line x1="14" y1="66" x2="60" y2="66" stroke="#2a2622" stroke-width="1"/>
  <rect x="70" y="58" width="46" height="22" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <line x1="70" y1="66" x2="116" y2="66" stroke="#2a2622" stroke-width="1"/>
  <text x="65" y="105" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">UML / SysML</text>
  <text x="65" y="120" text-anchor="middle" font-size="10" fill="#6b6660">safety-critical, large teams</text>
</g>
<!-- Spectrum axis -->
<line x1="70" y1="235" x2="650" y2="235" stroke="#6b6660" stroke-width="1.5"/>
<polygon points="650,235 642,231 642,239" fill="#6b6660"/>
<text x="70" y="252" text-anchor="start" font-size="10" fill="#6b6660">lower overhead</text>
<text x="650" y="252" text-anchor="end" font-size="10" fill="#6b6660">higher precision</text>
<circle cx="135" cy="235" r="5" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
<circle cx="360" cy="235" r="7" fill="#8b3a1f" stroke="#2a2622" stroke-width="1.5"/>
<circle cx="585" cy="235" r="5" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
<text x="360" y="285" text-anchor="middle" font-size="10" font-style="italic" fill="#6b6660">The middle option earns its keep by living in the repository as text, beside the code.</text>
</svg>
</div><figcaption>Three notations along a spectrum from informal to industrial. C4 with Mermaid is the chosen middle ground — enough precision to be useful, light enough to stay in version control.</figcaption></figure>

Picking the right notation is partly about precision and partly about overhead. Different tools sit at different points on that trade-off.

### Whiteboard sketches

A pen on a whiteboard is the right tool for early brainstorming. No formalism, no syntax to check, no learning curve. The trade-off is zero precision and zero versioning. Sketches are excellent for thinking; they are inadequate as the final architecture document.

A useful image: a napkin sketch over coffee captures an idea. It does not survive being passed across the office.

### UML and SysML

The **Unified Modeling Language**, abbreviated UML, has been the industry standard for modelling software systems since the late 1990s. It defines fourteen diagram types: class diagrams, sequence diagrams, state machine diagrams, activity diagrams, deployment diagrams, and many more.

**SysML** extends UML with diagrams for systems engineering — requirements traceability, physical constraints, hardware-software integration. UML and SysML are used in aerospace, defence, and automotive industries with heavyweight tools like Cameo Systems Modeler and Eclipse Papyrus.

Both have a steep learning curve and heavy tooling. For a small team working over a few weeks, the formalism adds more overhead than value. The strengths of UML and SysML — strict typing, simulation, automated code generation — pay off for a 200-engineer Airbus programme, not for a course project.

A useful image: UML is an industrial CNC machine. SysML is a CNC machine with extra accessories. Both are remarkable. Neither is what you want for a weekend furniture project, where a screwdriver and a level get the job done.

### The C4 model

The **C4 model** was proposed by Simon Brown specifically because UML felt too complex for most teams. C4 has exactly four levels of abstraction.

| Level | What it shows |
|---|---|
| 1. Context | The system as a single box surrounded by users and external systems |
| 2. Containers | The deployable units inside (processes, Docker containers, databases) |
| 3. Components | The internal structure of one container |
| 4. Code | Class-level detail (rarely drawn manually) |

C4 is technology-agnostic and focuses on structure and relationships. Levels 1 and 2 are sufficient for most software projects. Level 3 is useful when one container is complex enough to warrant its own diagram, such as an AI agent with multiple internal subsystems.

A useful image: C4 is Google Maps. Zoom out to see the country (Context). Zoom in to see the city (Containers). Zoom in further to see the streets (Components). Zoom all the way in to see individual buildings (Code). Same map, four levels of detail, each appropriate to a different question.

### Diagramming tools

The recommended diagramming tool for this course is **Mermaid**. Diagrams are written as plain text inside Markdown files, versioned with git alongside the code, and render automatically when the Markdown is viewed on GitHub.

The biggest advantage of Mermaid is that the diagrams live in the same repository as the code, get reviewed in the same pull requests, and stay synchronised with the system. A diagram pasted as a PNG into a Word document drifts out of sync within a week.

A useful image: a recipe taped to the inside of the kitchen cupboard, edited as the cook refines the recipe. The dish stays the recipe. The recipe stays the dish.

Other tools have specific niches. **draw.io** is open and free for more complex visual layouts. **Excalidraw** produces hand-drawn-looking diagrams useful for sketches and brainstorming. Neither matches Mermaid's advantage of living inside the repository as text.

### Comparing the choices

| Aspect | Whiteboard | C4 + Mermaid | UML / SysML |
|---|---|---|---|
| Learning curve | None | About an hour | Days to weeks |
| Precision | Low | Medium | High |
| Tooling | Pen | Markdown editor | Cameo, Papyrus |
| Best for | Brainstorming | Project documentation | Safety-critical, large teams |

For this course, C4 with Mermaid is the chosen middle ground.

---

## Part 6 — The Five Viewpoints in Practice

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 340" role="img" aria-label="The five viewpoints in practice" class="dgm">
<text x="360" y="32" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Five viewpoints, five questions</text>
  <g transform="translate(10,80)">
    <rect width="140" height="180" rx="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(0,30)"><rect x="46" y="22" width="48" height="36" rx="4" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
         <circle cx="22" cy="40" r="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
         <circle cx="118" cy="40" r="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
         <line x1="28" y1="40" x2="46" y2="40" stroke="#2a2622" stroke-width="1.2"/>
         <line x1="94" y1="40" x2="112" y2="40" stroke="#2a2622" stroke-width="1.2"/></g>
    <text x="70" y="138" text-anchor="middle" font-size="13" font-weight="600" fill="#2a2622">Context</text>
    <text x="70" y="156" text-anchor="middle" font-size="10.5" fill="#6b6660" font-style="italic">What touches it?</text>
  </g>
  <g transform="translate(150,80)">
    <rect width="140" height="180" rx="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(0,30)"><rect x="14" y="22" width="28" height="20" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
         <rect x="56" y="22" width="28" height="20" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
         <rect x="98" y="22" width="28" height="20" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
         <rect x="35" y="50" width="28" height="20" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
         <rect x="77" y="50" width="28" height="20" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
         <line x1="28" y1="42" x2="49" y2="50" stroke="#2a2622" stroke-width="1"/>
         <line x1="70" y1="42" x2="49" y2="50" stroke="#2a2622" stroke-width="1"/>
         <line x1="70" y1="42" x2="91" y2="50" stroke="#2a2622" stroke-width="1"/>
         <line x1="112" y1="42" x2="91" y2="50" stroke="#2a2622" stroke-width="1"/></g>
    <text x="70" y="138" text-anchor="middle" font-size="13" font-weight="600" fill="#2a2622">Functional</text>
    <text x="70" y="156" text-anchor="middle" font-size="10.5" fill="#6b6660" font-style="italic">What are the parts?</text>
  </g>
  <g transform="translate(290,80)">
    <rect width="140" height="180" rx="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(0,30)"><g>
    <ellipse cx="70" cy="26" rx="23" ry="6" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <path d="M 47 26 L 47 66 Q 70 74 93 66 L 93 26" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="70" cy="26" rx="23" ry="6" fill="none" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="70" cy="36" rx="21" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
    <ellipse cx="70" cy="48" rx="20" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
  </g>
         <path d="M 20 70 Q 50 90 120 70" fill="none" stroke="#8b3a1f" stroke-width="1.5"/>
         <polygon points="120,70 113,67 116,75" fill="#8b3a1f"/></g>
    <text x="70" y="138" text-anchor="middle" font-size="13" font-weight="600" fill="#2a2622">Information</text>
    <text x="70" y="156" text-anchor="middle" font-size="10.5" fill="#6b6660" font-style="italic">What data flows?</text>
  </g>
  <g transform="translate(430,80)">
    <rect width="140" height="180" rx="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(0,30)"><circle cx="20" cy="32" r="5" fill="#8b3a1f"/>
         <circle cx="20" cy="60" r="5" fill="#8b3a1f"/>
         <line x1="20" y1="22" x2="20" y2="78" stroke="#2a2622" stroke-width="1" stroke-dasharray="2 2"/>
         <line x1="120" y1="22" x2="120" y2="78" stroke="#2a2622" stroke-width="1" stroke-dasharray="2 2"/>
         <line x1="20" y1="36" x2="120" y2="42" stroke="#2a2622" stroke-width="1.2"/>
         <polygon points="120,42 113,40 115,46" fill="#2a2622"/>
         <line x1="120" y1="56" x2="20" y2="62" stroke="#2a2622" stroke-width="1.2"/>
         <polygon points="20,62 27,60 25,66" fill="#2a2622"/></g>
    <text x="70" y="138" text-anchor="middle" font-size="13" font-weight="600" fill="#2a2622">Behavioural</text>
    <text x="70" y="156" text-anchor="middle" font-size="10.5" fill="#6b6660" font-style="italic">What happens in order?</text>
  </g>
  <g transform="translate(570,80)">
    <rect width="140" height="180" rx="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(0,30)"><rect x="14" y="20" width="50" height="56" rx="3" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
         <rect x="76" y="20" width="50" height="56" rx="3" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
         <line x1="22" y1="34" x2="56" y2="34" stroke="#2a2622" stroke-width="0.8"/>
         <line x1="22" y1="44" x2="50" y2="44" stroke="#2a2622" stroke-width="0.8"/>
         <line x1="22" y1="54" x2="56" y2="54" stroke="#2a2622" stroke-width="0.8"/>
         <line x1="84" y1="34" x2="118" y2="34" stroke="#2a2622" stroke-width="0.8"/>
         <line x1="84" y1="44" x2="112" y2="44" stroke="#2a2622" stroke-width="0.8"/>
         <line x1="84" y1="54" x2="118" y2="54" stroke="#2a2622" stroke-width="0.8"/></g>
    <text x="70" y="138" text-anchor="middle" font-size="13" font-weight="600" fill="#2a2622">Deployment</text>
    <text x="70" y="156" text-anchor="middle" font-size="10.5" fill="#6b6660" font-style="italic">What runs where?</text>
  </g>
</svg>
</div><figcaption>Each viewpoint asks a different question — and each catches a kind of design error the others would hide.</figcaption></figure>

Each viewpoint introduced in Part 4 looks like something specific when drawn. The following examples apply each viewpoint to building control.

### Context view (C4 Level 1)

The context diagram shows the entire system as a single box, surrounded by every person and every external system it interacts with. Nothing internal is shown.

```
   ┌───────────────────┐                ┌──────────────────┐
   │ Building Manager  │── monitors ──► │                  │
   └───────────────────┘                │  AUTONOMOUS      │
                                        │  BUILDING        │
   ┌───────────────────┐                │  CONTROL         │
   │ Building Occupants│── presence ──► │  SYSTEM          │
   └───────────────────┘                │                  │
                                        └────┬─────────────┘
                                             │
                              reads / commands│   LLM
                                             │   inference
                                             ▼
                                    ┌─────────────┐    ┌─────────────┐
                                    │   BuildSim  │    │ GPU Server  │
                                    │   Server    │    │ (LLM)       │
                                    └─────────────┘    └─────────────┘
```

The diagram answers one question only — who or what touches the system from outside. A useful image: a country on a globe, with arrows showing where its trade goes. Nothing about cities is shown yet.

### Functional view (C4 Level 2)

The container diagram opens up the single box from the context view. Inside are the deployable units — each box typically maps to one Docker container.

```
   ┌───────────────────────────────────────────────────────────┐
   │ FIRE DETECTION SYSTEM                                     │
   │                                                           │
   │   Fire Simulator ──drives──► Smoke Sensors, Temp Sensors  │
   │                                       │                   │
   │                                  POST │ values            │
   │                                       ▼                   │
   │   BuildSim API                                            │
   │                                       │ store             │
   │                                       ▼                   │
   │                              InfluxDB ◄──── query ───  Anomaly Detector
   │                                                              │ alerts
   │                                                              ▼
   │                                                          Safety Agent ──reasoning──► GPU
   │                                                              │ commands
   │                                                              ▼
   │                                                      Sprinkler Actuator ──PUT──► BuildSim
   └───────────────────────────────────────────────────────────┘
```

Two different use cases lead to two different container diagrams. A fire-detection system needs sensors, an anomaly detector, a safety agent, sprinkler actuators, and an LLM. An HVAC optimisation system needs temperature and occupancy sensors, an MQTT broker, a time-series database, a forecasting model, an HVAC controller, and HVAC actuators. The same physical building, the same BuildSim, but two different functional decompositions.

A useful image: zooming into the country on the globe and seeing its cities. Each city is one container. Lines between them are roads — the interfaces.

### Component view (C4 Level 3)

When one container is complex enough, it gets its own diagram showing its internal structure. A safety agent might contain a BuildSim API client, a set of tool definitions, an agent memory, a reasoning chain implementing the ReAct pattern, and a safety guardrail module. Each is internal to the agent and invisible from the outside.

A useful image: zooming further into one city and seeing its streets. The streets are not visible at the city-level view, but they exist.

### Behavioral view (sequence diagram)

A **sequence diagram** is a comic strip with software components as the characters. It shows one specific scenario as a series of messages, in order.

```
   Fire simulator    ►  Smoke sensor:        smoke = 0.3 (rising)
   Smoke sensor      ►  BuildSim:            PUT sensor value
   Smoke sensor      ►  InfluxDB:            INSERT reading
   Anomaly detector  ►  InfluxDB:            SELECT recent readings
   Anomaly detector                          autoencoder flags anomaly
   Anomaly detector  ►  Safety agent:        alert in A2306
   Safety agent      ►  GPU (LLM):           "smoke anomaly, what action?"
   GPU (LLM)         ►  Safety agent:        "activate sprinkler, unlock fire doors"
   Safety agent      ►  Sprinkler actuator:  Sprinklers ON
   Sprinkler actuator►  BuildSim:            PUT actuator state = on
   Fire simulator                            sprinkler ON, smoke decreasing
```

Each line is one frame of the comic. The order of the lines is the order of events. Sequence diagrams are the right tool when the question is "what happens, in what order, for this scenario?" They expose timing issues that no static component diagram can.

A useful image: a film strip. Each frame is one moment in time. Played in order, the strip tells a story.

### State machine diagram

A **state machine** shows every state a component can be in and every event that moves it between states. The lifecycle of a sensor process might look like this:

```
   [start]
      │
      ▼
   Starting ────process starts────► Registering
                                         │
                                         ├──registered──► Running
                                         │                  │
                                         │                  ├──push reading every 5s──► (Running)
                                         │                  │
                                         │                  ├──connection lost────────► Reconnecting
                                         │                  │
                                         │                  └──SIGTERM────────────────► Stopped
                                         │
                                         └──BuildSim unavailable──► RetryRegister
                                                                          │
                                                                          └──wait 5s──► Registering
   Stopped ──► [end]
```

A useful image: a traffic light. It has states (red, yellow, green) and events that flip it (timer expires, pedestrian button pressed). The state machine is the rulebook for the lamp.

Every transition in the diagram is a piece of code that handles a specific event. If the diagram is complete, the implementation can be written almost mechanically from it.

### Information view (requirements table)

The information viewpoint is sometimes a table rather than a diagram. The requirements table lists every requirement with its ID, type, statement, priority, and acceptance criterion.

| ID | Type | Requirement | Priority | Acceptance criterion |
|---|---|---|---|---|
| FR-01 | Functional | Detect fire within 30 seconds | Must | Anomaly detector flags within 30 s |
| FR-02 | Functional | Activate sprinklers in affected rooms | Must | Actuator state changes to "on" |
| FR-03 | Functional | Compute evacuation routes avoiding fire | Must | Route excludes fire rooms |
| NFR-01 | Non-functional | Recover from sensor crash within 60 s | Must | New reading within 60 s of kill |
| NFR-02 | Non-functional | False positive rate below 5 % | Should | Evaluated on 24 h of normal data |
| REG-01 | Regulatory | Fire doors close per BBR timing | Must | Door responds within 5 s |

Every requirement traces to a test case. FR-01 maps to a test called `test_fire_detection_latency`.

A useful image: an invoice with line items. Each line is a deliverable. The total at the bottom is the contract.

---

## Part 7 — The Architecture Document

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 320" role="img" aria-label="Contents of the architecture document" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">The architecture document is the contract between design and code</text>
<!-- Central document -->
<g transform="translate(290,55)">
  <rect x="10" y="10" width="140" height="80" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <rect x="5" y="5" width="140" height="80" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <rect x="0" y="0" width="140" height="80" rx="3" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
  <text x="70" y="20" text-anchor="middle" font-size="9.5" font-weight="700" fill="#2a2622">architecture.md</text>
  <line x1="14" y1="32" x2="126" y2="32" stroke="#6b6660" stroke-width="0.8"/>
  <line x1="14" y1="42" x2="126" y2="42" stroke="#6b6660" stroke-width="0.8"/>
  <line x1="14" y1="52" x2="100" y2="52" stroke="#6b6660" stroke-width="0.8"/>
  <line x1="14" y1="62" x2="126" y2="62" stroke="#6b6660" stroke-width="0.8"/>
  <line x1="14" y1="72" x2="80" y2="72" stroke="#6b6660" stroke-width="0.8"/>
</g>
<!-- connector lines from doc bottom to each panel top -->
<line x1="320" y1="145" x2="105" y2="175" stroke="#8b3a1f" stroke-width="1.2" stroke-linecap="round" opacity="0.8"/>
<line x1="345" y1="145" x2="275" y2="175" stroke="#8b3a1f" stroke-width="1.2" stroke-linecap="round" opacity="0.8"/>
<line x1="375" y1="145" x2="445" y2="175" stroke="#8b3a1f" stroke-width="1.2" stroke-linecap="round" opacity="0.8"/>
<line x1="400" y1="145" x2="615" y2="175" stroke="#8b3a1f" stroke-width="1.2" stroke-linecap="round" opacity="0.8"/>
<!-- Panel 1: Diagrams -->
<g transform="translate(30,175)">
  <rect width="150" height="110" rx="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
  <rect x="50" y="14" width="50" height="20" rx="2" fill="#f3ece6" stroke="#2a2622" stroke-width="1"/>
  <circle cx="22" cy="24" r="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <circle cx="128" cy="24" r="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <line x1="28" y1="24" x2="50" y2="24" stroke="#6b6660" stroke-width="1"/>
  <line x1="100" y1="24" x2="122" y2="24" stroke="#6b6660" stroke-width="1"/>
  <rect x="20" y="46" width="30" height="16" rx="2" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <rect x="60" y="46" width="30" height="16" rx="2" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <rect x="100" y="46" width="30" height="16" rx="2" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <text x="75" y="85" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Diagrams</text>
  <text x="75" y="100" text-anchor="middle" font-size="9.5" fill="#6b6660">C4 · sequence · deployment</text>
</g>
<!-- Panel 2: Specifications -->
<g transform="translate(200,175)">
  <rect width="150" height="110" rx="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
  <rect x="15" y="14" width="120" height="50" rx="3" fill="#f4ead9" stroke="#6b6660" stroke-width="0.8"/>
  <text x="22" y="27" font-size="7.5" font-family="ui-monospace, monospace" fill="#2a2622">{ "temp": "float",</text>
  <text x="22" y="39" font-size="7.5" font-family="ui-monospace, monospace" fill="#2a2622">  "unit": "°C",</text>
  <text x="22" y="51" font-size="7.5" font-family="ui-monospace, monospace" fill="#2a2622">  "ts": "ISO-8601" }</text>
  <text x="75" y="85" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Specifications</text>
  <text x="75" y="100" text-anchor="middle" font-size="9.5" fill="#6b6660">schemas · APIs · state machines</text>
</g>
<!-- Panel 3: Test plan -->
<g transform="translate(370,175)">
  <rect width="150" height="110" rx="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
  <g transform="translate(22,18)">
    <rect width="11" height="11" rx="2" fill="#e2ebde" stroke="#2a2622" stroke-width="1"/>
    <path d="M 2.5 5.5 L 5 8 L 8.5 3" stroke="#2a2622" stroke-width="1.4" fill="none" stroke-linecap="round" stroke-linejoin="round"/>
    <line x1="18" y1="6" x2="100" y2="6" stroke="#6b6660" stroke-width="1"/>
  </g>
  <g transform="translate(22,34)">
    <rect width="11" height="11" rx="2" fill="#e2ebde" stroke="#2a2622" stroke-width="1"/>
    <path d="M 2.5 5.5 L 5 8 L 8.5 3" stroke="#2a2622" stroke-width="1.4" fill="none" stroke-linecap="round" stroke-linejoin="round"/>
    <line x1="18" y1="6" x2="90" y2="6" stroke="#6b6660" stroke-width="1"/>
  </g>
  <g transform="translate(22,50)">
    <rect width="11" height="11" rx="2" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
    <line x1="18" y1="6" x2="100" y2="6" stroke="#6b6660" stroke-width="1"/>
  </g>
  <text x="75" y="85" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Test plan</text>
  <text x="75" y="100" text-anchor="middle" font-size="9.5" fill="#6b6660">one row per requirement</text>
</g>
<!-- Panel 4: Repository -->
<g transform="translate(540,175)">
  <rect width="150" height="110" rx="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.2"/>
  <g font-family="ui-monospace, monospace" font-size="8" fill="#2a2622">
    <text x="14" y="22">project/</text>
    <text x="22" y="34">├ docs/</text>
    <text x="22" y="46">├ sensor-process/</text>
    <text x="22" y="58">├ ai-agent/</text>
    <text x="22" y="70">└ docker-compose.yml</text>
  </g>
  <text x="75" y="85" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Repository</text>
  <text x="75" y="100" text-anchor="middle" font-size="9.5" fill="#6b6660">one folder per container</text>
</g>
<text x="360" y="307" text-anchor="middle" font-size="10" font-style="italic" fill="#6b6660">Approved before any code is written — the design that the implementation must satisfy.</text>
</svg>
</div><figcaption>The architecture document carries diagrams, specifications, a test plan, and a repository structure — each binding the design to the running system.</figcaption></figure>

The architecture document is the contract between design and implementation. It is reviewed and approved before any code is written. It contains the five viewpoints described above, plus the specifications below.

### Diagrams and tables

The document includes:

- A **context diagram** (C4 Level 1) showing the system boundary, users, and external systems.
- A **container diagram** (C4 Level 2) showing every deployable unit and how they connect.
- A **requirements table** mapping each requirement to a priority, acceptance criterion, and test case.
- A **data flow diagram** showing how sensor data moves from measurement through storage to decision.
- One or more **sequence diagrams** for the key scenarios.
- A **deployment diagram** showing what runs where.

### Specifications

Beyond the diagrams, the design specification covers four additional artefacts.

**Data models.** JSON schemas for every message that crosses a boundary between components. Each schema specifies field names, types, units, and constraints. A useful image: a customs declaration form. Every field on the form has a defined meaning, and every package crossing the border must complete it correctly.

**API contracts.** Every REST endpoint and every MQTT topic in the system, with request and response shapes, status codes, QoS levels, and retention semantics. The plug-and-socket of the design — the prongs that have to match between components.

**State machines.** For the AI agent and for every component whose lifecycle matters, a state machine specifying every state and every transition. Especially important for error handling — the state machine is where you specify what happens when the network drops, when a sensor crashes, when a command fails.

**ML model specifications.** For each model, the input feature vector, the output, the source of training data, the training procedure, and the evaluation metrics that determine whether the model is good enough to deploy.

### Test plan

The test plan is written **before** implementation, not after. It links each requirement to one or more test cases. Each test case describes the initial state of the system, the stimulus that drives the test, the expected response, and the pass-or-fail criteria.

A useful image: a wedding-planner's checklist. Every item on the wishlist has a corresponding "yes/no — was it delivered?" tick. The wedding cannot be declared successful until every box is ticked.

### Repository structure

A consistent repository structure keeps the design and the code in sync. One folder per deployable container, plus a docs folder for the architecture document, plus a top-level docker-compose file that brings everything up.

```
project-root/
├── docs/
│   ├── architecture.md
│   ├── requirements.md
│   └── test-plan.md
├── sensor-process/
│   ├── Dockerfile
│   └── src/
├── ai-agent/
│   ├── Dockerfile
│   └── src/
├── actuator-process/
│   ├── Dockerfile
│   └── src/
├── docker-compose.yml
└── README.md
```

The structure mirrors the C4 container diagram. If the diagram shows five containers and the repository has one folder with everything jammed together, the design and the code have parted ways.

A useful image: the kitchen of a restaurant. Each station (grill, salad, dessert) has its own workspace, equipment, and chef. The menu (the design) is reflected in the kitchen layout (the code structure). A messy kitchen means slow, error-prone service.

---

## Part 8 — A Worked Example: Fire Detection End-to-End

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 320" role="img" aria-label="Fire detection — end-to-end pipeline" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Fire detection: from smoke to sprinkler</text>
<!-- chain of components -->
  <g transform="translate(30,80)">
    <rect width="100" height="110" rx="10" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(50,42)"><g transform="translate(0,0)">
    <circle cx="0" cy="0" r="14" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="4.199999999999999" fill="#8b3a1f"/>
    <path d="M -11.2 -14 Q -16.799999999999997 -19.599999999999998 -11.2 -25.2" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 11.2 -14 Q 16.799999999999997 -19.599999999999998 11.2 -25.2" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g></g>
    <text x="50" y="92" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Smoke sensor</text>
    <text x="50" y="106" text-anchor="middle" font-size="9.5" fill="#6b6660">PUBLISH</text>
  </g>
  <g transform="translate(166,80)">
    <rect width="100" height="110" rx="10" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="22" y="32" width="56" height="36" rx="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
         <text x="50" y="56" text-anchor="middle" font-size="10" font-weight="700" fill="#2a2622">MQTT</text>
    <text x="50" y="92" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">MQTT broker</text>
    <text x="50" y="106" text-anchor="middle" font-size="9.5" fill="#6b6660">topic: sensors/#</text>
  </g>
  <g transform="translate(302,80)">
    <rect width="100" height="110" rx="10" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(50,52)"><g>
    <ellipse cx="0" cy="-17" rx="20" ry="6" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <path d="M -20 -17 L -20 17 Q 0 25 20 17 L 20 -17" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="0" cy="-17" rx="20" ry="6" fill="none" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="0" cy="-7" rx="18" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
    <ellipse cx="0" cy="5" rx="17" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
  </g></g>
    <text x="50" y="92" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Time-series DB</text>
    <text x="50" y="106" text-anchor="middle" font-size="9.5" fill="#6b6660">INSERT</text>
  </g>
  <g transform="translate(438,80)">
    <rect width="100" height="110" rx="10" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(50,52)">
        <polyline points="-20,-8 -10,4 0,-6 10,8 20,-14 30,2" fill="none" stroke="#8b3a1f" stroke-width="2"/>
        <circle cx="20" cy="-14" r="4" fill="#8b3a1f"/>
        </g>
    <text x="50" y="92" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Anomaly model</text>
    <text x="50" y="106" text-anchor="middle" font-size="9.5" fill="#6b6660">detect spike</text>
  </g>
  <g transform="translate(574,80)">
    <rect width="100" height="110" rx="10" fill="#f0d9d1" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(50,52)"><g><circle cx="0" cy="0" r="20" fill="#8b3a1f" opacity="0.9"/>
    <text x="0" y="5" text-anchor="middle" font-size="14" font-weight="700" fill="#fdfbf7">AI</text></g></g>
    <text x="50" y="92" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">LLM safety agent</text>
    <text x="50" y="106" text-anchor="middle" font-size="9.5" fill="#6b6660">reason · decide</text>
  </g>
<!-- arrows between -->
<line x1="130" y1="135" x2="159" y2="135" stroke="#8b3a1f" stroke-width="1.8" stroke-linecap="round"/>
<polygon points="166,135 159,139 159,131" fill="#8b3a1f"/><line x1="266" y1="135" x2="295" y2="135" stroke="#8b3a1f" stroke-width="1.8" stroke-linecap="round"/>
<polygon points="302,135 295,139 295,131" fill="#8b3a1f"/><line x1="402" y1="135" x2="431" y2="135" stroke="#8b3a1f" stroke-width="1.8" stroke-linecap="round"/>
<polygon points="438,135 431,139 431,131" fill="#8b3a1f"/><line x1="538" y1="135" x2="567" y2="135" stroke="#8b3a1f" stroke-width="1.8" stroke-linecap="round"/>
<polygon points="574,135 567,139 567,131" fill="#8b3a1f"/>
<!-- loop back to building/actuator -->
<path d="M 624 200 Q 624 250 510 250 L 110 250 Q 70 250 70 200" fill="none" stroke="#8b3a1f" stroke-width="2" stroke-dasharray="6 4"/>
<polygon points="70,200 66,210 74,210" fill="#8b3a1f"/>
<text x="360" y="270" text-anchor="middle" font-size="11" font-weight="600" fill="#8b3a1f">command sprinklers → smoke drops → next loop</text>
<text x="360" y="290" text-anchor="middle" font-size="10" font-style="italic" fill="#6b6660">close the loop · the world responds · sensor reads again</text>
</svg>
</div><figcaption>Every component of the fire-detection scenario, in the order they participate. The dashed arrow closes the loop — the sprinkler changes the world, the next sensor reading reflects it.</figcaption></figure>

To make the abstract MBSE process concrete, follow a fire-detection system through every step.

### Step 1 — Requirements

The system must detect fire within 30 seconds, activate sprinklers in the affected rooms, compute evacuation routes that avoid the fire, recover from a sensor crash within 60 seconds, keep false positives below 5 percent, and ensure fire doors close within the timing specified by the Swedish BBR fire code. Each becomes a row in the requirements table with a unique ID, a priority, an acceptance criterion, and a target test case.

### Step 2 — Functional decomposition

"Detect fire" decomposes into eight operations: collect smoke, CO, and temperature readings continuously; validate the readings; compute features (rolling means, gradients); apply an anomaly model; escalate positive readings to a language-model agent for severity assessment; decide actions (sprinkler, fire-door, evacuation route); apply commands through actuator processes; log the decision and the outcome.

Each operation maps to a software component.

### Step 3 — Architecture (Container diagram)

The components: a fire simulator that drives the physical model, sensor processes that read from it, an MQTT broker that decouples sensors from consumers, a time-series database that stores history, an anomaly detector running the Isolation Forest model, a safety agent that calls the LLM and chooses actions, sprinkler and fire-door actuator processes, and a dashboard process that visualises building state. Each is one container; everything is described in a single `docker-compose.yml`.

### Step 4 — Interfaces

Every connection is specified exactly. Sensors publish to MQTT topics named `sensors/{level}/{room}/{type}` with the payload `{ts, sensor_id, room, level, type, unit, value}` at QoS 1. A consumer subscribes to `sensors/#` and writes each reading to the `readings` hypertable in TimescaleDB. The anomaly detector queries with `SELECT value FROM readings WHERE room = $1 AND type = 'smoke' AND ts > now() - INTERVAL '60 seconds'`. The agent issues actuator commands via `PUT /api/actuators/{id}/state` with body `{"state":"on"}`.

Nothing is left to interpretation.

### Step 5 — Behaviour

A sequence diagram walks through the fire scenario: smoke rises in A2306, the sensor publishes, the consumer inserts, the anomaly detector flags, the agent calls the LLM, the agent commands the sprinkler, the sprinkler updates both BuildSim and the simulator, and the next sensor cycle shows reduced smoke. A state machine for the sensor process specifies registration, normal operation, network-failure handling, and reconnection.

### Step 6 — Validation

Every requirement maps to a test:

| Requirement | Test | Pass criterion |
|---|---|---|
| FR-01 | `test_fire_detection_latency` | Less than 30 s from smoke spike to actuator command |
| FR-02 | `test_sprinkler_fires` | Actuator state changes to "on" within 5 s of detection |
| NFR-01 | `test_sensor_crash_recovery` | New reading within 60 s of killing the sensor |
| NFR-02 | `test_false_positives` | Below 5 percent on 24 hours of normal data |

Each test runs against the actual deployed system, not a mock.

The entire fire-detection system — requirements through validation — is captured in roughly twenty pages of structured artefacts: tables, diagrams, and JSON schemas. None of it depends on any reader interpreting prose.

---

## Part 9 — Vocabulary Reference

Every term used in this chapter, defined.

| Term | Definition |
|---|---|
| **CPS (Cyber-Physical System)** | A system in which computation and physical processes are tightly coupled through a continuous feedback loop |
| **Embedded intelligence** | Software that runs inside or beside the physical system it controls, not in a remote data centre |
| **Edge** | The portion of a network physically close to the data source; the opposite of cloud |
| **Sensor** | A component that measures a physical quantity and reports it as data |
| **Actuator** | A component that performs a physical action when commanded |
| **Feedback loop** | A control structure in which the output of a process influences the next input |
| **Automation** | Behaviour produced by a fixed rule that does not adapt to context |
| **Autonomy** | Behaviour that adapts to context using data, prediction, and learning |
| **MBSE (Model-Based Systems Engineering)** | A methodology that replaces prose documents with structured, precise models |
| **Requirement** | A testable statement of what the system must do |
| **Functional requirement** | A requirement about behaviour: what the system does |
| **Non-functional requirement** | A requirement about quality: how well the system does it |
| **Regulatory requirement** | A requirement imposed by an external standard or law |
| **Functional decomposition** | Breaking a high-level requirement into smaller operations that satisfy it |
| **Architecture** | The high-level structure of components and their relationships |
| **Interface** | A precise specification of how two components communicate |
| **Sequence diagram** | A diagram showing the messages exchanged between components over time, for one scenario |
| **State machine** | A diagram showing the discrete states of a component and the events that cause transitions |
| **Validation** | Demonstrating that a requirement has been met by a design element and a test |
| **Viewpoint** | A perspective on the system that filters out everything not relevant to one stakeholder concern |
| **IEEE 42010** | The international standard defining architecture descriptions through multiple viewpoints |
| **ArchiMate** | An enterprise architecture framework with three layers: business, application, technology |
| **4+1 View Model** | An influential set of architectural views proposed by Kruchten in 1995 |
| **C4 Model** | A simple architectural notation with four levels: Context, Containers, Components, Code |
| **UML (Unified Modeling Language)** | A standard modeling language with fourteen diagram types |
| **SysML** | An extension of UML for systems engineering, used in aerospace and defence |
| **Mermaid** | A text-based diagramming tool that renders inside Markdown |
| **Container (Docker)** | A packaged, isolated runtime that bundles a service with all its dependencies |
| **Docker Compose** | A tool for running multiple containers together, configured by a YAML file |
| **Architecture document** | The contract between design and implementation, produced by the MBSE process |

---

## Part 10 — Summary in Five Sentences

1. A building is a cyber-physical system in which a continuous feedback loop joins the physical world and the software that observes and changes it; the loop must be fast enough that the physics cannot run away, reliable enough to survive component failures, and correct enough to avoid harm.
2. The right place for the software's brain is the edge, close to the physical system, because latency, reliability, privacy, and bandwidth all favour local processing over remote.
3. Useful systems are autonomous, not merely automated: they reason about context, learn from data, and adapt over time rather than following a single fixed rule.
4. The path from idea to working code runs through Model-Based Systems Engineering, in which structured artefacts — requirements tables, architecture diagrams, interface specifications, behaviour models, and validation matrices — replace prose documents that cannot be analysed or tested.
5. A complete architecture description requires multiple viewpoints (context, functional, information, behavioural, deployment), because no single diagram can serve every stakeholder, and each viewpoint catches errors the others hide.

These five ideas are the foundation for everything that follows.
