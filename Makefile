TARGET=./build
ARCHS=amd64 386
LDFLAGS="-s -w"
GCFLAGS="all=-trimpath=${GOPATH}/src"
ASMFLAGS="all=-trimpath=${GOPATH}/src"

.DEFAULT_GOAL := all

all: clean windows linux darwin

windows:
	@mkdir -p ${TARGET} ; \
	for GOARCH in ${ARCHS}; do \
		echo "Building for windows $${GOARCH} ..." ; \
		GOOS=windows GOARCH=$${GOARCH} go build -ldflags=${LDFLAGS} -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} -o ${TARGET}/goshell-windows-$${GOARCH}.exe ; \
	done;

linux:
	@mkdir -p ${TARGET} ; \
	for GOARCH in ${ARCHS}; do \
		echo "Building for linux $${GOARCH} ..." ; \
		GOOS=linux GOARCH=$${GOARCH} go build -ldflags=${LDFLAGS} -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} -o ${TARGET}/goshell-linux-$${GOARCH} ; \
	done;

darwin:
	@mkdir -p ${TARGET} ; \
	for GOARCH in ${ARCHS}; do \
		echo "Building for darwin $${GOARCH} ..." ; \
		GOOS=darwin GOARCH=$${GOARCH} go build -ldflags=${LDFLAGS} -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} -o ${TARGET}/goshell-darwin-$${GOARCH} ; \
	done;

clean:
	@rm -rf ${TARGET}/*
