# Build the Lecture 1 slides (LTU beamer theme + D2 diagrams).
#
#   make            # render diagrams, then compile intro.pdf
#   make diagrams   # only (re)render the D2 diagrams to figures/*.pdf
#   make clean      # remove LaTeX intermediates
#   make distclean  # also remove intro.pdf and rendered diagrams

DOC      := intro
PDFLATEX := pdflatex -interaction=nonstopmode -halt-on-error
D2       := d2
# theme 200 (Dark Mauve): light lines on a dark panel, reads on the navy slides
D2FLAGS  := --theme 200 --pad 20

# one figures/<name>.pdf per diagrams/<name>.d2
SRC  := $(wildcard diagrams/*.d2)
FIGS := $(patsubst diagrams/%.d2,figures/%.pdf,$(SRC))

.PHONY: all diagrams clean distclean
all: $(DOC).pdf
	@echo "==> $(DOC).pdf ready"

diagrams: $(FIGS)

figures/%.pdf: diagrams/%.d2
	@mkdir -p figures
	$(D2) $(D2FLAGS) $< $@

# The screenshot figures/buildsim.png is a static asset, not generated here.
# Run pdflatex twice so \cite labels settle.
$(DOC).pdf: $(DOC).tex $(FIGS) $(wildcard *.sty)
	$(PDFLATEX) $(DOC).tex
	$(PDFLATEX) $(DOC).tex
	@rm -f $(DOC).aux $(DOC).log $(DOC).nav $(DOC).out $(DOC).snm $(DOC).toc $(DOC).vrb

clean:
	rm -f $(DOC).aux $(DOC).log $(DOC).nav $(DOC).out $(DOC).snm $(DOC).toc $(DOC).vrb \
	      $(DOC).bbl $(DOC).bcf $(DOC).blg $(DOC).run.xml

distclean: clean
	rm -f $(DOC).pdf $(FIGS)
