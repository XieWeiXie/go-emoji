BINARY=go-emoji

VERSION=1.0.0

BUILD=`date +%FT%T%z`

default:
	go build -o ${BINARY}  -tags=jsoniter

install:
	go mod vendor

prod:
	go build -o ${BINARY} -tags=jsoniter

beta:
	go build -o ${BINARY}  -tags=jsoniter

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY:  clean run install dev prod test
