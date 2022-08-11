# Anyver

**Anyver** is a command line interface (CLI) tool to select different versions of an application
to run globally on a Linux machine.

## Overview

Imagine you have a single binary executable file CLI application called `automagic`. It's on
version 1.0.0, but 1.0.1 has just been released and you want to have both on your system.

With **Anyver**, you can quickly select one of them to run globally. Just put in a directory the 
binaries with meaningful names, like `automagic-1.0.0` and `automagic-1.0.1`, for example. Then run
`anyver configure --directory <directory full path>` to configure the path for this directory
and run `anyver configure --name automagic` to set `automagic` as the name for the application.

By running `anyver set`, you will see both versions listed and will be able to select the one you want.
**Anyver** will just create a symbolic link to your home directory under `/.local/bin/`, which
should be in your PATH environment variable to work globally.

## Installation

Just download the binary from the releases page and use it. To make it available globally on your system,
just put it under your PATH environment variable.

You can also clone this repository and run `go build`, if you have Go installed, which will create the `anyver`
binary.

## Commands

```
completion  Generate the autocompletion script for the specified shell
configure   Configure different options for this tool.
help        Help about any command
list        List binaries available on directory set.
set         Select the binary you want from a list.
```

### Configure flags

```
-d, --directory string   Set the directory where your binaries are.
-n, --name string        Set the global name for your binary.
```

## Supported operating systems

Currently this application runs only on Linux.
