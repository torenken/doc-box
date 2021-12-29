include ./Makefile.Common

NAME := doc-box
LAMBDA_DIR := ./cmd
BUILD_DIR := ./build

# List of aws lambda functions (see cmd folder)
ALL_LAMBDAS = \
	createDocument

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
	aws lambda update-function-code --function-name ${NAME}-$* --zip-file fileb://${BUILD_DIR}/$*/bootstrap.zip --no-cli-pager

$(ALL_LAMBDAS):
	$(MAKE) build-$@
	$(MAKE) zip-$@
	$(MAKE) deploy-$@