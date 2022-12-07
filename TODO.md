# make

reusable make file includes, with standardised variables names

## TODO

Fly and Providence

- caddy is put on, so we can run many examples inside it.
- provide is neeed to be created a simple wbe page.
    - git source
    - introduction to say what is he use of the gui.
    - deep source: what is this dependent on at runtime ? 
        - server URL 
            - can ts providence. Its basically like go mod graph !!


NATS

- need easy to setup defaults
- need GIO GUI
 - Web version pipes the commands to the server, runs them there and returns then.
 - Non web version does them locally.
 - https://github.com/choria-io/appbuilder seems to do that !!
 - NATS CTL was built with appbuilder..
 - at https://github.com/search?q=org%3Achoria-io+natsctl&type=code


Make this golang based ? SO that it can be used with Decky. 
The whole architecture of Decky relies on running small workers pipelined together with some AST.
The makefiles are decent, but we really want just golang instead of makefiles, so that we can version and get compile time and runtime erors. 

Also the big question is if the same thng runs the the Designer and Runtime of Decky.
My gut instinct is to keep it similar to how a developer works with simple CLI replacing each Makefile and what it does.
Templates are good.
Makefiles often instal sometime, and this has become a generic pattern in my makefile. E.g Variable Overrides, which are really like a a go config with overrides from Env. 
The Print command is really nice 
The use of Folders and paths is also nice, because its clear what is going on.

https://github.com/go-task/task
examples
- https://github.com/omniskop/vitrum/blob/main/Taskfile.yml



Examples at https://github.com/gedw99/gio-make-examples
- should i move them here ? yes but only if they are not part of the assets. Assets has 2 git examples there are enough for basics.

PR's
- often i want to git clone someones PR, and to find the branch can be tedious. Make it simple to find the branch.
- maybe even list all PR's and the branch is what we want, then you can just modify your main makefile to pull that branch.

BUILD binaries and release them
- Get Github CI working off make, and then get the the CI to place the binaries into a github release
- Sometimes you want to place the binaries onto a server somewhere too.

go module usage and checkponting.. Not really useful right now..
- go-mod-upgrade works, but github.com/tailscale/depaware looks better because it also tells you about build tags and cgo.
    - also tells you what depends on what properly.
    - try it on gio as this uses both.
    - ex: https://github.com/tailscale/tailscale/blob/main/Makefile#L11

GIT tagging
- Need tagging added. We have the ability to clone by tag already.
- If the consumer has a source code dependency they can just use go-mod-upgrade.
- If the consumer has a binary dependency, then they can just manage it from their makefile.

CI for Latest versus Non latest
- CI_TIP should always pull in binary or source dependencies from the latest, so we instatly see when things break between projects.
- CI_DEV & CI_PROD should manually change their dependencies.

VSCODE 
- Need ability to debug without any settings in the project, and having to open it.
- flags dont seem to work
- DOCKER: Ops has a helper for debugging a docker with golang project inside. We should formalise this into vscode, since it relies on vscode and docker.

DOCKER
- Reality is that we need to use docker for easy upgrading of the servers.
- https://github.com/tailscale/mkctr looks really nice for pure golang projects like mine, and it supports muliple binaries ( arm, non arm, etc )
    - ex: https://github.com/tailscale/tailscale/blob/01a9906bf839f04733057f321265382a52faa389/build_docker.sh#L35


GCLOUD
- Need ability to easily deploy to CloudRun, and map to Domain.
    - graphjin has this ?
- Versioning is also important so that the name of something on gcloud shows what git SHA ( and / or git tag)

GO
- obfuscation using gable