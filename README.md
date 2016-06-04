# watch [![Travis-CI](https://api.travis-ci.org/ostera/watch.svg)](https://travis-ci.org/ostera/watch)
> ⌚ A Go alternative to GNU's watch – very useful for autorunning things!

## Installation

```
brew install ostera-watch
```

## Usage

```
Usage: watch [options] cmd

Options:

  -i, --interval <n>  interval period with unit, as 1s or 1ms
                      defaults to 1s

  -v, --version       outputs version number

  -h, --help          this help page
```

## Internals

`watch` relies on the current path and commands passed in to create a temporary file that will hold
the desired output. This files are created on `/var/tmp/watches/<id>`.

It will then procedd to clear the screen, print out the new output, and go back to sleep.

Upon termination, the corresponding file will be removed.
