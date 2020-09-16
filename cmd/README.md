# cmd Folder

## Overview

This folder holds the application main entry point for the project. The directory
name should match the name of the project so when it's installed or build, the
resulting binary matches. For instances: `cmd/simple-service` will result in
`simple-service` as the binary name.

This folder allows the repository to have multiple commands or binaries in the
same project. Just simply add a new folder, and a main.go in it.