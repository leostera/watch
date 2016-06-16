# watch [![Travis-CI](https://api.travis-ci.org/ostera/watch.svg)](https://travis-ci.org/ostera/watch)
> ⌚ A portable Go alternative to GNU's watch – very useful for autorunning things!

## Installation

```
go get github.com/ostera/watch

brew install https://raw.githubusercontent.com/ostera/homebrew-core/master/Formula/go-watch.rb
```

From source just run `make` and put the `watch` executable somewhere handy.

## Usage

```
~ λ watch

   Usage: watch [options] <cmd>

   Sample: watch -i=100ms make

   Options:

     -c, --clear                clear screen between command runs
     -i, --interval             interval in seconds or ms, defaulting to 1s
     -x, --exit                 exit on failure

     -h, --help                 this help page
     -v, --version              print out version

```

## Motivation

The main pain point was re-running tests/builds, choosing to clear the screen or not,
and having UNICODE/Emoji support.

I started off using OS X's watch, and eventually moving to GNU's watch. Unfortunately
neither has sub-second resolution, nor work with Emojis.

So I ended up moving to Visionmedia's watch. Now this last one doesn't clear the screen.

The middleground was a small, poor-man's watch, written in a few lines of bash:

```bash
function ww {
  readonly OUTPUT_FILE="/var/tmp/ww_files/$(pwd | sed 's / _ g')-$1"
  while true; do
    eval $* > $OUTPUT_FILE;
    clear;
    cat $OUTPUT_FILE;
    sleep 1;
  done
}
```

Hacky, but it worked pretty well! Now it was about time I wrote it in a way that
was easy to install, share, test, port, and contribute to :)

## Contributing

Build once, then use `./watch` to continuously build itself:

```
repos/watch λ ./watch make
/usr/local/bin/go vet
/usr/local/bin/go fmt
/usr/local/bin/go build -o ./watch
/usr/local/bin/go test
PASS
ok      _/Users/leostera/repos/watch    0.013s
/usr/local/bin/go test -bench .
PASS
BenchmarkRunSuccessfully-4           500           3335180 ns/op
BenchmarkRunExit-4                   500           3356888 ns/op
ok      _/Users/leostera/repos/watch    4.028s
exit: 0
```
