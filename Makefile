.PHONY: test build clean


NAME = go-print-file
BINARY_OUT = bin
IMAGE_OUT = image

all: test build gen

test:
	go test .

build:
	@mkdir -p ${BINARY_OUT}
	CGO_ENABLED=0 go build -o ${BINARY_OUT}/${NAME} .
	strip ${BINARY_OUT}/${NAME}

gen:
	@mkdir -p ${IMAGE_OUT}
	${BINARY_OUT}/${NAME} -o ${IMAGE_OUT}/image.png



clean:
	rm -rf ${BINARY_OUT} ${IMAGE_OUT}

