# Final report template (LaTeX) — fill-in skeleton

The **blank skeleton** of the architecture document, the structure to fill in for the
final report. It carries the same 13 sections as the worked example, each with a short
grey guidance note (what the section must contain) and bracketed slots to replace.

Use it alongside the other two documents:

- `../final_report_guide/` — explains each section and the ideas behind it (the *why*).
- `../final_report_example/` — a complete version at the grade-5 standard (the *what good looks like*).

## Files
- `final_report_template.tex` — the template **source** (start here).
- `final_report_template.pdf` — the compiled skeleton.

## How to use
1. Replace every `[bracketed slot]` with the project's own content.
2. Insert the five diagrams where the slots are marked (context, container, component,
   dynamic, deployment). See course notes Chapter 5 for rendering diagrams with D2.
3. Fill the tables (summary, requirements, design decisions, test plan).
4. Replace the pgfplots placeholder in Section 11 with measured results.
5. **Delete the grey guidance notes** and run the final checklist before submitting.

## Build
```
make            # build the PDF (runs pdflatex twice for the table of contents)
make clean      # remove LaTeX intermediates
make distclean  # also remove the PDF
```
The template compiles with `pdflatex` alone, no diagram toolchain is needed until the
team adds its own figures.
