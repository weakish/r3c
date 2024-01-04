NAME=r3c

PREFIX?=/usr/local
BINDIR=${PREFIX}/bin


${NAME}: ${NAME}.go
	@go build

install: ${NAME}
	@mkdir -p ${BINDIR}
	@install -c -m 755 ${NAME}  ${BINDIR}/${NAME}

uninstall:
	@rm -f ${bindir}/${NAME}

clean:
	@rm -f ${NAME}
