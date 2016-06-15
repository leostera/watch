package main

import (
	"runtime"
	"testing"
)

var platform struct {
	cmd  string
	flag string
}

func init() {
	if runtime.GOOS == "windows" {
		platform.cmd = "cmd"
		platform.flag = "/C"
	} else {
		platform.cmd = "sh"
		platform.flag = "-c"
	}
}

/******************************************************************************/
// Fixtures
/******************************************************************************/

func fixtureCmd() string       { return platform.cmd }
func fixtureArgs() []string    { return []string{platform.flag, "exit 0"} }
func fixtureBadArgs() []string { return []string{platform.flag, "exit 2"} }

func fixtureCmdSlice(args []string) []string {
	return append([]string{fixtureCmd()}, args...)
}

/******************************************************************************/
// Tests
/******************************************************************************/

func TestSuffixToInterval1MS(t *testing.T) {
	val, _ := suffixToInterval("MS", "1MS")
	if val != 1 {
		t.Fatalf("%f should be 1", val)
	}
}

func TestSuffixToInterval1ms(t *testing.T) {
	val, _ := suffixToInterval("ms", "1ms")
	if val != 1 {
		t.Fatalf("%f should be 1", val)
	}
}

func TestSuffixToInterval1S(t *testing.T) {
	val, _ := suffixToInterval("S", "1S")
	if val != 1000.0 {
		t.Fatalf("%f should be 100.0", val)
	}
}

func TestSuffixToInterval1s(t *testing.T) {
	val, _ := suffixToInterval("s", "1s")
	if val != 1000.0 {
		t.Fatalf("%f should be 1000.0", val)
	}
}

func TestRunSuccessfully(t *testing.T) {
	ok := run(fixtureCmdSlice(fixtureArgs()))
	if ok != 0 {
		t.Fatalf("%d should be 0", ok)
	}
}

func TestRunExit(t *testing.T) {
	err := run(fixtureCmdSlice(fixtureBadArgs()))
	if err != 2 {
		t.Fatalf("%d should not be 2", err)
	}
}

/******************************************************************************/
// Benchmarks
/******************************************************************************/

func BenchmarkRunSuccessfully(b *testing.B) {
	cmd := fixtureCmdSlice(fixtureArgs())
	for n := 0; n < b.N; n++ {
		run(cmd)
	}
}

func BenchmarkRunExit(b *testing.B) {
	cmd := fixtureCmdSlice(fixtureBadArgs())
	for n := 0; n < b.N; n++ {
		run(cmd)
	}
}

func BenchmarkIntervalToTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		intervalToTime(1)
	}
}

func BenchmarkSuffixToInterval(b *testing.B) {
	for n := 0; n < b.N; n++ {
		suffixToInterval("MS", "1MS")
	}
}
