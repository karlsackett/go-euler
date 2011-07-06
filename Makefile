include $(GOROOT)/src/Make.inc

TARG=
GOFILES=\
	p01.go\
OBJECTS=$(SOURCES:.go=.${O})

%.${O} : %.go
	${O}g -o $@ $<
% : %.${O}
	${O}l -o $@ $<

#include $(GOROOT)/src/Make.cmd

