# https://github.com/charmbracelet/


### STANDARD TEMPLATE below...

MAKE_FSPATH=$(PWD)/../../../gio-make
include $(MAKE_FSPATH)/help.mk
include $(MAKE_FSPATH)/env.mk
include $(MAKE_FSPATH)/vscode.mk
include $(MAKE_FSPATH)/go.mk
include $(MAKE_FSPATH)/gio.mk
include $(MAKE_FSPATH)/caddy.mk

# using a standard templates for testing
#include $(PWD)/.templates/env/git/.env

CADDY_SRC_FSPATH=$(PWD)

all: this-dep-os this-dep this-test this-build

this-print:
	$(MAKE) env-print
	$(MAKE) vscode-print
	$(MAKE) go-print
	$(MAKE) gio-print
	$(MAKE) caddy-print
	
dep-os:
	$(MAKE) env-dep
	$(MAKE) vscode-dep
	$(MAKE) go-dep
	$(MAKE) gio-dep
	$(MAKE) caddy-dep

this-dep: this-dep-delete
	git clone https://github.com/charmbracelet/soft-serve
	git clone https://github.com/charmbracelet/charm
	git clone https://github.com/charmbracelet/glow
this-dep-delete:
	rm -rf soft-serve
	rm -rf charm
	rm -rf glow

this-build: soft-build charm-build

this-serve:
	$(MAKE) caddy-run

### SOFT

SOFT_LIB_NAME=soft-serve

SOFT_GO_SRC_NAME=soft-server
SOFT_GO_SRC_FSPATH=$(PWD)/$(SOFT_LIB_NAME)/cmd/soft
SOFT_GO_MAKE=GO_SRC_NAME=$(SOFT_GO_SRC_NAME) GO_SRC_FSPATH=$(SOFT_GO_SRC_FSPATH)

# no used yet...
GIO_SRC_NAME=gioclient
GIO_SRC_FSPATH=$(PWD)/gio-client

soft-print:
	$(MAKE) $(SOFT_GO_MAKE) go-print
soft-build:
	$(MAKE) $(SOFT_GO_MAKE) go-build
soft-buildrun: soft-build
	$(MAKE) $(SOFT_GO_MAKE) go-buildrun
	# Starting SSH server on :23231
soft-test:
	#git clone ssh://github.com/charmbracelet/soft-serve


### CHARM

CHARM_LIB_NAME=charm

CHARM_GO_SRC_NAME=charm
CHARM_GO_SRC_FSPATH=$(PWD)/$(CHARM_LIB_NAME)
CHARM_GO_MAKE=GO_SRC_NAME=$(CHARM_GO_SRC_NAME) GO_SRC_FSPATH=$(CHARM_GO_SRC_FSPATH)

# no used yet...
GIO_SRC_NAME=gioclient
GIO_SRC_FSPATH=$(PWD)/gio-client

charm-print:
	$(MAKE) $(CHARM_GO_MAKE) go-print
charm-build:
	$(MAKE) $(CHARM_GO_MAKE) go-build
charm-buildrun: charm-build
	$(MAKE) $(CHARM_GO_MAKE) go-buildrun
	# Starting SSH server on :23231
charm-test:
	#git clone ssh://github.com/charmbracelet/charm-serve






	