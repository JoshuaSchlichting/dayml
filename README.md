# *dayml*

`dayml` is a simple tool for parsing TODO task status from daily notes written in `.yml` files.

## *`.yml` files*
`.yml` files contain _at least_ the following:
```
<yyymmdd>:
  todo:
    a todo objective: <bool representing completed status>
    another task: true # completed
    yet another task: false # yet to be completed
```
This means that you could have other tags or data in the document, but it will not be looked at by `dayml` as of the time of writing this.

## *Usage*
`dayml -f <filename>`

Or, execute `dayml` from a directory with a `.dayml.yml` file in it.



> This repo was initialized via `cobra-cli init dayml`.