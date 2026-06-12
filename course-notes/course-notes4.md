# Agentic AI for Autonomous Systems

Standalone notes for the fourth chapter of D7065E.

---

## Part 1 — From a Switch to a Thinker: Three Generations of Automation

<figure class="diagram">
<img src="figures/course-notes4-fig01.png" alt="Three generations of automation">
<figcaption><em>From fixed rules to ML models to LLM-driven agents — each step on the ladder adapts to more situations than the one before. The trade-off is always the same: more flexible behaviour, harder to predict.</em></figcaption>
</figure>

The history of building automation is the history of decisions getting smarter. Three generations of automation describe where the field has been and where it is now.

Before walking through the generations, it is worth fixing the central term. The word *agent* did not arrive with large language models; it has a four-decade history in artificial intelligence. The classical definition is a system situated in an environment that exhibits **autonomy** (it operates without direct intervention), **reactivity** (it perceives and responds to changes in its environment), **pro-activeness** (it takes initiative in pursuit of goals), and **social ability** (it interacts with other agents) ([Wooldridge & Jennings, 1995](#wooldridge1995)). This generalises into the notion of a *rational agent* — one that selects actions expected to maximise its performance measure given its percept history — specified with the **PEAS** framework (Performance measure, Environment, Actuators, Sensors) for describing an agent's task environment ([Russell & Norvig, 2021](#russell2021)). A building control agent maps onto PEAS directly: the performance measure combines energy cost, comfort, and safety; the environment is the building and its occupants; the actuators are HVAC units, dampers, and door locks; the sensors are the temperature, CO2, smoke, and occupancy instruments. What has changed recently is not the *definition* of an agent but the *reasoning engine* available to implement one. Keeping that distinction in mind prevents a common confusion: an LLM is not an agent; an LLM is one possible brain for an agent.

### Generation 1: Rule-based automation

The oldest and simplest. A rule says: `if temperature > 25°C then turn on the AC`. A programmable logic controller (a small industrial computer) can evaluate thousands of these rules per millisecond. The behaviour is deterministic, predictable, and easy to verify. Engineers love rules because rules don't surprise them.

Picture this: a thermostat with a dial. You set the dial, the thermostat compares the room temperature against the dial, and the heater clicks on or off. There is no thinking happening. There is no context. The thermostat doesn't know it's three in the morning, or that the building is empty, or that today is a public holiday. It only knows the number.

Rules work beautifully until reality stops fitting into the rules. A rule that turns on the AC at 25°C is correct on most days. It is wrong at 3 a.m. when nobody is in the building. It is wrong on a holiday when the office is closed. It is wrong when a faulty ventilation duct is causing the heat — running the AC will only make the duct fault worse. Real buildings have hundreds of interacting rules, each trying to cover an edge case, and the interactions between them become impossible to maintain.

### Generation 2: Machine-learning inference

A learned model replaces hand-crafted rules. Instead of `if temperature > 25°C`, the model takes a list of features — temperature, occupancy, time of day, day of week, weather forecast, recent HVAC commands — and outputs a prediction. The model is trained on thousands or millions of examples of how the building behaves, and it learns patterns no rule-writer would have thought to capture.

A useful image: a doctor who has seen ten thousand patients. The doctor doesn't follow a single rule like "if temperature > 38°C, prescribe antibiotics." They take in a constellation of symptoms — temperature, blood pressure, the patient's history, the time of year — and recognise the pattern. The diagnosis comes from pattern recognition, not from a flowchart.

ML inference is much better than rules at handling situations the engineer didn't anticipate. But it is still, fundamentally, a function: input features in, prediction out. The model does not reason about *why*. It cannot explain itself in plain language. It cannot handle situations far outside its training data.

This limitation has a precise name in the statistical learning literature: a trained model interpolates within the distribution of its training data, and its guarantees weaken sharply under *distribution shift* — when the deployed environment differs from the data the model was fitted on ([Russell & Norvig, 2021](#russell2021)). A model trained on a building's behaviour in winter may be confidently wrong in summer; a model trained before a renovation may be confidently wrong after it. The trade-off relative to rules is therefore symmetrical: rules are verifiable but brittle in the face of unanticipated cases, while learned models are flexible within their training distribution but opaque and unreliable outside it.

### Generation 3: Agentic AI

Agentic AI goes one step further. An **agent** observes the current situation, reasons about its context in natural language, plans a sequence of actions, calls tools to execute the plan, observes the outcome, and adjusts its approach. The brain of the agent is typically a large language model (an LLM).

The technical lineage behind this generation is short but worth knowing. The transformer architecture made it practical to train very large sequence models on text ([Vaswani et al., 2017](#vaswani2017)). Scaling those models up revealed an unexpected capability: **in-context learning**, the ability to perform a new task from a few examples or instructions placed in the prompt, without any retraining ([Brown et al., 2020](#brown2020)). Prompting a model to produce explicit intermediate reasoning steps — *chain-of-thought* prompting — then turned out to substantially improve performance on multi-step problems ([Wei et al., 2022](#wei2022)). Interleaving that reasoning with actions on an external environment closed the loop, producing the ReAct pattern that underlies most modern agents ([Yao et al., 2023](#yao2023)), and models can even learn *when and how* to call external tools ([Schick et al., 2023](#schick2023)). Each of these steps converts the LLM from a text generator into something that can participate in the perceive–reason–act cycle described long before LLMs existed ([Wooldridge & Jennings, 1995](#wooldridge1995)).

Think of it this way: instead of a doctor following a flowchart, imagine a senior consultant. The consultant asks the patient questions, weighs trade-offs out loud, considers what would happen if they chose treatment A versus treatment B, runs a test to gather more information, and explains the reasoning to the patient in plain English. The consultant might encounter a case they have never seen before and still produce a sensible plan by reasoning from first principles.

An LLM-based agent can:
- Handle a situation it has never encountered in training.
- Explain its decisions in human language.
- Ask for clarification when the input is ambiguous.
- Adjust its plan after observing the result of an action.

These capabilities come with a cost that must be stated just as clearly: an LLM agent is slower than either predecessor by several orders of magnitude, its behaviour is stochastic rather than deterministic, and its failure modes (Part 5) are qualitatively new. The flexibility–predictability trade-off in the figure above is not a slogan; it is the central engineering tension of this chapter.

These three generations are not in competition. A real building uses all three: hard-wired rules for the most safety-critical functions (because rules are fast and verifiable), ML models for real-time inference (because models are smarter than rules), and LLM agents for higher-level reasoning that needs context and explanation (because agents can think). The art is knowing which to use where.

---

## Part 2 — The Agent Loop

<figure class="diagram">
<img src="figures/course-notes4-fig02.png" alt="The agent loop">
<figcaption><em>The agent loop is the engine that turns a model into a behaviour. Perceive, think, act, observe — and round again.</em></figcaption>
</figure>

Every agent, no matter the framework or the model, runs the same loop. Six steps, repeated forever.

The loop is not an invention of the LLM era. It is a refinement of the classical *sense–plan–act* cycle from robotics and of the classical agent function, in which an agent maps a percept sequence to an action through some internal program ([Russell & Norvig, 2021](#russell2021)). What the six-step formulation adds is the explicit separation of *reasoning* from *planning* (useful when the reasoning is auditable natural language) and the explicit *update* step, which gives the agent state that persists across cycles. In the classical taxonomy, an agent without the update step is a *simple reflex agent* — it can only react to the current percept — while an agent that maintains and revises internal state is a *model-based* agent, and one that evaluates candidate actions against goals and utilities is a *goal-based* or *utility-based* agent ([Russell & Norvig, 2021](#russell2021)). The LLM agents in this course are, in those terms, utility-based agents whose internal model happens to be expressed in natural language.

<figure class="diagram">
<img src="figures/course-notes4-fig08.png" alt="The six-step agent loop">
<figcaption><em>Perceive, reason, plan, act, observe, update — and the update step is what makes the next cycle smarter than this one.</em></figcaption>
</figure>

A useful image: a chess player thinking through a move. They look at the board (perceive). They reason about what's threatening and what's possible. They plan a sequence — capture the knight, then move the rook. They make the first move (act). They watch what the opponent does (observe). They update their mental notes about the opponent's style (update). Then the loop starts again with the new board state.

### Perceive

The agent gathers fresh information about its environment. For a building agent, this means querying the time-series database for recent sensor readings, checking the dashboard's pending alerts, looking at what other agents have done recently.

An analogy: walking into a room and noticing the temperature, the people present, the open door, the smell of food. You're not deciding anything yet — you're just taking it in.

### Reason

The agent thinks through what the perceptions mean. With an LLM-based agent, this is literally the LLM generating a "Thought" — a short natural-language sentence explaining its interpretation.

Picture this: muttering to yourself as you walk through the room. "Hot in here. Three people. Door's open to the hot corridor — that's probably why. They're working, not just passing through — I don't want to interrupt them."

### Plan

The agent decides on a sequence of actions that should achieve its goal. The plan might be one action ("close the door") or several ("close the door, increase ventilation, send a notification").

A useful image: drawing up a small to-do list. "Step one: close the door. Step two: check if it helped in five minutes."

### Act

The agent executes the plan by calling tools — functions that affect the world. A tool might command an actuator, query a database, send an alert, or compute a route.

Think of it this way: actually walking over and closing the door. Up until now everything has been in your head. Now you put your hand on the handle.

### Observe

After acting, the agent checks what happened. Did the actuator state change? Did the sensor reading respond? Was the alert sent?

A useful image: stepping back and watching the temperature display for a minute. Did closing the door cool the room?

### Update

The agent writes the result of this loop into its memory or audit log, so the next loop has the benefit of this experience. Without the update step, every loop starts from a blank slate.

An analogy: jotting in your notebook, "closed the door in room A2306 at 14:32, room temperature dropped 1.2°C in five minutes, conclude that the corridor was indeed the heat source." Tomorrow, you'll skip the diagnostic and close the door faster.

The update step is the one that turns inference into learning. Many beginner agents skip it, and they become elaborate functions that don't actually get smarter over time. A well-designed agent always closes the loop with a structured note that the next perceive step can read.

Two practical design questions follow from the loop. First, *how often should it run?* A loop that runs every second burns tokens and money on cycles where nothing has changed; a loop that runs every hour misses events. The usual answer is event-driven triggering: the loop sleeps until an anomaly detector, a schedule, or a human request wakes it. Second, *what bounds the loop?* An unbounded reasoning loop can call tools forever — a runaway agent is a real failure mode, not a hypothetical one. Production agents therefore cap the number of tool calls per cycle and the wall-clock time per cycle, and treat hitting either cap as an error worth logging. Both questions are instances of the same trade-off introduced in Part 1: more deliberation buys better decisions at the price of latency and cost.

---

## Part 3 — Choosing the Right Tool: Rules vs ML vs Agents

A useful question to ask before any decision is "how fast must this be, and how much context does it need?" The answer points to one of three approaches.

| Decision type | Latency budget | Example | Best approach |
|---|---|---|---|
| Safety-critical | Under 100 ms | Fire suppression, emergency unlock | Hard-coded rule, possibly with ML on the edge |
| Real-time control | Under 1 second | Adjust temperature setpoint | Edge ML inference |
| Operational reasoning | 1 to 30 seconds | Optimise HVAC schedule for the next hour | LLM agent |
| Strategic planning | Minutes | Plan tomorrow's energy schedule | LLM agent + optimiser |
| Human-facing explanation | Interactive | Explain why heating changed | LLM |

Picture this: a hospital emergency department. The triage nurse follows a strict rule ("airway, breathing, circulation") because seconds count. The resident on call uses pattern recognition from training. The senior consultant reasons through complex cases. Three layers, three levels of cost and speed, each used where it fits.

A well-designed building control system uses all three layers in the same way:

- **Rules** for the fastest safety functions, where the cost of being slow is human harm and the cost of being wrong is catastrophic.
- **ML models** for real-time inference, where the situation is well-defined enough to train on and the latency budget rules out anything slower.
- **LLM agents** for higher-level reasoning, where context, trade-offs, and explanation matter more than raw speed.

The table is a heuristic, not a law, and the boundaries move with technology: local models (Part 5) push LLM reasoning into the one-second band, and learned controllers trained with reinforcement learning can in principle move sophisticated control into the real-time band ([Sutton & Barto, 2018](#sutton2018)). But the *structure* of the argument is stable. Every decision in a cyber-physical system has a latency budget set by the physics of the process and a context requirement set by the ambiguity of the situation, and the engineering task is to place each decision in the cheapest layer that satisfies both. Placing a decision too high wastes money and adds latency; placing it too low produces a brittle rule that fails on the edge cases the higher layers exist to handle.

### Pointers into the machine-learning toolbox

The course project requires a data-driven component — typically anomaly detection on sensor streams or short-horizon prediction — and the right move is almost never to invent a method. The techniques below are the standard, well-understood starting points, all covered by the standard survey of the anomaly-detection field ([Chandola et al., 2009](#chandola2009)), and all available as mature library implementations through scikit-learn ([Pedregosa et al., 2011](#pedregosa2011)) or the PyTorch/Keras deep-learning stacks. Use the library implementation, cite the original paper for the method you chose, and spend your effort on feature design, evaluation, and integration with the agent — not on re-implementing algorithms from scratch.

| Technique | Original paper | Library entry point | Typical building use |
|---|---|---|---|
| Isolation Forest | [Liu et al., 2008](#liu2008) | scikit-learn `IsolationForest` | Flagging anomalous sensor readings in multivariate streams |
| Local Outlier Factor | [Breunig et al., 2000](#breunig2000) | scikit-learn `LocalOutlierFactor` | Detecting readings that are abnormal relative to their local neighbourhood |
| One-class SVM | [Schölkopf et al., 2001](#scholkopf2001) | scikit-learn `OneClassSVM` | Learning the boundary of "normal" operation from fault-free data |
| Autoencoder | [Goodfellow et al., 2016](#goodfellow2016), ch. 14 | PyTorch / Keras | Reconstruction-error anomaly detection on high-dimensional sensor vectors |
| LSTM | [Hochreiter & Schmidhuber, 1997](#hochreiter1997) | PyTorch / Keras | Time-series prediction (load, temperature) and prediction-error anomaly detection |
| RL / MPC control | [Sutton & Barto, 2018](#sutton2018) | Gymnasium-style environments, custom | Learning or optimising HVAC control policies in simulation |

For most projects, an Isolation Forest scoring each sensor window is the right first model: it is fast, needs no labelled faults, and is exactly the pre-filter that the LLM-escalation pattern of Part 5 assumes.

---

## Part 4 — Why LLMs Make Good Controllers

Large language models bring four capabilities to building control that neither rules nor traditional ML can match. Each is worth understanding because it shapes when an LLM is the right tool.

All four capabilities trace back to the same underlying property: a sufficiently large transformer ([Vaswani et al., 2017](#vaswani2017)) trained on broad text acquires the ability to follow instructions and adapt to new tasks purely from its prompt — the few-shot, in-context learning behaviour ([Brown et al., 2020](#brown2020)). For a control engineer, in-context learning means the "training data" for a specific building can be supplied at inference time: the system prompt describes the building, its policies, and its quirks, and the model conditions its reasoning on that description without any retraining. That is a fundamentally different deployment model from Generation 2, where adapting to a new building means collecting data and refitting.

### Contextual reasoning

An LLM can reason about trade-offs that involve many objectives at once. Consider this: "Energy cost is currently high because of peak pricing, but three occupants are complaining about the heat, and the outdoor temperature is forecast to drop in two hours. The optimal action is to tolerate the discomfort for 45 minutes, then pre-cool aggressively when the price drops." No rule system can express this kind of multi-dimensional weighing. No ML classifier can either, because the relevant facts are different every time.

A useful image: a chess grandmaster commenting on a position. "I'd normally trade queens here, but my opponent has shown they prefer endgames, and I'm slightly down on time — so I'll keep the queens and look for a tactical opportunity instead." The reasoning weaves together many considerations that no fixed rule could capture.

There is empirical substance behind the analogy. The quality of an LLM's multi-step reasoning improves markedly when the model is prompted to externalise its intermediate steps rather than jump to an answer ([Wei et al., 2022](#wei2022)) — which is precisely why the agent designs later in this chapter insist that every decision be preceded by an explicit, logged Thought. The intermediate reasoning is not decoration; it is both a performance mechanism and an audit artefact.

### Natural-language understanding

A building manager doesn't think in terms of CO2 sensor IDs. They think in terms of "the meeting room on the third floor feels stuffy." An LLM agent can take that sentence, translate it into a sensor query, diagnose the issue, and act — without anyone having to look up the sensor's ID.

Think of it this way: a smart assistant who understands "the room that gets afternoon sun" without being told which compass direction that is.

### Explainability

When an LLM makes a decision, it can explain it in a paragraph. "I turned off HVAC zone 4 because the CO2 sensor in that zone has been stuck at 450 ppm for the last 30 minutes despite three occupants being present according to the booking system, which suggests a sensor fault rather than genuine air quality. I have created a maintenance ticket." A traditional ML model cannot produce this. It outputs a number, not an explanation.

A useful image: the difference between a thermometer that beeps and a thermometer that says, "I'm beeping because the patient's fever rose 0.8°C in the last hour, faster than the usual recovery curve, so we should check for infection."

One caution belongs here, because it is a frequent source of over-trust: the explanation an LLM produces is a *plausible account* generated by the same process that produced the decision, not a guaranteed trace of an internal computation. It is enormously useful for operators and auditors, and far better than no explanation at all — but it should be read the way one reads a colleague's account of their own reasoning: informative, usually honest, occasionally a rationalisation. This is one more reason the architecture validates actions independently (Part 7) rather than trusting the story that accompanies them.

### Handling novel situations

A fire alarm in a building hosting a film shoot that uses artificial smoke. A CO2 sensor that has been painted over by accident. A building manager who manually overrode the HVAC last week and forgot. LLMs handle these because they reason from common sense. Rule systems and ML models fail because none of these are in the rules or the training data.

An analogy: a new employee on their first day, facing a situation nobody trained them for. They reason from first principles: "this seems unusual; let me check with someone before acting." That kind of cautious common sense is exactly what an LLM brings.

---

## Part 5 — The Limits of LLMs (and How to Work Around Them)

LLMs are powerful but flawed. Four problems matter for building control, and each has a known mitigation.

### Hallucination

LLMs sometimes produce plausible-sounding but incorrect outputs. An LLM might suggest activating an actuator that doesn't exist, claim a sensor is showing a value it isn't, or invent a regulation. Hallucination is not malice — the LLM is doing exactly what it was trained to do, which is produce text that *looks* right.

Hallucination is a studied phenomenon, not folklore, and the literature draws a distinction that matters for control systems: *intrinsic* hallucination, where the output contradicts the source information the model was given, and *extrinsic* hallucination, where the output asserts something that cannot be verified from the source at all ([Ji et al., 2023](#ji2023)). An agent that misreports a sensor value present in its context is hallucinating intrinsically; an agent that invents an actuator is hallucinating extrinsically. The conclusion of that literature is sobering and useful: hallucination is a property of how these models are trained and decoded, it cannot be fully eliminated by prompting, and robust systems must therefore be designed so that hallucinated outputs are *caught*, not merely made less likely ([Ji et al., 2023](#ji2023)).

Picture this: a confident but unreliable witness in court. They speak with conviction. They're sometimes completely correct. They're sometimes completely wrong, and they don't know the difference.

Three mitigations:

- **Constrain the output to a JSON schema.** The LLM is forced to produce structured output with valid field types. It can't invent random tools or actuators because those aren't in the schema.
- **Validate every tool-call argument before execution.** The LLM might say "activate sprinkler in room ZZZ." The validation layer rejects this because room ZZZ doesn't exist. The LLM never reaches the actuator.
- **Never trust the LLM's claims about sensor values.** Always read ground truth from the API or database. If the LLM says "the temperature is 32°C," check the database. The LLM is the reasoner, not the source of truth.

### Latency

A cloud LLM API call takes 500 to 3000 milliseconds. That is far too slow for any decision on the critical safety path.

A useful image: imagine a thermostat that has to consult its lawyer before clicking on. The lawyer's advice might be excellent, but it arrives after the room has overheated.

Mitigations:

- **Use LLMs only for deliberative reasoning.** Decisions that can tolerate seconds of latency — HVAC scheduling, anomaly classification, planning — are fine.
- **Use ML models for real-time inference.** The ML model runs in milliseconds. When the ML model flags an anomaly, *then* the LLM is invoked to reason about whether it's a true positive.
- **Cache common reasoning patterns.** "Smoke above 0.7 plus rising temperature" is a frequent input. Cache the LLM's response.
- **Use smaller local models for routine decisions.** A local model like Llama 3.2 3B running on a GPU server takes 100-200 ms instead of 1000-2000 ms.

### Cost

Calling a cloud LLM API for every sensor reading is economically unsustainable. A building with 100 sensors reading every five seconds would make 1.7 million LLM calls per day. Even at one cent per call, that's $17,000 a day.

Think of it this way: you would not call a senior lawyer to review every email you send. You'd ask the lawyer about the hard ones and handle the rest yourself.

Mitigations:

- **Invoke the LLM only for novel or anomalous situations.** A simple ML or rule filter handles routine cases. The LLM is reserved for the unusual.
- **Use the ML model as a filter that escalates to the LLM only when needed.** This is the most common production pattern. Inference runs on every reading; the LLM runs only on the few that look interesting.

### Local models

A local model — running on the lab's GPU server or on a developer's laptop — solves both the latency and the cost problem. The model is slower and less capable than a frontier cloud model, but it's free and fast.

**Ollama** (https://ollama.com) is the standard tool for running open-source LLMs (Llama, Gemma, Mistral, Qwen) locally with a simple, OpenAI-compatible API. The course's lab GPU server runs Ollama. For routine reasoning in a building control system, a local model is entirely adequate.

The trade-off is capability for control. A small local model reasons less reliably than a frontier model — it follows complex instructions less faithfully, makes more tool-call formatting errors, and handles fewer simultaneous considerations. But it has properties a cloud API can never offer: deterministic availability (no rate limits, no outages outside your control), data locality (sensor data never leaves the building's network, which matters for occupancy data under privacy regulation), and a fixed, sunk cost. A sensible production architecture uses a local model for the high-volume routine reasoning and reserves a frontier model for the rare, genuinely hard cases — the same escalation logic as the ML-filter pattern, applied one level up.

A useful image: having a librarian in the office building. They don't know everything that the New York Public Library would know, but they're available immediately and free.

### Structured output

The single most useful trick for working with LLMs reliably is to ask for JSON instead of free text. With OpenAI-compatible APIs (including Ollama), `response_format={"type": "json_object"}` forces the model to produce JSON. With Anthropic's API, `tool_choice="required"` forces a tool call that has a defined schema.

An analogy: asking a job applicant to fill in a form instead of writing a letter. The form has named fields and types. Even a poor candidate produces something parseable. A letter, by contrast, can be free-form and impossible to process automatically.

---

## Part 6 — Giving the Agent Hands: Tool Use and Function Calling

<figure class="diagram">
<img src="figures/course-notes4-fig03.png" alt="Tool use — giving the agent hands">
<figcaption><em>Tools let the agent reach into systems it can't operate by itself. Each tool is a typed function the model can call when it needs to act.</em></figcaption>
</figure>

An LLM, on its own, can only read text and produce text. It can't read a sensor, write to a database, or command an actuator. **Tool calling** — also called function calling — is the mechanism that extends the LLM with the ability to invoke external functions.

The idea that a language model can decide to call external functions was first demonstrated as a learned capability: a model can teach itself, from a handful of demonstrations, *when* a calculator, a search engine, or a calendar would improve its answer, and emit the call in-line ([Schick et al., 2023](#schick2023)). Modern APIs industrialise the same idea — instead of the model learning tool use implicitly, the developer declares a catalogue of typed functions, and the model is trained to emit well-formed calls against that catalogue. The conceptual point survives the engineering change: tool use turns a closed-book text generator into an open-book system that can defer to authoritative external sources, which is also one of the most effective structural defences against the extrinsic hallucination discussed in Part 5.

Picture this: a brilliant consultant who is locked in a room with only a phone. They can think and talk. They can call you and ask you to do things. You go and do those things, then come back and report what happened.

The protocol works in five steps:

1. The system defines a set of tools (functions) with names, descriptions, and parameter schemas.
2. The user sends a message to the LLM, along with the tool definitions.
3. The LLM reasons about what to do. If it needs to act, it generates a structured **tool call** — a JSON object with the function name and the arguments.
4. The calling code executes that function and gets a result.
5. The result is returned to the LLM as a **tool result** message. The LLM then either calls another tool or produces a final answer.

This loop — Thought, Action, Observation, Thought, Action, Observation — is the foundation of the **ReAct pattern**, introduced in 2022 ([Yao et al., 2023](#yao2023)). ReAct stands for "Reasoning + Acting", and it is by far the most common pattern for LLM agents.

A worked example:

```
   Thought: The smoke sensor in A2306 reads 0.82, which is above the alert
            threshold of 0.7. I should check if there are corroborating
            signals before activating sprinklers.
   Action:  read_sensors(room="A2306", sensor_types=["temperature", "co2"])
   Observ:  temperature=38.5°C (normal is 21°C), co2=890 ppm (normal is 450)
   Thought: Temperature is significantly elevated and CO2 is rising. Combined
            with smoke, this is consistent with a fire, not a false alarm.
            I should activate sprinklers and unlock evacuation doors.
   Action:  set_actuator(actuator_id="sprinkler-A2306", state="on",
                         reason="fire detected: smoke=0.82, temp=38.5C")
   Observ:  {success: true, state: "on"}
   Action:  set_actuator(actuator_id="doors-floor2", state="unlocked",
                         reason="fire evacuation")
   Observ:  {success: true, state: "unlocked"}
   Thought: Sprinklers activated and evacuation doors unlocked. I should
            alert the building manager.
   Action:  create_alert(severity="critical", message="Fire detected in
                         A2306. Sprinklers activated. Doors unlocked.")
```

The key insight is that the reasoning (Thought) is explicit. Every step has a sentence explaining why the next action is justified. For safety-critical systems this is the difference between a system you can audit and one you can't. Every decision the agent makes can be replayed and explained.

A useful image: a detective's case notebook. Every observation, every deduction, every action is written down in order. If the case goes to court, the notebook is the evidence.

---

## Part 7 — Designing Good Tools

Tool design is more important than agent design. A well-designed tool catalogue lets even a small LLM reason effectively. A badly-designed catalogue confuses even a frontier model. Four principles drive good tool design.

The deeper reason is that the tool catalogue *is* the agent's action space. In the rational-agent framing, an agent can only be as good as the actions available to it and the percepts those actions return ([Russell & Norvig, 2021](#russell2021)); in a tool-calling agent, both are defined entirely by the tool schemas and their results. Every principle below is therefore a special case of a single idea: design the action space so that the right behaviour is easy to express, the wrong behaviour is hard to express, and every action's consequences are visible.

### One tool per concern

Don't write one big `control_building()` tool that does everything. Write narrow tools, each with a single job: `read_sensors`, `set_actuator`, `find_route`, `highlight_rooms`, `query_history`, `create_alert`. Each is small, validated, and easy to test.

Think of it this way: a Swiss Army knife with named tools versus a single blob marked "do something useful." When the LLM needs to cut wire, it picks the wire cutter. When it needs to open a bottle, it picks the bottle opener. Each tool's purpose is obvious from its name.

### Descriptive names and docstrings

The LLM decides which tool to call by reading its name and description. A tool called `get_data()` will be misused. A tool called `read_room_sensors(room_id, sensor_type, last_minutes)` will be used correctly. Spend time on the description — it is literally what teaches the LLM how to use the tool.

A useful image: jars in a kitchen pantry. If they're all labelled "stuff," the cook will pick at random. If each one is labelled clearly — "sea salt, fine" or "bicarbonate of soda" — the cook will reach for the right one.

### Validate inputs before execution

Never trust tool-call arguments unconditionally. Check that the room exists, that the sensor type is valid, that numeric arguments are within sensible ranges. If validation fails, return an error message — the LLM can read it and try again with corrected arguments.

An analogy: a bouncer at a club door. The bouncer doesn't care that the person is wearing a tuxedo. They check the guest list. If the name isn't on the list, the person doesn't get in, regardless of how confident they appear.

### Return rich, informative results

When a tool returns a result, the LLM uses that result to reason about what to do next. A bare result like `{"state": "on"}` is the bare minimum. A rich result is far more useful:

```json
{
  "state": "on",
  "previous_state": "off",
  "last_changed_at": "2026-04-27T13:42:05Z",
  "commanded_by": "safety_agent",
  "applied_within_ms": 230
}
```

The LLM can now reason about how recently the actuator changed, who else has commanded it, and whether the response time was normal.

Picture this: a good waiter doesn't just bring you the dish. They tell you what's in it and what it goes well with. The extra context lets you order better the next time.

### A complete example

A tool definition that follows all four principles, in OpenAI/Ollama function-calling format:

```python
tools = [{
    "type": "function",
    "function": {
        "name": "read_sensors",
        "description": (
            "Read current sensor values for a specific room. "
            "Returns temperature (°C), humidity (%), CO2 (ppm), "
            "smoke (0-1), and occupancy state."),
        "parameters": {
            "type": "object",
            "properties": {
                "room_id": {
                    "type": "string",
                    "description": "Room identifier, e.g. 'A2306'"
                },
                "sensor_types": {
                    "type": "array",
                    "items": {"type": "string"},
                    "description": ("List of sensor types to read. "
                                    "Options: temperature, humidity, "
                                    "co2, smoke, occupancy")
                }
            },
            "required": ["room_id"]
        }
    }
}]
```

---

## Part 8 — The Model Context Protocol (MCP)

A new agent project usually starts with the team defining their own tool format, their own way of registering tools, their own way of exposing them to the LLM. After a few projects, everyone has invented the same wheel a slightly different way.

The **Model Context Protocol**, abbreviated MCP, was proposed by Anthropic in 2024 to standardise this ([Anthropic, 2024](#anthropic2024); specification at https://modelcontextprotocol.io). It defines a single protocol for connecting AI models to external data sources and tools, so that any MCP-compatible model can discover and use tools from any MCP-compatible server. The economic argument is the classic one for any interface standard: without a standard, connecting *M* models to *N* tool providers requires on the order of M×N bespoke integrations; with a standard, it requires M+N implementations of one protocol. The trade-off, as with every standard, is a layer of indirection and protocol machinery that a single hard-wired integration would not need — which is why MCP pays off most when tools are shared across multiple agents and applications, and least for a one-off script.

The architecture has three components:

- **MCP Host** — the application that runs the LLM (Claude Code, an AI agent application).
- **MCP Client** — the protocol client inside the host that manages server connections.
- **MCP Server** — a server that exposes tools, resources, and prompts via the MCP protocol.

A useful image: think of MCP as a universal remote control standard. Before MCP, every TV maker shipped its own remote, and you needed a dozen remotes for your living room. MCP is HDMI-CEC for AI tools — one standard so that any device works with any controller.

For building control, an **MCP server** wraps the BuildSim API as MCP tools. The BuildSim MCP server in the course repository exposes 17 tools: `get_building`, `get_rooms`, `list_equipment`, `add_sensor`, `set_actuator_state`, `highlight_rooms`, `find_route`, and more. Any MCP-compatible agent — Claude, GPT-4, a local Ollama model with an MCP wrapper — can discover and use those tools automatically. No custom integration code.

The protocol itself uses JSON-RPC 2.0 over stdio or HTTP with server-sent events. Implementing an MCP server means defining tool schemas, implementing tool handlers (which usually wrap REST API calls), exposing a `tools/list` endpoint that returns the catalogue, and exposing a `tools/call` endpoint that executes a chosen tool.

The MCP Python SDK handles the protocol plumbing. Only the tool implementations have to be written.

---

## Part 9 — Agent Frameworks: What They Buy You and What They Cost

Agent frameworks provide scaffolding for building LLM agents: tool execution loops, memory management, state machines, multi-agent coordination, and integrations with many LLM providers. The right choice depends on the complexity of the use case.

Two general observations before the catalogue. First, every framework embodies a trade-off between *abstraction* and *transparency*: the framework saves you from writing the agent loop, but it also hides the loop, and when an agent misbehaves the debugging happens inside someone else's control flow. For a safety-relevant system, being able to point at the exact line where a tool call is validated is worth a great deal. Second, this layer of the stack is young and moves quickly — APIs change between minor versions, and the right habit is to treat the framework as replaceable scaffolding around stable concepts (the agent loop, ReAct, tool schemas) rather than as the architecture itself. The concepts in this chapter will outlive any particular library listed below.

### No framework

For simple agents that call one or two tools in a fixed sequence, plain Python with direct LLM API calls is often clearest. A 50-line Python script that reads sensors, constructs a prompt, calls the LLM, parses the tool call, executes the tool, and loops is simpler to understand, test, and debug than the same logic expressed in a framework's domain-specific language.

Think of it this way: cooking a single fried egg. You don't need a chef. You need a frying pan and basic technique.

### LangChain

The most widely used LLM application framework (https://python.langchain.com). Provides chains (sequential tool calls), agents (LLM-driven tool selection), memory (conversation history), and a huge ecosystem of integrations. Best when many integrations are needed quickly — connectors to databases, vector stores, cloud LLM providers, and so on. The cost of the breadth is well known in practice: many layers of abstraction for what is sometimes a single API call, and a fast-moving API surface.

A useful image: a well-stocked kitchen. Many tools, many ingredients, many recipes already half-prepared. Great for many dishes; overkill if you only ever fry eggs.

### LangGraph

A graph-based agent framework from the LangChain team (https://langchain-ai.github.io/langgraph/). Agents are defined as nodes in a directed graph with cycles; the LLM decides which edge to follow at each step. Better than plain LangChain for complex workflows with branching, looping, and human-in-the-loop checkpoints, because the set of reachable states is explicit in the graph rather than implicit in the model's behaviour — a property that matters when you need to argue that an agent *cannot* reach a forbidden state.

```python
from typing import TypedDict

from langgraph.graph import StateGraph, END
from langgraph.prebuilt import ToolNode

class AgentState(TypedDict):
    messages: list
    sensor_readings: dict
    pending_actions: list

workflow = StateGraph(AgentState)
workflow.add_node("read_sensors", sensor_reading_node)
workflow.add_node("reason", llm_reasoning_node)
workflow.add_node("act", ToolNode(building_tools))

workflow.set_entry_point("read_sensors")
workflow.add_edge("read_sensors", "reason")
workflow.add_conditional_edges("reason", should_act,
                               {"act": "act", "done": END})
workflow.add_edge("act", "reason")

agent = workflow.compile()
```

An analogy: a board game with a movement map. Each square has rules for what moves are legal. The agent rolls the dice and moves; the framework enforces the rules of the board.

### CrewAI

A framework for role-based multi-agent systems (https://docs.crewai.com). Agents are defined with **roles** (architect, designer, builder), **goals**, and **backstories**. **Tasks** have descriptions and expected outputs. **Crews** assign tasks to agents, and agents can delegate to each other.

Picture this: a film production team. The director, the cinematographer, the editor — each has a defined role, each can collaborate with the others, and the producer (the crew) coordinates them.

### AutoGen

Microsoft's conversational multi-agent framework ([Wu et al., 2023](#wu2023); https://microsoft.github.io/autogen/). Agents exchange messages in a structured conversation; a coordinator directs the flow. The accompanying paper frames multi-agent conversation as a general programming model for LLM applications and is worth reading for its catalogue of conversation patterns. Good for tasks where multiple agents critique each other — code review, report writing, multi-step debugging.

A useful image: a panel discussion. The moderator picks who speaks. Each panellist has a perspective. The conversation produces a richer answer than any single panellist alone.

### Pydantic AI

A newer framework built around Pydantic for structured outputs and type safety (https://ai.pydantic.dev). Good for agents that need reliable, typed tool-call arguments and responses. Less feature-rich than LangGraph but simpler to use for straightforward agents.

Think of it this way: building furniture with a precision template. Less flexible than freehand work, but every joint fits.

### Choosing among them

A reasonable default for this course: no framework or Pydantic AI for a single agent with a handful of tools; LangGraph when the workflow genuinely has branches, loops, and checkpoints that deserve to be explicit; CrewAI or AutoGen only when the design truly calls for multiple communicating agents. The common beginner mistake is to reach for the most powerful framework first — the equivalent of hiring a film crew to fry an egg — and then spend the project debugging the framework instead of the agent.

---

## Part 10 — The ReAct Pattern in Detail

<figure class="diagram">
<img src="figures/course-notes4-fig04.png" alt="ReAct pattern — Thought · Action · Observation">
<figcaption><em>In the ReAct pattern, the agent alternates between reasoning (Thought) and acting (Action), with each Observation feeding the next Thought. The trace is also the explanation.</em></figcaption>
</figure>

The ReAct pattern — Reasoning + Acting — was introduced in 2022, building directly on the chain-of-thought result ([Wei et al., 2022](#wei2022); [Yao et al., 2023](#yao2023)). Chain-of-thought showed that explicit intermediate reasoning improves an LLM's answers; ReAct's contribution was to ground that reasoning in an external environment, so that each reasoning step can trigger an action and each action's result can correct the reasoning. Grounding of this kind reduces the hallucination that pure chain-of-thought suffers from — the model can *look things up* instead of confabulating them — while pure acting without reasoning fares worse on tasks requiring planning ([Yao et al., 2023](#yao2023)). The pattern is short, clear, and has become the de facto template for LLM agents.

The pattern interleaves three step types:

- **Thought**: the LLM's reasoning, in natural language.
- **Action**: a tool call.
- **Observation**: the result of the tool call.

Repeated until the LLM produces a final answer instead of another action.

The key insight is that the **reasoning is explicit and auditable**. For a safety-critical building control system, this is invaluable. After a fire incident, you can replay every Thought, every Action, every Observation, and see exactly what the agent did and why.

A useful image: a pilot's flight log. Every decision is timestamped, with the reason. If the plane lands safely, nobody looks at the log. If something goes wrong, the log is the first thing the investigators read.

The interleaving also helps the LLM stay on track. A pure plan-then-execute pattern is fragile: the plan is made once, and if reality doesn't match the plan, the agent fails. ReAct adapts continuously — every Observation feeds back into the next Thought, and the plan evolves.

The pattern has costs worth naming. Each Thought–Action–Observation cycle is a full model invocation, so a ten-step ReAct trace costs roughly ten times the latency and tokens of a single-shot answer, and the growing trace consumes context window. ReAct is therefore the right pattern for deliberative, tool-heavy tasks where adaptivity and auditability dominate — exactly the operational-reasoning band of the table in Part 3 — and the wrong pattern for the high-frequency, low-latency decisions that belong to rules and ML models.

---

## Part 11 — Multi-Agent Systems

<figure class="diagram">
<img src="figures/course-notes4-fig05.png" alt="Multi-agent collaboration">
<figcaption><em>Specialised agents own a slice of the problem and communicate through a shared bus. A coordinator resolves conflicts; an auditor records every decision.</em></figcaption>
</figure>

A single agent that controls every aspect of a building is fragile. It is hard to test, hard to update, and has a single point of failure. Multi-agent systems decompose the control problem into specialised agents that coordinate.

Multi-agent systems are an established research field with a literature that predates LLMs by decades. The central problems — how agents should represent each other, communicate, negotiate, and coordinate without a central controller — were identified early, and *social ability* appears in the classical agent definition precisely because realistic agents share their environment with other agents ([Wooldridge & Jennings, 1995](#wooldridge1995)). The LLM era changes the implementation — agents can now negotiate in natural language, as systems like AutoGen demonstrate ([Wu et al., 2023](#wu2023)) — but not the problem structure: everything in this part and the next is a modern instance of questions that field has studied since the early 1990s.

### Why multiple agents

Four reasons.

**Separation of concerns.** A **safety agent** handles fires and emergencies with highest priority. An **energy agent** optimises consumption. A **comfort agent** maintains temperature and air quality. Each agent has a clear responsibility and can be developed, tested, and updated independently.

**Parallel reasoning.** Multiple agents can reason simultaneously. The energy agent is continuously optimising the HVAC schedule while the comfort agent is processing an occupant feedback event, and they do not block each other.

**Robustness.** If the comfort agent crashes, the safety agent keeps running. Single-agent architectures have a single point of failure.

**Specialisation.** A small, fast model may be sufficient for routine energy optimisation, while a larger model is used only for complex safety decisions. Different agents use different models, optimising cost and latency.

An analogy: a hospital. The cardiologist handles hearts. The nephrologist handles kidneys. The emergency physician handles whatever walks in. Each specialist is excellent at their narrow domain. The patient benefits from all of them.

### How agents talk to each other

Three communication patterns.

**Shared state (the blackboard pattern).** All agents read and write to a shared data store (the time-series database, a Redis instance, the BuildSim API). Agents observe each other's actions through state changes.

Picture this: detectives in a precinct office sharing a whiteboard. Each adds clues; each sees what the others have added. Simple and decoupled, but with the risk that two detectives independently chase the same lead.

**Message passing.** Agents send structured messages to each other through a message queue (MQTT, Redis Streams). Explicit, auditable, decoupled.

A useful image: a customer-service team where every escalation is a ticket. The customer-service rep tickets a manager; the manager tickets engineering; engineering tickets the customer-service rep with a solution. Every step is recorded.

**Shared context in the LLM.** For LLM-based agents, the entire conversation history (including other agents' decisions) is included in each agent's context. Each agent sees what the others have decided. Expensive in tokens but produces well-informed decisions.

Think of it this way: a war-room briefing where every participant hears every report. Slow but coherent.

---

## Part 12 — Conflict Resolution and the Pathologies Multi-Agent Brings

The moment a system has multiple agents with different objectives, conflicts become possible. The energy agent wants the HVAC off. The comfort agent wants it on. Who wins?

This is not an accident of bad design; it is structural. Decomposing one global objective (run the building well) into per-agent objectives (minimise energy, maximise comfort, guarantee safety) necessarily creates agents whose locally optimal actions disagree, and the coordination and negotiation mechanisms below are the classical answers from the multi-agent literature ([Wooldridge & Jennings, 1995](#wooldridge1995)). The choice among them is itself a trade-off between simplicity, flexibility, and the quality of the compromise reached.

### Conflict-resolution strategies

**Priority-based.** Each agent has a fixed priority. The safety agent's priority is 1000 — it always wins. Energy is 10. Comfort is 20. When agents conflict, the higher-priority agent's command is executed. Simple, predictable, but inflexible.

A useful image: emergency services priority. When the fire alarm rings, the fire brigade takes over the road. Nothing else competes.

**Auction-based.** Agents bid for actuator control with a numeric priority. The agent with the highest bid wins. Bids can be dynamic — the safety agent bids 1000 only when smoke is detected, 0 otherwise. This gives flexibility based on context.

An analogy: an auction house. The dynamic bid lets a normally-quiet bidder suddenly out-shout everyone when something special comes up.

**Consensus.** Agents negotiate a compromise. The energy agent proposes a temperature setpoint; the comfort agent counter-proposes; they converge. More complex but produces better outcomes when neither agent needs to win completely.

Picture this: two negotiators meeting in the middle. Each starts from their preferred outcome and walks halfway. Both leave slightly unhappy and entirely satisfied.

**Hierarchical (supervisor pattern).** A supervisor agent receives all proposed actions from subordinate agents, reasons about conflicts, and produces a final unified action set. The supervisor has the most context but is a bottleneck.

A useful image: a project manager arbitrating between team members. Slower than letting them work it out, but produces consistent decisions.

### Pathologies — failure modes that don't exist with a single agent

Multi-agent systems can fail in ways that single-agent systems cannot.

**Oscillation.** Agent A turns on the HVAC, which raises the temperature; Agent B turns it off to save energy; temperature drops; Agent A turns it on again. The loop repeats indefinitely.

Think of it this way: two thermostats with different setpoints, wired into the same room. Each fights the other forever. Neither room nor energy bill ever stabilise.

The prevention is **hysteresis** — refusing to change state unless the trigger condition has been true for at least N seconds. A shared decision log so each agent knows what the others recently decided also helps.

**Deadlock.** Agent A waits for Agent B to release an actuator; Agent B waits for Agent A to release a different one. Neither can proceed.

A useful image: two cars at a four-way stop, both waiting for the other to go.

Prevention: avoid circular dependencies in actuator ownership; use timeout-based locks.

**Thrashing.** Multiple agents issue conflicting commands in rapid succession, causing actuators to change state faster than the physical system can respond.

An analogy: flipping a light switch ten times a second. The bulb never has time to warm up; it just flickers.

Prevention: rate-limit actuator commands; enforce a minimum time between state changes for any actuator.

---

## Part 13 — Safety and Trust in Agentic AI

<figure class="diagram">
<img src="figures/course-notes4-fig06.png" alt="Safety guardrails — concentric layers">
<figcaption><em>An autonomous agent is wrapped in concentric layers of policy. Authority limits what it can touch; policy says what it must never do; humans approve the rest; an audit log records everything.</em></figcaption>
</figure>

An AI agent controlling a building can cause real harm. It can leave emergency doors locked during a fire, over-cool a server room, or cause energy spikes that trip a breaker. Safety in agentic AI is not abstract — it is an engineering requirement, and there are well-established techniques to enforce it.

The research community anticipated these problems before LLM agents existed. The "concrete problems in AI safety" were catalogued for any learning system that acts in the world, and three of the categories map directly onto building control ([Amodei et al., 2016](#amodei2016)). **Negative side effects**: an agent optimising one objective damages something outside that objective — the energy agent that saves electricity by shutting down ventilation in an occupied room. **Reward hacking**: an agent satisfies the letter of its objective while defeating its purpose — an agent rewarded for keeping a temperature *reading* in range that learns the cheapest action is to influence the sensor rather than the room. **Safe exploration**: an agent that tries actions to learn their effects must not try them on systems where a bad experiment is irreversible — fire suppression is not a place to explore. Each mechanism in this part is best understood as a defence against one or more of these failure classes: guardrails bound side effects, ground-truth validation and audit trails make reward hacking detectable, and human-in-the-loop approval removes irreversible actions from the agent's exploration space.

### Guardrails

**Guardrails** are constraints applied before an action is executed. The LLM never sees the guardrail; the validation code runs between the tool call and the actuator. Examples:

- Never lock all exits simultaneously, even if the LLM reasons that it should.
- Smoke suppression can only be deactivated by a human operator.
- Actuator commands above a certain rate are rejected.
- If any action would result in all HVAC units being off in winter, human confirmation is required.

Picture this: bumpers on a bowling lane. The LLM is the ball. The bumpers don't change the ball's spin or speed, but they catch a wild throw before it ends up in the gutter.

### Human-in-the-loop (HITL)

For high-stakes, irreversible actions — activating fire suppression in a particular zone, unlocking emergency exits — the agent proposes the action and waits for a human to approve. HITL adds latency, but the cost of a wrong decision is high enough that the latency is worth it.

A useful image: a customer-service rep needing manager approval to issue a refund over £500. Slows things down, but stops large mistakes.

### Audit trail

Every agent decision, tool call, and reasoning step must be logged immutably. When something goes wrong, reconstructing exactly what the agent decided and why is essential. The ReAct pattern is naturally audit-friendly: every Thought, every Action, every Observation is already a structured record.

Think of it this way: a bank ledger. Every transaction has a timestamp, an account, an amount, and a reference. If a discrepancy is found, the ledger is the source of truth.

### Graceful degradation

When the AI agent is unavailable — the LLM API is down, the agent process has crashed, the network is degraded — the system must fall back to a safe state without human intervention. All actuators remain in their current state. Rule-based fallbacks handle the most critical safety functions. Alerts notify operators that AI control is inactive. The system can still be operated manually via the BuildSim API.

A useful image: a hospital's manual procedures when the IT system goes down. Paper charts. Phone calls. Slower, less informed, but the patients still get cared for.

### Testing safety

Adversarial testing is essential. Inject sensor faults: what does the agent do when the smoke sensor reads 0 for ten minutes? Simulate network outages: what happens when the LLM API returns an error? Send conflicting commands: what does the coordinator do? Test with novel situations: what if a sensor reports a value the agent has never seen before?

An analogy: a fire drill in an office. You don't wait for a real fire to find out whether the evacuation plan works.

Tests for these scenarios are written **before** deployment, not after. They are part of the validation step of MBSE.

---

## Part 14 — A Worked Example: Fire-Detection Agent

<figure class="diagram">
<img src="figures/course-notes4-fig07.png" alt="Fire-detection agent in action">
<figcaption><em>The fire-detection agent in action: a sensor spike triggers reasoning, the agent calls actuator and notification tools, and the next sensor reading closes the loop.</em></figcaption>
</figure>

To make all of this concrete, walk a fire-detection agent through one complete cycle.

### Setup

- Smoke sensors per room, publishing readings every 5 seconds to MQTT.
- A consumer writing every reading to TimescaleDB.
- An anomaly model (Isolation Forest; [Liu et al., 2008](#liu2008)) that scores each sensor every 30 seconds.
- The safety agent, a ReAct agent using Ollama (Llama 3.2 3B locally) for reasoning.
- Six tools: `read_sensors`, `set_actuator`, `highlight_rooms`, `find_route`, `query_history`, `create_alert`.
- Guardrails: known actuator IDs only; legal states only; a `reason` field of at least 5 characters is required.
- Audit trail: every Thought, Action, Observation written to a `decisions` table in TimescaleDB.

### The trigger

Smoke level in room A2306 spikes to 0.85. The anomaly model scores 0.91, well above threshold. The safety agent wakes for its perceive cycle.

### Perceive

The agent calls `read_sensors(room="A2306", last_seconds=30)`. The result:

```json
{
  "room": "A2306",
  "sensors": {
    "smoke":       {"current": 0.85, "mean": 0.62, "trend": +0.41},
    "temperature": {"current": 38.5, "mean": 32.1, "trend": +6.4},
    "co2":         {"current": 890,  "mean": 670,  "trend": +220}
  }
}
```

### Reason

The agent generates a Thought:

> "Smoke at 0.85, temperature rising sharply (+6.4 °C in 30 seconds), CO2 also rising. This is a high-confidence fire signal across three independent sensor types. I should activate sprinklers and unlock the evacuation doors. Before acting, I want to know the safest evacuation route."

### Plan and act

The agent issues a tool call: `find_route(from_room="A2306", to_room="exit-stair-east")`. The result includes a path with five waypoints and a distance of 42 metres.

Then `set_actuator(actuator_id="sprinkler-A2306", state="on", reason="fire detected, smoke 0.85, temp 38.5, CO2 890, three signals concur")`. The guardrail checks that the actuator exists (yes), the state is legal (yes), and the reason is at least 5 characters (yes). The command goes through.

Then `set_actuator(actuator_id="firedoor-A2306-east", state="unlocked", reason="evacuation route to east stair")`.

Then `highlight_rooms(rooms=[{"room_id": 552, "color": "#ff0000"}, ...])` to make the fire zone visible on the 3D viewer.

Then `create_alert(severity="critical", message="Fire detected in A2306. Sprinklers active. Evacuation east via stair-east.")`.

### Observe

The agent calls `read_sensors` again 30 seconds later. Smoke level has dropped to 0.45. Temperature still high but stable. The sprinkler is working.

### Update

The full sequence — Thought, Actions, Observations — is written to the audit log. The agent's next cycle starts with this history available.

### When something goes wrong

If Ollama is unreachable when the agent tries to reason, the agent falls back to a pre-coded rule: `if smoke > 0.7 then sprinkler on`. The decision is logged with `step='fallback'` so the audit trail records that the LLM was bypassed. A separate watchdog process enforces hard safety rules independent of the agent: smoke above 0.9 always triggers sprinklers, regardless of what any agent does.

---

## Part 15 — Vocabulary Reference

Every term used in this chapter, defined.

| Term | Definition |
|---|---|
| **Agent** | A system that perceives, reasons, plans, acts, observes, and updates — using an LLM, ML model, or rules as its reasoning engine |
| **Agentic AI** | An AI system built around an agent that uses an LLM to reason about its situation and choose actions |
| **The Agent Loop** | Perceive → Reason → Plan → Act → Observe → Update — the six-step cycle every agent runs |
| **LLM (Large Language Model)** | A neural network trained on text that can generate plausible text in response to prompts |
| **Tool calling** | The mechanism by which an LLM invokes external functions through structured tool-call messages |
| **Function calling** | Synonym for tool calling, used by OpenAI |
| **ReAct pattern** | Reasoning + Acting — interleaving Thought, Action, Observation in the agent's output |
| **Thought** | The LLM's natural-language reasoning step, logged for audit |
| **Action** | A tool call invoked by the LLM |
| **Observation** | The result of a tool call, returned to the LLM as input for the next step |
| **Hallucination** | The LLM producing plausible but incorrect output |
| **Structured output** | Forcing the LLM to produce JSON that matches a schema, rather than free-form text |
| **Local model** | An LLM running on local hardware (Ollama) instead of a cloud API |
| **MCP (Model Context Protocol)** | An open standard for connecting LLMs to external tools, proposed by Anthropic in 2024 |
| **MCP Host** | The application that runs the LLM |
| **MCP Client** | The protocol client managing connections to MCP servers |
| **MCP Server** | A server that exposes tools, resources, and prompts via MCP |
| **Tool (in an agent)** | A function the LLM can call to read or change the world |
| **Tool schema** | The JSON-schema definition of a tool's name, description, and parameters |
| **LangChain** | The most widely used LLM application framework, with chains, agents, memory, and many integrations |
| **LangGraph** | A graph-based agent framework from the LangChain team for complex workflows |
| **CrewAI** | A framework for role-based multi-agent systems |
| **AutoGen** | Microsoft's conversational multi-agent framework |
| **Pydantic AI** | A framework focused on structured outputs and type safety |
| **Multi-agent system** | An architecture with several specialised agents that coordinate to control one system |
| **Coordination layer** | The component that resolves conflicts between agents |
| **Priority-based resolution** | Each agent has a fixed priority; highest priority wins conflicts |
| **Auction-based resolution** | Agents bid for control; highest bid wins |
| **Consensus resolution** | Agents negotiate a compromise rather than one winning outright |
| **Hierarchical (supervisor)** | A supervisor agent arbitrates decisions made by subordinates |
| **Oscillation** | A pathology in which two agents flip an actuator back and forth |
| **Deadlock** | A pathology in which two agents wait for each other forever |
| **Thrashing** | A pathology in which actuators change state faster than the physical system can respond |
| **Hysteresis** | Refusing to change state until the trigger condition has been true for a minimum time |
| **Guardrail** | A validation rule applied before an actuator command is executed |
| **Human-in-the-loop (HITL)** | Requiring human approval before high-stakes actions are executed |
| **Audit trail** | The immutable log of every agent decision, tool call, and reasoning step |
| **Graceful degradation** | The ability to fall back to a safe state when the AI agent is unavailable |
| **Adversarial testing** | Deliberately injecting faults to test how the agent handles them |
| **Ollama** | A tool for running open-source LLMs locally with a simple API |

---

## Part 16 — Summary in Five Sentences

1. Three generations of automation — rules, ML inference, and agentic AI — sit on a continuum from fast-and-rigid to slow-and-thoughtful, and a well-designed system uses all three where each fits.
2. Every agent runs the same six-step loop — Perceive, Reason, Plan, Act, Observe, Update — and the Update step is what turns inference into learning.
3. LLMs are powerful controllers because they handle context, novel situations, and natural language, but they hallucinate, are slow, and are expensive — every one of those weaknesses has a known mitigation through structured output, validation, ML pre-filtering, and local models.
4. Tools are the agent's hands, and the four design principles — one tool per concern, descriptive names, input validation, rich results — matter more than the choice of framework or model.
5. Multi-agent systems are stronger and more failure-resilient than single agents, but they require a conflict-resolution strategy (priority, auction, consensus, or hierarchy) and explicit mitigations for the pathologies they create (oscillation, deadlock, thrashing).

These five ideas are the foundation for every agent in the rest of the course.

---

## Part 17 — References

### Literature

- <a id="amodei2016"></a>Amodei, D., Olah, C., Steinhardt, J., Christiano, P., Schulman, J., & Mané, D. (2016). Concrete Problems in AI Safety. *arXiv:1606.06565*.
- <a id="breunig2000"></a>Breunig, M. M., Kriegel, H.-P., Ng, R. T., & Sander, J. (2000). LOF: Identifying Density-Based Local Outliers. *Proceedings of the ACM SIGMOD International Conference on Management of Data*.
- <a id="brown2020"></a>Brown, T., et al. (2020). Language Models are Few-Shot Learners. *Advances in Neural Information Processing Systems (NeurIPS)*.
- <a id="chandola2009"></a>Chandola, V., Banerjee, A., & Kumar, V. (2009). Anomaly Detection: A Survey. *ACM Computing Surveys*, 41(3).
- <a id="goodfellow2016"></a>Goodfellow, I., Bengio, Y., & Courville, A. (2016). *Deep Learning*. MIT Press. https://www.deeplearningbook.org
- <a id="hochreiter1997"></a>Hochreiter, S., & Schmidhuber, J. (1997). Long Short-Term Memory. *Neural Computation*, 9(8).
- <a id="ji2023"></a>Ji, Z., et al. (2023). Survey of Hallucination in Natural Language Generation. *ACM Computing Surveys*, 55(12).
- <a id="liu2008"></a>Liu, F. T., Ting, K. M., & Zhou, Z.-H. (2008). Isolation Forest. *Proceedings of the IEEE International Conference on Data Mining (ICDM)*.
- <a id="pedregosa2011"></a>Pedregosa, F., et al. (2011). Scikit-learn: Machine Learning in Python. *Journal of Machine Learning Research*, 12.
- <a id="russell2021"></a>Russell, S., & Norvig, P. (2021). *Artificial Intelligence: A Modern Approach* (4th ed.). Pearson.
- <a id="schick2023"></a>Schick, T., Dwivedi-Yu, J., Dessì, R., Raileanu, R., Lomeli, M., Zettlemoyer, L., Cancedda, N., & Scialom, T. (2023). Toolformer: Language Models Can Teach Themselves to Use Tools. *Advances in Neural Information Processing Systems (NeurIPS)*.
- <a id="scholkopf2001"></a>Schölkopf, B., Platt, J. C., Shawe-Taylor, J., Smola, A. J., & Williamson, R. C. (2001). Estimating the Support of a High-Dimensional Distribution. *Neural Computation*, 13(7).
- <a id="sutton2018"></a>Sutton, R. S., & Barto, A. G. (2018). *Reinforcement Learning: An Introduction* (2nd ed.). MIT Press.
- <a id="vaswani2017"></a>Vaswani, A., Shazeer, N., Parmar, N., Uszkoreit, J., Jones, L., Gomez, A. N., Kaiser, Ł., & Polosukhin, I. (2017). Attention Is All You Need. *Advances in Neural Information Processing Systems (NeurIPS)*.
- <a id="wei2022"></a>Wei, J., Wang, X., Schuurmans, D., Bosma, M., Ichter, B., Xia, F., Chi, E., Le, Q., & Zhou, D. (2022). Chain-of-Thought Prompting Elicits Reasoning in Large Language Models. *Advances in Neural Information Processing Systems (NeurIPS)*.
- <a id="wooldridge1995"></a>Wooldridge, M., & Jennings, N. R. (1995). Intelligent Agents: Theory and Practice. *The Knowledge Engineering Review*, 10(2).
- <a id="wu2023"></a>Wu, Q., et al. (2023). AutoGen: Enabling Next-Gen LLM Applications via Multi-Agent Conversation. *arXiv:2308.08155*.
- <a id="yao2023"></a>Yao, S., Zhao, J., Yu, D., Du, N., Shafran, I., Narasimhan, K., & Cao, Y. (2023). ReAct: Synergizing Reasoning and Acting in Language Models. *International Conference on Learning Representations (ICLR)*. (Preprint: arXiv:2210.03629, 2022.)

### Software, frameworks, and online resources

- <a id="anthropic2024"></a>Anthropic (2024). Introducing the Model Context Protocol. https://modelcontextprotocol.io
- AutoGen — Microsoft's conversational multi-agent framework. https://microsoft.github.io/autogen/
- CrewAI — role-based multi-agent framework. https://docs.crewai.com
- LangChain — LLM application framework. https://python.langchain.com
- LangGraph — graph-based agent framework. https://langchain-ai.github.io/langgraph/
- Ollama — local LLM runtime with an OpenAI-compatible API. https://ollama.com
- Pydantic AI — type-safe agent framework built on Pydantic. https://ai.pydantic.dev
- scikit-learn — machine learning in Python ([Pedregosa et al., 2011](#pedregosa2011)). https://scikit-learn.org — `IsolationForest`: https://scikit-learn.org/stable/modules/generated/sklearn.ensemble.IsolationForest.html
