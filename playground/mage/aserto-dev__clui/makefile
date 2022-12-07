# https://github.com/aserto-dev/clui

REPO_NAME=clui

print:

dep-git:
	git clone git@github.com:aserto-dev/clui.git
	@echo $(REPO_NAME) >> .gitignore

dep-git-delete:
	rm -rf $(REPO_NAME)

run:
	cd $(REPO_NAME)/example && go run .