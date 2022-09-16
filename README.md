# Anyver

**Anyver** is a command line interface (CLI) tool to select different versions of an application
to run globally on a Linux machine.

## Overview

Imagine you have a single binary executable file CLI application called `automagic`. It's on
version 1.0.0, but 1.0.1 has just been released and you want to have both on your system.

With **Anyver**, you can quickly select one of them to run globally. Just put in a directory the 
binaries with meaningful names, like `automagic-1.0.0` and `automagic-1.0.1`, for example. Then run
`anyver add automagic --directory <directory full path>` to add this tool and configure the path for
it.

Then run `anyver set` to select one of the binaries. You will see both versions listed and will be able 
to select the one you want. **Anyver** will just create a symbolic link to your home 
directory under `/.local/bin/`, which should be in your `PATH` environment variable to work globally.

## Installation

Just download the binary from the releases page and use it. To make it available globally on your system,
just put it under your `PATH` environment variable.

You can also clone this repository and run `go build`, if you have Go installed, which will create the `anyver`
binary.

## Commands

```
add         Add a command line interface (CLI) tool for usage
change      Change the command line interface (CLI) tool for usage
completion  Generate the autocompletion script for the specified shell
help        Help about any command
list        List binaries available on directory set
remove      Remove a command line interface (CLI) tool from configuration
set         Select the binary you want from a list
```

## Supported operating systems

Currently this application runs only on Linux.
