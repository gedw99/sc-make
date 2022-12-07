MAKE_FSPATH=$(PWD)/../..

include $(MAKE_FSPATH)/sgit.mk
include $(MAKE_FSPATH)/env.mk
include $(MAKE_FSPATH)/help.mk

SGIT_TEMPLATE_REPO_URL=https://github.com/audrenbdb/goinstall
print:
	$(MAKE) env-print
	$(MAKE) sgit-print

init:
	$(MAKE) sgit-run

dep-os:
	$(MAKE) env-dep
