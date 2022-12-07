MAKE_FSPATH=$(PWD)/../..

include $(MAKE_FSPATH)/env.mk
include $(MAKE_FSPATH)/vscode.mk
include $(MAKE_FSPATH)/go.mk
include $(MAKE_FSPATH)/gio.mk
include $(MAKE_FSPATH)/caddy.mk

# using a standard templates for testing
include $(PWD)/.env


GO_SRC_NAME=server
GO_SRC_FSPATH=$(PWD)

GIO_SRC_NAME=gioclient
GIO_SRC_FSPATH=$(PWD)/gio-client

CADDY_SRC_FSPATH=$(PWD)

all: dep-os this-test this-build

print:
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
	
this-test:
	# gets a env value
	@echo $(GIT_ORIGIN_USER)

	$(MAKE) env-templates-print
	$(MAKE) env-templates-dep

	$(MAKE) caddy-templates-print
	$(MAKE) caddy-templates-dep
	
this-build:
	$(MAKE) go-build
	$(MAKE) gio-build


this-serve:
	$(MAKE) caddy-run

