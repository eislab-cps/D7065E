# Example report (LaTeX) — the expected grade-5 standard

A complete, worked architecture document for the BuildSim use case
**energy-aware ventilation & climate control**. It is the quality bar for your own
report: requirements trace through design to validation (MBSE), the system is
reasoned about as a whole (systems thinking), and the evaluation shows measured
evidence plus an honest critique (critical thinking).

## Files
- `final_report_example.pdf` — the compiled report.
- `final_report_example.tex` — the report **source** (edit this). Performance plots
  are native `pgfplots` (no external tool).
- `diagrams/*.d2` — the architecture/behaviour diagrams as **D2** source.
- `figures/figN.png` — the diagrams rendered from the `.d2` files.

## Rebuild
Compile the LaTeX (run twice for the table of contents and pgfplots):
```
pdflatex final_report_example.tex
```
Re-render a diagram after editing its D2 source (D2 is a Go tool: `go install oss.terrastruct.com/d2@latest`):
```
d2 --layout elk --pad 25 diagrams/container.d2 figures/fig2.png
```
Diagram → figure map: context=fig1, container=fig2, sensor-sim=fig3, pipeline=fig4,
sequence=fig5, state=fig6, deployment=fig7. The three evaluation charts are `pgfplots`
inside the `.tex`.
