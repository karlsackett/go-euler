# Copyright 2011 Karl Sackett <karl.sackett@gmail.com>
# Use of this source code is governed by the MIT
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=\
	p01\
	p02\

%.${O}: %.go
	${O}g -o $@ $<
%: %.${O}
	${O}l -o $@ $<

all: $(TARG)

clean:
	rm -rf *.o *.a *.[$(OS)] [$(OS)].out

nuke: clean
	rm -f $(TARG)

