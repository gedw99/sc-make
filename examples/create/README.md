# Create


gio-make create --git https://github.com/gioui/gio-example --template git

If no args, then crates an empty one

if with --git, then modifies the root make file with the correct value.
- defaults to the git template.

## Template

first we need to get a root makefile from the "assets/project/template" folder based on the name passe din.

gio-make create --git https://github.com/gioui/gio-example --template git

## git

Then we need to modify that root makefile with the correct git args.

