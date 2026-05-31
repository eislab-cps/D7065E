# Top-level build for all course documents.
#
#   make           # build everything (reports + template + proposals)
#   make reports   # only the two D2/LaTeX reports
#   make proposals # only the two short proposal PDFs (LaTeX only)
#   make template  # only the fill-in report template (LaTeX only)
#   make clean     # remove LaTeX intermediates everywhere
#   make distclean # also remove generated PDFs and rendered figures

REPORTS   := final_report_example final_report_guide
PROPOSALS := proposal_example proposal_guide
TEMPLATE  := final_report_template
SUBDIRS   := $(REPORTS) $(PROPOSALS) $(TEMPLATE)

.PHONY: all reports proposals template clean distclean $(SUBDIRS)

# default target: build every document (each as its own PDF)
all: reports template proposals

reports: $(REPORTS)
proposals: $(PROPOSALS)
template: $(TEMPLATE)

# each subdirectory has its own Makefile
$(SUBDIRS):
	$(MAKE) -C $@

clean:
	@for d in $(SUBDIRS); do $(MAKE) -C $$d clean; done

distclean:
	@for d in $(SUBDIRS); do $(MAKE) -C $$d distclean; done
