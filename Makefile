# Makefile for thingful/httpmock
#
# Targets:
# 	test: runs tests
# 	clean: cleans build artefacts
# 	coverage: generates console coverage report
# 	html: generates html coverage report
#
GOCMD=go
GOTEST=$(GOCMD) test -v
GOCOVER=$(GOCMD) tool cover
ARTEFACT_DIR=./build

.PHONY: test
test:
	mkdir -p $(ARTEFACT_DIR)
	$(GOTEST) -coverprofile=$(ARTEFACT_DIR)/cover.out .

.PHONY: clean
clean:
	rm -rf $(ARTEFACT_DIR)

.PHONY: coverage
coverage: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -func=$(ARTEFACT_DIR)/cover.out

.PHONY: html
html: test
	mkdir -p $(ARTEFACT_DIR)
	$(GOCOVER) -html=$(ARTEFACT_DIR)/cover.out -o $(ARTEFACT_DIR)/coverage.html
