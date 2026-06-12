# D7065E course notes — export the draw.io figures
#
#   make            export changed .drawio files to figures/*.png
#   make distclean  remove generated PNGs

DRAWIO := /Applications/draw.io.app/Contents/MacOS/draw.io

SRC := $(wildcard diagrams/*.drawio)
PNG := $(patsubst diagrams/%.drawio,figures/%.png,$(SRC))

.PHONY: all distclean

all: $(PNG)

# 2x scale (-s 2), 12px border (-b 12); the navy background is set in each file
figures/%.png: diagrams/%.drawio
	@mkdir -p figures
	$(DRAWIO) -x -f png -s 2 -b 12 -o $@ $< >/dev/null 2>&1
	@echo "exported $@"

distclean:
	rm -f $(PNG)
