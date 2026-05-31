# Diagrams as Code, Generating Architecture Diagrams

A hands-on tutorial for D7065E.

---

## Part 1 — Why Diagrams, and Why as Code

The earlier chapters argued that the architecture is the specification: the system is modelled before it is built. Most of that model is diagrams, context views, container views, component views, and sequence diagrams. So the practical question is how those diagrams should be made.

There are two broad ways to make a diagram. The first is to **draw it by hand** in a graphical editor, dragging boxes and arrows around a canvas until the picture looks right. The second is to **describe it in text** and let a tool compute the layout and render the picture. The second approach is usually called *diagrams as code*, and for an engineering course built around version control and reproducibility it is the better default.

The reasons are the same reasons code lives in Git rather than in a pile of screenshots.

**Diffable.** A text diagram is a few lines of plain text. When the architecture changes, the diff shows exactly which nodes and edges changed. A hand-drawn diagram is an opaque binary or a tangle of XML where moving one box rewrites half the file.

**Reproducible.** The same source always renders to the same picture. There is no slow drift where one diagram uses one box style and the next uses another because a week passed between drawing them.

**Automatable.** Rendering is a command, so it can run in a `Makefile` or a CI pipeline. The diagrams in the lab assignment are built exactly this way, one `make` rebuilds every figure and every PDF.

**Layout is handled automatically.** A good tool computes positions and routes the arrows. Time is spent deciding *what connects to what*, which is the architectural question, not nudging boxes by a few pixels, which is not.

The cost is that a hand editor gives pixel-perfect control and a text tool does not. For architecture diagrams that is the right trade: the goal is a clear, consistent, maintainable model, not a bespoke illustration.

---

## Part 2 — The Landscape of Tools

Diagrams-as-code is a crowded field. The table below sketches the main options and what they are good for, so the choice for this course is an informed one rather than a default.

| Tool | Language / runtime | Strengths | Watch out for |
|---|---|---|---|
| **Draw.io / diagrams.net** | GUI (XML files) | Free, familiar, full manual control | Not text-first; diffs are noisy; layout is manual |
| **Mermaid** | JavaScript | Renders inline on GitHub and in many wikis; easy to start | Limited layout control; large diagrams get messy |
| **PlantUML** | Java | Very broad UML and C4 support; mature | Needs a Java runtime; dated default styling |
| **Graphviz (DOT)** | C | Powerful, battle-tested layout engine | Low-level; verbose for rich diagrams |
| **Structurizr** | Java / DSL | Purpose-built for the C4 model | Java toolchain; more setup than a single binary |
| **D2** | Go (single binary) | Modern syntax, good layouts and themes, no runtime deps | Younger project; smaller ecosystem |
| **Excalidraw** | GUI | Quick, hand-drawn look for sketches | Manual; not a model that can be regenerated |
| **TikZ / pgfplots** | LaTeX | Precise figures and charts inside documents | Steep; better for charts than box-and-line architecture |

For this course the recommended tool is **D2**. It is a single Go binary with no Java dependency, the syntax is readable, the layout engines and themes produce consistent results, and it covers everything the C4 model needs. **Mermaid** is a reasonable lighter alternative when a diagram needs to render directly inside a GitHub README; **pgfplots** is the right choice for evaluation *charts* (it is used for exactly that in the report example). The rest of this tutorial is a working introduction to D2.

---

## Part 3 — Using D2

### Installing

D2 is distributed as a single binary. With a Go toolchain installed:

```
go install oss.terrastruct.com/d2@latest
```

This drops the `d2` binary in `$(go env GOPATH)/bin`, so make sure that directory is on the `PATH`. There is also an install script and Homebrew formula on the project site for those who prefer not to use Go.

### The basics: shapes, connections, labels

A D2 file is a set of nodes and the connections between them. A node is declared by naming it; an arrow is `->`. Text after a colon is a label.

```d2
direction: right

manager: Facility Manager {
  shape: person
}
control: Ventilation Control
buildsim: BuildSim {
  style.stroke-dash: 4
}

manager -> control: monitors
control -> buildsim: reads sensors, writes actuators
```

Rendered with the course theme, that produces:

<figure class="diagram">
<img src="figures/diagrams-as-code-fig01.svg" alt="A small D2 context diagram: a facility manager monitoring a ventilation control system that talks to BuildSim">
<figcaption><em>A small context-style diagram. Three nodes, two connections, and the layout is computed automatically. The dashed node marks an external system.</em></figcaption>
</figure>

A few things are worth noticing. `direction: right` lays the diagram out left-to-right. `shape: person` is one of D2's built-in shapes; others include `cylinder` (a datastore), `queue` (a message broker), and `rectangle` (the default). `style.stroke-dash: 4` draws a dashed border, a useful convention for "this is outside the system being designed".

### Containers and boundaries

Nesting nodes inside a `{ }` block groups them. This is exactly how a C4 boundary is drawn: the components live inside a dashed box that represents the system.

```d2
direction: right

control: Ventilation Control {
  style.stroke-dash: 4
  style.fill: transparent

  sensors: Sensor Service
  decider: Autonomous Decider
  history: History {
    shape: cylinder
  }

  sensors -> history: store readings
  decider -> history: query recent
}

buildsim: BuildSim

control.sensors -> buildsim: read state
control.decider -> buildsim: write commands
```

<figure class="diagram">
<img src="figures/diagrams-as-code-fig02.svg" alt="A D2 container diagram: a dashed Ventilation Control boundary containing a sensor service, an autonomous decider, and a history datastore, all interacting with BuildSim">
<figcaption><em>A container-style diagram. The dashed, transparent box is the system boundary; its inner nodes are the containers. Connections can cross the boundary by qualifying the node name, <code>control.sensors</code>.</em></figcaption>
</figure>

The `style.fill: transparent` keeps the boundary box from painting a solid colour over the diagram, so the inner containers stand out, a small detail that makes C4 diagrams much more readable.

### Layout engines and themes

D2 separates *what* the diagram contains from *how* it is laid out. Two flags control the look:

- `--layout` chooses the layout engine. `dagre` (the default) is fast and good for label-dense diagrams; `elk` handles larger, nested diagrams more cleanly. A third engine, `tala`, is available separately.
- `--theme` chooses the colour theme by number. This course standardises on theme **4, Cool Classics**, so every diagram across the notes and the lab shares one visual language.

A typical render command:

```
d2 --theme 4 --layout elk --pad 25 control.d2 figures/control.svg
```

`--pad 25` adds a margin so the diagram does not crowd the page edge. The output format is chosen by the **file extension**: `.svg` for vector output (sharp at any size, ideal for these notes) and `.png` for a raster image (needed when a tool downstream cannot consume SVG).

To generate a PNG instead, change the extension:

```
d2 --theme 4 --layout elk --pad 25 control.d2 figures/control.png
```

A higher-resolution PNG is produced by scaling the render up:

```
d2 --theme 4 --layout elk --pad 25 --scale 2 control.d2 figures/control.png
```

The first PNG export downloads a small headless browser that D2 uses to rasterise the SVG; this happens once and is then cached, so later renders are fast. SVG output has no such dependency. The two example figures above are SVG, generated with these commands (swap the `.svg` extension for `.png` to get raster versions):

```
d2 --theme 4 --layout dagre --pad 25 context.d2   figures/diagrams-as-code-fig01.svg
d2 --theme 4 --layout elk   --pad 25 container.d2  figures/diagrams-as-code-fig02.svg
```

### Automating with Make

Because rendering is a single command, it belongs in a build file. The lab assignment renders every `.d2` source to a matching figure with a pattern rule, the shape of which is:

```make
D2FLAGS := --theme 4 --pad 25

figures/%.svg: diagrams/%.d2
	d2 $(D2FLAGS) --layout elk $< $@
```

Typing `make` then re-renders only the diagrams whose source changed, and rebuilds the documents that include them. The diagram is part of the build, not a manual export step that is easy to forget.

### Practical tips

- **Keep labels short.** Layout quality drops fast when nodes carry sentences. Put the detail in the surrounding text, not in the box.
- **Use containers for every C4 boundary.** Nesting is the natural way to express "these parts belong to this system".
- **Make boundary boxes transparent.** `style.fill: transparent` on a dashed container keeps the inside readable.
- **Pick one theme and one layout engine and stay with them.** Consistency across a document is worth more than the perfect look of any single diagram.
- **Let the engine route the arrows.** If a diagram looks tangled, the fix is usually fewer connections or a clearer grouping, not manual repositioning, which D2 does not offer anyway.

---

## Part 4 — D2 and the C4 Model

The C4 model from Chapter 1 maps cleanly onto D2's building blocks, which is much of why it is a good fit for this course.

| C4 element | D2 construct |
|---|---|
| Person (actor) | a node with `shape: person` |
| Software system / container / component | a plain node, labelled with its responsibility |
| System or container **boundary** | a nested `{ }` container, usually dashed and transparent |
| Datastore | a node with `shape: cylinder` |
| Message broker / queue | a node with `shape: queue` |
| Relationship | a connection `->` with a label describing the interaction |
| Dynamic (sequence) diagram | a node with `shape: sequence_diagram` |

Every diagram in the lab assignment is drawn this way. Both the report guide (`lab-assignment/final_report_guide/`) and the worked example (`lab-assignment/final_report_example/`) render their figures from D2 sources with the same theme and Makefile pattern shown above. The example uses exactly these constructs to draw the full C4 set, Context, Container, Component, and a Dynamic diagram, plus a deployment view. Reading those `.d2` sources alongside this tutorial is the fastest way to see the conventions applied to a complete system. The point is not the notation for its own sake: a diagram that can be diffed, reviewed, and regenerated is part of the engineering record, in the same way the code is.

---

## References

- **D2** — language documentation, install instructions, themes, and layout engines: <https://d2lang.com>
- **D2 source** — the project repository: <https://github.com/terrastruct/d2>
- **The C4 model** — Simon Brown's notation for visualising software architecture: <https://c4model.com>
- **Mermaid** — text-based diagrams that render inline on GitHub: <https://mermaid.js.org>
- **PlantUML** — text-based UML and C4 diagrams (Java): <https://plantuml.com>
- **Graphviz** — the DOT language and layout engines: <https://graphviz.org>
- **Structurizr** — a C4-specific DSL and tooling: <https://structurizr.com>
- **diagrams.net (draw.io)** — the graphical editor: <https://www.drawio.com>
- **pgfplots** — LaTeX charts, used for the evaluation plots in the report example: <https://ctan.org/pkg/pgfplots>
