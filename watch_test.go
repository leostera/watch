package main

import(
  "testing"
)

func fixCmdSlice() []string {
  return []string { "echo", "hello" }
}

func fixCmdString() string {
  return "echo hello"
}

func TestBuildArgs(t *testing.T) {
  str := buildArgs(fixCmdSlice())
  if str != fixCmdString() {
    t.Fatalf("%s should be %s", str, fixCmdString())
  }
}
