# https://github.com/joeblew99/gio-make

GIT_REPO_URL=https://github.com/joeblew99/gio-make

MAKE_FSPATH=$(PWD)
include ./assets/make/help.mk

GIO_MAKE_BIN_NAME=gio-make
GIO_MAKE_BIN_FSPATH=$(GOPATH)/bin
GIO_MAKE_BIN=$(GIO_MAKE_BIN_FSPATH)/$(GIO_MAKE_BIN_NAME)

print:
	@echo 
	@echo --- GIO-MAKE ---
	@echo 'GIT_REPO_URL:            $(GIT_REPO_URL)'
	@echo 'GIO_MAKE_BIN_NAME:       $(GIO_MAKE_BIN_NAME)'
	@echo 'GIO_MAKE_BIN_FSPATH:     $(GIO_MAKE_BIN_FSPATH)'
	@echo 'GIO_MAKE_BIN:            $(GIO_MAKE_BIN_NAME) installed at : $(shell which $(GIO_MAKE_BIN))'


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
	go build -o $(GIO_MAKE_BIN) .
	$(MAKE) print

## deletes from go bin
install-delete:
	rm -rf $(GIO_MAKE_BIN)
	$(MAKE) print



### TEST

TEST_FSPATH=$(PWD)/_test

## test help 
test-help:
	$(GIO_MAKE_BIN_NAME) -h

## test version 
test-version:
	$(GIO_MAKE_BIN_NAME) version


## list outputs the embedded files
test-list:
	$(GIO_MAKE_BIN_NAME) list

test-create-help:
	$(GIO_MAKE_BIN_NAME) create -h


## test create with no agrs
test-create: test-create-delete
	mkdir -p $(TEST_FSPATH)/create
	cd $(TEST_FSPATH)/create && $(GIO_MAKE_BIN_NAME) create

## test create will create a project from an URL.
test-create-git: test-create-delete
	mkdir -p $(TEST_FSPATH)/create
	cd $(TEST_FSPATH)/create && $(GIO_MAKE_BIN_NAME) create --git https://github.com/gioui/gio-example --template git

## test create
test-create-delete:
	rm -rf $(TEST_FSPATH)/create
	


	