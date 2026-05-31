# D7065E — Lab assignment documents

Course material for the project in **D7065E, Embedded Intelligence at the Edge**.
Each deliverable comes as a *guide* (how to write it, and why) and a *worked example*
(what a strong answer looks like); the final report also has a fill-in *template*:

| Document | Folder | What it is |
|---|---|---|
| **Project proposal (Week 2)** | `proposal_guide/` | Explains what the short proposal must contain. |
| | `proposal_example/` | A deliberately provisional, filled-in proposal. |
| **Final report (architecture document)** | `final_report_guide/` | Explains each section, and the MBSE/C4 ideas behind it. |
| | `final_report_example/` | A complete worked report at the grade-5 standard. |
| | `final_report_template/` | The blank skeleton to fill in, guidance notes and slots. |

The use case throughout the examples is **energy-aware ventilation & climate
control** on a BuildSim office floor.

## Building

From this directory:

```
make            # build everything (reports + template + proposals)
make reports    # only the two D2/LaTeX reports
make template   # only the fill-in report template
make proposals  # only the two short proposal PDFs
make clean      # remove LaTeX intermediates everywhere
make distclean  # also remove generated PDFs and rendered figures
```

`make` produces one PDF per document:

- `final_report_guide/report_guide.pdf`
- `final_report_example/final_report_example.pdf`
- `final_report_template/final_report_template.pdf`
- `proposal_guide/proposal_guide.pdf`
- `proposal_example/proposal_example.pdf`

Each subfolder has its own `Makefile`; the top-level one recurses into them.

## Toolchain

- **LaTeX** (`pdflatex`) for all documents. Each is compiled twice so the table of
  contents and `pgfplots` charts settle.
- **D2** for the architecture and behaviour diagrams (a Go tool, no Java). The
  report Makefiles render each `diagrams/*.d2` to `figures/*.png` (same base name)
  with a uniform theme:

  ```
  d2 --theme 4 --layout elk --pad 25 diagrams/container.d2 figures/container.png
  ```

  Install with `go install oss.terrastruct.com/d2@latest` and make sure
  `$(go env GOPATH)/bin` is on the `PATH`.
- **pgfplots** (native LaTeX) for the evaluation charts in the report example, so
  there is no external charting tool.

## Diagrams (C4 model)

The report diagrams follow the [C4 model](https://c4model.com): System Landscape,
Context (L1), Container (L2), Component (L3), and a Dynamic diagram, plus a
deployment view. Level 4 (Code) is intentionally not drawn by hand, per C4's own
advice. Editing a `.d2` file and re-running `make` re-renders the figure and
rebuilds the affected PDF.
