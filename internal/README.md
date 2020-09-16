# internal Folder

## Overview

This folder holds all of the private library code used in this service. This
means it's not meant for use with other projects. Public usable code should be
put into a `pkg/` folder.

Most code will go into the `pkg/` folder.

## Test Code

Test code shoud live in the same directory as the actual code. Not in a separate
directory. Go is smart enough to not add the `_test.go` files to its resutlant
artifact.
