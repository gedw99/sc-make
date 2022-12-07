# https://github.com/create-go-app/cli

REPO_NAME=cli
REPO_FSPATH=$(PWD)/$(REPO_NAME)
REPO=https://github.com/create-go-app/$(REPO_NAME)

BIN_NAME=cli
BIN_FSPATH=$(PWD)/_bin
BIN=$(BIN_FSPATH)/$(BIN_NAME)
# add bin to system path ( https://stackoverflow.com/questions/8941110/how-i-could-add-dir-to-path-in-makefile ) 
export PATH:=$(BIN_FSPATH):$(PATH)
PATH_CMD=export PATH:=$(BIN_FSPATH):$(PATH); cli

print:

	@echo 
	@echo REPO:			$(REPO)

	@echo BIN:			$(BIN)
	@echo BIN_NAME:		$(BIN_NAME)
	@echo PATH_CMD:		$(PATH_CMD)


dep: dep-delete
	git clone $(REPO)

dep-delete:
	rm -rf $(REPO_FSPATH)

build:
	cd $(REPO_FSPATH)/cmd/cgapp && go build -o $(BIN) .


run-export:
	which $(BIN_NAME)

run-help:
	$(PATH_CMD)

TEST_FSPATH=$(PWD)/_test

test-init:
	mkdir -p $(TEST_FSPATH)
test-create: test-init
	cd $(TEST_FSPATH)/create && $(PATH_CMD) create
	
