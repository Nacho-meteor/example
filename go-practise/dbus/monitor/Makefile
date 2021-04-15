PROG_TEST=dbus-test

PREFIX=/usr
PWD=$(shell pwd)
GOCODE=/usr/share/gocode
GOPATH_DIR=gopath
GOPKG_PREFIX=dbus-test
export GO111MODULE=off

all: build

prepare:
		@if [ ! -d ${GOPATH_DIR}/src/${GOPKG_PREFIX} ]; then \
				mkdir -p ${GOPATH_DIR}/src/${GOPKG_PREFIX}; \
				ln -sf ${PWD}/pkg ${GOPATH_DIR}/src/${GOPKG_PREFIX}; \
		fi

build: prepare
		env GOPATH=${PWD}/${GOPATH_DIR}:${GOCODE} go build -o ${PWD}/${PROG_TEST} ./

install:
		mkdir -pv ${DESTDIR}${PREFIX}/bin/
		cp -rf ${PROG_TEST} ${DESTDIR}${PREFIX}/bin/

uninstall:
		rm -rf ${PREFIX}/bin/${PROG_TEST}

clean:
		rm -rf ${GOPATH_DIR} ${PROG_TEST}

rebuild: clean build