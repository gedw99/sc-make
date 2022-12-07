# TASK
# https://github.com/go-task/task


TASK_BIN=task
TASK_BIN_VERSION=v3.16.0

# Override variables
# where to put the standard templates
TASK_SRC_TEMPLATES_FSPATH=$(PWD)/.templates/task
# where to look for TASKFile to use at runtime 
TASK_SRC_FSPATH=$(PWD)/

# Computed variables
# PERFECT :) https://www.systutorials.com/how-to-get-a-makefiles-directory-for-including-other-makefiles/
_TASK_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
_TASK_TEMPLATES_SOURCE=$(_TASK_SELF_DIR)/templates/task
_TASK_TEMPLATES_TARGET=$(TASK_SRC_TEMPLATES_FSPATH)

## task print, outputs all variables needed to run task
task-print:
	@echo
	@echo --- TASK ---

	@echo Deps:
	@echo TASK_BIN: 					$(TASK_BIN) installed TASK at : $(shell which $(TASK_BIN))
	@echo TASK_BIN_VERSION: 			$(TASK_BIN_VERSION)

	@echo
	@echo Override variables:
	@echo TASK_SRC_FSPATH: 				$(TASK_SRC_FSPATH)
	
	@echo
	@echo Computed variables:
	@echo _TASK_SELF_DIR:				$(_TASK_SELF_DIR)
	@echo _TASK_TEMPLATES_SOURCE: 		$(_TASK_TEMPLATES_SOURCE)
	@echo _TASK_TEMPLATES_TARGET: 		$(_TASK_TEMPLATES_TARGET)
	@echo

## prints the templates 
task-templates-print:
	@echo
	@echo listing templates ...
	cd $(_TASK_TEMPLATES_SOURCE) && $(ENV_TREE_BIN) -a -C


## task dep installs the task and mkcert binary to the go bin
## cand copies the templates up into your templates working directory
# Useful where you want to grab them and customise.

## installs task
task-dep:
	@echo
	@echo installing task tool
	go install github.com/go-task/task/v3/cmd/task@$(TASK_BIN_VERSION)
	@echo installed gio [ $(TASK_BIN_VERSION)   at : $(shell which $(TASK_BIN))


	$(MAKE) task-templates-dep

## installs the task templates into your project
task-templates-dep:
	@echo
	@echo installing task templates to your project....
	mkdir -p $(_TASK_TEMPLATES_TARGET)
	cp -r $(_TASK_TEMPLATES_SOURCE)/* $(_TASK_TEMPLATES_TARGET)/
	@echo installed task templates  at : $(_TASK_TEMPLATES_TARGET)

## task run runs task using your Caddyfile
task-run:
	cd $(TASK_SRC_FSPATH) && $(TASK_BIN) fmt -overwrite
	open https://localhost:8443
	cd $(TASK_SRC_FSPATH) && $(TASK_BIN) run -watch