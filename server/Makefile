# install
GOMOCK_VERSION = v1.6.0 
SQLBOILER_VERSION = v4.14.0

go-install:
	go install github.com/golang/mock/mockgen@${GOMOCK_VERSION}
	go install github.com/volatiletech/sqlboiler/v4@${SQLBOILER_VERSION}
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@${SQLBOILER_VERSION}

# protobuf make
BIN_DIR = bin
SERVER_DIR = server

ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	SHELL_VERSION = $(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim())
	OS = $(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim())
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
	CHECK_DIR_CMD = if (!(Test-Path $@)) { $$e = [char]27; Write-Error "$$e[31mDirectory $@ doesn't exist$${e}[0m" }
	HELP_CMD = Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
	RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
	RM_RF_CMD = ${RM_F_CMD} -Recurse
	SERVER_BIN = ${SERVER_DIR}.exe
else
	SHELL := bash
	SHELL_VERSION = $(shell echo $$BASH_VERSION)
	UNAME := $(shell uname -s)
	VERSION_AND_ARCH = $(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		OS = macos ${VERSION_AND_ARCH}
	else ifeq ($(UNAME),Linux)
		OS = linux ${VERSION_AND_ARCH}
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
	CHECK_DIR_CMD = test -d ./proto || (echo "\033[31mDirectory proto doesn't exist\033[0m" && false)
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
	SERVER_BIN = ${SERVER_DIR}
endif

.DEFAULT_GOAL := help
.PHONY: help

# user: $@ ## Generate Pbs and build for user
# income_and_expenditure: $@ ## Generate Pbs and build for income_and_expenditure
# user_tag: $@ ## Generate Pbs and build for user_tag
# admin_tag: $@ ## Generate Pbs and build for admin_tag
# group: $@ ## Generate Pbs and build for group
## protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. --grpc-gateway_opt=module=${PACKAGE} --grpc-gateway_out=. --openapiv2_opt=allow_merge=true --openapiv2_out=${OPENAPI_OUT_DIR} ${PROTO_DIR}/*.proto

test: ## Launch tests
	go test ./...

clean: clean_greet clean_calculator clean_blog ## Clean generated files
	${RM_F_CMD} ssl/*.crt
	${RM_F_CMD} ssl/*.csr
	${RM_F_CMD} ssl/*.key
	${RM_F_CMD} ssl/*.pem
	${RM_RF_CMD} ${BIN_DIR}

rebuild: clean all ## Rebuild the whole project

bump: all ## Update packages version
	go get -u ./... 

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

help: ## Show this help
	@${HELP_CMD}

## build
BUILD_TAGS_PRODUCTION='production'
BUILD_TAGS_DEVELOPMENT='development unittest'

build-base:
	go build -o ./build/${BIN_NAME} -tags '$(BUILD_TAGS) go' -installsuffix go -ldflags '-s -w' ./cmd/server.go

build-development:
	$(MAKE) build-base BUILD_TAGS=${BUILD_TAGS_DEVELOPMENT} BIN_NAME=dev

build-production-linux:
	$(MAKE) build-base BUILD_TAGS=${BUILD_TAGS_PRODUCTION} CGO_ENABLED=0 GOOS=linux GOARCH=amd64 BIN_NAME=prod

sqlboiler:
	sqlboiler mysql -c sqlboiler_local.toml

TEST_FLAGS ?= -cover -race
TEST_TAGS ?= ""
COVER_PROFILE ?= ""

docker-test:
	docker compose run backend go test ./... ${TEST_FLAGS} -tags=${TEST_TAGS} -coverprofile=${COVER_PROFILE}
