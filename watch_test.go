package main

import(
  "testing"
)

func fixtureCmd() string { return "exit" }
func fixtureArgs() []string { return []string { "0" } }
func fixtureBadArgs() []string { return []string { "1" } }

func fixtureCmdSlice(args []string) []string {
  return append( []string { fixtureCmd() }, args...)
}

func TestBuildArgs(t *testing.T) {
  args := fixtureArgs()
  cmd := fixtureCmdSlice(args)
  str := buildArgs(cmd)
  cmd_str := buildArgs(cmd)
  if str != cmd_str {
    t.Fatalf("%s should be %s", str, cmd_str)
  }
}

func TestRunSuccessfully(t *testing.T) {
  ok := run(fixtureCmdSlice(fixtureArgs()))
  if ok != 0 {
    t.Fatalf("%s should be 0", ok)
  }
}

func TestRunExit(t *testing.T) {
  err := run(fixtureCmdSlice(fixtureBadArgs()))
  if err != 1 {
    t.Fatalf("%s should not be 1", err)
  }
}

func BenchmarkBuildArgs(b *testing.B) {
  cmd := fixtureCmdSlice(fixtureArgs())
  for n := 0; n < b.N; n++ {
    buildArgs(cmd)
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

func BenchmarkWrapForShell(b *testing.B) {
  cmd := buildArgs(fixtureCmdSlice(fixtureBadArgs()))
  for n := 0; n < b.N; n++ {
    wrapForShell(cmd)
  }
}

func BenchmarkGetShell(b *testing.B) {
  for n := 0; n < b.N; n++ {
    getShell()
  }
}

func BenchmarkIntervalToTime(b *testing.B) {
  for n := 0; n < b.N; n++ {
    intervalToTime(1)
  }
}
