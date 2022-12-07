# https://github.com/create-go-app/cli

REPO_NAME=cli
REPO_FSPATH=$(PWD)/$(REPO_NAME)
REPO=https://github.com/create-go-app/$(REPO_NAME)

BIN_NAME=cli
BIN_FSPATH=$(PWD)/_bin
BIN=$(BIN_FSPATH)/$(BIN_NAME)
# add bin to system path ( https://stackoverflow.com/questions/8941110/how-i-could-add-dir-to-path-in-makefile ) 
export PATH:=$(BIN_FSPATH):$(PATH)
BIN_PATH_CMD=export PATH:=$(BIN_FSPATH):$(PATH); cli

print:

	@echo 
	@echo REPO:				$(REPO)

	@echo BIN:				$(BIN)
	@echo BIN_NAME:			$(BIN_NAME)
	@echo BIN_PATH_CMD:		$(BIN_PATH_CMD)


dep: dep-delete
	git clone $(REPO)

dep-delete:
	rm -rf $(REPO_FSPATH)

build: build-delete
	cd $(REPO_FSPATH)/cmd/cgapp && go build -ldflags="-s -w" -o $(BIN) .
build-delete:
	rm -rf $(BIN_FSPATH)

run-help:
	$(BIN_PATH_CMD)

TEST_FSPATH=$(PWD)/_test

test-dep: test-delete
	mkdir -p $(TEST_FSPATH)
test-delete:
	rm -rf $(TEST_FSPATH)

test-new: test-dep
	rm -rf $(TEST_FSPATH)/new 
	mkdir -p $(TEST_FSPATH)/new 
	cd $(TEST_FSPATH)/new && $(BIN_PATH_CMD) new

test-create: test-dep
	mkdir -p $(TEST_FSPATH)/create 
	cd $(TEST_FSPATH)/create && $(BIN_PATH_CMD) create
	cd $(TEST_FSPATH)/create && $(BIN_PATH_CMD) deploy
	

	
