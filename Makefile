include ./Makefile.Common

AWS_PROFILE := # configure aws profile here
NAME := doc-box
LAMBDA_DIR := ./cmd
BUILD_DIR := ./build

# list of aws lambda functions (see cmd folder)
ALL_LAMBDAS = \
	createDocument \
	attachDocument

.DEFAULT_GOAL := all

.PHONY: all
all: goimpi golint gotest build

.PHONY: clean
clean:
	@rm -rf ${BUILD_DIR}

.PHONY: build
build: $(addprefix build-, $(ALL_LAMBDAS))

.PHONY: build-%
build-%:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ${BUILD_DIR}/$*/bootstrap ${LAMBDA_DIR}/$*

.PHONY: zip
zip: build $(addprefix zip-, $(ALL_LAMBDAS))

.PHONY: zip-%
zip-%:
	zip -j ${BUILD_DIR}/$*/bootstrap.zip ${BUILD_DIR}/$*/bootstrap

deploy-%:
	aws lambda update-function-code --function-name ${NAME}-$* --zip-file fileb://${BUILD_DIR}/$*/bootstrap.zip --no-cli-pager --profile ${AWS_PROFILE}

$(ALL_LAMBDAS):
	$(MAKE) build-$@
	$(MAKE) zip-$@
	$(MAKE) deploy-$@

.PHONY: unit-tests-with-cover
unit-tests-with-cover:
	@echo "running go unit test + coverage in `pwd`"
	@$(GO_TEST) $(GO_TEST_OPT_WITH_COVERAGE) $(ALL_BUSINESS_SRC_DIRS)
	go tool cover -html=coverage.txt -o coverage.html

.PHONY: integration-tests-with-cover
integration-tests-with-cover:
	@if [ -z "${AWS_SESSION_TOKEN}" ]; then \
		echo "Missing aws secret token. Please run \"assume-role\" with your profile."; \
		exit 1; \
	fi
	@echo "running go integration test coverage in `pwd`"
	$(GO_TEST) $(GO_TEST_OPT_WITH_INTEGRATION) $(ALL_BUSINESS_SRC_DIRS)
	go tool cover -html=integration-coverage.txt -o integration-coverage.html

.PHONY: execute-qs-check
execute-qs-check: goimpi golint unit-tests-with-cover build syft-and-grype

.PHONY: show-env
show-env:
	@echo "codebuild environment -> node version: `node --version`, yarn version: `yarn --version`, `go version`"