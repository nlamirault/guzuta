# Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

APP="guzuta"
EXE="scaleway"

SHELL = /bin/bash

DIR = $(shell pwd)
GO_PATH = .

DOCKER = docker

GB=$(GOPATH)/bin/gb
GODEP= $(GOPATH)/bin/godep
GOLINT= $(GOPATH)/bin/golint
ERRCHECK= $(GOPATH)/bin/errcheck
GOVER= $(GOPATH)/bin/gover
GOVERALLS= $(GOPATH)/bin/goveralls

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

SRC=src/github.com/nlamirault/guzuta

VERSION=$(shell \
        grep "const Version" $(SRC)/version/version.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

PACKAGE=$(APP)-$(VERSION)
ARCHIVE=$(PACKAGE).tar

all: help

help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@echo -e "$(WARN_COLOR)init$(NO_COLOR)    :  Install requirements"
	@echo -e "$(WARN_COLOR)deps$(NO_COLOR)    :  Install dependencies"
	@echo -e "$(WARN_COLOR)build$(NO_COLOR)   :  Make all binaries"
	@echo -e "$(WARN_COLOR)test$(NO_COLOR)    :  Launch unit tests"
	@echo -e "$(WARN_COLOR)style$(NO_COLOR)   :  Check golang style"
	@echo -e "$(WARN_COLOR)clean$(NO_COLOR)   :  Cleanup"
	@echo -e "$(WARN_COLOR)reset$(NO_COLOR)   :  Remove all dependencies"
	@echo -e "$(WARN_COLOR)release$(NO_COLOR) :  Make a new release"

clean:
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm -f $(EXE) $(APP)-*.tar.gz coverage.out guzuta.test gover.coverprofile

.PHONY: init
init:
	@echo -e "$(OK_COLOR)[$(APP)] Install requirements$(NO_COLOR)"
	@go get -u github.com/golang/glog
	@go get -u github.com/constabulary/gb/...
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/kisielk/errcheck

deps:
	@echo -e "$(OK_COLOR)[$(APP)] Install dependancies$(NO_COLOR)"
	@GOPATH=$(GO_PATH) $(GODEP) restore

build:
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	$(GB) build all

doc:
	@GOPATH=$(GO_PATH) godoc -v -http=:6060

fmt:
	@echo -e "$(OK_COLOR)[$(APP)] Launch fmt $(NO_COLOR)"
	@go fmt ./...

errcheck:
	@echo -e "$(OK_COLOR)[$(APP)] Launch errcheck $(NO_COLOR)"
	@$(ERRCHECK) ./...

vet:
	@echo -e "$(OK_COLOR)[$(APP)] Launch vet $(NO_COLOR)"
	@go vet github.com/nlamirault/guzuta

lint:
	@echo -e "$(OK_COLOR)[$(APP)] Launch golint $(NO_COLOR)"
	@$(GOLINT) src/github.com/nlamirault/guzuta/...

style: fmt vet lint

test:
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@$(GB) test all

coverage:
	@echo -e "$(OK_COLOR)[$(APP)] Launch code coverage $(NO_COLOR)"
	@go test ./... -cover

coveralls:
	@GOPATH=$(GO_PATH) go get github.com/mattn/goveralls
	@GOPATH=$(GO_PATH) $(DIR)/addons/coverage --coveralls

release: clean build
	@echo -e "$(OK_COLOR)[$(APP)] Make archive $(VERSION) $(NO_COLOR)"
	@rm -fr $(PACKAGE) && mkdir $(PACKAGE)
	@cp -r $(EXE) $(PACKAGE)
	@tar cf $(ARCHIVE) $(PACKAGE)
	@gzip $(ARCHIVE)
	@rm -fr $(PACKAGE)
	@addons/github.sh $(VERSION)

# for go-projectile
gopath:
	@echo ${GOPATH}
