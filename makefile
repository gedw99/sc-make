# https://github.com/gedw99/sc-make

GIT_REPO_URL=https://github.com/gedw99/sc-make

MAKE_FSPATH=$(PWD)
include ./assets/make/help.mk

SC_MAKE_BIN_NAME=sc-make
SC_MAKE_BIN_FSPATH=$(GOPATH)/bin
SC_MAKE_BIN=$(SC_MAKE_BIN_FSPATH)/$(SC_MAKE_BIN_NAME)

print:
	@echo 
	@echo --- GIO-MAKE ---
	@echo 'GIT_REPO_URL:           $(GIT_REPO_URL)'
	@echo 'SC_MAKE_BIN_NAME:       $(SC_MAKE_BIN_NAME)'
	@echo 'SC_MAKE_BIN_FSPATH:     $(SC_MAKE_BIN_FSPATH)'
	@echo 'SC_MAKE_BIN:            $(SC_MAKE_BIN_NAME) installed at : $(shell which $(SC_MAKE_BIN))'


### GIT

COMMIT_MESSAGE='added more ...'
## git-commit
git-commit:
	git add --all
	git commit -am $(COMMIT_MESSAGE)

## git-pull
git-pull:
	git pull 

## git-push
git-push:
	git push 
	open $(GIT_REPO_URL)



### INSTALL 

## install to go bin
install:
	go generate -v ./...
	#cd assets && go run generate.go
	go build -o $(SC_MAKE_BIN) .
	$(MAKE) print

## deletes from go bin
install-delete:
	rm -rf $(SC_MAKE_BIN)
	$(MAKE) print



### TEST

TEST_FSPATH=$(PWD)/_test

## test help 
test-help:
	$(SC_MAKE_BIN_NAME) -h

## test version 
test-version:
	$(SC_MAKE_BIN_NAME) version


## list outputs the embedded files
test-list:
	$(SC_MAKE_BIN_NAME) list

test-create-help:
	$(SC_MAKE_BIN_NAME) create -h


## test create with no agrs
test-create: test-create-delete
	mkdir -p $(TEST_FSPATH)/create
	cd $(TEST_FSPATH)/create && $(SC_MAKE_BIN_NAME) create

## test create will create a project from an URL.
test-create-git: test-create-delete
	mkdir -p $(TEST_FSPATH)/create
	cd $(TEST_FSPATH)/create && $(SC_MAKE_BIN_NAME) create --git https://github.com/gioui/gio-example --template git

## test create
test-create-delete:
	rm -rf $(TEST_FSPATH)/create
	


	