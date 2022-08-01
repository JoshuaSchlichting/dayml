# *dayml*

`dayml` is a simple tool for parsing TODO task status from daily notes written in `.yml` files.

## *`.yml` files*
`.yml` files contain _at least_ the following:
```
<yyyymmdd>:
  todo:
    a todo objective: <bool representing completed status>
    another task: true # completed
    yet another task: false # yet to be completed
```
This means that you could have other tags or data in the document, but it will not be looked at by `dayml` as of the time of writing this.

## *Usage*
`dayml -f <filename>`

Or, execute `dayml` from a directory with a `.dayml.yml` file in it.

## *Installing*
Executing `./build.sh` will output the file to `./bin/dayml_<os>_<arch>`

Executing `./build.sh --install` will output the file to `./bin/dayml_<os>_<arch>` and also copy the binary to `/usr/local/bin`. 

> This repo was initialized via `cobra-cli init dayml`.