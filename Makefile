include $(GOROOT)/src/Make.inc

TARG=
SOURCES=\
	p01.go\
	p02.go\
OBJECTS=$(SOURCES:.go=.${O})

%.${O}: %.go
	${O}g -o $@ $<
%: %.${O}
	${O}l -o $@ $<

#include $(GOROOT)/src/Make.cmd

