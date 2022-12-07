# GIO-MAKE

TODO: move examples into here...

Reusable git and make templates for my projects.

it has:

- golang
- gio
- caddy

## Install 

Compile golang bin into binary called **gio-make**.

The following commands are available.

**gio-make-clone**

```gio-make-clone --urlhttps://github.com/sebastianpacurar/gioui-experiment --mode single``` 

Clones a github repo and setups up the make configuration.

- url:
  - the git repo url.
- template:
  - single.
  - multi.

logic:
- need the makefile to be a template, so we can modify it at runtime.
  - pull from github or embedded ? embedded is best.
  - so get embedding working.
- calculate fodler name from URL ( github__org__name)
- create folder
- copy ( or de embed) and files 
- find replace the make file template

---

**create-go-app**

Provides ability to copy existing github projects and confgiure them with frontend and backends.

NOTE: This is for later. Its designed for users that want to pull examples of systems i have make.

Based on https://github.com/create-go-app/cli/blob/master/.goreleaser.yml

## Make

1. sgit.mk will template out everything if you give it a git URI.

Designed to be used above a git repo to manage all the remotes and branches for busy developers.

puts all the makefiles you need for gio and golang development with defaults.


2. Backend and frontend URLS

Asks you for the backend and frontend URLS and git pulls them for you.




Is versioned in that gio-make is committed to a repo and you pull it as a submodule.

you can also inject makefile into a gitrepo as needed.

env variables driving things uses dotenv, and github secrets. See TOOD....

## Install

go install github.com/gedw99/gio-make

this has all the mame files and templates embeded.

## Usage

from a directory where oyu want to start a project run 

``` gio-make <GIT-URL>``` 

## Why ?


These makefile follow certain conventions to:

- Ensure CI and Local Dev work the same. A github workflow calls the exact same makefiles that you use localy. No Split brain. 

- Provide conventions that are extensible. Each .mk does a specific thing, allowing you to put any logic or customisations in your own root makefile where you need to do anything differently.

## Examples

Various example are in. https://github.com/gedw99/gio-make-examples

GIO
- https://github.com/gedw99/gio-make-examples/gioui__gio-example


## Conventions

There are some conventions in all this that you can follow to stay out of trouble.



## Print variable

Each makefile includes has a ```<prefix>-print``` that echo's the used variables back to you. Use it to see what MAKE variables you need in your makefile.

For example:

In [goi.mk](go.mk), you have the following variables only to worry about:

```GOI_SRC_NAME``` 

```GOI_SRC_FSPATH```


So, those variables need to exist in your own **makefile**

For example:

```GIO_SRC_NAME=goi-projectname``` 

```GIO_SRC_FSPATH=$(PWD)$(GIO_SRC_NAME)```


## Dep

Each make file includes has a ```<prefix>-dep``` that ensures any dependency it needs are installed.

This ensures both CI and Local Dev always have the dependencies and they are the same in each environment.

It also will put templates into your root folder that aid in working with that technology. 

For example in [gio.mk](gio.mk), it will ensure that the gogio compile tool is installed. The version can be reflected from the go.modules per project.

It will also put templates into your root that are useful when working with gio projects.

TODO: look into incorporating bingo. EX: https://github.com/grafana/grafana/tree/main/.bingo


## Creating your own .mk files

Use a Prefix to isolate each make target of each makefile respectively.

For example:

In [go.mk](go.mk), you have:
- ```go-build```
- ```go-run```, etc

in [git.mk](git.mk) you have:
- ```git-repo-clone-origin```
- ```git-repo-clone-upstream```, etc