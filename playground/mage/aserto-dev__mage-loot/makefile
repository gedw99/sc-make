# https://github.com/aserto-dev/mage-loot

# All works well
# Combine with gotasks for the final high level control, because magefiles are too labouriosu at the top level.

REPO_NAME=mage-loot

BIN_NAME=static-output

print:

dep-git:
	git clone git@github.com:aserto-dev/mage-loot.git
	@echo $(REPO_NAME) >> .gitignore

dep-tools:
	go install github.com/magefile/mage@latest
dep-tools-delete:
	rm -f $(shell which mage)

dep-git-delete:
	rm -rf $(REPO_NAME)

mod:
	cd $(REPO_NAME) && go mod tidy
	cd $(REPO_NAME) && go mod verify

## THIS - only works if compiled. Use go-compile first :)

# Override this is you want 
RUN_TARGET=??
this-dep: go-compile

this-run:
	# make RUN_TARGET=-l this-run
	cd $(REPO_NAME)/magefiles && ./static-output $(RUN_TARGET)
this-help: this-dep
	cd $(REPO_NAME)/magefiles && ./static-output -h
this-list:
	cd $(REPO_NAME)/magefiles && ./static-output -l

this-deps:
	cd $(REPO_NAME)/magefiles && ./static-output -h deps
	cd $(REPO_NAME)/magefiles && ./static-output -v deps
this-lint:
	cd $(REPO_NAME)/magefiles && ./static-output -h lint
	cd $(REPO_NAME)/magefiles && ./static-output -v lint


## GO - only works in go installed

# Override this is you want 
RUN_TARGET=??
go-run:
	# make RUN_TARGET=-l go-run
	cd $(REPO_NAME) && go run mage.go $(RUN_TARGET)
go-list:
	cd $(REPO_NAME) && go run mage.go -l
go-deps:
	cd $(REPO_NAME) && go run mage.go deps
go-lint:
	cd $(REPO_NAME) && go run mage.go lint
go-test:
	cd $(REPO_NAME) && go run mage.go test
go-compile:
	cd $(REPO_NAME) && go run mage.go -compile ./$(BIN_NAME)
go-compile-delete:
	rm -f $(REPO_NAME)/magefiles/$(BIN_NAME)


### MAGE - only works if mage installed. Just needed this to test higher level this. Dont need it.

mage-list:
	cd $(REPO_NAME) && mage -l
mage-run:
	cd $(REPO_NAME) && mage $(RUN_TARGET)
mage-compile:
	cd $(REPO_NAME) && mage -compile ./$(BIN_NAME)
