# AI-Driven Development & Test-Driven Generation

## The Specification-First Workflow

### Why the Old Workflow Breaks Down with AI

The traditional software development workflow goes roughly: understand the problem, write code, write tests, discover the design was wrong, rewrite. Tests are treated as a verification afterthought, and the design often emerges from the code rather than driving it. This is painful enough for human developers, but it breaks down completely when AI is writing the code.

AI coding tools are extraordinarily good at generating code that looks correct. They produce clean, idiomatic, well-formatted code that passes a superficial review. The problem is that they generate code to satisfy the prompt, not the problem. If the prompt is vague ("write a sensor process"), the AI makes assumptions, often reasonable ones, sometimes disastrously wrong ones. When the code fails, it is not clear whether the AI misunderstood the prompt, the prompt was ambiguous, or the design was wrong in the first place.

The solution is to flip the workflow: write the specification first, then write the tests, then generate the code. The specification is precise enough that the AI has no room for dangerous assumptions. The tests encode the expected behaviour so precisely that "correct" is a binary outcome. The AI becomes a contractor who must satisfy a clearly written contract, not an artist interpreting a vague brief.

### The Architecture Document as Prompt

The MBSE artifacts produced in week 1 are not bureaucratic overhead; they are the inputs to the AI development process:

- **Requirements** → acceptance tests (each requirement becomes a test case)
- **Component diagram** → project structure (each container becomes a directory with a Dockerfile)
- **Sequence diagrams** → integration tests (the sequence diagram is the integration test scenario)
- **API contracts** → interface tests (the JSON schema is the assertion)
- **Data models** → unit tests (test that every field is present, typed, and in range)
- **State machines** → state transition tests (test every state and every transition)

A well-written architecture document dramatically reduces the effort of implementation. An AI that has the architecture document, sequence diagrams, and API contracts as context can generate a working skeleton in minutes. An AI working from a vague description will generate something that looks plausible but fails on the first real test.

> **Key principle:** The specification is more important than ever when working with AI. AI amplifies the quality of the specification: a good spec becomes excellent code; a bad spec becomes confidently wrong code.

```mermaid
%%{init: {"theme": "neutral"}}%%
flowchart LR
    SPEC[Architecture\nDocument] --> TESTS[Write Tests\nfrom spec]
    TESTS --> PROMPT[Craft AI\nPrompt]
    SPEC --> PROMPT
    PROMPT --> GEN[AI Generates\nCode]
    GEN --> RUN[Run Tests]
    RUN -->|fail| DEBUG[Debug /\nRefine Prompt]
    DEBUG --> PROMPT
    RUN -->|pass| REVIEW[Code Review\n+ Integration]
    REVIEW -->|issues| DEBUG
    REVIEW -->|approved| DONE[Working Component]
```

## Test-Driven Generation

Test-driven generation combines classical Test-Driven Development with AI code generation. The idea is simple but powerful: instead of writing both the tests and the implementation, the developer writes only the tests and lets the AI generate the code that passes them. The tests become the specification, the contract that the AI must satisfy. If the AI produces code that looks correct but fails a test, the bug is caught immediately. If the tests pass, the code is correct by definition, regardless of whether a human would have written it differently.

This matters especially for AI-generated code because AI tools are good at producing plausible-looking code that may contain subtle errors. Without tests, these errors can go undetected until production. With tests written first, every error is caught at generation time.

### The Practice

[Test-Driven Development](https://www.amazon.com/Test-Driven-Development-Kent-Beck/dp/0321146530) was introduced by Kent Beck in 2003. The core discipline is: write a failing test, write the minimum code to make it pass, then refactor. With AI code generation, this workflow becomes: write a failing test, give the test to the AI as context ("implement code that makes this test pass"), run the tests, and if they fail, refine the prompt or tests and regenerate.

This catches hallucinated APIs early (the test fails with an import error), wrong logic early (assertion errors), and missing error handling (write a test for the failure case and the AI is forced to handle it). The key insight is that writing a good test is often easier than writing good code, and AI is better at generating code from a clear specification (the test) than from a vague description.

### Test Types for CPS

**Unit tests** test a single function or class in isolation. For building control:

```python
# Example: test that a smoke sensor reading is correctly classified
def test_smoke_alert_threshold():
    classifier = SmokeClassifier(threshold=0.7)
    assert classifier.classify(0.65) == SensorState.NOMINAL
    assert classifier.classify(0.75) == SensorState.ALERT
    assert classifier.classify(1.0) == SensorState.ALERT

def test_smoke_classifier_rejects_invalid_input():
    classifier = SmokeClassifier(threshold=0.7)
    with pytest.raises(ValueError):
        classifier.classify(-0.1)  # Smoke level cannot be negative
    with pytest.raises(ValueError):
        classifier.classify(1.5)   # Smoke level cannot exceed 1.0
```

**Integration tests** test that two or more components work together correctly. They require a running BuildSim instance (or a mock of it):

```python
# Example: test that sensor process stores readings in the database
def test_sensor_process_stores_readings(buildsim_mock, db_connection):
    sensor_process = SensorProcess(api_url="http://localhost:8080", db=db_connection)
    buildsim_mock.set_sensor("smoke-A2306", value=0.82)
    sensor_process.poll_once()
    readings = db_connection.query("SELECT * FROM readings WHERE sensor_id='smoke-A2306'")
    assert len(readings) == 1
    assert readings[0]["value"] == pytest.approx(0.82)
```

**Behavioural tests** test the end-to-end response to a scenario. These are the most valuable tests for CPS, verifying the entire system against a requirement:

```python
# Requirement R-FIRE-01: fire detected → sprinklers on within 30 seconds
def test_fire_detection_activates_sprinklers(building_system, buildsim_client):
    buildsim_client.set_sensor("smoke-A2306", value=0.85)
    buildsim_client.set_sensor("temperature-A2306", value=45.0)
    time.sleep(30)  # wait up to 30 seconds
    actuator_state = buildsim_client.get_actuator("sprinkler-A2306")
    assert actuator_state["state"] == "on"
```

**Property-based tests** verify invariants that must always hold:

```python
# Property: actuator state is always one of the valid states
@given(smoke_level=st.floats(0.0, 1.0), temperature=st.floats(10.0, 80.0))
def test_agent_always_produces_valid_actuator_command(smoke_level, temperature):
    agent = SafetyAgent()
    command = agent.decide({"smoke": smoke_level, "temperature": temperature})
    assert command.actuator_id in VALID_ACTUATOR_IDS
    assert command.state in ["on", "off", "auto"]
```

[Hypothesis](https://hypothesis.readthedocs.io/en/latest/) is the standard Python library for property-based testing.

### Testing AI Agent Behaviour

Testing an LLM-based agent is harder than testing a deterministic function, because the LLM response varies between calls. Three strategies are available.

**Mock the LLM:** replace the LLM client with a mock that returns deterministic responses. This tests that the agent correctly interprets the LLM output and takes the right action, without testing the LLM itself.

```python
def test_agent_activates_sprinklers_on_fire_decision(mock_llm):
    mock_llm.set_response('{"action": "activate_sprinkler", "zone": "A2306", "reason": "fire detected"}')
    agent = SafetyAgent(llm=mock_llm)
    agent.process_alert({"smoke": 0.85, "room": "A2306"})
    assert mock_llm.last_tool_call == ("activate_sprinkler", {"zone": "A2306"})
```

**Record and replay:** in a staging environment, run the real LLM and record all requests and responses. In CI, replay the recorded responses. This captures real LLM behaviour while making tests deterministic and fast.

**Scenario tests with evaluation:** define a set of scenarios (sensor states) and expected outputs (actuator commands or reasoning categories). Run the real LLM, and use a second LLM or a classifier to evaluate whether the response is acceptable. This is the approach used in [LLM evaluation frameworks](https://docs.confident-ai.com/).

## AI Coding Tools

### CLI-Based Tools

| Tool | Description | Link |
|------|-------------|------|
| **Claude Code** | Anthropic's CLI agent — reads the codebase, edits files, runs tests, uses tools | [claude.ai/code](https://claude.ai/code) |
| **Aider** | Terminal pair programmer — edits files in a repo, git integration, supports many models | [aider.chat](https://aider.chat/) |

**Claude Code** is particularly powerful for this course because it can read an entire architecture document, understand the context of the project, and make targeted edits across multiple files while running tests to verify the result. It operates in a conversation that maintains context across a session.

**Aider** integrates tightly with git, commits after each successful change, and supports a wide range of models (including local Ollama models). It is excellent for incremental changes to an existing codebase.

### IDE-Integrated Tools

| Tool | Description | Link |
|------|-------------|------|
| **GitHub Copilot** | Inline code completion and chat in VS Code/JetBrains | [github.com/features/copilot](https://github.com/features/copilot) |
| **Cursor** | AI-first IDE with codebase-aware chat, edit, and generation; based on VS Code | [cursor.com](https://cursor.com/) |
| **Windsurf** | AI IDE with "Cascade" agent that can plan and execute multi-file changes | [windsurf.com](https://windsurf.com/) |
| **Google Gemini Code Assist** | AI coding assistant, IDE plugin for VS Code/JetBrains, integrated with Google Cloud | [cloud.google.com/gemini](https://cloud.google.com/gemini/docs/codeassist/overview) |

### Chat and API-Based Tools

| Tool | Description | Link |
|------|-------------|------|
| **ChatGPT / OpenAI Codex** | Code generation via chat interface or API, supports function calling and code interpreter | [platform.openai.com](https://platform.openai.com/) |
| **Google Gemini** | Multimodal AI with code generation, available via web chat and API | [gemini.google.com](https://gemini.google.com/) |

**Cursor** is the most popular AI IDE for professional software development in 2025. Its `@codebase` feature indexes an entire repository and enables questions about the code as well as changes with full codebase context. The [Cursor documentation](https://docs.cursor.com/) is the best starting point.

### How to Use Them Effectively

**Start with context.** Before asking an AI to generate code, provide the architecture document, the relevant sequence diagrams, the API contract for the interface being implemented, and any existing code that should serve as a pattern. An AI with good context produces dramatically better code than one working from a vague description.

**Be specific.** Compare: "write a sensor process" (vague) versus "implement a Python process that: (1) connects to the BuildSim WebSocket at `ws://localhost:8080/stream`, (2) receives sensor reading messages in the format `{sensor_id, value, timestamp}`, (3) validates that `value` is within the physical range for the sensor type, (4) inserts each valid reading into a TimescaleDB table `readings(sensor_id, value, timestamp)` using asyncpg, (5) reconnects automatically on disconnect with exponential backoff. Use the asyncio library. Here is the database schema: ..." (specific).

**Iterate.** AI rarely produces perfect code on the first attempt. Generate, review, run tests, identify gaps, refine the prompt. This cycle takes 3–5 iterations for a non-trivial component.

**Do not trust blindly.** AI-generated code must be read, understood, and tested. Every line of code in the repository is the responsibility of the developer, not the AI that generated it.

**Use AI for boilerplate.** Dockerfiles, docker-compose configuration, `requirements.txt`, API client boilerplate, data model classes, logging setup, and health check endpoints are exactly the kind of repetitive, structured code that AI generates reliably.

**Keep control of architecture.** The structure of the system is a design decision that belongs to the developer, not the AI. AI fills in the implementation. Allowing the AI to refactor architecture risks optimising for code cleanliness at the expense of deliberate design intent.

## Prompting for Code Generation

### Prompt Engineering for Developers

Prompt engineering for code generation is not about clever tricks; it is about being precise. The more precisely a requirement is specified, the better the result. The [Prompt Engineering Guide](https://www.promptingguide.ai/) covers the principles in detail.

**Specification as prompt.** Copying the relevant section of the architecture document directly into the prompt works because that document was written to be precise. Example:

```
Here is the architecture for the sensor ingestion component from my design document:

[paste your sensor process architecture section]

Implement this component in Python. Use the asyncio library for concurrency.
Follow this data model: [paste your data model].
The BuildSim API is documented here: [paste API docs excerpt].
```

**Few-shot prompting.** Showing one complete, working example and asking for a similar one: "Here is a working temperature sensor process [paste code]. Now implement an equivalent process for CO2 sensors. The CO2 sensor readings are in ppm (0–5000) and the alert threshold is 1000 ppm."

**Test-first prompting.** Providing the test file and asking for the implementation: "Here are the pytest tests for the smoke classifier [paste tests]. Write the `SmokeClassifier` class that makes all tests pass. Do not modify the tests."

**Constraint prompting.** Explicit constraints prevent common AI failure modes. Examples:
- "Do not use any external libraries beyond what is listed in requirements.txt"
- "All database queries must use parameterised statements — no string formatting"
- "All configuration values must come from environment variables, not hardcoded"
- "Every function must have a docstring and type annotations"

**Incremental generation.** For complex systems, generating one component at a time reduces errors: first the data models (no dependencies), then the database client (depends on data models), then the sensor process (depends on database client and BuildSim API client), then the AI agent (depends on database client and tool definitions), then the actuator process (depends on tool definitions), and finally the docker-compose.yml (depends on all of the above).

**Review prompting.** After generating code, asking the AI to review it with a structured prompt catches many issues: "Review this code for: (1) security issues, (2) race conditions, (3) missing error handling, (4) places where configuration should be externalised, (5) places where logging would help debugging. Output a list of issues with suggested fixes."

### What Makes a Good Prompt

A good code generation prompt contains: what the component does (one clear sentence), input and output (types, formats, example values), dependencies (which libraries, APIs, databases it interacts with), constraints (what it must not do or use), context (existing code it must integrate with), and tests (the tests it must pass).

## Debugging and Verifying AI-Generated Code

### Common Failure Modes

AI-generated code fails in predictable ways. Knowing the patterns allows faster identification of issues.

**Hallucinated APIs.** The AI calls a function or method that does not exist, or uses a real function with the wrong signature. Running the code immediately catches this: a hallucinated API fails on import or first call.

**Plausible but wrong logic.** The code looks correct and passes superficial review but fails on edge cases, such as a sensor averaging function that divides by zero when the list is empty, or a timestamp parser that fails on timestamps with microseconds. Tests covering edge cases (empty inputs, boundary values, malformed data) catch these.

**Missing error handling.** AI tends to write happy-path code. What happens when the database is unavailable? When the BuildSim WebSocket disconnects? When a sensor reading is missing a required field? Explicit requirements in the prompt address this: "add error handling for: database connection failure, WebSocket disconnect, malformed input. Log errors with context. Use exponential backoff for retries."

**Hardcoded values.** AI frequently hardcodes URLs, thresholds, and credentials that should be configuration. Searching generated code for string literals and numbers and moving them to environment variables or a configuration file is a necessary step.

**Security issues.** SQL injection via string formatting, secrets stored in code, no input validation on incoming data. Explicit constraints in the prompt and a security-focused review prompt after generation address these.

### Verification Strategy

The verification sequence is: run the tests (written before generating the code), read the code (understanding every line before accepting responsibility for it), fault injection (killing a process, sending malformed data, simulating a network outage, and observing system behaviour), checking against the specification (verifying the code matches the sequence diagrams and handles every state in the state machine), and review prompting (asking the AI to find its own bugs with a structured review prompt).

### The Human's Role in AI-Driven Development

AI handles boilerplate, data model generation, API client generation, test fixture generation, Dockerfile generation, and repetitive patterns. The developer handles architectural decisions, security review, integration testing, debugging subtle logic errors, deciding what the system should do, and taking responsibility for the result.

The division is roughly: AI writes the code; the developer specifies, reviews, tests, and owns it.

## Containerisation and CI with AI

### Letting AI Handle DevOps Boilerplate

Docker configuration, CI pipeline setup, and deployment configuration are exactly the kind of structured, repetitive work that AI generates reliably. A good prompt:

```
Generate a Dockerfile for this Python service:
- Base image: python:3.12-slim
- Install dependencies from requirements.txt
- Run as a non-root user
- Health check: HTTP GET /health every 30 seconds
- Graceful shutdown on SIGTERM
- Entry point: python -m sensor_process

Also generate a docker-compose.yml for these services:
[paste your container diagram description]
Include: restart policies, health check dependencies, a shared network, and environment variables from .env files.
```

AI-generated CI pipelines (GitHub Actions) are particularly useful for this course. A prompt like "generate a GitHub Actions workflow that: runs pytest on push, builds Docker images, and fails if any test fails" produces a working `.github/workflows/ci.yml` in seconds.

### What Must Be Decided by the Developer

AI does not know the operational requirements of a specific system. The decisions that remain with the developer include: which containers are stateful (needing persistent volumes) versus stateless, how containers communicate (shared network, which ports to expose), what the restart policy should be (always restart safety components; not necessarily the dashboard), what secrets need to be injected (database passwords, API keys) and how, and which services depend on which (the AI agent should not start before the database is ready).

These decisions belong in the architecture document. The `docker-compose.yml` is the executable form of the container diagram, and they should remain consistent with each other.

## Recommended Reading

- "Prompt Engineering Guide" — [promptingguide.ai](https://www.promptingguide.ai/) — comprehensive, free, practical; read the "techniques" section
- Claude Code documentation — [docs.anthropic.com/en/docs/claude-code](https://docs.anthropic.com/en/docs/claude-code) — reference for the tool used in this course
- Aider usage guide — [aider.chat/docs/usage.html](https://aider.chat/docs/usage.html) — good alternative for terminal-native workflow
- Cursor documentation — [docs.cursor.com](https://docs.cursor.com/) — the most popular AI IDE; read the "context" section
- Beck, K. "Test-Driven Development: By Example" (Addison-Wesley) — the foundational TDD reference; the principles apply directly to AI-generated code
- "pytest documentation" — [docs.pytest.org](https://docs.pytest.org/) — the standard Python testing framework; know fixtures, parametrize, and conftest.py
- "Hypothesis: Property-Based Testing for Python" — [hypothesis.readthedocs.io](https://hypothesis.readthedocs.io/en/latest/) — invaluable for testing CPS invariants
- Fowler, M. "Mocks Aren't Stubs" — [martinfowler.com/articles/mocksArentStubs.html](https://martinfowler.com/articles/mocksArentStubs.html) — essential background for testing AI agents with mocks
- "How to Write Better Prompts for Code Generation" — [simonwillison.net](https://simonwillison.net/) — Simon Willison's blog covers practical AI development in depth
