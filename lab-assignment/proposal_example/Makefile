# Build the LaTeX proposal (no diagrams).
#
# Usage:
#   make            # build the PDF
#   make clean      # remove LaTeX intermediate files
#   make distclean  # also remove the PDF

PDFLATEX := pdflatex -interaction=nonstopmode -halt-on-error
DOC      := proposal_example

.PHONY: all clean distclean
all: $(DOC).pdf
	@echo "==> $(DOC).pdf ready"

$(DOC).pdf: $(DOC).tex
	$(PDFLATEX) $(DOC).tex
	$(PDFLATEX) $(DOC).tex
	@rm -f $(DOC).aux $(DOC).log $(DOC).out $(DOC).toc

clean:
	rm -f $(DOC).aux $(DOC).log $(DOC).out $(DOC).toc

distclean: clean
	rm -f $(DOC).pdf
