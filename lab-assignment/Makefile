# Top-level build for all course documents.
#
#   make           # build everything (report example + guide + proposals)
#   make reports   # only the two D2/LaTeX reports
#   make proposals # only the two short proposal PDFs (LaTeX only)
#   make clean     # remove LaTeX intermediates everywhere
#   make distclean # also remove generated PDFs and rendered figures

REPORTS   := final_report_example final_report_guide
PROPOSALS := proposal_example proposal_guide

.PHONY: all reports proposals clean distclean $(REPORTS) $(PROPOSALS)

# default target: build the two reports and the two proposals (each as its own PDF)
all: reports proposals

reports: $(REPORTS)
proposals: $(PROPOSALS)

# each subdirectory has its own Makefile
$(REPORTS) $(PROPOSALS):
	$(MAKE) -C $@

clean:
	@for d in $(REPORTS) $(PROPOSALS); do $(MAKE) -C $$d clean; done

distclean:
	@for d in $(REPORTS) $(PROPOSALS); do $(MAKE) -C $$d distclean; done
