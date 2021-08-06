TEST?=$$(go list ./... | grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
PKG_NAME=influxdb-v2

default: build

initialize:
	@sh -c "scripts/config_influxdb.sh" -timeout 10m

build: fmtcheck
	go install
	
clean:
	@rm terraform-provider-influxdb-v2
	@go clean -testcache

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

testacc: fmtcheck fmt
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m


fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

fmt:
	gofmt -w $(GOFMT_FILES)

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)
    
stop-influx:
	docker stop tf_acc_tests_influxdb
	docker rm tf_acc_tests_influxdb

.PHONY: build test initialize testacc vet fmt fmtcheck errcheck test-compile
