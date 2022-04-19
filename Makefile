BINARY="aes"
VERSION=1.0
BUILD=`date +%FT%T%z`

PACKAGES=`go list ./... | grep -v /`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`


default:
	@go build -o ${BINARY}

run: default
	./${BINARY}

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
	  	echo "Please run 'make fmt' and commit the result:"; \
        echo "$${diff}"; \
        exit 1; \
	fi;

install:
	@govendor sync -v

test:
	@go test -cpu=1,2,4 -v tags integration ./...

vet:
	@go vet ${VETPACKAGES}

docker:
	@docker build -t zhangyan/xxxxx:latest .

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default run list fmt fmt-check install test vet docker clean