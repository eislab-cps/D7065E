# Build the D2 diagrams and the LaTeX guide.
#
# Requires:
#   * d2     -- `go install oss.terrastruct.com/d2@latest`
#               (then make sure "$(go env GOPATH)/bin" is on your PATH)
#   * TeX Live with pdflatex
#
# Usage:
#   make            # render diagrams (if changed) and build the PDF
#   make diagrams   # only (re)render the D2 diagrams
#   make clean      # remove LaTeX intermediate files
#   make distclean  # also remove the PDF and the rendered figures

D2       ?= d2
D2FLAGS  := --theme 4 --pad 25
PDFLATEX := pdflatex -interaction=nonstopmode -halt-on-error
DOC      := report_guide

# context + landscape lay out better with dagre; everything else with elk
DAGRE := context landscape
ELK   := container component dynamic pipeline state deployment

DAGRE_FIGS := $(DAGRE:%=figures/%.png)
ELK_FIGS   := $(ELK:%=figures/%.png)
FIGS       := $(DAGRE_FIGS) $(ELK_FIGS)

.PHONY: all diagrams clean distclean
all: $(DOC).pdf
	@echo "==> $(DOC).pdf ready"

diagrams: $(FIGS)

$(DAGRE_FIGS): figures/%.png: diagrams/%.d2 | figures
	$(D2) $(D2FLAGS) --layout dagre $< $@

$(ELK_FIGS): figures/%.png: diagrams/%.d2 | figures
	$(D2) $(D2FLAGS) --layout elk $< $@

figures:
	mkdir -p figures

# run twice for the table of contents
$(DOC).pdf: $(DOC).tex $(FIGS)
	$(PDFLATEX) $(DOC).tex
	$(PDFLATEX) $(DOC).tex
	@rm -f $(DOC).aux $(DOC).log $(DOC).out $(DOC).toc

clean:
	rm -f $(DOC).aux $(DOC).log $(DOC).out $(DOC).toc

distclean: clean
	rm -f $(DOC).pdf $(FIGS)
