include config.mk

r3c: r3c.go
	@go build

install: r3c
	@mkdir -p ${PREFIX}/bin
	@install r3c  ${PREFIX}/bin

uninstall:
	@rm -f ${PREFIX}/bin/r3c

clean:
	@rm -f r3c