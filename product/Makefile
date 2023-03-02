CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=route256/product/cmd/app

all: format build lint

build: generate

lint: install-lint
	${LINTBIN} run

precommit: format build lint
	echo "OK"

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

generate:
	mkdir -p pkg/product
	protoc -I api/product -I ../vendor-proto \
	--go_out=pkg/product --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=../bin/protoc-gen-go \
	--go-grpc_out=pkg/product --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=../bin/protoc-gen-go-grpc \
	api/product/product_service.proto