# Makefile for devops tools
#
#  This Makefile will have the following goals:
#  default: Just compile the tools that need compilation
#  test: Should run the tests of every tool, being shell script, interpreted or compiled code
#  install: Install the tools inside the environment
#

include Makefile.vars


all: build

.DEFAULT_GOAL: build
.PHONY: build
build:
	@echo build:
	go get 
	go build $(LDFLAGS) -o $(TOOLNAME)

.PHONY: test
test:
		go test


.PHONY: clean
clean:
	if [ -f ${TOOLNAME} ] ; then rm ${TOOLNAME} ; fi
