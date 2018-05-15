.PHONY: build clean
NOW=$(shell date '+%Y-%m-%d_%H:%M:%S')
REV?=$(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags '-X rmqmonitor/version.Build=${NOW}@${REV} -w -s'
BINARY="spiderQ"
build:
	go build -o ./${BINARY} ${LDFLAGS}
clean:
	if test -f ${BINARY}; then \
	rm -f ${BINARY}; fi