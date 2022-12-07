MAKE_FSPATH=$(PWD)/../..

include $(MAKE_FSPATH)/help.mk
include $(MAKE_FSPATH)/env.mk
include $(MAKE_FSPATH)/vscode.mk
include $(MAKE_FSPATH)/git.mk
include $(MAKE_FSPATH)/go.mk
include $(MAKE_FSPATH)/gio.mk
include $(MAKE_FSPATH)/caddy.mk
include $(MAKE_FSPATH)/fly.mk
# using a standard templates for testing
include $(PWD)/.env


GIT_NAME=gio-example
GIT_ORG=gioui
GIT_SERVER=github.com/$(GIT_ORG)/$(GIT_NAME)
GIT_BRANCH=main

# !! toggle this so that you can pick which git remote you want to work on !!
#GIT_WORK_FSPATH=$(PWD)/$(GIT_REPO_ORIGIN_NAME)
GIT_WORK_FSPATH=$(PWD)/$(GIT_REPO_UPSTREAM_NAME)

LIB_NAME=$(GIT_NAME)

GO_SRC_NAME=.
GO_SRC_FSPATH=$(GIT_WORK_FSPATH)/$(GIT_NAME)/$(GO_SRC_NAME)

GIO_SRC_NAME=example
GIO_SRC_FSPATH=$(GIT_WORK_FSPATH)/$(GIT_NAME)/$(GIO_SRC_NAME)

CADDY_SRC_FSPATH=$(PWD)


all: dep-os this-test this-build

print:
	$(MAKE) env-print
	$(MAKE) git-print
	$(MAKE) vscode-print
	$(MAKE) go-print
	$(MAKE) gio-print
	$(MAKE) caddy-print
	
dep-os:
	$(MAKE) env-dep
	$(MAKE) git-dep
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

