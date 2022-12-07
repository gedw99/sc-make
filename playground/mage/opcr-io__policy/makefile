# https://github.com/opcr-io/policy

include ./../../../assets/make/help.mk

REPO_NAME=policy


# wrapper for mage. Mage is not needed to be installed
GO_MAGE_CMD=cd $(REPO_NAME) && go run mage.go
MAGE_BIN_NAME=$(REPO_NAME)-mage
MAGE_BIN_CMD=cd $(REPO_NAME) && ./$(MAGE_BIN_NAME)
BIN_NAME=$(REPO_NAME)
BIN_CMD=cd $(REPO_NAME) && ./$(BIN_NAME)


print:
	@echo ""
	@echo "REPO_NAME:      $(REPO_NAME)"
	
	@echo ""
	@echo "GO_MAGE_CMD:    $(GO_MAGE_CMD)"
	@echo "MAGE_BIN_NAME:  $(MAGE_BIN_NAME)"
	@echo "MAGE_BIN_CMD:   $(MAGE_BIN_CMD)"
	@echo "BIN_NAME:       $(BIN_NAME)"
	@echo ""

dep-git:
	git clone git@github.com:opcr-io/policy.git
	@echo $(REPO_NAME) >> .gitignore

dep-git-delete:
	rm -rf $(REPO_NAME)

mod:
	cd $(REPO_NAME) && go mod tidy
	cd $(REPO_NAME) && go mod verify

mod-upgrade:
	cd $(REPO_NAME) && go-mod-upgrade

## list all mage targets
go-l:
	$(GO_MAGE_CMD) -v -l
T=?
## run a mage target (ex: " make T=-l go-run ")
go-t:
	# make T=-l go-run
	$(GO_MAGE_CMD) $(T)

## compile magefiles to a binary. Useful for end users. Can them be a dep for a higher level project.
go-c:
	$(GO_MAGE_CMD) -compile ./$(MAGE_BIN_NAME)






## list any target
mage-l: 
	$(MAGE_BIN_CMD) -l

# make T=-l go-run: lists targets
# make T=deps -d mage-run: runs deps with debug printing
# make T=deps -d mage-run: runs deps
## run any mage target (ex: " make T=-l go-run ")
mage-t:
	$(MAGE_BIN_CMD) $(T) 



this-t:
	@echo "run the bin"
	

