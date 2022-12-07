
include ./../assets/make/help.mk
include ./../assets/make/env.mk

MAKE_FSPATH=./make
include $(MAKE_FSPATH)/task.mk

print:


dep-tools:

clean:
	rm -rf test01
test01: clean
	mkdir -p test01 
	cd test01 && gio-make create --git https://github.com/gioui/gio-example --template git