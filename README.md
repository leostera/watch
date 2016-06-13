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

The `Makefile` includes useful targets. The one run by Travis is just `make`, which looks like this

```
repos/watch λ make
/usr/local/bin/go build
/usr/local/bin/go test
PASS
ok      _/Users/leostera/repos/watch    0.014s
/usr/local/bin/go test -bench .
PASS
BenchmarkRunSuccessfully-4           300           4206227 ns/op
BenchmarkRunExit-4                   300           4217427 ns/op
BenchmarkIntervalToTime-4       2000000000               0.67 ns/op
BenchmarkSuffixToInterval-4     30000000                52.7 ns/op
ok      _/Users/leostera/repos/watch    6.445s
```
