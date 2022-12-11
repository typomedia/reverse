# `revert` - A GNU `rev` Clone

Reverse lines characterwise. Written in [Go](https://go.dev/) 1.19.

## Usage

    Usage: reverse [options] [file ...]

    Reverse lines characterwise.
    
    Options:
    -h, --help     display this help
    -V, --version  display version

## Examples

    echo "hello world" > test.txt
    reverse test.txt
    dlrow olleH

## Build

    make build

## Requirements

Only needed for embedding a resource icon on Windows:

    go install github.com/akavel/rsrc@latest
    rsrc -ico go.ico
    make build

---
Copyright © 2022 Typomedia Foundation. All rights reserved.
