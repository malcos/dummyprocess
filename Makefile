export GO111MODULE=on

COVERAGE_FILE=coverage.out

EXECUTABLE=dummy
NATIVE=$(EXECUTABLE)$(if $(filter windows,$(shell go env GOHOSTOS)),.exe)
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
ALL=$(NATIVE) $(LINUX) $(DARWIN) $(WINDOWS)

build: clean $(NATIVE)

clean:
	rm -f $(ALL) $(COVERAGE_FILE)

all: clean $(ALL)

test:
	go test -count=1 ./...

coverage:
	go test -cover -coverprofile $(COVERAGE_FILE) -count=1 ./...

viewcoverage: coverage
	go tool cover -html=$(COVERAGE_FILE)

$(NATIVE):
	go build -o $@ ./cmd

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -o $@ ./cmd

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -o $@ ./cmd

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -o $@ ./cmd
