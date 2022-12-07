# todo

## boot strapper

This is written in golang and so availabel locally.
It then pulls the make files from github.

Takes a github url and scaffolds out all the needed make files, so that really fast to start a project.

- needs a makefile go-template, so we can replace the correct git attributes like server, org, repo.

- might need to distinguish if its an origin or upstream on the commandline flags.

## env

loads https://github.com/joho/godotenv and then looks for environments files in the file system.

this allows you to have one .env file that all your git repos use.

you can then load the into your ci secrets. https://docs.github.com/en/actions/learn-github-actions/environment-variables


## github

Puts the correct github workflos and others that works with the other files into the .github folder, so that you have automatic CI.

- see the gio-printer for owkring examples, along with its makefile that is designed to work with it for mono repos.


## vscode

Puts the correct vscode file inthe repo.

- this needs launch.json go template so that the names of the commands are replaced.
- maybe also add the ability to launch the debugger from the makefile. Not sure how to do this.

## inject

Injects the makefiles into the public git repo, so that its easy to put automate a repo.

## ops - ansible deployment

I tend to reply on NATS for everything in my deployments. lets not get into why - long story.

But for a fresh server, you want certain things to be bootstrapped onto the server that are not part of oyur main binary.

One way is ansible roles.
https://github.com/create-go-app/cli/tree/master/pkg/registry/roles

This system uses ansible todeploy stuff for you. Its nats based and complex though.
https://github.com/choria-io
