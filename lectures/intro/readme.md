# D7065E — Lecture 1: Course Introduction

LTU-themed Beamer deck with draw.io diagrams, following the same structure and
pedagogy as the D7024E course-introduction lecture.

## Layout

- `intro.tex` — the deck
- `diagrams/*.drawio` — diagram sources (draw.io)
- `media/diagrams/*.png` — generated diagram exports (via `make diagrams`)
- `media/buildsim.png` — BuildSim screenshot (static asset)
- `media/template/` — LTU theme art
- `material/` — source notes for course facts (grading, lab, oral exam, use cases)
- `beamer*LTU.sty` — LTU Beamer theme

## Build

```
make            # export changed diagrams + compile intro.pdf
make diagrams   # only re-export .drawio -> media/diagrams/*.png
make watch      # build and open the PDF
```

Requires `pdflatex` and the draw.io desktop app
(`/Applications/draw.io.app`) for diagram export.

## Diagram style

Dark-slide palette (transparent PNG export on the navy `mainColor #032040`):
node fill `#0B3B66`, stroke/subtle `#89A5BE`, accent `#FF8247`,
light text `#C9D6E2`/white, ok `#7FE3A1`, bad `#E05252`, Helvetica.

## Theme license

Beamer theme (c) 2024 Malte Kerl, MIT license. Media directory content is owned by LTU.
