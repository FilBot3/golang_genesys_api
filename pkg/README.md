# pkg Folder

## Overview

This folder should hold the Publically reusable code such as API Clients or
Utility functions which may be handy for other projects, but don't qualify for
its own repository. For things that are specific and usable by only this
application, then put it into the `internal/` folder.

Most code will probably go in here.

## Test Code

Test code shoud live in the same directory as the actual code. Not in a separate
directory. Go is smart enough to not add the `_test.go` files to its resutlant
artifact.
