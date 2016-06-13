package main

import(
  "testing"
)

func fixtureCmd() string { return "ls" }
func fixtureArgs() []string { return []string { "." } }
func fixtureBadArgs() []string { return []string { "wat" } }

func fixtureCmdSlice(args []string) []string {
  return append( []string { fixtureCmd() }, args...)
}

func TestRunSuccessfully(t *testing.T) {
  ok := run(fixtureCmdSlice(fixtureArgs()))
  if ok != 0 {
    t.Fatalf("%s should be 0", ok)
  }
}

func TestRunExit(t *testing.T) {
  err := run(fixtureCmdSlice(fixtureBadArgs()))
  if err != 2 {
    t.Fatalf("%s should not be 2", err)
  }
}

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
