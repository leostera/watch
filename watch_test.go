package main

import(
  "testing"
)

func fixtureCmdSlice() []string {
  return []string { "echo", "hello" }
}

func fixtureErrorCmdSlice() []string {
  return []string { "exit", "1" }
}

func fixtureCmdString() string {
  return "echo hello"
}

func TestBuildArgs(t *testing.T) {
  str := buildArgs(fixtureCmdSlice())
  if str != fixtureCmdString() {
    t.Fatalf("%s should be %s", str, fixtureCmdString())
  }
}

func TestRunSuccessfully(t *testing.T) {
  ok := run(fixtureCmdSlice())
  if ok != 0 {
    t.Fatalf("%s should be 0", ok)
  }
}

func TestRunExit(t *testing.T) {
  err := run(fixtureErrorCmdSlice())
  if err != 1 {
    t.Fatalf("%s should not be 1", err)
  }
}
