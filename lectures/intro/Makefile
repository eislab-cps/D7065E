# D7065E Lecture 1 — build the diagrams and the presentation
#
#   make            build everything (diagrams that changed + PDF)
#   make diagrams   re-export changed .drawio files to media/diagrams/*.png
#   make pdf        compile the Beamer deck
#   make watch      open the PDF
#   make clean      remove LaTeX aux files
#   make distclean  also remove generated PNGs and the PDF

DRAWIO   := /Applications/draw.io.app/Contents/MacOS/draw.io
LATEX    := pdflatex -interaction=nonstopmode
DECK     := intro

SRC_DIAGRAMS := $(wildcard diagrams/*.drawio)
PNG_DIAGRAMS := $(patsubst diagrams/%.drawio,media/diagrams/%.png,$(SRC_DIAGRAMS))

.PHONY: all diagrams pdf watch clean distclean

all: diagrams pdf

diagrams: $(PNG_DIAGRAMS)

# navy background from the .drawio files, 2x scale (-s 2), 12px border (-b 12)
media/diagrams/%.png: diagrams/%.drawio
	@mkdir -p media/diagrams
	$(DRAWIO) -x -f png -s 2 -b 12 -o $@ $< >/dev/null 2>&1
	@echo "exported $@"

pdf: $(DECK).pdf

$(DECK).pdf: $(DECK).tex $(PNG_DIAGRAMS) $(wildcard beamer*LTU.sty)
	$(LATEX) $(DECK).tex >/dev/null || ($(LATEX) $(DECK).tex | grep -A2 "^!"; exit 1)
	@$(MAKE) --no-print-directory clean
	@echo "built $(DECK).pdf"

watch: pdf
	open $(DECK).pdf

clean:
	rm -f $(DECK).aux $(DECK).log $(DECK).nav $(DECK).out $(DECK).snm $(DECK).toc $(DECK).vrb

distclean: clean
	rm -f $(DECK).pdf $(PNG_DIAGRAMS)
