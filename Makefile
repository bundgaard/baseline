SHELL:=/bin/bash

# AWS RESOURCES SHARED WITH OTHER MAKEFILES

PROJECT_STACK_NAME?=
BUILD_STACK_NAME?=


# STATIC FUNCTIONS 

GIT_BRANCH?=$(shell git symbolic-ref HEAD 2>/dev/null)
GIT_COMMIT?=$(shell git rev-list HEAD -1)

# LOAD MAKEFILES
include Go.mk


# END LOAD MAKEFILES
hello:
	$(call gobuild,$@,./cmd/hello,$(GIT_COMMIT),$(subst refs/heads/,,$(GIT_BRANCH)))

build: hello


.PHONY=clean
clean: hello
	rm -rf $<