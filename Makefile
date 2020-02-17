SHELL:=/bin/bash

# AWS RESOURCES SHARED WITH OTHER MAKEFILES

PROJECT_STACK_NAME?=
BUILD_STACK_NAME?=

# GITLAB CI
CI_REGISTRY ?=
CI_COMMIT_SHA ?=
CI_PROJECT_PATH ?=

# STATIC FUNCTIONS 

GIT_BRANCH?=$(shell git symbolic-ref HEAD 2>/dev/null)
GIT_COMMIT?=$(shell git rev-list HEAD -1)

# LOAD MAKEFILES
include Go.mk Docker.mk Cloudformation.mk

# END LOAD MAKEFILES
hello:
	@$(call gobuild,$@,./cmd/hello,$(GIT_COMMIT),$(subst refs/heads/,,$(GIT_BRANCH)))

build: hello



## Phony non existing target, as we cannot show we have anything, when building Docker images
.PHONY=deploy
deploy:
	@$(call docker-build, tag1 tag2 tag3)


.PHONY=clean
clean: hello
	rm -rf $<