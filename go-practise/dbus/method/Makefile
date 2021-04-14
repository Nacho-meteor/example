PROG_EXAMPLE_DAEMON=dbus-example

PROG_FILE_CONF=com.practise.Example.conf
PROG_FILE_SYSTEMD=dbus-example.service

DIR_CONF=misc/conf
DIR_SYSTEMD=misc/systemd

DIR_EXAMPLE_DAEMON=cmd/

PREFIX=/usr

PWD=$(shell pwd)
GOCODE=/usr/share/gocode
GOPATH_DIR=gopath
GOPKG_PREFIX=uos-resource-manager
export GO111MODULE=off

all: build


build: 
		go build -o ${PWD}/${PROG_EXAMPLE_DAEMON} ./${DIR_EXAMPLE_DAEMON}

install:
		cp -rf ${PROG_EXAMPLE_DAEMON} ${PREFIX}/sbin
		cp -rf ${DIR_CONF}/${PROG_FILE_CONF} ${PREFIX}/share/dbus-1/system.d/
		cp -rf ${DIR_SYSTEMD}/${PROG_FILE_SYSTEMD} ${PREFIX}/lib/systemd/system/

uninstall:
		rm ${PREFIX}/sbin/${PROG_EXAMPLE_DAEMON}
		rm ${PREFIX}/share/dbus-1/system.d/${PROG_FILE_CONF}
		rm ${PREFIX}/lib/systemd/system/${PROG_FILE_SYSTEMD}

clean:
		rm -rf ${GOPATH_DIR} ${PROG_EXAMPLE_DAEMON}

rebuild: clean build