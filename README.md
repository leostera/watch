# watch [![Travis-CI](https://api.travis-ci.org/ostera/watch.svg)](https://travis-ci.org/ostera/watch)
> ⌚ A portable Go alternative to GNU's watch – very useful for autorunning things!

## Installation

```
brew install ostera-watch
```

From source just run `make` and put the `watch` executable somewhere handy.

## Usage

```
~ λ watch

   Usage: watch [options] <cmd>

   Sample: watch -i=100ms make

   Options:

     -i, --interval             interval in seconds or ms, defaulting to 1s
     -x, --exit                 exit on failure
     -v, --version              print out version
     -h, --help                 this help page

```

## Contributing

Build once, then use `./watch` to continuously build itself:

```
repos/watch λ ./watch make
/usr/local/bin/go vet
/usr/local/bin/go fmt
/usr/local/bin/go build
/usr/local/bin/go test
PASS
ok      _/Users/leostera/repos/watch    0.015s
/usr/local/bin/go test -bench .
PASS
BenchmarkRunSuccessfully-4           300           4697357 ns/op
BenchmarkRunExit-4                   300           4388672 ns/op
BenchmarkIntervalToTime-4       2000000000               0.70 ns/op
BenchmarkSuffixToInterval-4     50000000                42.1 ns/op
ok      _/Users/leostera/repos/watch    8.128s
```
