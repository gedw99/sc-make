# https://gist.githubusercontent.com/tdegrunt/045f6b3377f3f7ffa408/raw/259357f481f3c58aaf4255347636d8994e67cd49/walk.go

print:

test-run:test-run-clean
	clear
	cd test && go build .
	cd test && ./test

test-run-clean:
	cd test && rm -rf ./makefile

git-get-run:
	git clone git@github.com:FalcoSuessgott/gitget.git
	gitget -h